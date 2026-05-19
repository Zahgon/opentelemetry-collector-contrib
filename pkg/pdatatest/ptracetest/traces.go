// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ptracetest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/ptracetest"

import (
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// CompareTraces compares each part of two given Traces and returns
// an error if they don't match. The error describes what didn't match.
func CompareTraces(expected, actual ptrace.Traces, options ...CompareTracesOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching resources so that each can only be matched once

// CompareResourceSpans compares each part of two given ResourceSpans and returns
// an error if they don't match. The error describes what didn't match.
func CompareResourceSpans(expected, actual ptrace.ResourceSpans) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching scope logs so that each record can only be matched once

// CompareScopeSpans compares each part of two given SpanSlices and returns
// an error if they don't match. The error describes what didn't match.
func CompareScopeSpans(expected, actual ptrace.ScopeSpans) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching spans so that each span can only be matched once

// CompareSpan compares each part of two given Span and returns
// an error if they don't match. The error describes what didn't match.
func CompareSpan(expected, actual ptrace.Span) error { _ = "STUB: not implemented"; return nil }

// compareSpanEventSlice compares each part of two given SpanEventSlice and returns
// an error if they don't match. The error describes what didn't match.
func compareSpanEventSlice(expected, actual ptrace.SpanEventSlice) (errs error) {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching span events so that each span event can only be matched once

// CompareSpanEvent compares each part of two given SpanEvent and returns
// an error if they don't match. The error describes what didn't match.
func CompareSpanEvent(expected, actual ptrace.SpanEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// compareSpanLinkSlice compares each part of two given SpanLinkSlice and returns
// an error if they don't match. The error describes what didn't match.
func compareSpanLinkSlice(expected, actual ptrace.SpanLinkSlice) (errs error) {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching span links so that each span link can only be matched once

// CompareSpanLink compares each part of two given SpanLink and returns
// an error if they don't match. The error describes what didn't match.
func CompareSpanLink(expected, actual ptrace.SpanLink) error { _ = "STUB: not implemented"; return nil }
