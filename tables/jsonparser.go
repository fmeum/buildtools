/*
Copyright 2017 Google LLC

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

// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/tables].
package tables

import (
	"github.com/bazel-contrib/buildtools/v10/tables"
)

type Definitions = tables.Definitions

// ParseJSONDefinitions reads and parses JSON table definitions from file.
func ParseJSONDefinitions(file string) (Definitions, error) {
	return tables.ParseJSONDefinitions(file)
}

// ParseAndUpdateJSONDefinitions reads definitions from file and merges or
// overrides the values in memory.
func ParseAndUpdateJSONDefinitions(file string, merge bool) error {
	return tables.ParseAndUpdateJSONDefinitions(file, merge)
}
