// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ptraceutiltest // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector/internal/ptraceutiltest"

import "go.opentelemetry.io/collector/pdata/ptrace"

// TestTraces returns a ptrace.Traces with a uniform structure where resources, scopes, spans,
// and spanevents are identical across all instances, except for one identifying field.
//
// Identifying fields:
// - Resources have an attribute called "resourceName" with a value of "resourceN".
// - Scopes have a name with a value of "scopeN".
// - Spans have a name with a value of "spanN".
// - Span Events have an attribute "spanEventName" with a value of "spanEventN".
//
// Example: TestTraces("AB", "XYZ", "MN", "1234") returns:
//
//	resourceA, resourceB
//	    each with scopeX, scopeY, scopeZ
//	        each with spanM, spanN
//	            each with spanEvent1, spanEvent2, spanEvent3, spanEvent4
//
// Each byte in the input string is a unique ID for the corresponding element.
func NewTraces(resourceIDs, scopeIDs, spanIDs, spanEventIDs string) ptrace.Traces {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces)
}

func NewTracesFromOpts(resources ...ptrace.ResourceSpans) ptrace.Traces {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces)
}

func Resource(id string, scopes ...ptrace.ScopeSpans) ptrace.ResourceSpans {
	_ = "STUB: not implemented"
	return *new(ptrace.ResourceSpans)
}

func Scope(id string, spans ...ptrace.Span) ptrace.ScopeSpans {
	_ = "STUB: not implemented"
	return *new(ptrace.ScopeSpans)
}

func Span(id string, ses ...ptrace.SpanEvent) ptrace.Span {
	_ = "STUB: not implemented"
	return *new(ptrace.Span)
}

func SpanEvent(id string) ptrace.SpanEvent {
	_ = "STUB: not implemented"
	return *new(ptrace.SpanEvent)
}
