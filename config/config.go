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

// Package config provides environment specific configuration elements to unused_deps
//
// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/config].
package config

import (
	"github.com/bazel-contrib/buildtools/v10/config"
)

var (
	// DefaultBuildTool should be used to build and extract deps
	//
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/config.DefaultBuildTool].
	// Assigning to it has no effect, assign to the v10 variable instead.
	DefaultBuildTool = config.DefaultBuildTool

	// BuildToolHelp message for above
	//
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/config.BuildToolHelp].
	// Assigning to it has no effect, assign to the v10 variable instead.
	BuildToolHelp = config.BuildToolHelp

	// ExtraActionFileNameHelp help for extra actions
	//
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/config.ExtraActionFileNameHelp].
	// Assigning to it has no effect, assign to the v10 variable instead.
	ExtraActionFileNameHelp = config.ExtraActionFileNameHelp

	// DefaultBinDir is used to query for Bazel's bazel-bin directory.
	//
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/config.DefaultBinDir].
	// Assigning to it has no effect, assign to the v10 variable instead.
	DefaultBinDir = config.DefaultBinDir

	// DefaultOutputBase is used to query for Bazel's output base directory.
	//
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/config.DefaultOutputBase].
	// Assigning to it has no effect, assign to the v10 variable instead.
	DefaultOutputBase = config.DefaultOutputBase

	// DefaultOutputPath is used to query for Bazel's bazel-out directory.
	//
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/config.DefaultOutputPath].
	// Assigning to it has no effect, assign to the v10 variable instead.
	DefaultOutputPath = config.DefaultOutputPath

	// DefaultExtraBuildFlags is internal-only
	//
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/config.DefaultExtraBuildFlags].
	// Assigning to it has no effect, assign to the v10 variable instead.
	DefaultExtraBuildFlags = config.DefaultExtraBuildFlags
)
