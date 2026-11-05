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

// FileType represents a type of a file (default (for .bzl files), BUILD, or WORKSPACE).
// Certain formatting or refactoring rules can be applied to several file types, so they support
// bitwise operations: `type1 | type2` can represent a scope (e.g. BUILD and WORKSPACE files) and
// `scope & fileType` can be used to check whether a file type belongs to a scope.
type FileType = build.FileType

const (
	// TypeDefault represents general Starlark files
	TypeDefault = build.TypeDefault

	// TypeBuild represents BUILD files
	TypeBuild = build.TypeBuild

	// TypeWorkspace represents WORKSPACE files
	TypeWorkspace = build.TypeWorkspace

	// TypeBzl represents .bzl files
	TypeBzl = build.TypeBzl

	// TypeModule represents MODULE.bazel and *.MODULE.bazel files
	TypeModule = build.TypeModule

	// TypeRepo represents REPO.bazel files
	TypeRepo = build.TypeRepo

	// TypeVendor represents VENDOR.bazel files
	TypeVendor = build.TypeVendor
)

// ParseBuild parses a file, marks it as a BUILD file and returns the corresponding parse tree.
//
// The filename is used only for generating error messages.
func ParseBuild(filename string, data []byte) (*File, error) {
	return build.ParseBuild(filename, data)
}

// ParseWorkspace parses a file, marks it as a WORKSPACE file and returns the corresponding parse tree.
//
// The filename is used only for generating error messages.
func ParseWorkspace(filename string, data []byte) (*File, error) {
	return build.ParseWorkspace(filename, data)
}

// ParseModule parses a file, marks it as a MODULE.bazel file and returns the corresponding parse tree.
//
// The filename is used only for generating error messages.
func ParseModule(filename string, data []byte) (*File, error) {
	return build.ParseModule(filename, data)
}

// ParseRepo parses a file, marks it as a REPO.bazel file and returns the corresponding parse tree.
//
// The filename is used only for generating error messages.
func ParseRepo(filename string, data []byte) (*File, error) {
	return build.ParseRepo(filename, data)
}

// ParseVendor parses a file, marks it as a VENDOR.bazel file and returns the corresponding parse tree.
//
// The filename is used only for generating error messages.
func ParseVendor(filename string, data []byte) (*File, error) {
	return build.ParseVendor(filename, data)
}

// ParseBzl parses a file, marks it as a .bzl file and returns the corresponding parse tree.
//
// The filename is used only for generating error messages.
func ParseBzl(filename string, data []byte) (*File, error) {
	return build.ParseBzl(filename, data)
}

// ParseDefault parses a file, marks it as a generic Starlark file and returns the corresponding parse tree.
//
// The filename is used only for generating error messages.
func ParseDefault(filename string, data []byte) (*File, error) {
	return build.ParseDefault(filename, data)
}

// Parse parses the input data and returns the corresponding parse tree.
//
// Uses the filename to detect the formatting type (build, workspace, or default) and calls
// ParseBuild, ParseWorkspace, or ParseDefault correspondingly.
func Parse(filename string, data []byte) (*File, error) {
	return build.Parse(filename, data)
}

// ParseError contains information about the error encountered during parsing.
type ParseError = build.ParseError
