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

package warn

import (
	"github.com/bazel-contrib/buildtools/v10/warn"
	"github.com/bazelbuild/buildtools/build"
)

// NotLoadedFunctionUsageCheck checks whether there's a usage of a given not imported function in the file
// and adds a load statement if necessary.
func NotLoadedFunctionUsageCheck(f *build.File, fileReader *FileReader, globals []string, loadFrom string) []*LinterFinding {
	return warn.NotLoadedFunctionUsageCheck(f, fileReader, globals, loadFrom)
}

// NotLoadedSymbolUsageCheck checks whether there's a usage of a given not imported function in the file
// and adds a load statement if necessary.
func NotLoadedSymbolUsageCheck(f *build.File, fileReader *FileReader, globals []string, loadFrom string) []*LinterFinding {
	return warn.NotLoadedSymbolUsageCheck(f, fileReader, globals, loadFrom)
}

// NativeCcRulesWarning produces a warning for missing loads of cc rules
func NativeCcRulesWarning(rule string) func(f *build.File, fileReader *FileReader) []*LinterFinding {
	return warn.NativeCcRulesWarning(rule)
}

// NativeCcToolchainRulesWarning produces a warning for missing loads of cc toolchain rules
func NativeCcToolchainRulesWarning(rule string) func(f *build.File, fileReader *FileReader) []*LinterFinding {
	return warn.NativeCcToolchainRulesWarning(rule)
}

// NativeCcSymbolsWarning produces a warning for missing loads of cc top-level symbols
func NativeCcSymbolsWarning(symbol string, bzlfile string) func(f *build.File, fileReader *FileReader) []*LinterFinding {
	return warn.NativeCcSymbolsWarning(symbol, bzlfile)
}

// NativeJavaRulesWarning produces a warning for missing loads of java rules
func NativeJavaRulesWarning(rule string) func(f *build.File, fileReader *FileReader) []*LinterFinding {
	return warn.NativeJavaRulesWarning(rule)
}

// NativeJavaToolchainRulesWarning produces a warning for missing loads of java toolchain rules
func NativeJavaToolchainRulesWarning(rule string) func(f *build.File, fileReader *FileReader) []*LinterFinding {
	return warn.NativeJavaToolchainRulesWarning(rule)
}

// NativeJavaSymbolsWarning produces a warning for missing loads of java top-level symbols
func NativeJavaSymbolsWarning(symbol string, bzlfile string) func(f *build.File, fileReader *FileReader) []*LinterFinding {
	return warn.NativeJavaSymbolsWarning(symbol, bzlfile)
}

// NativeProtoRulesWarning produces a warning for missing loads of proto rules
func NativeProtoRulesWarning(rule string) func(f *build.File, fileReader *FileReader) []*LinterFinding {
	return warn.NativeProtoRulesWarning(rule)
}

// NativeShellRulesWarning produces a warning for missing loads of shell rules
func NativeShellRulesWarning(rule string) func(f *build.File, fileReader *FileReader) []*LinterFinding {
	return warn.NativeShellRulesWarning(rule)
}
