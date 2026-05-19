// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/sampling"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/pkg/samplingpolicy"
)

type traceStateFilter struct {
	key     string
	logger  *zap.Logger
	matcher func(string) bool
}

var _ samplingpolicy.Evaluator = (*traceStateFilter)(nil)

// NewTraceStateFilter creates a policy evaluator that samples all traces with
// the given value by the specific key in the trace_state.
func NewTraceStateFilter(settings component.TelemetrySettings, key string, values []string) samplingpolicy.Evaluator {
	_ = "STUB: not implemented"
	// initialize the exact value map
	return *new(samplingpolicy.Evaluator)
}

// the key-value pair("=" will take one character) in trace_state can't exceed 256 characters

// Evaluate looks at the trace data and returns a corresponding SamplingDecision.
func (tsf *traceStateFilter) Evaluate(_ context.Context, _ pcommon.TraceID, trace *samplingpolicy.TraceData) (samplingpolicy.Decision, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision), nil
}

func (*traceStateFilter) IsStateful() bool { _ = "STUB: not implemented"; return false }
