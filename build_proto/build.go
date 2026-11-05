// All exported identifiers of this package forward to [github.com/bazel-contrib/buildtools/v10/build_proto].
package build_proto

import (
	"github.com/bazel-contrib/buildtools/v10/build_proto"
)

type FilesetEntry_SymlinkBehavior = build_proto.FilesetEntry_SymlinkBehavior

const (
	FilesetEntry_COPY = build_proto.FilesetEntry_COPY

	FilesetEntry_DEREFERENCE = build_proto.FilesetEntry_DEREFERENCE
)

// Enum value maps for FilesetEntry_SymlinkBehavior.
var (
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/build_proto.FilesetEntry_SymlinkBehavior_name].
	// Assigning to it has no effect, assign to the v10 variable instead.
	FilesetEntry_SymlinkBehavior_name = build_proto.FilesetEntry_SymlinkBehavior_name

	// This is a copy of [github.com/bazel-contrib/buildtools/v10/build_proto.FilesetEntry_SymlinkBehavior_value].
	// Assigning to it has no effect, assign to the v10 variable instead.
	FilesetEntry_SymlinkBehavior_value = build_proto.FilesetEntry_SymlinkBehavior_value
)

type Attribute_Discriminator = build_proto.Attribute_Discriminator

const (
	Attribute_INTEGER = build_proto.Attribute_INTEGER

	Attribute_STRING = build_proto.Attribute_STRING

	Attribute_LABEL = build_proto.Attribute_LABEL

	Attribute_OUTPUT = build_proto.Attribute_OUTPUT

	Attribute_STRING_LIST = build_proto.Attribute_STRING_LIST

	Attribute_LABEL_LIST = build_proto.Attribute_LABEL_LIST

	Attribute_OUTPUT_LIST = build_proto.Attribute_OUTPUT_LIST

	Attribute_DISTRIBUTION_SET = build_proto.Attribute_DISTRIBUTION_SET

	Attribute_LICENSE = build_proto.Attribute_LICENSE

	Attribute_STRING_DICT = build_proto.Attribute_STRING_DICT

	Attribute_FILESET_ENTRY_LIST = build_proto.Attribute_FILESET_ENTRY_LIST

	Attribute_LABEL_LIST_DICT = build_proto.Attribute_LABEL_LIST_DICT

	Attribute_STRING_LIST_DICT = build_proto.Attribute_STRING_LIST_DICT

	Attribute_BOOLEAN = build_proto.Attribute_BOOLEAN

	Attribute_TRISTATE = build_proto.Attribute_TRISTATE

	Attribute_INTEGER_LIST = build_proto.Attribute_INTEGER_LIST

	Attribute_UNKNOWN = build_proto.Attribute_UNKNOWN

	Attribute_LABEL_DICT_UNARY = build_proto.Attribute_LABEL_DICT_UNARY

	Attribute_SELECTOR_LIST = build_proto.Attribute_SELECTOR_LIST

	Attribute_LABEL_KEYED_STRING_DICT = build_proto.Attribute_LABEL_KEYED_STRING_DICT

	Attribute_DEPRECATED_STRING_DICT_UNARY = build_proto.Attribute_DEPRECATED_STRING_DICT_UNARY
)

// Enum value maps for Attribute_Discriminator.
var (
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/build_proto.Attribute_Discriminator_name].
	// Assigning to it has no effect, assign to the v10 variable instead.
	Attribute_Discriminator_name = build_proto.Attribute_Discriminator_name

	// This is a copy of [github.com/bazel-contrib/buildtools/v10/build_proto.Attribute_Discriminator_value].
	// Assigning to it has no effect, assign to the v10 variable instead.
	Attribute_Discriminator_value = build_proto.Attribute_Discriminator_value
)

type Attribute_Tristate = build_proto.Attribute_Tristate

const (
	Attribute_NO = build_proto.Attribute_NO

	Attribute_YES = build_proto.Attribute_YES

	Attribute_AUTO = build_proto.Attribute_AUTO
)

// Enum value maps for Attribute_Tristate.
var (
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/build_proto.Attribute_Tristate_name].
	// Assigning to it has no effect, assign to the v10 variable instead.
	Attribute_Tristate_name = build_proto.Attribute_Tristate_name

	// This is a copy of [github.com/bazel-contrib/buildtools/v10/build_proto.Attribute_Tristate_value].
	// Assigning to it has no effect, assign to the v10 variable instead.
	Attribute_Tristate_value = build_proto.Attribute_Tristate_value
)

type Target_Discriminator = build_proto.Target_Discriminator

const (
	Target_RULE = build_proto.Target_RULE

	Target_SOURCE_FILE = build_proto.Target_SOURCE_FILE

	Target_GENERATED_FILE = build_proto.Target_GENERATED_FILE

	Target_PACKAGE_GROUP = build_proto.Target_PACKAGE_GROUP

	Target_ENVIRONMENT_GROUP = build_proto.Target_ENVIRONMENT_GROUP
)

// Enum value maps for Target_Discriminator.
var (
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/build_proto.Target_Discriminator_name].
	// Assigning to it has no effect, assign to the v10 variable instead.
	Target_Discriminator_name = build_proto.Target_Discriminator_name

	// This is a copy of [github.com/bazel-contrib/buildtools/v10/build_proto.Target_Discriminator_value].
	// Assigning to it has no effect, assign to the v10 variable instead.
	Target_Discriminator_value = build_proto.Target_Discriminator_value
)

type AllowedRuleClassInfo_AllowedRuleClasses = build_proto.AllowedRuleClassInfo_AllowedRuleClasses

const (
	AllowedRuleClassInfo_ANY = build_proto.AllowedRuleClassInfo_ANY

	AllowedRuleClassInfo_SPECIFIED = build_proto.AllowedRuleClassInfo_SPECIFIED
)

// Enum value maps for AllowedRuleClassInfo_AllowedRuleClasses.
var (
	// This is a copy of [github.com/bazel-contrib/buildtools/v10/build_proto.AllowedRuleClassInfo_AllowedRuleClasses_name].
	// Assigning to it has no effect, assign to the v10 variable instead.
	AllowedRuleClassInfo_AllowedRuleClasses_name = build_proto.AllowedRuleClassInfo_AllowedRuleClasses_name

	// This is a copy of [github.com/bazel-contrib/buildtools/v10/build_proto.AllowedRuleClassInfo_AllowedRuleClasses_value].
	// Assigning to it has no effect, assign to the v10 variable instead.
	AllowedRuleClassInfo_AllowedRuleClasses_value = build_proto.AllowedRuleClassInfo_AllowedRuleClasses_value
)

type License = build_proto.License

type StringDictEntry = build_proto.StringDictEntry

type LabelDictUnaryEntry = build_proto.LabelDictUnaryEntry

type LabelListDictEntry = build_proto.LabelListDictEntry

type LabelKeyedStringDictEntry = build_proto.LabelKeyedStringDictEntry

type StringListDictEntry = build_proto.StringListDictEntry

type FilesetEntry = build_proto.FilesetEntry

// Default values for FilesetEntry fields.
const Default_FilesetEntry_SymlinkBehavior = build_proto.Default_FilesetEntry_SymlinkBehavior

type Attribute = build_proto.Attribute

type Rule = build_proto.Rule

type ConfiguredRuleInput = build_proto.ConfiguredRuleInput

type RuleSummary = build_proto.RuleSummary

type PackageGroup = build_proto.PackageGroup

type EnvironmentGroup = build_proto.EnvironmentGroup

type SourceFile = build_proto.SourceFile

type GeneratedFile = build_proto.GeneratedFile

type Target = build_proto.Target

type QueryResult = build_proto.QueryResult

type AllowedRuleClassInfo = build_proto.AllowedRuleClassInfo

type AttributeDefinition = build_proto.AttributeDefinition

type AttributeValue = build_proto.AttributeValue

type RuleDefinition = build_proto.RuleDefinition

type BuildLanguage = build_proto.BuildLanguage

type Attribute_SelectorEntry = build_proto.Attribute_SelectorEntry

type Attribute_Selector = build_proto.Attribute_Selector

type Attribute_SelectorList = build_proto.Attribute_SelectorList

type AttributeValue_DictEntry = build_proto.AttributeValue_DictEntry

// This is a copy of [github.com/bazel-contrib/buildtools/v10/build_proto.File_build_proto_build_proto].
// Assigning to it has no effect, assign to the v10 variable instead.
var File_build_proto_build_proto = build_proto.File_build_proto_build_proto
