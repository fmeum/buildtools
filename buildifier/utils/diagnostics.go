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

package utils

import (
	"github.com/bazel-contrib/buildtools/v10/buildifier/utils"
	"github.com/bazelbuild/buildtools/warn"
)

// Diagnostics contains diagnostic information returned by formatter and linter
type Diagnostics = utils.Diagnostics

// FileDiagnostics contains diagnostics information for a file
type FileDiagnostics = utils.FileDiagnostics

// NewDiagnostics returns a new Diagnostics object
func NewDiagnostics(fileDiagnostics ...*FileDiagnostics) *Diagnostics {
	return utils.NewDiagnostics(fileDiagnostics...)
}

// NewFileDiagnostics returns a new FileDiagnostics object
func NewFileDiagnostics(filename string, warnings []*warn.Finding) *FileDiagnostics {
	return utils.NewFileDiagnostics(filename, warnings)
}

// InvalidFileDiagnostics returns a new FileDiagnostics object for an invalid file
func InvalidFileDiagnostics(filename string) *FileDiagnostics {
	return utils.InvalidFileDiagnostics(filename)
}
