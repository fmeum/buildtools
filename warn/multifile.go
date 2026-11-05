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
)

// FileReader is a class that can read an arbitrary Starlark file
// from the repository and cache the results.
type FileReader = warn.FileReader

// NewFileReader creates and initializes a FileReader instance with a
// custom readFile function that can read an arbitrary file in the
// repository using a path relative to the workspace root
// (OS-independent, with forward slashes).
func NewFileReader(readFile func(string) ([]byte, error)) *FileReader {
	return warn.NewFileReader(readFile)
}
