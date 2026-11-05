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

package build

import (
	"github.com/bazel-contrib/buildtools/v10/build"
)

// A Rule represents a single BUILD rule.
type Rule = build.Rule

// NewRule is a simple constructor for Rule.
func NewRule(call *CallExpr) *Rule {
	return build.NewRule(call)
}

// Strings returns expr as a []string.
// If expr is not a list of string literals,
// Strings returns a nil slice instead.
// If expr is an empty list of string literals,
// returns a non-nil empty slice.
// (this allows differentiating between these two cases)
func Strings(expr Expr) []string {
	return build.Strings(expr)
}
