// Package buildifier provides a Buildifier which doesn't call os.exec.
//
// There are some programming frameworks which consider os.exec dangerous.
//
// The package is imported just for its side-effects. For example:
//
// import _ "github.com/bazelbuild/buildtools/edit/safe/buildifier"
//
// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/edit/safe].
package buildifier

import (
	_ "github.com/bazel-contrib/buildtools/v10/edit/safe"
)
