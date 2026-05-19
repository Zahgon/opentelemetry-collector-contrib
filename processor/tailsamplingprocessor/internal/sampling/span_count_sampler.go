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

type spanCount struct {
	logger   *zap.Logger
	minSpans int64
	maxSpans int64
}

var _ samplingpolicy.Evaluator = (*spanCount)(nil)

// NewSpanCount creates a policy evaluator sampling traces with more than one span per trace
func NewSpanCount(settings component.TelemetrySettings, minSpans, maxSpans int32) samplingpolicy.Evaluator {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator)
}

// Evaluate looks at the trace data and returns a corresponding SamplingDecision.
func (c *spanCount) Evaluate(_ context.Context, _ pcommon.TraceID, traceData *samplingpolicy.TraceData) (samplingpolicy.Decision, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision), nil
}

// IsStateful determines if an evaluator can be used for ingest time decisions.
// In the case of a span count evaluator that is only possible if no max span
// count is set as a trace can always receive more spans and become NotSampled.
func (c *spanCount) IsStateful() bool { _ = "STUB: not implemented"; return false }
