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

// Package file provides utility file operations.
//
// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/file].
package file

import (
	"github.com/bazel-contrib/buildtools/v10/file"
)

// ReadFile can be updated from the caller to change the API
// for reading a file.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/file.ReadFile].
// Assigning to it has no effect, assign to the v10 variable instead.
var ReadFile = file.ReadFile

// WriteFile can be updated from the caller to change the API
// for writing a file.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/file.WriteFile].
// Assigning to it has no effect, assign to the v10 variable instead.
var WriteFile = file.WriteFile

// WriteFileMode can be updated from the caller to change the API
// for writing a file with a specific mode.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/file.WriteFileMode].
// Assigning to it has no effect, assign to the v10 variable instead.
var WriteFileMode = file.WriteFileMode

// OpenReadFile can be updated from the caller to change the API
// for opening a file.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/file.OpenReadFile].
// Assigning to it has no effect, assign to the v10 variable instead.
var OpenReadFile = file.OpenReadFile

// DisableSymlinkSafety disables symlink safety checks for WriteFile and WriteFileMode.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/file.DisableSymlinkSafety].
// Assigning to it has no effect, assign to the v10 variable instead.
var DisableSymlinkSafety = file.DisableSymlinkSafety
