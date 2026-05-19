// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ptracetest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/ptracetest"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// CompareTracesOption can be used to mutate expected and/or actual traces before comparing.
type CompareTracesOption interface {
	applyOnTraces(expected, actual ptrace.Traces)
}

type compareTracesOptionFunc func(expected, actual ptrace.Traces)

func (f compareTracesOptionFunc) applyOnTraces(expected, actual ptrace.Traces) {
	_ = "STUB: not implemented"
	return

	// IgnoreResourceAttributeValue is a CompareTracesOption that removes a resource attribute
	// from all resources.
}

func IgnoreResourceAttributeValue(attributeName string) CompareTracesOption {
	_ = "STUB: not implemented"
	return *new(CompareTracesOption)
}

func maskTracesResourceAttributeValue(traces ptrace.Traces, attributeName string) {
	_ = "STUB: not implemented"
	return
}

// IgnoreResourceEntityRefs is a CompareTracesOption that clears entity references
// on all resources.
func IgnoreResourceEntityRefs() CompareTracesOption {
	_ = "STUB: not implemented"
	return *new(CompareTracesOption)
}

func maskTracesResourceEntityRefs(traces ptrace.Traces) { _ = "STUB: not implemented"; return }

// IgnoreResourceSpansOrder is a CompareTracesOption that ignores the order of resource traces/metrics/logs.
func IgnoreResourceSpansOrder() CompareTracesOption {
	_ = "STUB: not implemented"
	return *new(CompareTracesOption)
}

func sortResourceSpansSlice(rms ptrace.ResourceSpansSlice) { _ = "STUB: not implemented"; return }

// IgnoreScopeSpansOrder is a CompareTracesOption that ignores the order of instrumentation scope traces/metrics/logs.
func IgnoreScopeSpansOrder() CompareTracesOption {
	_ = "STUB: not implemented"
	return *new(CompareTracesOption)
}

func sortScopeSpansSlices(ts ptrace.Traces) { _ = "STUB: not implemented"; return }

// IgnoreSpansOrder is a CompareTracesOption that ignores the order of spans.
func IgnoreSpansOrder() CompareTracesOption {
	_ = "STUB: not implemented"
	return *new(CompareTracesOption)
}

func sortSpanSlices(ts ptrace.Traces) { _ = "STUB: not implemented"; return }

// IgnoreSpanID is a CompareTracesOption that clears SpanID fields on all spans.
func IgnoreSpanID() CompareTracesOption {
	_ = "STUB: not implemented"
	return *new(CompareTracesOption)
}

func maskSpanID(traces ptrace.Traces, spanID pcommon.SpanID) { _ = "STUB: not implemented"; return }

// IgnoreSpanAttributeValue is a CompareTracesOption that clears value of the span attribute.
func IgnoreSpanAttributeValue(attributeName string) CompareTracesOption {
	_ = "STUB: not implemented"
	return *new(CompareTracesOption)
}

func maskSpanAttributeValue(traces ptrace.Traces, attributeName string) {
	_ = "STUB: not implemented"
	return
}

// IgnoreScopeSpanInstrumentationScopeName is a CompareTracesOption that clears value of the scope span instrumentation scope name.
func IgnoreScopeSpanInstrumentationScopeName() CompareTracesOption {
	_ = "STUB: not implemented"
	return *new(CompareTracesOption)
}

func maskScopeSpanInstrumentationScopeName(traces ptrace.Traces) { _ = "STUB: not implemented"; return }

// IgnoreScopeSpanInstrumentationScopeVersion is a CompareTracesOption that clears value of the scope span instrumentation scope version.
func IgnoreScopeSpanInstrumentationScopeVersion() CompareTracesOption {
	_ = "STUB: not implemented"
	return *new(CompareTracesOption)
}

func maskScopeSpanInstrumentationScopeVersion(traces ptrace.Traces) {
	_ = "STUB: not implemented"
	return
}

// IgnoreScopeSpanInstrumentationScopeAttributeValue is a CompareTracesOption that clears value of the scope span instrumentation scope name.
func IgnoreScopeSpanInstrumentationScopeAttributeValue(attributeName string) CompareTracesOption {
	_ = "STUB: not implemented"
	return *new(CompareTracesOption)
}

func maskScopeSpanInstrumentationScopeAttributeValue(traces ptrace.Traces, attributeName string) {
	_ = "STUB: not implemented"
	return
}

// IgnoreStartTimestamp is a CompareTracesOption that clears StartTimestamp fields on all spans.
func IgnoreStartTimestamp() CompareTracesOption {
	_ = "STUB: not implemented"
	return *new(CompareTracesOption)
}

func maskStartTimestamp(traces ptrace.Traces, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

// IgnoreEndTimestamp is a CompareTracesOption that clears EndTimestamp fields on all spans.
func IgnoreEndTimestamp() CompareTracesOption {
	_ = "STUB: not implemented"
	return *new(CompareTracesOption)
}

func maskEndTimestamp(traces ptrace.Traces, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

// IgnoreTraceID is a CompareTracesOption that clears TraceID fields on all spans.
func IgnoreTraceID() CompareTracesOption {
	_ = "STUB: not implemented"
	return *new(CompareTracesOption)
}

func maskTraceID(traces ptrace.Traces, traceID pcommon.TraceID) { _ = "STUB: not implemented"; return }
