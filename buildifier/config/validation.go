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

package config

import (
	"github.com/bazel-contrib/buildtools/v10/buildifier/config"
)

// ValidateInputType validates the value of --type
func ValidateInputType(inputType *string) error {
	return config.ValidateInputType(inputType)
}

// ValidateFormat validates the value of --format
func ValidateFormat(format, mode *string) error {
	return config.ValidateFormat(format, mode)
}

// ValidateModes validates flags --mode, --lint, and -d
func ValidateModes(mode, lint *string, dflag *bool, additionalModes ...string) error {
	return config.ValidateModes(mode, lint, dflag, additionalModes...)
}

// ValidateWarnings validates the value of the --warnings flag
func ValidateWarnings(warnings *string, allWarnings, defaultWarnings *[]string) ([]string, error) {
	return config.ValidateWarnings(warnings, allWarnings, defaultWarnings)
}
