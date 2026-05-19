// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

/*
	Encapsulates iteration over the Spans inside ptrace.Traces from the underlying representation.
	Everyone is doing the same kind of iteration and checking over a set traces.
*/

// TraceVisitor interface defines a iteration callback when walking through traces
type TraceVisitor interface {
	// Called for each tuple of Resource, InstrumentationScope, and Span
	// If Visit returns false, the iteration is short-circuited
	visit(resource pcommon.Resource, scope pcommon.InstrumentationScope, span ptrace.Span) (ok bool)
}

// accept method is called to start the iteration process
func accept(traces ptrace.Traces, v TraceVisitor) { _ = "STUB: not implemented"; return }

// Walk each ResourceSpans instance

// instrumentation library is optional
