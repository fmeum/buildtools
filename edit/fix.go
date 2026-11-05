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

package edit

import (
	"github.com/bazel-contrib/buildtools/v10/edit"
	"github.com/bazelbuild/buildtools/build"
)

// AllRuleFixes is a list of all Buildozer fixes that can be applied on a rule.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/edit.AllRuleFixes].
// Assigning to it has no effect, assign to the v10 variable instead.
var AllRuleFixes = edit.AllRuleFixes

// FileLevelFixes is a list of all Buildozer fixes that apply on the whole file.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/edit.FileLevelFixes].
// Assigning to it has no effect, assign to the v10 variable instead.
var FileLevelFixes = edit.FileLevelFixes

// FixRule aims to fix errors in BUILD files, remove deprecated features, and
// simplify the code.
func FixRule(f *build.File, pkg string, rule *build.Rule, fixes []string) *build.File {
	return edit.FixRule(f, pkg, rule, fixes)
}

// FixFile fixes everything it can in the BUILD file.
func FixFile(f *build.File, pkg string, fixes []string) *build.File {
	return edit.FixFile(f, pkg, fixes)
}
