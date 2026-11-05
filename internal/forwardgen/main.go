// Command forwardgen generates the forwarding packages of this module.
//
// For every exported top-level identifier of a package of the old module
// (github.com/bazelbuild/buildtools at the commit before the rename) it emits:
//   - types:     type T = v10pkg.T
//   - functions: func F(...) { return v10pkg.F(...) }
//   - constants: const C = v10pkg.C
//   - variables: var V = v10pkg.V
//
// It also checks that the v10 API is a superset of the old API (modulo the
// module path) and reports every difference it finds as a PROBLEM line.
//
// This directory is a separate module, so it is not part of
// github.com/bazelbuild/buildtools. To regenerate the forwarding packages,
// delete them and run, from the root of this branch:
//
//	git worktree add --detach /tmp/old eaa4d125b423
//	git worktree add --detach /tmp/v10 v10.0.0
//	(cd internal/forwardgen && go run . -old /tmp/old -new /tmp/v10 -out ../..)
//
// The main packages (buildifier, buildozer, unused_deps) are not generated:
// they are copies of their v10 sources.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/tools/go/packages"
)

const (
	oldMod = "github.com/bazelbuild/buildtools"
	newMod = "github.com/bazel-contrib/buildtools/v10"
)

var problems []string

func problem(format string, args ...any) {
	problems = append(problems, fmt.Sprintf(format, args...))
}

func load(dir string) map[string]*packages.Package {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax |
			packages.NeedTypes | packages.NeedTypesInfo | packages.NeedImports |
			packages.NeedDeps,
		Dir:        dir,
		BuildFlags: []string{"-mod=mod"},
	}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		log.Fatal(err)
	}
	m := map[string]*packages.Package{}
	for _, p := range pkgs {
		m[p.PkgPath] = p
	}
	return m
}

// qualifier maps both module paths onto the old one so that type strings
// from the two universes can be compared.
func qualifier(p *types.Package) string {
	path := p.Path()
	if path == newMod || strings.HasPrefix(path, newMod+"/") {
		path = oldMod + strings.TrimPrefix(path, newMod)
	}
	return path
}

func typeString(t types.Type) string { return types.TypeString(t, qualifier) }

func main() {
	oldDir := flag.String("old", "", "checkout of the old module")
	newDir := flag.String("new", "", "checkout of the v10 module")
	outDir := flag.String("out", "", "directory to write the forwarding packages to")
	flag.Parse()

	oldPkgs := load(*oldDir)
	newPkgs := load(*newDir)

	var paths []string
	for p := range oldPkgs {
		paths = append(paths, p)
	}
	sort.Strings(paths)
	for _, path := range paths {
		op := oldPkgs[path]
		if op.Name == "main" {
			continue
		}
		if len(op.Errors) > 0 {
			log.Fatalf("%s: %v", path, op.Errors)
		}
		rel := strings.TrimPrefix(path, oldMod)
		np := newPkgs[newMod+rel]
		if np == nil {
			problem("%s: no v10 package", path)
			continue
		}
		checkPackage(op, np)
		generatePackage(op, np, filepath.Join(*outDir, filepath.FromSlash(strings.TrimPrefix(rel, "/"))))
	}

	for _, p := range problems {
		fmt.Println("PROBLEM:", p)
	}
}

// checkPackage reports exported API of op that is missing or different in np.
func checkPackage(op, np *packages.Package) {
	oldScope := op.Types.Scope()
	newScope := np.Types.Scope()
	for _, name := range oldScope.Names() {
		oo := oldScope.Lookup(name)
		if !oo.Exported() {
			continue
		}
		no := newScope.Lookup(name)
		where := op.PkgPath + "." + name
		if no == nil {
			problem("%s: missing in v10", where)
			continue
		}
		if fmt.Sprintf("%T", oo) != fmt.Sprintf("%T", no) {
			problem("%s: kind changed from %T to %T", where, oo, no)
			continue
		}
		if a, b := typeString(oo.Type()), typeString(no.Type()); a != b {
			problem("%s: type changed from %s to %s", where, a, b)
		}
		switch oo := oo.(type) {
		case *types.Const:
			if a, b := oo.Val().ExactString(), no.(*types.Const).Val().ExactString(); a != b {
				problem("%s: value changed from %s to %s", where, a, b)
			}
		case *types.TypeName:
			checkType(where, oo.Type(), no.Type())
		}
	}
}

func checkType(where string, ot, nt types.Type) {
	if a, b := typeString(ot.Underlying()), typeString(nt.Underlying()); a != b {
		// Report field and interface differences in detail.
		switch ou := ot.Underlying().(type) {
		case *types.Struct:
			nu, ok := nt.Underlying().(*types.Struct)
			if !ok {
				problem("%s: underlying type changed from %s to %s", where, a, b)
				break
			}
			nf := map[string]*types.Var{}
			for i := 0; i < nu.NumFields(); i++ {
				nf[nu.Field(i).Name()] = nu.Field(i)
			}
			for i := 0; i < ou.NumFields(); i++ {
				f := ou.Field(i)
				if !f.Exported() {
					continue
				}
				g := nf[f.Name()]
				if g == nil {
					problem("%s.%s: field missing in v10", where, f.Name())
				} else if typeString(f.Type()) != typeString(g.Type()) {
					problem("%s.%s: field type changed from %s to %s", where, f.Name(), typeString(f.Type()), typeString(g.Type()))
				}
			}
			problem("%s: struct changed (informational): %s -> %s", where, a, b)
		case *types.Interface:
			problem("%s: interface changed from %s to %s", where, a, b)
		default:
			problem("%s: underlying type changed from %s to %s", where, a, b)
		}
	}
	oms := types.NewMethodSet(types.NewPointer(ot))
	nms := types.NewMethodSet(types.NewPointer(nt))
	if _, ok := ot.Underlying().(*types.Interface); ok {
		oms = types.NewMethodSet(ot)
		nms = types.NewMethodSet(nt)
	}
	for i := 0; i < oms.Len(); i++ {
		sel := oms.At(i)
		if !sel.Obj().Exported() {
			continue
		}
		nsel := nms.Lookup(sel.Obj().Pkg(), sel.Obj().Name())
		if nsel == nil {
			// Lookup with the new package for unexported names; exported
			// names are package independent.
			for j := 0; j < nms.Len(); j++ {
				if nms.At(j).Obj().Name() == sel.Obj().Name() {
					nsel = nms.At(j)
				}
			}
		}
		if nsel == nil {
			problem("%s.%s: method missing in v10", where, sel.Obj().Name())
			continue
		}
		if a, b := typeString(sel.Type()), typeString(nsel.Type()); a != b {
			problem("%s.%s: method type changed from %s to %s", where, sel.Obj().Name(), a, b)
		}
	}
}

// docLines converts a comment group to "//"-prefixed lines, dropping
// directives.
func docLines(cg *ast.CommentGroup) []string {
	if cg == nil {
		return nil
	}
	text := cg.Text()
	text = strings.TrimRight(text, "\n")
	if text == "" {
		return nil
	}
	var lines []string
	for _, l := range strings.Split(text, "\n") {
		switch {
		case l == "":
			lines = append(lines, "//")
		case strings.HasPrefix(l, "\t"):
			lines = append(lines, "//"+l)
		default:
			lines = append(lines, "// "+l)
		}
	}
	return lines
}

type gen struct {
	buf     bytes.Buffer
	fset    *token.FileSet
	pkgName string // name of the v10 package in generated code
	newPath string
}

func (g *gen) p(format string, args ...any) {
	fmt.Fprintf(&g.buf, format, args...)
	g.buf.WriteByte('\n')
}

func (g *gen) node(n ast.Node) string {
	var b bytes.Buffer
	if err := printer.Fprint(&b, g.fset, n); err != nil {
		log.Fatal(err)
	}
	return b.String()
}

func (g *gen) doc(indent string, cg *ast.CommentGroup, name string) {
	g.docKind(indent, cg, name, false)
}

func (g *gen) docKind(indent string, cg *ast.CommentGroup, name string, isVar bool) {
	lines := docLines(cg)
	if isVar {
		if len(lines) > 0 {
			lines = append(lines, "//")
		}
		lines = append(lines, fmt.Sprintf("// This is a copy of [%s.%s].", g.newPath, name),
			"// Assigning to it has no effect, assign to the v10 variable instead.")
	}
	for _, l := range lines {
		g.p("%s%s", indent, l)
	}
}

func exported(name string) bool { return ast.IsExported(name) }

func (g *gen) funcDecl(d *ast.FuncDecl) {
	if d.Recv != nil || !exported(d.Name.Name) {
		return
	}
	if d.Type.TypeParams != nil {
		problem("%s.%s: generic function needs manual handling", g.newPath, d.Name.Name)
	}
	var params, args []string
	n := 0
	variadic := false
	for _, f := range d.Type.Params.List {
		typ := g.node(f.Type)
		if _, ok := f.Type.(*ast.Ellipsis); ok {
			variadic = true
		}
		names := f.Names
		if len(names) == 0 {
			names = []*ast.Ident{{Name: "_"}}
		}
		var ns []string
		for _, id := range names {
			name := id.Name
			if name == "_" {
				name = fmt.Sprintf("p%d", n)
			}
			if name == g.pkgName {
				problem("%s.%s: parameter %s shadows the v10 package", g.newPath, d.Name.Name, name)
			}
			n++
			ns = append(ns, name)
			args = append(args, name)
		}
		params = append(params, strings.Join(ns, ", ")+" "+typ)
	}
	if variadic {
		args[len(args)-1] += "..."
	}
	results := ""
	if d.Type.Results != nil && len(d.Type.Results.List) > 0 {
		var rs []string
		for _, f := range d.Type.Results.List {
			var ns []string
			for _, id := range f.Names {
				ns = append(ns, id.Name)
			}
			r := g.node(f.Type)
			if len(ns) > 0 {
				r = strings.Join(ns, ", ") + " " + r
			}
			rs = append(rs, r)
		}
		if len(rs) == 1 && len(d.Type.Results.List[0].Names) == 0 {
			results = " " + rs[0]
		} else {
			results = " (" + strings.Join(rs, ", ") + ")"
		}
	}
	g.doc("", d.Doc, d.Name.Name)
	g.p("func %s(%s)%s {", d.Name.Name, strings.Join(params, ", "), results)
	call := fmt.Sprintf("%s.%s(%s)", g.pkgName, d.Name.Name, strings.Join(args, ", "))
	if results != "" {
		g.p("\treturn %s", call)
	} else {
		g.p("\t%s", call)
	}
	g.p("}")
	g.p("")
}

func (g *gen) genDecl(d *ast.GenDecl) {
	switch d.Tok {
	case token.IMPORT:
		return
	case token.TYPE:
		for _, s := range d.Specs {
			s := s.(*ast.TypeSpec)
			if !exported(s.Name.Name) {
				continue
			}
			if s.TypeParams != nil {
				problem("%s.%s: generic type needs manual handling", g.newPath, s.Name.Name)
			}
			doc := s.Doc
			if doc == nil && len(d.Specs) == 1 {
				doc = d.Doc
			}
			g.doc("", doc, s.Name.Name)
			g.p("type %s = %s.%s", s.Name.Name, g.pkgName, s.Name.Name)
			g.p("")
		}
	case token.CONST, token.VAR:
		kw := d.Tok.String()
		type item struct {
			doc     *ast.CommentGroup
			comment *ast.CommentGroup
			names   []string
		}
		var items []item
		for _, s := range d.Specs {
			s := s.(*ast.ValueSpec)
			var names []string
			for _, id := range s.Names {
				if exported(id.Name) {
					names = append(names, id.Name)
				}
			}
			if len(names) > 0 {
				items = append(items, item{s.Doc, s.Comment, names})
			}
		}
		if len(items) == 0 {
			return
		}
		isVar := d.Tok == token.VAR
		if len(items) == 1 && len(items[0].names) == 1 && (!d.Lparen.IsValid() || len(d.Specs) == 1) {
			it := items[0]
			doc := it.doc
			if doc == nil {
				doc = d.Doc
			}
			g.docKind("", doc, it.names[0], isVar)
			line := fmt.Sprintf("%s %s = %s.%s", kw, it.names[0], g.pkgName, it.names[0])
			if it.comment != nil {
				line += " // " + strings.TrimSpace(it.comment.Text())
			}
			g.p("%s", line)
			g.p("")
			return
		}
		if lines := docLines(d.Doc); len(lines) > 0 {
			for _, l := range lines {
				g.p("%s", l)
			}
		}
		g.p("%s (", kw)
		for i, it := range items {
			for j, name := range it.names {
				var doc *ast.CommentGroup
				if j == 0 {
					doc = it.doc
				}
				if i > 0 || j > 0 {
					if doc != nil || len(it.names) == 1 {
						g.p("")
					}
				}
				g.docKind("\t", doc, name, isVar)
				line := fmt.Sprintf("\t%s = %s.%s", name, g.pkgName, name)
				if it.comment != nil && j == len(it.names)-1 {
					line += " // " + strings.TrimSpace(it.comment.Text())
				}
				g.p("%s", line)
			}
		}
		g.p(")")
		g.p("")
	}
}

// leadingComment returns the license header of f, if any.
func leadingComment(f *ast.File) *ast.CommentGroup {
	if len(f.Comments) == 0 {
		return nil
	}
	c := f.Comments[0]
	if c == f.Doc || c.Pos() > f.Package {
		return nil
	}
	first := c.List[0].Text
	if strings.HasPrefix(first, "//go:build") || strings.HasPrefix(first, "// +build") ||
		strings.HasPrefix(first, "//line ") || strings.HasPrefix(first, "// Code generated") {
		return nil
	}
	return c
}

func generatePackage(op, np *packages.Package, dir string) {
	fset := op.Fset
	// Pick a name for the v10 import that doesn't clash with imports used in
	// signatures.
	pkgName := np.Name
	for _, f := range op.Syntax {
		for _, imp := range f.Imports {
			path, _ := strconv.Unquote(imp.Path.Value)
			name := ""
			if imp.Name != nil {
				name = imp.Name.Name
			} else if ip := op.Imports[path]; ip != nil {
				name = ip.Name
			}
			if name == pkgName {
				pkgName = "v10" + np.Name
			}
		}
	}

	var docFile string
	for _, f := range op.Syntax {
		if f.Doc != nil {
			docFile = fset.File(f.Pos()).Name()
		}
	}

	wrote := false
	for i, f := range op.Syntax {
		filename := fset.File(f.Pos()).Name()
		g := &gen{fset: fset, pkgName: pkgName, newPath: np.PkgPath}
		for _, d := range f.Decls {
			switch d := d.(type) {
			case *ast.FuncDecl:
				g.funcDecl(d)
			case *ast.GenDecl:
				g.genDecl(d)
			}
		}
		body := g.buf.String()
		isDocFile := filename == docFile || (docFile == "" && i == 0)
		if body == "" && !isDocFile {
			continue
		}

		// Assemble the file.
		var out bytes.Buffer
		if lc := leadingComment(f); lc != nil {
			for _, c := range lc.List {
				out.WriteString(c.Text)
				out.WriteByte('\n')
			}
			out.WriteByte('\n')
		}
		if isDocFile {
			pkgDoc := f.Doc
			if pkgDoc != nil && strings.HasPrefix(pkgDoc.Text(), "Generated file") {
				pkgDoc = nil
			}
			lines := docLines(pkgDoc)
			if len(lines) > 0 {
				lines = append(lines, "//")
			}
			lines = append(lines, fmt.Sprintf("// All exported identifiers of this package forward to [%s].", np.PkgPath))
			for _, l := range lines {
				out.WriteString(l + "\n")
			}
		}
		fmt.Fprintf(&out, "package %s\n\n", op.Name)

		// Imports: the v10 package plus every import of the old file that
		// the generated code still references.
		used := map[string]bool{}
		if body != "" {
			bf, err := parser.ParseFile(token.NewFileSet(), "", "package x\n"+body, parser.SkipObjectResolution)
			if err != nil {
				log.Fatalf("%s: %v\n%s", filename, err, body)
			}
			ast.Inspect(bf, func(n ast.Node) bool {
				if sel, ok := n.(*ast.SelectorExpr); ok {
					if id, ok := sel.X.(*ast.Ident); ok {
						used[id.Name] = true
					}
				}
				return true
			})
		}
		var imports []string
		v10Name := ""
		if pkgName != np.Name {
			v10Name = pkgName + " "
		}
		if used[pkgName] {
			imports = append(imports, fmt.Sprintf("%s%q", v10Name, np.PkgPath))
		} else if isDocFile {
			// Keep side effects (init functions) of the v10 package.
			imports = append(imports, fmt.Sprintf("_ %q", np.PkgPath))
		}
		for _, imp := range f.Imports {
			path, _ := strconv.Unquote(imp.Path.Value)
			name := ""
			explicit := ""
			if imp.Name != nil {
				name = imp.Name.Name
				explicit = name + " "
			} else if ip := op.Imports[path]; ip != nil {
				name = ip.Name
			}
			if used[name] {
				imports = append(imports, fmt.Sprintf("%s%q", explicit, path))
			}
		}
		if len(imports) > 0 {
			out.WriteString("import (\n")
			for _, imp := range imports {
				out.WriteString("\t" + imp + "\n")
			}
			out.WriteString(")\n\n")
		}
		out.WriteString(body)

		src, err := format.Source(out.Bytes())
		if err != nil {
			log.Fatalf("%s: %v\n%s", filename, err, out.String())
		}
		base := filepath.Base(filename)
		base = strings.Replace(base, ".gen.pb.go", ".go", 1)
		base = strings.Replace(base, ".gen.go", ".go", 1)
		base = strings.Replace(base, ".y.go", ".go", 1)
		dst := filepath.Join(dir, base)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			log.Fatal(err)
		}
		if err := os.WriteFile(dst, src, 0o644); err != nil {
			log.Fatal(err)
		}
		wrote = true
	}
	if !wrote {
		problem("%s: nothing written", op.PkgPath)
	}
}
