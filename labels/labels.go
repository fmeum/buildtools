/*
 * Copyright 2020 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package labels contains helper functions for working with labels.
//
// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/labels].
package labels

import (
	"github.com/bazel-contrib/buildtools/v10/labels"
)

// Label represents a Bazel target label.
type Label = labels.Label

// Parse parses an absolute Bazel label (eg. //devtools/buildozer:rule)
// and returns the corresponding Label object.
func Parse(target string) Label {
	return labels.Parse(target)
}

// ParseRelative parses a label `input` which may be absolute or relative.
// If it's relative then it's considered to belong to `pkg`
func ParseRelative(input, pkg string) Label {
	return labels.ParseRelative(input, pkg)
}

// Shorten rewrites labels to use the canonical form (the form
// recommended by build-style).
// "//foo/bar:bar" => "//foo/bar", or ":bar" if the label belongs to pkg
func Shorten(input, pkg string) string {
	return labels.Shorten(input, pkg)
}

// Equal returns true if label1 and label2 are equal. The function
// takes care of the optional ":" prefix and differences between long-form
// labels and local labels (relative to pkg).
func Equal(label1, label2, pkg string) bool {
	return labels.Equal(label1, label2, pkg)
}
