/*
Copyright 2020 Google LLC

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

package build

import (
	"github.com/bazel-contrib/buildtools/v10/build"
)

// StopTraversalError is a special error that tells the walker to not traverse
// further and visit child nodes of the current node.
type StopTraversalError = build.StopTraversalError

// Walk walks the expression tree v, calling f on all subexpressions
// in a preorder traversal.
//
// The stk argument is the stack of expressions in the recursion above x,
// from outermost to innermost.
func Walk(v Expr, f func(x Expr, stk []Expr)) {
	build.Walk(v, f)
}

// WalkPointers is the same as Walk but calls the callback function with pointers to nodes.
func WalkPointers(v Expr, f func(x *Expr, stk []Expr)) {
	build.WalkPointers(v, f)
}

// WalkInterruptable is the same as Walk but allows traversal to be interrupted.
func WalkInterruptable(v Expr, f func(x Expr, stk []Expr) error) {
	build.WalkInterruptable(v, f)
}

// Edit walks the expression tree v, calling f on all subexpressions
// in a preorder traversal. If f returns a non-nil value, the tree is mutated.
// The new value replaces the old one.
//
// The stk argument is the stack of expressions in the recursion above x,
// from outermost to innermost.
func Edit(v Expr, f func(x Expr, stk []Expr) Expr) Expr {
	return build.Edit(v, f)
}

// EditChildren is similar to Edit but doesn't visit the initial node, instead goes
// directly to its children.
func EditChildren(v Expr, f func(x Expr, stk []Expr) Expr) {
	build.EditChildren(v, f)
}

// WalkOnce calls f on every child of v.
func WalkOnce(v Expr, f func(x *Expr)) {
	build.WalkOnce(v, f)
}

// WalkStatements traverses sub statements (not all nodes)
func WalkStatements(v Expr, f func(x Expr, stk []Expr) error) {
	build.WalkStatements(v, f)
}
