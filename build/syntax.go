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

// Package build implements parsing and printing of BUILD files.
//
// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/build].
package build

import (
	"github.com/bazel-contrib/buildtools/v10/build"
)

// A Position describes the position between two bytes of input.
type Position = build.Position

// An Expr represents an input element.
type Expr = build.Expr

// A Comment represents a single # comment.
type Comment = build.Comment

// Comments collects the comments associated with an expression.
type Comments = build.Comments

// A File represents an entire BUILD or .bzl file.
type File = build.File

// A CommentBlock represents a top-level block of comments separate
// from any rule.
type CommentBlock = build.CommentBlock

// An Ident represents an identifier.
type Ident = build.Ident

// An TypedIdent represents an identifier with type annotation: "foo: int".
//
// May occur as a statement ("var statement") or as the LHS of AssignExpr when used as a statement.
type TypedIdent = build.TypedIdent

// A TypeAppExpr represents a type application: a type_name followed by a type_list: struct[{"name": str}, ...], typing.Sequence[int].
type TypeAppExpr = build.TypeAppExpr

// An EllipsisExpr represents the '...' token in a type_arg.
type EllipsisExpr = build.EllipsisExpr

// BranchStmt represents a `pass`, `break`, or `continue` statement.
type BranchStmt = build.BranchStmt

// A LiteralExpr represents a literal number.
type LiteralExpr = build.LiteralExpr

// A StringExpr represents a single literal string.
type StringExpr = build.StringExpr

// An End represents the end of a parenthesized or bracketed expression.
// It is a place to hang comments.
type End = build.End

// A CallExpr represents a function call expression: X(List).
type CallExpr = build.CallExpr

// A DotExpr represents a field selector: X.Name.
type DotExpr = build.DotExpr

// A Comprehension represents a list comprehension expression: [X for ... if ...].
type Comprehension = build.Comprehension

// A ForClause represents a for clause in a list comprehension: for Var in Expr.
type ForClause = build.ForClause

// An IfClause represents an if clause in a list comprehension: if Cond.
type IfClause = build.IfClause

// A KeyValueExpr represents a dictionary entry: Key: Value.
type KeyValueExpr = build.KeyValueExpr

// A DictExpr represents a dictionary literal: { List }.
type DictExpr = build.DictExpr

// A ListExpr represents a list literal: [ List ].
type ListExpr = build.ListExpr

// A SetExpr represents a set literal: { List }.
type SetExpr = build.SetExpr

// A TupleExpr represents a tuple literal: (List)
type TupleExpr = build.TupleExpr

// A UnaryExpr represents a unary expression: Op X.
type UnaryExpr = build.UnaryExpr

// A BinaryExpr represents a binary expression: X Op Y.
type BinaryExpr = build.BinaryExpr

// An AssignExpr represents a binary expression with `=`: LHS = RHS.
type AssignExpr = build.AssignExpr

// A ParenExpr represents a parenthesized expression: (X).
type ParenExpr = build.ParenExpr

// A SliceExpr represents a slice expression: expr[from:to] or expr[from:to:step] .
type SliceExpr = build.SliceExpr

// An IndexExpr represents an index expression: X[Y].
type IndexExpr = build.IndexExpr

// A Function represents the common parts of LambdaExpr and DefStmt
type Function = build.Function

// A LambdaExpr represents a lambda expression: lambda Var: Expr.
type LambdaExpr = build.LambdaExpr

// ConditionalExpr represents the conditional: X if TEST else ELSE.
type ConditionalExpr = build.ConditionalExpr

// A LoadStmt loads another module and binds names from it:
// load(Module, "x", y="foo").
//
// The AST is slightly unfaithful to the concrete syntax here because
// Skylark's load statement, so that it can be implemented in Python,
// binds some names (like y above) with an identifier and some (like x)
// without.  For consistency we create fake identifiers for all the
// strings.
type LoadStmt = build.LoadStmt

// A TypeAliasStmt represents a type alias definition: "type Foo = int" or "type Foo[T, U] = dict[T, U]".
type TypeAliasStmt = build.TypeAliasStmt

// A DefStmt represents a function definition expression: def foo(List):.
type DefStmt = build.DefStmt

// A ReturnStmt represents a return statement: return f(x).
type ReturnStmt = build.ReturnStmt

// A ForStmt represents a for loop block: for x in range(10):.
type ForStmt = build.ForStmt

// An IfStmt represents an if-else block: if x: ... else: ... .
// `elif`s are treated as a chain of `IfStmt`s.
type IfStmt = build.IfStmt
