/*
Copyright 2016 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package edit provides high-level auxiliary functions for AST manipulation
// on BUILD files.
//
// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/edit].
package edit

import (
	"github.com/bazel-contrib/buildtools/v10/edit"
	"github.com/bazelbuild/buildtools/build"
	"regexp"
)

var (
	// ShortenLabelsFlag if true converts added labels to short form , e.g. //foo:bar => :bar
	//
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/edit.ShortenLabelsFlag].
	// Assigning to it has no effect, assign to the v10 variable instead.
	ShortenLabelsFlag = edit.ShortenLabelsFlag

	// DeleteWithComments if true a list attribute will be be deleted in ListDelete, even if there is a comment attached to it
	//
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/edit.DeleteWithComments].
	// Assigning to it has no effect, assign to the v10 variable instead.
	DeleteWithComments = edit.DeleteWithComments
)

// InterpretLabelForWorkspaceLocation returns the name of the BUILD file to
// edit, the full package name, and the rule. It takes a workspace-rooted
// directory to use.
func InterpretLabelForWorkspaceLocation(root, target string) (buildFile, repo, pkg, rule string) {
	return edit.InterpretLabelForWorkspaceLocation(root, target)
}

// InterpretLabel returns the name of the BUILD file to edit, the full
// package name, and the rule. It uses the pwd for resolving workspace file paths.
func InterpretLabel(target string) (buildFile string, pkg string, rule string) {
	return edit.InterpretLabel(target)
}

// InterpretLabelWithRepo returns the name of the BUILD file to edit, repo name,
// the full package name, and the rule. It uses the pwd for resolving workspace
// file paths.
func InterpretLabelWithRepo(target string) (buildFile string, repo string, pkg string, rule string) {
	return edit.InterpretLabelWithRepo(target)
}

// ExprToRule returns a Rule from an Expr.
// The boolean is false iff the Expr is not a function call, or does not have
// the expected kind.
func ExprToRule(expr build.Expr, kind string) (*build.Rule, bool) {
	return edit.ExprToRule(expr, kind)
}

// ExistingPackageDeclaration returns the package declaration, or nil if there is none.
func ExistingPackageDeclaration(f *build.File) *build.Rule {
	return edit.ExistingPackageDeclaration(f)
}

// PackageDeclaration returns the package declaration. If it doesn't
// exist, it is created at the top of the BUILD file, after optional
// docstring, comments, and load statements.
func PackageDeclaration(f *build.File) *build.Rule {
	return edit.PackageDeclaration(f)
}

// RemoveEmptyPackage removes empty package declarations from the file, i.e.:
//
//	package()
//
// This might appear because of a buildozer transformation (e.g. when removing a package
// attribute). Removing it is required for the file to be valid.
func RemoveEmptyPackage(f *build.File) *build.File {
	return edit.RemoveEmptyPackage(f)
}

// RemoveEmptyUseRepoCalls removes empty use repo declarations from the file.
func RemoveEmptyUseRepoCalls(f *build.File) *build.File {
	return edit.RemoveEmptyUseRepoCalls(f)
}

// InsertAfter inserts an expression after index i.
func InsertAfter(i int, stmt []build.Expr, expr build.Expr) []build.Expr {
	return edit.InsertAfter(i, stmt, expr)
}

// IndexOfLast finds the index of the last expression of a specific kind.
func IndexOfLast(stmt []build.Expr, Kind string) int {
	return edit.IndexOfLast(stmt, Kind)
}

// InsertAfterLastOfSameKind inserts an expression after the last expression of the same kind.
func InsertAfterLastOfSameKind(stmt []build.Expr, expr *build.CallExpr) []build.Expr {
	return edit.InsertAfterLastOfSameKind(stmt, expr)
}

// InsertAtEnd inserts an expression at the end of a list, before trailing comments.
func InsertAtEnd(stmt []build.Expr, expr build.Expr) []build.Expr {
	return edit.InsertAtEnd(stmt, expr)
}

// FindRuleByName returns the rule in the file that has the given name.
// If the name is "__pkg__", it returns the global package declaration.
func FindRuleByName(f *build.File, name string) *build.Rule {
	return edit.FindRuleByName(f, name)
}

// IndexOfRuleByName returns the index (in f.Stmt) of the CallExpr which defines a rule named `name`, or -1 if it doesn't exist.
func IndexOfRuleByName(f *build.File, name string) (int, *build.Rule) {
	return edit.IndexOfRuleByName(f, name)
}

// FindExportedFile returns the first exports_files call which contains the
// file 'name', or nil if not found
func FindExportedFile(f *build.File, name string) *build.Rule {
	return edit.FindExportedFile(f, name)
}

// DeleteRule returns the AST without the specified rule
func DeleteRule(f *build.File, rule *build.Rule) *build.File {
	return edit.DeleteRule(f, rule)
}

// DeleteRuleByName returns the AST without the rules that have the
// given name.
func DeleteRuleByName(f *build.File, name string) *build.File {
	return edit.DeleteRuleByName(f, name)
}

// DeleteRuleByKind removes the rules of the specified kind from the AST.
// Returns an updated copy of f.
func DeleteRuleByKind(f *build.File, kind string) *build.File {
	return edit.DeleteRuleByKind(f, kind)
}

// AllLists returns all the lists concatenated in an expression.
// For example, in: glob(["*.go"]) + [":rule"]
// the function will return [[":rule"]].
func AllLists(e build.Expr) []*build.ListExpr {
	return edit.AllLists(e)
}

// AllSelects returns all the selects concatenated in an expression.
func AllSelects(e build.Expr) []*build.CallExpr {
	return edit.AllSelects(e)
}

// FirstList works in the same way as AllLists, except that it
// returns only one list, or nil.
func FirstList(e build.Expr) *build.ListExpr {
	return edit.FirstList(e)
}

// AllStrings returns all the string literals concatenated in an expression.
// For example, in: "foo" + x + "bar"
// the function will return ["foo", "bar"].
func AllStrings(e build.Expr) []*build.StringExpr {
	return edit.AllStrings(e)
}

// ListFind looks for a string in the list expression (which may be a
// concatenation of lists). It returns the element if it is found. nil
// otherwise.
func ListFind(e build.Expr, item string, pkg string) *build.StringExpr {
	return edit.ListFind(e, item, pkg)
}

// ContainsComments returns whether the expr has a comment that includes str.
func ContainsComments(expr build.Expr, str string) bool {
	return edit.ContainsComments(expr, str)
}

// ContainsDisableComment checks whether the expr has a comment that disables
// a warning, including when multiple warnings are disabled on the same line.
func ContainsDisableComment(expr build.Expr, warning string) bool {
	return edit.ContainsDisableComment(expr, warning)
}

// RemoveEmptySelectsAndConcatLists iterates the tree in order to turn
// empty selects into empty lists and adjacent lists are concatenated
func RemoveEmptySelectsAndConcatLists(e build.Expr) build.Expr {
	return edit.RemoveEmptySelectsAndConcatLists(e)
}

// ComputeIntersection returns the intersection of the two lists given as parameters;
// if the containing elements are not build.StringExpr, the result will be nil.
func ComputeIntersection(list1, list2 []build.Expr) []build.Expr {
	return edit.ComputeIntersection(list1, list2)
}

// SelectListsIntersection returns the intersection of the lists of strings inside
// the dictionary argument of the select expression given as a parameter
func SelectListsIntersection(sel *build.CallExpr, pkg string) (intersection []build.Expr) {
	return edit.SelectListsIntersection(sel, pkg)
}

// ResolveAttr extracts common elements of the lists inside select dictionaries
// and adds them at attribute level rather than select level, as well as turns
// empty selects into empty lists and concatenates adjacent lists
func ResolveAttr(r *build.Rule, attr, pkg string) {
	edit.ResolveAttr(r, attr, pkg)
}

// SelectDelete removes the item from all the lists which are values
// in the dictionary of every select
func SelectDelete(e build.Expr, item, pkg string, deleted **build.StringExpr) {
	edit.SelectDelete(e, item, pkg, deleted)
}

// RemoveFromList removes one element from a ListExpr and stores
// the deleted StringExpr at the address pointed by the last parameter
func RemoveFromList(li *build.ListExpr, item, pkg string, deleted **build.StringExpr) {
	edit.RemoveFromList(li, item, pkg, deleted)
}

// ListDelete deletes the item from a list expression in e and returns
// the StringExpr deleted, or nil otherwise.
func ListDelete(e build.Expr, item, pkg string) (deleted *build.StringExpr) {
	return edit.ListDelete(e, item, pkg)
}

// ListAttributeDelete deletes string item from list attribute attr, deletes attr if empty,
// and returns the StringExpr deleted, or nil otherwise.
func ListAttributeDelete(rule *build.Rule, attr, item, pkg string) *build.StringExpr {
	return edit.ListAttributeDelete(rule, attr, item, pkg)
}

// ListReplace replaces old with value in all lists in e and returns a Boolean
// to indicate whether the replacement was successful.
func ListReplace(e build.Expr, old, value, pkg string) bool {
	return edit.ListReplace(e, old, value, pkg)
}

// ListSubstitute replaces strings matching a regular expression in all lists
// in e and returns a Boolean to indicate whether the replacement was
// successful.
func ListSubstitute(e build.Expr, oldRegexp *regexp.Regexp, newTemplate string) bool {
	return edit.ListSubstitute(e, oldRegexp, newTemplate)
}

// AddValueToList adds a value to a list. If the expression is
// not a list, a list with a single element is appended to the original
// expression.
func AddValueToList(oldList build.Expr, pkg string, item build.Expr, sorted bool) build.Expr {
	return edit.AddValueToList(oldList, pkg, item, sorted)
}

// AddValueToListAttribute adds the given item to the list attribute identified by name and pkg.
func AddValueToListAttribute(r *build.Rule, name string, pkg string, item build.Expr, vars *map[string]*build.AssignExpr) {
	edit.AddValueToListAttribute(r, name, pkg, item, vars)
}

// MoveAllListAttributeValues moves all values from list attribute oldAttr to newAttr,
// and deletes oldAttr.
func MoveAllListAttributeValues(rule *build.Rule, oldAttr, newAttr, pkg string, vars *map[string]*build.AssignExpr) error {
	return edit.MoveAllListAttributeValues(rule, oldAttr, newAttr, pkg, vars)
}

// DictionarySet looks for the key in the dictionary expression. If value is not nil,
// it replaces the current value with it. In all cases, it returns the current value.
func DictionarySet(dict *build.DictExpr, key string, value build.Expr) build.Expr {
	return edit.DictionarySet(dict, key, value)
}

// DictionaryGet looks for the key in the dictionary expression, and returns the
// current value. If it is unset, it returns nil.
func DictionaryGet(dict *build.DictExpr, key string) build.Expr {
	return edit.DictionaryGet(dict, key)
}

// DictionaryDelete looks for the key in the dictionary expression. If the key exists,
// it removes the key-value pair and returns it. Otherwise it returns nil.
func DictionaryDelete(dict *build.DictExpr, key string) (deleted build.Expr) {
	return edit.DictionaryDelete(dict, key)
}

// RenameAttribute renames an attribute in a rule.
func RenameAttribute(r *build.Rule, oldName, newName string) error {
	return edit.RenameAttribute(r, oldName, newName)
}

// EditFunction is a wrapper around build.Edit. The callback is called only on
// functions 'name'.
func EditFunction(v build.Expr, name string, f func(x *build.CallExpr, stk []build.Expr) build.Expr) build.Expr {
	return edit.EditFunction(v, name, f)
}

// UsedSymbols returns the set of symbols used in the BUILD file (variables, function names).
func UsedSymbols(stmt build.Expr) map[string]bool {
	return edit.UsedSymbols(stmt)
}

// UsedTypes returns the set of types used in the BUILD file (variables, function names).
func UsedTypes(stmt build.Expr) map[string]bool {
	return edit.UsedTypes(stmt)
}

// NewLoad creates a new LoadStmt node
func NewLoad(location string, from, to []string) *build.LoadStmt {
	return edit.NewLoad(location, from, to)
}

// AppendToLoad appends symbols to an existing load statement
// Returns true if the statement was acually edited (if the required symbols haven't been
// loaded yet)
func AppendToLoad(load *build.LoadStmt, from, to []string) bool {
	return edit.AppendToLoad(load, from, to)
}

// InsertLoad inserts a load statement at the top of the list of statements.
// The load statement is constructed using a string location and two slices of from- and to-symbols.
// The function panics if the slices aren't of the same length. Symbols that are already loaded
// from the given filepath are ignored. If stmts already contains a load for the
// location in arguments, appends the symbols to load to it.
func InsertLoad(stmts []build.Expr, location string, from, to []string) []build.Expr {
	return edit.InsertLoad(stmts, location, from, to)
}

// ReplaceLoad removes load statements for passed to-symbols and replaces them with a new
// load at the top of the list of statements. The new load statement is constructed using
// a string location and two slices of from- and to-symbols. If stmts already contains a
// load for the location in arguments, appends the symbols to load to it.
// The function panics if the slices aren't of the same lentgh.
func ReplaceLoad(stmts []build.Expr, location string, from, to []string) []build.Expr {
	return edit.ReplaceLoad(stmts, location, from, to)
}

// ParseLabel parses a Blaze label (eg. //devtools/buildozer:rule), and returns
// the repo name ("" for the main repo), package (with leading slashes trimmed)
// and rule name (e.g. ["", "devtools/buildozer", "rule"]).
// Deprecated; use `labels.Parse` instead
func ParseLabel(target string) (string, string, string) {
	return edit.ParseLabel(target)
}

// ShortenLabel rewrites labels to use the canonical form (the form
// recommended by build-style).
// Doesn't do anything if `--shorten_label=false` flag is provided.
// Use `labels.Shorten` to shorten labels unconditionally.
func ShortenLabel(label, pkg string) string {
	return edit.ShortenLabel(label, pkg)
}

// LabelsEqual returns true if label1 and label2 are equal.
// Deprecated; use `labels.Equal` instead
func LabelsEqual(label1, label2, pkg string) bool {
	return edit.LabelsEqual(label1, label2, pkg)
}
