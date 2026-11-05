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

package tables

import (
	"github.com/bazel-contrib/buildtools/v10/tables"
)

// IsLabelArg contains a list of named arguments to a rule call that are
// considered to have a value that can be treated as a label or list of labels.
// There is a separate denylist for rule-specific exceptions.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.IsLabelArg].
// Assigning to it has no effect, assign to the v10 variable instead.
var IsLabelArg = tables.IsLabelArg

// LabelDenylist is the list of call arguments that cannot be
// shortened, because they are not interpreted using the same
// rules as for other labels.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.LabelDenylist].
// Assigning to it has no effect, assign to the v10 variable instead.
var LabelDenylist = tables.LabelDenylist

// By default, edit.types.IsList consults lang.TypeOf to determine if an arg is a list.
// You may override this using IsListArg. Specifying a name here overrides any value
// in lang.TypeOf.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.IsListArg].
// Assigning to it has no effect, assign to the v10 variable instead.
var IsListArg = tables.IsListArg

// IsSortableListArg contains a list of named arguments to a rule call that are
// considered to be a sortable list . There is a separate denylist for
// rule-specific exceptions.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.IsSortableListArg].
// Assigning to it has no effect, assign to the v10 variable instead.
var IsSortableListArg = tables.IsSortableListArg

// SortableDenylist records specific rule arguments that must not be reordered.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.SortableDenylist].
// Assigning to it has no effect, assign to the v10 variable instead.
var SortableDenylist = tables.SortableDenylist

// SortableAllowlist records specific rule arguments that are guaranteed to be reorderable
// (format: "rule_name.attribute_name").
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.SortableAllowlist].
// Assigning to it has no effect, assign to the v10 variable instead.
var SortableAllowlist = tables.SortableAllowlist

// NamePriority maps an argument name to its sorting priority.
//
// NOTE(bazel-team): These are the old buildifier rules. It is likely that this table
// will change, perhaps swapping in a separate table for each call,
// derived from the order used in the Build Encyclopedia.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.NamePriority].
// Assigning to it has no effect, assign to the v10 variable instead.
var NamePriority = tables.NamePriority

// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.StripLabelLeadingSlashes].
// Assigning to it has no effect, assign to the v10 variable instead.
var StripLabelLeadingSlashes = tables.StripLabelLeadingSlashes

// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.ShortenAbsoluteLabelsToRelative].
// Assigning to it has no effect, assign to the v10 variable instead.
var ShortenAbsoluteLabelsToRelative = tables.ShortenAbsoluteLabelsToRelative

// AndroidNativeRules lists all Android rules that are being migrated from Native to Starlark.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.AndroidNativeRules].
// Assigning to it has no effect, assign to the v10 variable instead.
var AndroidNativeRules = tables.AndroidNativeRules

// AndroidLoadPath is the load path for the Starlark Android Rules.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.AndroidLoadPath].
// Assigning to it has no effect, assign to the v10 variable instead.
var AndroidLoadPath = tables.AndroidLoadPath

// CcLoadPathPrefix is the load path for the Starlark C++ Rules.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.CcLoadPathPrefix].
// Assigning to it has no effect, assign to the v10 variable instead.
var CcLoadPathPrefix = tables.CcLoadPathPrefix

// JavaLoadPathPrefix is the load package for the Starlark Java Rules.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.JavaLoadPathPrefix].
// Assigning to it has no effect, assign to the v10 variable instead.
var JavaLoadPathPrefix = tables.JavaLoadPathPrefix

// ShellLoadPathPrefix is the load package for the Starlark Shell Rules.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.ShellLoadPathPrefix].
// Assigning to it has no effect, assign to the v10 variable instead.
var ShellLoadPathPrefix = tables.ShellLoadPathPrefix

// PyNativeRules lists all Python rules that are being migrated from Native to Starlark.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.PyNativeRules].
// Assigning to it has no effect, assign to the v10 variable instead.
var PyNativeRules = tables.PyNativeRules

// PyLoadPath is the load path for the Starlark Python Rules.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.PyLoadPath].
// Assigning to it has no effect, assign to the v10 variable instead.
var PyLoadPath = tables.PyLoadPath

// ProtoLoadPathPrefix is the load path prefix for the Starlark Proto Rules.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.ProtoLoadPathPrefix].
// Assigning to it has no effect, assign to the v10 variable instead.
var ProtoLoadPathPrefix = tables.ProtoLoadPathPrefix

// ModuleToLegacyRepoName contains the mapping from module name to WORKSPACE repository name
// for modules with load fixes if those names are different.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.ModuleToLegacyRepoName].
// Assigning to it has no effect, assign to the v10 variable instead.
var ModuleToLegacyRepoName = tables.ModuleToLegacyRepoName

// IsModuleOverride contains the names of all Bzlmod module overrides available in MODULE.bazel.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.IsModuleOverride].
// Assigning to it has no effect, assign to the v10 variable instead.
var IsModuleOverride = tables.IsModuleOverride

// AllowedSymbolLoadLocations contains locations for loading rules that are allowed to be used.
//
// This is a copy of [github.com/bazel-contrib/buildtools/v10/tables.AllowedSymbolLoadLocations].
// Assigning to it has no effect, assign to the v10 variable instead.
var AllowedSymbolLoadLocations = tables.AllowedSymbolLoadLocations

// OverrideTables allows a user of the build package to override the special-case rules. The user-provided tables replace the built-in tables.
func OverrideTables(labelArg, denylist, listArg, sortableListArg, sortDenylist, sortAllowlist map[string]bool, namePriority map[string]int, stripLabelLeadingSlashes, shortenAbsoluteLabelsToRelative bool, symbolLoadLocation map[string][]string) {
	tables.OverrideTables(labelArg, denylist, listArg, sortableListArg, sortDenylist, sortAllowlist, namePriority, stripLabelLeadingSlashes, shortenAbsoluteLabelsToRelative, symbolLoadLocation)
}

// MergeTables allows a user of the build package to override the special-case rules. The user-provided tables are merged into the built-in tables.
func MergeTables(labelArg, denylist, listArg, sortableListArg, sortDenylist, sortAllowlist map[string]bool, namePriority map[string]int, stripLabelLeadingSlashes, shortenAbsoluteLabelsToRelative bool, symbolLoadLocation map[string][]string) {
	tables.MergeTables(labelArg, denylist, listArg, sortableListArg, sortDenylist, sortAllowlist, namePriority, stripLabelLeadingSlashes, shortenAbsoluteLabelsToRelative, symbolLoadLocation)
}
