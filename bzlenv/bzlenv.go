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

// Package bzlenv provides function to create and update a static environment.
//
// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/bzlenv].
package bzlenv

import (
	"github.com/bazel-contrib/buildtools/v10/bzlenv"
	"github.com/bazelbuild/buildtools/build"
)

// ValueKind describes how a binding was declared.
type ValueKind = bzlenv.ValueKind

// List of ValueKind values.
const (
	Builtin = bzlenv.Builtin // language builtin

	Imported = bzlenv.Imported // declared with load()

	Global = bzlenv.Global // declared with assignment on top-level

	Function = bzlenv.Function // declared with a def

	Parameter = bzlenv.Parameter // function parameter

	Local = bzlenv.Local // local variable, defined with assignment or as a loop variable

	TypeParameter = bzlenv.TypeParameter // type parameter of a type alias or a generic function
)

// NameInfo represents information about a symbol name.
type NameInfo = bzlenv.NameInfo

// Environment represents a static environment (e.g. information about all available symbols).
type Environment = bzlenv.Environment

// NewEnvironment creates a new empty Environment.
func NewEnvironment() *Environment {
	return bzlenv.NewEnvironment()
}

// CollectLValues returns the list of identifiers that are assigned (assuming that node is a valid
// LValue). For example, it returns `a`, `b` and `c` for the input `a, (b, c)`.
func CollectLValues(node build.Expr) []*build.Ident {
	return bzlenv.CollectLValues(node)
}

// WalkOnceWithEnvironment calls fct on every child of node, while maintaining the Environment of all available symbols.
func WalkOnceWithEnvironment(node build.Expr, env *Environment, fct func(e *build.Expr, env *Environment)) {
	bzlenv.WalkOnceWithEnvironment(node, env, fct)
}
