// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package validation // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter/internal/validation"

import (
	"regexp"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

const (
	// InvalidSpanSampleLimit is the maximum number of invalid span samples to collect
	// This is just to not pollute the log output a lot
	InvalidSpanSampleLimit = 5

	// Backend validation constants matching traces-gateway validation
	maxFutureNanos = uint64(1 * time.Hour / time.Nanosecond)
	maxPastNanos   = uint64(24 * time.Hour / time.Nanosecond)
)

// Regex patterns to match backend error messages
var (
	// "Invalid trace identifier: {hex}. A valid trace identifier is a 16-byte array with at least one non-zero byte."
	invalidTraceIDPattern = regexp.MustCompile(`Invalid trace identifier`)

	// "Invalid span identifier: {hex}. A valid span identifier is an 8-byte array with at least one non-zero byte."
	invalidSpanIDPattern = regexp.MustCompile(`Invalid span identifier`)

	// "Invalid span start time: {timestamp}. A span timestamp should be at most 24 hours back and 1 front."
	invalidStartTimePattern = regexp.MustCompile(`Invalid span start time`)

	// "Invalid span duration: start time {start_time_unix_nano} is greater than end time {end_time_unix_nano}."
	invalidDurationPattern = regexp.MustCompile(`Invalid span duration`)

	// "Application name is not available. A "cx.application.name" attribute must be specified for a resource."
	missingAppNamePattern = regexp.MustCompile(`Application name is not available`)

	// "Subsystem name is not available. A "cx.subsystem.name" attribute must be specified for a resource."
	missingSubsystemNamePattern = regexp.MustCompile(`Subsystem name is not available`)

	// "Span {span_id} is too big" - extracts the span ID from error message
	spanTooBigPattern = regexp.MustCompile(`Span ([0-9a-f]+) is too big`)
)

// Resource attribute keys to include in invalid span details for debugging
var resourceAttributeKeys = []string{
	"service.name",
	"service.namespace",
	"service.instance.id",
	"service.version",
	"service.group",
	"k8s.namespace.name",
	"k8s.pod.name",
	"k8s.deployment.name",
	"k8s.statefulset.name",
	"k8s.daemonset.name",
	"aws.ecs.cluster.name",
	"cx.application.name",
	"cx.subsystem.name",
}

// PartialSuccessErrorType represents types of partial success errors from the backend
type PartialSuccessErrorType string

// Partial success error types that can be detected in error messages
const (
	ErrorTypeInvalidDuration  PartialSuccessErrorType = "invalid_duration"
	ErrorTypeInvalidTraceID   PartialSuccessErrorType = "invalid_trace_id"
	ErrorTypeInvalidSpanID    PartialSuccessErrorType = "invalid_span_id"
	ErrorTypeInvalidStartTime PartialSuccessErrorType = "invalid_start_time"
	ErrorTypeNoAppName        PartialSuccessErrorType = "missing_application_name"
	ErrorTypeNoSubsystemName  PartialSuccessErrorType = "missing_subsystem_name"
	ErrorTypeSpanTooBig       PartialSuccessErrorType = "span_too_big"
)

// BaseSpanDetail contains common fields for all span validation details
// NOTE: JSON tags are required for proper serialization by zap logger
type BaseSpanDetail struct {
	TraceID                     string            `json:"trace_id"`
	SpanID                      string            `json:"span_id"`
	SpanName                    string            `json:"span_name"`
	ResourceAttributes          map[string]string `json:"resource_attributes,omitempty"`
	InstrumentationScopeName    string            `json:"instrumentation_scope_name,omitempty"`
	InstrumentationScopeVersion string            `json:"instrumentation_scope_version,omitempty"`
}

// InvalidSpanDetail contains debugging information for spans with invalid durations
type InvalidSpanDetail struct {
	SpanDetails       BaseSpanDetail `json:"span_details"`
	StartTimeUnixNano uint64         `json:"start_time_unix_nano"`
	EndTimeUnixNano   uint64         `json:"end_time_unix_nano"`
	DurationNano      int64          `json:"duration_nano,omitempty"`
}

// InvalidTraceIDDetail contains debugging information for invalid trace IDs
type InvalidTraceIDDetail struct {
	SpanDetails BaseSpanDetail `json:"span_details"`
}

// InvalidSpanIDDetail contains debugging information for invalid span IDs
type InvalidSpanIDDetail struct {
	SpanDetails BaseSpanDetail `json:"span_details"`
}

// InvalidStartTimeDetail contains debugging information for invalid start times
type InvalidStartTimeDetail struct {
	SpanDetails       BaseSpanDetail `json:"span_details"`
	StartTimeUnixNano uint64         `json:"start_time_unix_nano"`
}

// MissingAttributeDetail contains debugging information for missing required attributes
type MissingAttributeDetail struct {
	SpanDetails      BaseSpanDetail `json:"span_details"`
	MissingAttribute string         `json:"missing_attribute"`
}

// TooBigSpanDetail contains debugging information for spans that exceed size limits
type TooBigSpanDetail struct {
	SpanDetails         BaseSpanDetail `json:"span_details"`
	SerializedSizeBytes int            `json:"serialized_size_bytes"`
}

// BuildPartialSuccessLogFieldsForTraces creates log fields for partial success responses
// This is the main entry point for exporters to get validation details as log fields
// It returns ready-to-use zap fields that can be passed directly to the logger
func BuildPartialSuccessLogFieldsForTraces(errorMessage string, td ptrace.Traces, appNameAttr, subsystemNameAttr string) []zap.Field {
	_ = "STUB: not implemented"
	return nil
}

// Missing app or subsysten name should not happen because we are adding them in the
// exporter but checking just in case there is any bug

// collectInvalidDurationSpans scans traces for spans with start time > end time
func collectInvalidDurationSpans(td ptrace.Traces) []InvalidSpanDetail {
	_ = "STUB: not implemented"
	return nil
}

// Skip spans with zero timestamps - backend treats these as incomplete/in-progress

// collectInvalidTraceIDs scans traces for invalid trace IDs (all zeros)
func collectInvalidTraceIDs(td ptrace.Traces) []InvalidTraceIDDetail {
	_ = "STUB: not implemented"
	return nil
}

// collectInvalidSpanIDs scans traces for invalid span IDs (all zeros)
func collectInvalidSpanIDs(td ptrace.Traces) []InvalidSpanIDDetail {
	_ = "STUB: not implemented"
	return nil
}

// collectInvalidStartTimes scans traces for invalid start times (too far in past/future)
func collectInvalidStartTimes(td ptrace.Traces, now time.Time) []InvalidStartTimeDetail {
	_ = "STUB: not implemented"
	return nil
}

// collectMissingAttributes scans traces for missing required attributes
func collectMissingAttributes(td ptrace.Traces, attributeKey string) []MissingAttributeDetail {
	_ = "STUB: not implemented"
	return nil
}

// collectSpansByID finds specific spans by their span ID (from backend error message)
// Backend error format: "Span {span_id} is too big"
func collectSpansByID(td ptrace.Traces, spanIDHex string) []TooBigSpanDetail {
	_ = "STUB: not implemented"
	return nil
}

// collectInvalidSpans is a generic helper that iterates through spans and collects invalid ones
type spanCollectorFunc[T any] func(span ptrace.Span, scope pcommon.InstrumentationScope, resourceAttrs map[string]string) (T, bool)

func collectInvalidSpans[T any](td ptrace.Traces, collector spanCollectorFunc[T]) []T {
	_ = "STUB: not implemented"
	return nil
}

// createBaseSpanDetail creates the base span detail from a span
func createBaseSpanDetail(span ptrace.Span, scope pcommon.InstrumentationScope, resourceAttrs map[string]string) BaseSpanDetail {
	_ = "STUB: not implemented"
	return *new(BaseSpanDetail)
}

// summarizeResourceAttributes extracts specific resource attributes for debugging
func summarizeResourceAttributes(res pcommon.Resource, keys []string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
