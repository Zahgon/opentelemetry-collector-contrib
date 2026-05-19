// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/sampling"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/pkg/samplingpolicy"
)

// hasResourceOrSpanWithCondition iterates through all the resources and instrumentation library spans until any
// callback returns true.
func hasResourceOrSpanWithCondition(
	td ptrace.Traces,
	shouldSampleResource func(resource pcommon.Resource) bool,
	shouldSampleSpan func(span ptrace.Span) bool,
) samplingpolicy.Decision {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision)
}

// invertHasResourceOrSpanWithCondition iterates through all the resources and instrumentation library spans until any
// callback returns false.
func invertHasResourceOrSpanWithCondition(
	td ptrace.Traces,
	shouldSampleResource func(resource pcommon.Resource) bool,
	shouldSampleSpan func(span ptrace.Span) bool,
) samplingpolicy.Decision {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision)
}

// hasSpanWithCondition iterates through all the instrumentation library spans until any callback returns true.
func hasSpanWithCondition(td ptrace.Traces, shouldSample func(span ptrace.Span) bool) samplingpolicy.Decision {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision)
}

func hasInstrumentationLibrarySpanWithCondition(ilss ptrace.ScopeSpansSlice, check func(span ptrace.Span) bool, invert bool) bool {
	_ = "STUB: not implemented"
	return false
}

func SetAttrOnScopeSpans(data ptrace.Traces, attrName, attrKey string) {
	_ = "STUB: not implemented"
	return
}

func SetBoolAttrOnScopeSpans(data ptrace.Traces, attrName string, attrValue bool) {
	_ = "STUB: not implemented"
	return
}
