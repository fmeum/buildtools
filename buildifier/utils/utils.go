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

// Package utils contains shared methods that can be used by different implementations of
// buildifier binary
//
// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/buildifier/utils].
package utils

import (
	"github.com/bazel-contrib/buildtools/v10/buildifier/utils"
	"github.com/bazelbuild/buildtools/build"
	"github.com/bazelbuild/buildtools/warn"
)

// ExpandDirectories takes a list of file/directory names and returns a list with file names
// by traversing each directory recursively and searching for relevant Starlark files.
func ExpandDirectories(args *[]string) ([]string, error) {
	return utils.ExpandDirectories(args)
}

// GetParser returns a parser for a given file type
func GetParser(inputType string) func(filename string, data []byte) (*build.File, error) {
	return utils.GetParser(inputType)
}

// Lint calls the linter and returns a list of unresolved findings
func Lint(f *build.File, lint string, warningsList *[]string, verbose bool) []*warn.Finding {
	return utils.Lint(f, lint, warningsList, verbose)
}
