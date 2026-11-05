// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/deps_proto].
package deps_proto

import (
	"github.com/bazel-contrib/buildtools/v10/deps_proto"
)

type Dependency_Kind = deps_proto.Dependency_Kind

const (
	Dependency_EXPLICIT = deps_proto.Dependency_EXPLICIT

	Dependency_IMPLICIT = deps_proto.Dependency_IMPLICIT

	Dependency_UNUSED = deps_proto.Dependency_UNUSED

	Dependency_INCOMPLETE = deps_proto.Dependency_INCOMPLETE
)

// Enum value maps for Dependency_Kind.
var (
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/deps_proto.Dependency_Kind_name].
	// Assigning to it has no effect, assign to the v10 variable instead.
	Dependency_Kind_name = deps_proto.Dependency_Kind_name

	// This is a copy of [github.com/bazel-contrib/buildtools/v10/deps_proto.Dependency_Kind_value].
	// Assigning to it has no effect, assign to the v10 variable instead.
	Dependency_Kind_value = deps_proto.Dependency_Kind_value
)

type SourceLocation = deps_proto.SourceLocation

type Dependency = deps_proto.Dependency

type Dependencies = deps_proto.Dependencies

// This is a copy of [github.com/bazel-contrib/buildtools/v10/deps_proto.File_deps_proto_deps_proto].
// Assigning to it has no effect, assign to the v10 variable instead.
var File_deps_proto_deps_proto = deps_proto.File_deps_proto_deps_proto
