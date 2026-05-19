// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pprof // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/pprof"

import (
	"cmp"
	"errors"

	"github.com/google/pprof/profile"
	"go.opentelemetry.io/collector/pdata/pprofile"
)

// errNotFound is returned if something requested is not available
var errNotFound = errors.New("not found")

func ConvertPprofileToPprof(src *pprofile.Profiles) (*profile.Profile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Basic check that all profiles hold the same number of samples.

// Helper maps to avoid duplicates.

// Pre-populate all mappings from the mapping table to preserve all of them
// (not just those referenced by locations) and maintain their original order.

// Skip the zero-value mapping at index 0.

// As all profiles hold the same number of samples,
// assume they only differ on the sample type and therefore can
// be merged into a pprof profile.

// Convert profiles samples into pprof samples.

// By convention, pprof uses the last sample type as default, while OTel Profiles
// uses the first profile as default. Therefore, swap first and last.

// Swap first and last: first OTel profile becomes last pprof sample type

// pprof.Sample.label is skipped for the moment.

// Set pprof values that should be common across all profiles.
// By convention, pprof uses the last sample type as default, while OTel Profiles
// uses the first profile as default. Therefore, swap first and last.

// Swap first and last: first OTel profile becomes last pprof sample type

// allElementsSame checks if all elements in the provided slice are equal.
// It returns true if the slice contains at least one element and all elements are equal.
// It returns false otherwise.
func allElementsSame[T cmp.Ordered](in []T) bool { _ = "STUB: not implemented"; return false }

// getAttributeString walks the attribute_table and returns the string value for
// the given key.
// It returns errNotFound if the key can not be found in attribute_table.
func getAttributeString(dic pprofile.ProfilesDictionary, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// getAttributeBool walks the attribute_table for the given indices and returns the bool value
// for the given key.
// It returns errNotFound if the key can not be found in attribute_table.
func getAttributeBool(dic pprofile.ProfilesDictionary, attrIndices []int32, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// getAttributeStringWithPrefix walks the attribute_table and returns a string array
// for all keys that start with the prefix pprof.profile.comment.
// It returns errNotFound if the key can not be found in attribute_table.
func getAttributeStringWithPrefix(dic pprofile.ProfilesDictionary) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getStringFromIdx returns the string on idx.
func getStringFromIdx(dic pprofile.ProfilesDictionary, idx int) string {
	_ = "STUB: not implemented"
	return ""
}

// uint64Tobytes helps to convert uint64 to []byte.
func uint64Tobytes(v uint64) []byte { _ = "STUB: not implemented"; return nil }

// getFunctionHash returns a non-cryptographic hash as identifier for a Function.
func getFunctionHash(name, systemName, fileName string, startLine int64) uint64 {
	_ = "STUB: not implemented"
	return 0
}

//nolint:errcheck
//nolint:errcheck
//nolint:errcheck
//nolint:errcheck

// populateFunction helps to populate deduplicated functions.
func populateFunction(dst *profile.Profile, functionMap map[uint64]*profile.Function,
	name, systemName, fileName string, startLine int64,
) *profile.Function {
	_ = "STUB: not implemented"
	return nil
}

// getMappingHash returns a non-cryptographic hash as identifier for a Mapping.
func getMappingHash(start, limit, offset uint64, file string) uint64 {
	_ = "STUB: not implemented"
	return 0
}

//nolint:errcheck
//nolint:errcheck
//nolint:errcheck
//nolint:errcheck

// populateMapping helps to populate deduplicated mappings.
func populateMapping(dst *profile.Profile, mappingMap map[uint64]*profile.Mapping,
	start, limit, offset uint64, fileName string, dic pprofile.ProfilesDictionary, m pprofile.Mapping,
) *profile.Mapping {
	_ = "STUB: not implemented"
	return nil
}

// Extract mapping attributes from the OTel profile

// Extract BuildID from attributes

// Extract boolean mapping flags

// getLocationHash returns a non-cryptographic hash as identifier for a Location.
func getLocationHash(m *profile.Mapping, addr uint64, lines []profile.Line) uint64 {
	_ = "STUB: not implemented"
	return 0
}

//nolint:errcheck

//nolint:errcheck

//nolint:errcheck
//nolint:errcheck

//nolint:errcheck

// populateLocation helps to populate deduplicated locations.
func populateLocation(dst *profile.Profile, locationMap map[uint64]*profile.Location,
	m *profile.Mapping, addr uint64, lines []profile.Line, dic pprofile.ProfilesDictionary, loc pprofile.Location,
) *profile.Location {
	_ = "STUB: not implemented"
	return nil
}

// Extract location attributes from the OTel profile

// Extract IsFolded from attributes
