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

type latency struct {
	logger           *zap.Logger
	thresholdMs      int64
	upperThresholdMs int64
}

var _ samplingpolicy.Evaluator = (*latency)(nil)

// NewLatency creates a policy evaluator sampling traces with a duration greater than a configured threshold
func NewLatency(settings component.TelemetrySettings, thresholdMs, upperThresholdMs int64) samplingpolicy.Evaluator {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator)
}

// Evaluate looks at the trace data and returns a corresponding SamplingDecision.
func (l *latency) Evaluate(_ context.Context, _ pcommon.TraceID, traceData *samplingpolicy.TraceData) (samplingpolicy.Decision, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision), nil
}

// IsStateful determines if an evaluator can be used for ingest time decisions.
// In the case of a latency evaluator that is only possible if no upper
// threshold is set as a trace can always get longer in duration and become
// NotSampled.
func (l *latency) IsStateful() bool { _ = "STUB: not implemented"; return false }
