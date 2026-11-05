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

// Options represents choices about how buildozer should behave.
type Options = edit.Options

// NewOpts returns a new Options struct with some defaults set.
func NewOpts() *Options {
	return edit.NewOpts()
}

// Usage is a user-overridden func to print the program usage.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/edit.Usage].
// Assigning to it has no effect, assign to the v10 variable instead.
var Usage = edit.Usage

// CmdEnvironment stores the information the commands below have access to.
type CmdEnvironment = edit.CmdEnvironment

// AttrType is an enum representing a type of the attribute, e.g. expr or not
// provided type
type AttrType = edit.AttrType

// CommandInfo provides a command function and info on incoming arguments.
type CommandInfo = edit.CommandInfo

// AllCommands associates the command names with their function and number
// of arguments.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/edit.AllCommands].
// Assigning to it has no effect, assign to the v10 variable instead.
var AllCommands = edit.AllCommands

// SplitOnSpaces behaves like strings.Fields, except that spaces can be escaped.
// Also splits on linebreaks unless they are escaped too.
// " some dummy\\ string" -> ["some", "dummy string"]
func SplitOnSpaces(input string) []string {
	return edit.SplitOnSpaces(input)
}

// BuildFileNames is exported so that users that want to override it
// in scripts are free to do so.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/edit.BuildFileNames].
// Assigning to it has no effect, assign to the v10 variable instead.
var BuildFileNames = edit.BuildFileNames

// Buildifier formats the build file using the buildifier logic.
type Buildifier = edit.Buildifier

// RegisterBuildifier replaces the default buildifier with an
// alternative implementation.
//
// It may only be called once.
func RegisterBuildifier(b Buildifier) {
	edit.RegisterBuildifier(b)
}

// EditFile is a function that does any prework needed before editing a file.
// e.g. "checking out for write" from a locking source control repo.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/edit.EditFile].
// Assigning to it has no effect, assign to the v10 variable instead.
var EditFile = edit.EditFile

// Buildozer loops over all arguments on the command line fixing BUILD files.
func Buildozer(opts *Options, args []string) int {
	return edit.Buildozer(opts, args)
}

// ExecuteCommandsOnInlineFile executes the given commands on the given file content.
// Returns the new file content after applying the commands.
func ExecuteCommandsOnInlineFile(fileContent []byte, commands []string) ([]byte, error) {
	return edit.ExecuteCommandsOnInlineFile(fileContent, commands)
}
