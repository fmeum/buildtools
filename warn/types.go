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

package warn

import (
	"github.com/bazel-contrib/buildtools/v10/warn"
	"github.com/bazelbuild/buildtools/build"
)

// Type describes an expression type in Starlark.
type Type = warn.Type

// List of known types
const (
	Unknown = warn.Unknown

	Bool = warn.Bool

	Ctx = warn.Ctx

	CtxActions = warn.CtxActions

	CtxActionsArgs = warn.CtxActionsArgs

	Depset = warn.Depset

	Dict = warn.Dict

	Int = warn.Int

	None = warn.None

	String = warn.String

	List = warn.List

	Float = warn.Float
)

// DetectTypes tries to infer the type of expressions in the current file, using basic heuristics.
//
// Warning: the types inferred by the function might change in the future, as we update the
// heuristics.
func DetectTypes(f *build.File) map[build.Expr]Type {
	return warn.DetectTypes(f)
}
