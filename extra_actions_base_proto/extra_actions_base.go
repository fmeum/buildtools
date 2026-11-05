// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/extra_actions_base_proto].
package extra_actions_base_proto

import (
	"github.com/bazel-contrib/buildtools/v10/extra_actions_base_proto"
)

type ExtraActionSummary = extra_actions_base_proto.ExtraActionSummary

type DetailedExtraActionInfo = extra_actions_base_proto.DetailedExtraActionInfo

type ExtraActionInfo = extra_actions_base_proto.ExtraActionInfo

type EnvironmentVariable = extra_actions_base_proto.EnvironmentVariable

type SpawnInfo = extra_actions_base_proto.SpawnInfo

type CppCompileInfo = extra_actions_base_proto.CppCompileInfo

type CppLinkInfo = extra_actions_base_proto.CppLinkInfo

type JavaCompileInfo = extra_actions_base_proto.JavaCompileInfo

type PythonInfo = extra_actions_base_proto.PythonInfo

// Deprecated: Marked as deprecated in extra_actions_base_proto/extra_actions_base.proto.
type ExtraActionInfo_StringList = extra_actions_base_proto.ExtraActionInfo_StringList

type ExtraActionInfo_AspectDescriptor = extra_actions_base_proto.ExtraActionInfo_AspectDescriptor

type ExtraActionInfo_AspectDescriptor_StringList = extra_actions_base_proto.ExtraActionInfo_AspectDescriptor_StringList

// Extension fields to ExtraActionInfo.
var (
	// optional blaze.SpawnInfo spawn_info = 1003;
	//
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/extra_actions_base_proto.E_SpawnInfo_SpawnInfo].
	// Assigning to it has no effect, assign to the v10 variable instead.
	E_SpawnInfo_SpawnInfo = extra_actions_base_proto.E_SpawnInfo_SpawnInfo

	// optional blaze.CppCompileInfo cpp_compile_info = 1001;
	//
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/extra_actions_base_proto.E_CppCompileInfo_CppCompileInfo].
	// Assigning to it has no effect, assign to the v10 variable instead.
	E_CppCompileInfo_CppCompileInfo = extra_actions_base_proto.E_CppCompileInfo_CppCompileInfo

	// optional blaze.CppLinkInfo cpp_link_info = 1002;
	//
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/extra_actions_base_proto.E_CppLinkInfo_CppLinkInfo].
	// Assigning to it has no effect, assign to the v10 variable instead.
	E_CppLinkInfo_CppLinkInfo = extra_actions_base_proto.E_CppLinkInfo_CppLinkInfo

	// optional blaze.JavaCompileInfo java_compile_info = 1000;
	//
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/extra_actions_base_proto.E_JavaCompileInfo_JavaCompileInfo].
	// Assigning to it has no effect, assign to the v10 variable instead.
	E_JavaCompileInfo_JavaCompileInfo = extra_actions_base_proto.E_JavaCompileInfo_JavaCompileInfo

	// optional blaze.PythonInfo python_info = 1005;
	//
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/extra_actions_base_proto.E_PythonInfo_PythonInfo].
	// Assigning to it has no effect, assign to the v10 variable instead.
	E_PythonInfo_PythonInfo = extra_actions_base_proto.E_PythonInfo_PythonInfo
)

// This is a copy of [github.com/bazel-contrib/buildtools/v10/extra_actions_base_proto.File_extra_actions_base_proto_extra_actions_base_proto].
// Assigning to it has no effect, assign to the v10 variable instead.
var File_extra_actions_base_proto_extra_actions_base_proto = extra_actions_base_proto.File_extra_actions_base_proto_extra_actions_base_proto
