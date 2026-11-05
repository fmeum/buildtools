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

// DisableRewrites disables certain rewrites (for debugging).
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/build.DisableRewrites].
// Assigning to it has no effect, assign to the v10 variable instead.
var DisableRewrites = build.DisableRewrites

// AllowSort allows sorting of these lists even with sorting otherwise disabled (for debugging).
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/build.AllowSort].
// Assigning to it has no effect, assign to the v10 variable instead.
var AllowSort = build.AllowSort

// Rewriter controls the rewrites to be applied.
//
// If non-nil, the rewrites with the specified names will be run. If
// nil, a default set of rewrites will be used that is determined by
// the type (BUILD vs default starlark) of the file being rewritten.
type Rewriter = build.Rewriter

// Rewrite applies rewrites to a file
func Rewrite(f *File) {
	build.Rewrite(f)
}

// SortStringList sorts x, a list of strings.
// The list is broken by non-strings and by blank lines and comments into chunks.
// Each chunk is sorted in place.
func SortStringList(x Expr) {
	build.SortStringList(x)
}

// SortLoadArgs sorts a load statement arguments (lexicographically, but positional first)
func SortLoadArgs(load *LoadStmt) bool {
	return build.SortLoadArgs(load)
}
