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

// Package wspace provides a method to find the root of the bazel tree.
//
// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/wspace].
package wspace

import (
	"github.com/bazel-contrib/buildtools/v10/wspace"
	"os"
)

// IsRegularFile returns true if the path refers to a regular file after
// following symlinks.
func IsRegularFile(path string) bool {
	return wspace.IsRegularFile(path)
}

// FindWorkspaceRoot splits the current code context (the rootDir if present,
// the working directory if not.) It returns the path of the directory
// containing the WORKSPACE file, and the rest.
func FindWorkspaceRoot(rootDir string) (root string, rest string) {
	return wspace.FindWorkspaceRoot(rootDir)
}

// Find searches from the given dir and up for the file that satisfies a condition of `rootFiles`
// returning the directory containing it, or an error if none found in the tree.
func Find(dir string, rootFiles map[string]func(os.FileInfo) bool) (string, error) {
	return wspace.Find(dir, rootFiles)
}

// FindRepoBuildFiles parses the WORKSPACE to find BUILD files for non-Bazel
// external repositories, specifically those defined by one of these rules:
//
//	git_repository(), new_local_repository(), new_http_archive()
func FindRepoBuildFiles(root string) (map[string]string, error) {
	return wspace.FindRepoBuildFiles(root)
}

// SplitFilePath splits a file path into the workspace root, package name and label.
// Workspace root is determined as the last directory in the file path that
// contains a WORKSPACE (or WORKSPACE.bazel) file.
// Package and label are always separated with forward slashes.
// Returns empty strings if no WORKSPACE file is found.
func SplitFilePath(filename string) (workspaceRoot, pkg, label string) {
	return wspace.SplitFilePath(filename)
}
