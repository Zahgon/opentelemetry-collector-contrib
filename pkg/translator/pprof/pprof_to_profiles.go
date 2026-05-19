// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pprof // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/pprof"

import (
	"errors"

	"github.com/google/pprof/profile"
	"go.opentelemetry.io/collector/pdata/pprofile"
)

var (
	// errPprofInvalid is returned for invalid pprof data.
	errPprofInvalid = errors.New("invalid pprof data")
	// errIdxFormatInvalid is returned on invalid string formats for indices.
	errIdxFormatInvalid = errors.New("invalid format of attribute indices")
)

const (
	// noAttrUnit is an internal helper to indicate that no
	// unit is associated to this Attribute.
	noAttrUnit = int32(-1)
)

// attr is a helper struct to build pprofile.ProfilesDictionary.attribute_table.
type attr struct {
	keyStrIdx  int32
	valueHash  uint64
	unitStrIdx int32
}

// fn is a helper struct to build pprofile.ProfilesDictionary.function_table.
type fn struct {
	name       string
	systemName string
	fileName   string
	startLine  int64
}

// mm is a helper struct to build pprofile.ProfilesDictionary.mapping_table.
type mm struct {
	memoryStart    uint64
	memoryLimit    uint64
	fileOffset     uint64
	filenameStrIdx int32
	attrIdxs       string // List of consecutive increasing indices separated by a semicolon.
}

// loc is a helper struct to build pprofile.ProfilesDictionary.location_table.
type loc struct {
	mappingIdx int32
	address    uint64
	lines      string // String representation of a sorted list of lines.
	attrIdxs   string // List of consecutive increasing indices separated by a semicolon.
}

// stackKey is a string type used as a key in the stack table.
type stackKey string

// stack is a helper struct to build pprofile.ProfilesDictionary.stack_table.
type stack struct {
	id           int32
	locationIdxs []int32
}

// lookupTables is a helper struct around pprofile.ProfilesDictionary.
type lookupTables struct {
	mappingTable        map[mm]int32
	lastMappingTableIdx int32

	locationTable        map[loc]int32
	lastLocationTableIdx int32

	functionTable        map[fn]int32
	lastFunctionTableIdx int32

	// The concept of profiles.Link does not exit in pprof.
	// Therefore this helper table is skipped.

	stringTable        map[string]int32
	lastStringTableIdx int32

	attributeTable        map[attr]int32
	attributeHashToValue  map[uint64]any
	lastAttributeTableIdx int32

	stackTable        map[stackKey]stack
	lastStackTableIdx int32
}

// ConvertPprofToProfiles converts a pprof profile to OTLP profiles format.
func ConvertPprofToProfiles(src *profile.Profile) (*pprofile.Profiles, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initialize remaining lookup tables of pprofile.ProfilesDictionary in initLookupTables.

// Pre-populate all mappings from the original profile in order.
// This ensures all mappings are preserved and their IDs match the original order.

// Add envelope messages

// Use a dedicated pprofile.Profile for each sample type.
// By convention, pprof uses the last sample type as default, while OTel Profiles
// uses the first profile as default. Therefore, swap first and last.

// Swap first and last: last pprof sample type becomes first OTel profile

// pprof.Profile.sample_type

// pprof.Profile.sample

// pprof.Sample.location_id

// pprof.Sample.value

// pprof.Sample.label - this field is split into string and numeric labels.

// pprof.Profile.mapping
// As OTel pprofile manages its own ProfilesDictionary.mapping_table, there
// is no 1 to 1 mapping here.

// pprof.Profile.location
// As OTel pprofile manages its own ProfilesDictionary.location_table, there
// is no 1 to 1 mapping here.

// pprof.Profile.function
// As OTel pprofile manages its own ProfilesDictionary.function_table, there
// is no 1 to 1 mapping here.

// pprof.Profile.string_table
// As OTel pprofile manages its own ProfilesDictionary.string_table, there
// is no 1 to 1 mapping here.

// pprof.Profile.drop_frames

// pprof.Profile.keep_frames

// pprof.Profile.time_nanos

// pprof.Profile.duration_nanos

// pprof.Profile.period_type

// pprof.Profile.period

// pprof.Profile.comment

// Append a index to the attribute key, so that
// later src.Comments can be reconstructed correctly.

// pprof.Profile.default_sample_type
// As OTel pprofile uses a single Sample Type, it is implicit its default type.

// pprof.Profile.doc_url

// getIdxForFunction returns the corresponding index for the function.
// If the function does not yet exist in the cache, it will be added.
func (lts *lookupTables) getIdxForFunction(name, systemName, fileName string, startLine int64) int32 {
	_ = "STUB: not implemented"
	return 0
}

// getIdxForString returns the corresponding index for the string.
// If the string does not yet exist in the cache, it will be added.
func (lts *lookupTables) getIdxForString(s string) int32 { _ = "STUB: not implemented"; return 0 }

// getValueHash returns a unique hash for the given value.
func getValueHash(v any) uint64 { _ = "STUB: not implemented"; return 0 }

// Cast the slice to a byte slice for zero-allocation hashing if possible,
// but for safety/simplicity, we process the slice:

// Fallback for unexpected types (nil, etc.)

// getIdxForAttribute returns the corresponding index for the attribute.
// If the attribute does not yet exist in the cache, it will be added.
func (lts *lookupTables) getIdxForAttribute(key string, value any) int32 {
	_ = "STUB: not implemented"
	return 0
}

// getIdxForAttributeWithUnit returns the corresponding index for the attribute.
// If the attribute does not yet exist in the cache, it will be added.
func (lts *lookupTables) getIdxForAttributeWithUnit(key, unit string, value any) int32 {
	_ = "STUB: not implemented"
	return 0
}

// getIdxForMapping returns the correspoinding index for the mapping.
// If the mapping does not yet exist in the cache, it will be added.
func (lts *lookupTables) getIdxForMapping(start, limit, offset uint64, fnStrIdx int32, attrIdx []int32) int32 {
	_ = "STUB: not implemented"
	return 0
}

// getIdxForMMAttributes returns a list of indices to attributes related
// to the mapping.
func (lts *lookupTables) getIdxForMMAttributes(m *profile.Mapping) []int32 {
	_ = "STUB: not implemented"

	// pprof.Mapping.build_id
	// Assume all build_ids are GNU build IDs
	return nil
}

// pprof.Mapping.has_*

// getIdxForStack returns the corresponding index for a stack and its location
// indices. If the mapping does not yet exist in the cache, it will be added.
func (lts *lookupTables) getIdxForStack(locs []*profile.Location) int32 {
	_ = "STUB: not implemented"
	return 0
}

// Use makeStackKey instead of attrIdxToString to preserve stack order.

// getIdxForLocation returns the corresponding index for a location.
// If the location does not yet exist in the cache, it will be added.
func (lts *lookupTables) getIdxForLocation(l *profile.Location) int32 {
	_ = "STUB: not implemented"
	return 0

	// pprof.Location.is_folded
}

// pprof.Location.address

// pprof.Location.line

// pprof.Location.mapping_id

// initLookupTables returns a supporting elements to construct pprofile.ProfilesDictionary.
func initLookupTables() lookupTables { _ = "STUB: not implemented"; return *new(lookupTables) }

// mapping_table[0] must always be zero value (Mapping{}) and present.

// location_table[0] must always be zero value (Location{}) and present.

// function_table[0] must always be zero value (Function{}) and present.

// string_table[0] must always be "" and present.

// attribute_table[0] must always be zero value (KeyValueAndUnit{}) and present.

// stack_table[0] must always be zero value (Stack{}) and present.

// dumpLookupTables fills pprofile.ProfilesDictionary with the content of
// the supporting lookup tables.
func (lts *lookupTables) dumpLookupTables(dic pprofile.ProfilesDictionary) error {
	_ = "STUB: not implemented"
	return nil
}

// The concept of profiles.Link does not exist in pprof.
// Therefore LinkTable only holds an empty value to be compliant.

// attrIdxToString is a helper function to convert a list of indices
// into a string. This function modifies the input slice.
func attrIdxToString(indices []int32) string { _ = "STUB: not implemented"; return "" }

// stringToAttrIdx is a helper function to convert a string into
// a list of indices.
func stringToAttrIdx(indices string) ([]int32, error) { _ = "STUB: not implemented"; return nil, nil }

// makeStackKey is a helper function to convert a list of location indices
// into a stackKey, preserving their order (unlike attrIdxToString which sorts).
func makeStackKey(indices []int32) stackKey { _ = "STUB: not implemented"; return *new(stackKey) }

// linesToString is a helper function to convert a list of lines into a string.
func (lts *lookupTables) linesToString(lines []profile.Line) string {
	_ = "STUB: not implemented"
	return ""
}

// stringToLine is a helper function to convert a string into a list of lines.
func stringToLine(lines string) ([]pprofile.Line, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
