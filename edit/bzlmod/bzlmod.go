/*
Copyright 2023 Google LLC

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

// Package bzlmod contains functions for working with MODULE.bazel files.
//
// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/edit/bzlmod].
package bzlmod

import (
	"github.com/bazel-contrib/buildtools/v10/edit/bzlmod"
	"github.com/bazelbuild/buildtools/build"
)

// Proxies returns the names of extension proxies (i.e. the names of variables to which the result
// of a use_extension call is assigned) for the given extension with the given value of the
// dev_dependency attribute.
// Extension proxies created with "isolate = True" are ignored.
func Proxies(f *build.File, rawExtBzlFile string, extName string, dev bool) []string {
	return bzlmod.Proxies(f, rawExtBzlFile, extName, dev)
}

// AllProxies returns the names of all extension proxies (i.e. the names of variables to which the
// result of a use_extension call is assigned) corresponding to the same extension usage as the
// given proxy.
// For an isolated extension usage, a list containing only the given proxy is returned.
// For a non-isolated extension usage, the proxies of all non-isolated extension usages of the same
// extension with the same value for the dev_dependency parameter are returned.
// If the given proxy is not an extension proxy, nil is returned.
func AllProxies(f *build.File, proxy string) []string {
	return bzlmod.AllProxies(f, proxy)
}

// UseRepos returns the use_repo calls that use the given proxies.
func UseRepos(f *build.File, proxies []string) []*build.CallExpr {
	return bzlmod.UseRepos(f, proxies)
}

// NewUseRepo inserts and returns a new use_repo call after the last usage of any of the given
// proxies, where a usage is either a use_extension call or a tag definition.
func NewUseRepo(f *build.File, proxies []string) (*build.File, *build.CallExpr) {
	return bzlmod.NewUseRepo(f, proxies)
}

// AddRepoUsages adds the given repos to the given use_repo calls without introducing duplicate
// arguments.
// useRepos must not be empty.
// Keyword arguments are preserved but adding them is currently not supported.
func AddRepoUsages(useRepos []*build.CallExpr, repos ...string) {
	bzlmod.AddRepoUsages(useRepos, repos...)
}

// RemoveRepoUsages removes the given repos from the given use_repo calls.
// Repositories are identified via their names as exported by the module extension (i.e. the value
// rather than the key in the case of keyword arguments).
func RemoveRepoUsages(useRepos []*build.CallExpr, repos ...string) {
	bzlmod.RemoveRepoUsages(useRepos, repos...)
}

// ExtractModuleToApparentNameMapping collects the mapping of module names (e.g. "rules_go") to
// user-configured apparent names (e.g. "my_rules_go") from the repo's MODULE.bazel, if it exists.
// The given function is called with a repo-relative, slash-separated path and should return the
// content of the MODULE.bazel or *.MODULE.bazel file at that path, or nil if the file does not
// exist.
// See https://bazel.build/external/module#repository_names_and_strict_deps for more information on
// apparent names.
func ExtractModuleToApparentNameMapping(fileReader func(relPath string) *build.File) func(string) string {
	return bzlmod.ExtractModuleToApparentNameMapping(fileReader)
}
