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

// Package testutils provides some useful helpers for buildozer/buildifer tests.
//
// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/testutils].
package testutils

import (
	"github.com/bazel-contrib/buildtools/v10/testutils"
	"testing"
)

// Diff returns the output of running diff on b1 and b2.
func Diff(b1, b2 []byte) ([]byte, error) {
	return testutils.Diff(b1, b2)
}

// Tdiff logs the Diff output to t.Error.
func Tdiff(t *testing.T, a, b []byte) {
	testutils.Tdiff(t, a, b)
}
