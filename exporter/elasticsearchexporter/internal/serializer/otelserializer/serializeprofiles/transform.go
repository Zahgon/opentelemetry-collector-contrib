// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package serializeprofiles // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/serializer/otelserializer/serializeprofiles"

import (
	"errors"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/ebpf-profiler/libpf"
)

// Transform transforms a [pprofile.Profile] into our own
// representation, for ingestion into Elasticsearch
func Transform(dic pprofile.ProfilesDictionary, resource pcommon.Resource, scope pcommon.InstrumentationScope, profile pprofile.Profile) ([]StackPayload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// profileContainer is checked for nil inside stackPayloads().

// checkProfileType acts as safeguard to make sure only known profiles are
// accepted. Different kinds of profiles are currently not supported
// and mixing profiles will make profiling information unusable.
func checkProfileType(dic pprofile.ProfilesDictionary, profile pprofile.Profile) error {
	_ = "STUB: not implemented"
	return nil
}

// Make sure only on-CPU profiling data is accepted at the moment.
// This needs to match with
//nolint:lll
// https://github.com/open-telemetry/opentelemetry-ebpf-profiler/blob/a720d06a401cb23249c5066dc69e96384af99cf3/reporter/otlp_reporter.go#L531

// Make sure only on-CPU profiling data is accepted at the moment.
// This needs to match with
//nolint:lll
// https://github.com/open-telemetry/opentelemetry-ebpf-profiler/blob/a720d06a401cb23249c5066dc69e96384af99cf3/reporter/otlp_reporter.go#L536

// stackPayloads creates a slice of StackPayloads from the given ResourceProfiles,
// ScopeProfiles, and ProfileContainer.
func stackPayloads(dic pprofile.ProfilesDictionary, resource pcommon.Resource, scope pcommon.InstrumentationScope, profile pprofile.Profile) ([]StackPayload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The lowest sensical frequency is 1Hz.

// Set the stacktrace and stackframes to the payload.
// The docs only need to be written once.

// Artificial error frames can't be symbolized.

// Skip interpreted frames and already symbolized native frames (kernel, Golang is planned).

// Add one event per timestamp and its count value.

func unsymbolizedExecutables(executables map[libpf.FileID]struct{}) []UnsymbolizedExecutable {
	_ = "STUB: not implemented"
	return nil
}

func unsymbolizedLeafFrames(frameIDs map[frameID]struct{}) []UnsymbolizedLeafFrame {
	_ = "STUB: not implemented"
	return nil
}

// symbolizedFrames returns a slice of StackFrames that have symbols.
func symbolizedFrames(frames []StackFrame) []StackFrame { _ = "STUB: not implemented"; return nil }

func isFrameSymbolized(frame StackFrame) bool { _ = "STUB: not implemented"; return false }

func stackTraceEvent(dic pprofile.ProfilesDictionary, traceID string, sample pprofile.Sample, frequency int64,
	commonResourceAttrs map[string]string,
) StackTraceEvent {
	_ = "STUB: not implemented"
	return *new(StackTraceEvent)
}

// Elasticsearch v9.2+ doesn't read the count value any more.

// Use a project ID other than 1 to not conflict with ECH default value.

// Store event-specific attributes.

func stackTrace(stackTraceID string, frames []StackFrame, frameTypes []libpf.FrameType) StackTrace {
	_ = "STUB: not implemented"
	return *new(StackTrace)
}

// Up to 255 consecutive identical frame types are converted into 2 bytes (binary).
// We expect mostly consecutive frame types in a trace. Even if the encoding
// takes more than 32 bytes in single cases, the probability that the average base64 length
// per trace is below 32 bytes is very high.
// We expect resizing of buf to happen very rarely.

func stackFrames(dic pprofile.ProfilesDictionary, sample pprofile.Sample) ([]StackFrame, []libpf.FrameType, *frameID, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func getFrameID(dic pprofile.ProfilesDictionary, location pprofile.Location) *frameID {
	_ = "STUB: not implemented"
	// The MappingIndex is known to be valid.
	return nil
}

// Synthesize a file ID if the htlhash build ID is not available.

type attributable interface {
	AttributeIndices() pcommon.Int32Slice
}

// errMissingAttribute allows to differentiate errors handling the AttributeTable
// and indicates that a attribute was not included in the AttributeTable.
var errMissingAttribute = errors.New("missing attribute")

// getStringFromAttribute returns a string from one of attrIndices from the attribute table
// of the profile if the attribute key matches the expected attrKey.
func getStringFromAttribute(dic pprofile.ProfilesDictionary, record attributable, attrKey string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// getBuildID returns the Build ID for the given mapping. It checks for both
// old-style Build ID (stored with the mapping) and Build ID as attribute.
// If the build ID attribute is missing, returns a zero FileID and no error.
func getBuildID(dic pprofile.ProfilesDictionary, mapping pprofile.Mapping) (libpf.FileID, error) {
	_ = "STUB: not implemented"
	// Fetch build ID from profiles.attribute_table.
	return *new(libpf.FileID), nil
}

func executables(dic pprofile.ProfilesDictionary, mappings pprofile.MappingSlice) ([]ExeMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is true for interpreted languages like Python.

// No build ID was specified or could be fetched.

// stackTraceID creates a unique trace ID from the stack frames.
// For the OTEL profiling protocol, we have all required information in one wire message.
// But for the Elastic gRPC protocol, trace events and stack traces are sent separately, so
// that the host agent still needs to generate the stack trace IDs.
//
// The following code generates the same trace ID as the host agent.
// For ES 9.0.0, we could use a faster hash algorithm, e.g. xxh3, and hash strings instead
// of hashing binary data.
func stackTraceID(frames []StackFrame) (string, error) { _ = "STUB: not implemented"; return "", nil }

// reverse ordered frames, done in stackFrames()

// Using FormatUint() or putting AppendUint() into a function leads
// to escaping to heap (allocation).

// make instead of nil avoids a heap allocation

func getLocations(dic pprofile.ProfilesDictionary, stack pprofile.Stack) []pprofile.Location {
	_ = "STUB: not implemented"
	return nil
}

func getString(dic pprofile.ProfilesDictionary, index int) string {
	_ = "STUB: not implemented"
	return ""
}

func getFunction(dic pprofile.ProfilesDictionary, index int) pprofile.Function {
	_ = "STUB: not implemented"
	return *new(pprofile.Function)
}

// return empty function if index is out of bounds

func GetStartOfWeekFromTime(t time.Time) uint32 { _ = "STUB: not implemented"; return 0 }

func addEventHostData(data map[string]string, attrs pcommon.Map) { _ = "STUB: not implemented"; return }

func int64ToBytes(value int64) []byte { _ = "STUB: not implemented"; return nil }

func populateResourceData(dic pprofile.ProfilesDictionary, resource pcommon.Resource, scope pcommon.InstrumentationScope, profile pprofile.Profile) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
