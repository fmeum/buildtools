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
)

// IsList returns true for all attributes whose type is a list.
func IsList(attr string) bool {
	return edit.IsList(attr)
}

// IsIntList returns true for all attributes whose type is an int list.
func IsIntList(attr string) bool {
	return edit.IsIntList(attr)
}

// IsString returns true for all attributes whose type is a string or a label.
func IsString(attr string) bool {
	return edit.IsString(attr)
}

// IsStringDict returns true for all attributes whose type is a string dictionary.
func IsStringDict(attr string) bool {
	return edit.IsStringDict(attr)
}

// ContainsLabels returns true for all attributes whose type is a label or a label list.
func ContainsLabels(kind, attr string) bool {
	return edit.ContainsLabels(kind, attr)
}
