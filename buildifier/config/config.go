/*
Copyright 2022 Google LLC

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

// Package config provides configuration objects for buildifier
//
// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/buildifier/config].
package config

import (
	"github.com/bazel-contrib/buildtools/v10/buildifier/config"
)

// New constructs a Config with default values.
func New() *Config {
	return config.New()
}

// FindConfigPath locates the nearest buildifier configuration file.  First
// tries the value of the BUILDIFIER_CONFIG environment variable.  If no
// environment variable is defined, The configuration file will be resolved
// starting from the process cwd and searching up the file tree until a config
// file is (or isn't) found.
func FindConfigPath(rootDir string) string {
	return config.FindConfigPath(rootDir)
}

// Config is used to configure buildifier
type Config = config.Config

// ArrayFlags is a string slice that satisfies the flag.Value interface
type ArrayFlags = config.ArrayFlags

// Example creates a sample configuration file for the -config=example flag.
func Example() *Config {
	return config.Example()
}
