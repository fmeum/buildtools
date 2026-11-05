/*
Copyright 2021 Google LLC

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

// GetParamName extracts the param name from an item of function params.
func GetParamName(param Expr) (name, op string) {
	return build.GetParamName(param)
}

// GetParamIdent extracts the param identifier from an item of function params.
func GetParamIdent(param Expr) (ident *Ident, op string) {
	return build.GetParamIdent(param)
}

// GetTypes returns the list of types defined by the a given expression.
// Examples:
//
// List[tuple[bool, int]] should return [List, Tuple, bool, int]
// str should return str
func GetTypes(t Expr) []string {
	return build.GetTypes(t)
}

// IsMultiLine returns whether an Expr is multiline or not.
func IsMultiLine(param Expr) bool {
	return build.IsMultiLine(param)
}
