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

// Unquote unquotes the quoted string, returning the actual
// string value, whether the original was triple-quoted, and
// an error describing invalid input.
func Unquote(quoted string) (s string, triple bool, err error) {
	return build.Unquote(quoted)
}

// IsCorrectEscaping reports whether a string doesn't contain any incorrectly
// escaped sequences such as "\a".
func IsCorrectEscaping(value string) bool {
	return build.IsCorrectEscaping(value)
}
