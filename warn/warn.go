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

// Package warn implements functions that generate warnings for BUILD files.
//
// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/warn].
package warn

import (
	"github.com/bazel-contrib/buildtools/v10/warn"
	"github.com/bazelbuild/buildtools/build"
)

// LintMode is an enum representing a linter mode. Can be either "warn", "fix", or "suggest"
type LintMode = warn.LintMode

const (
	// ModeWarn means only warnings should be returned for each finding.
	ModeWarn = warn.ModeWarn

	// ModeFix means that all warnings that can be fixed automatically should be fixed and
	// no warnings should be returned for them.
	ModeFix = warn.ModeFix

	// ModeSuggest means that automatic fixes shouldn't be applied, but instead corresponding
	// suggestions should be attached to all warnings that can be fixed automatically.
	ModeSuggest = warn.ModeSuggest
)

// LinterFinding is a low-level warning reported by single linter/fixer functions.
type LinterFinding = warn.LinterFinding

// LinterReplacement is a low-level object returned by single fixer functions.
type LinterReplacement = warn.LinterReplacement

// A Finding is a warning reported by the analyzer. It may contain an optional suggested fix.
type Finding = warn.Finding

// A Replacement is a suggested fix. Text between Start and End should be replaced with Content.
type Replacement = warn.Replacement

// RuleWarningMap lists the warnings that run on a single rule.
// These warnings run only on BUILD files (not bzl files).
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/warn.RuleWarningMap].
// Assigning to it has no effect, assign to the v10 variable instead.
var RuleWarningMap = warn.RuleWarningMap

// FileWarningMap lists the warnings that run on the whole file.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/warn.FileWarningMap].
// Assigning to it has no effect, assign to the v10 variable instead.
var FileWarningMap = warn.FileWarningMap

// MultiFileWarningMap lists the warnings that run on the whole file, but may use other files.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/warn.MultiFileWarningMap].
// Assigning to it has no effect, assign to the v10 variable instead.
var MultiFileWarningMap = warn.MultiFileWarningMap

// HasDisablingComment checks if a node has a comment that disables a certain warning
func HasDisablingComment(expr build.Expr, warning string) bool {
	return warn.HasDisablingComment(expr, warning)
}

// DisabledWarning checks if the warning was disabled by a comment.
// The comment format is buildozer: disable=<warning>
func DisabledWarning(f *build.File, findingLine int, warning string) bool {
	return warn.DisabledWarning(f, findingLine, warning)
}

// FileWarnings returns a list of all warnings found in the file.
func FileWarnings(f *build.File, enabledWarnings []string, formatted *[]byte, mode LintMode, fileReader *FileReader) []*Finding {
	return warn.FileWarnings(f, enabledWarnings, formatted, mode, fileReader)
}

// FixWarnings fixes all warnings that can be fixed automatically.
func FixWarnings(f *build.File, enabledWarnings []string, verbose bool, fileReader *FileReader) {
	warn.FixWarnings(f, enabledWarnings, verbose, fileReader)
}

// AllWarnings is the list of all available warnings.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/warn.AllWarnings].
// Assigning to it has no effect, assign to the v10 variable instead.
var AllWarnings = warn.AllWarnings

// DefaultWarnings is the list of all warnings that should be used inside google3
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/warn.DefaultWarnings].
// Assigning to it has no effect, assign to the v10 variable instead.
var DefaultWarnings = warn.DefaultWarnings
