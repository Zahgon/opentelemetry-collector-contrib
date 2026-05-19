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

type alwaysSample struct {
	logger *zap.Logger
}

var _ samplingpolicy.Evaluator = (*alwaysSample)(nil)

// NewAlwaysSample creates a policy evaluator the samples all traces.
func NewAlwaysSample(settings component.TelemetrySettings) samplingpolicy.Evaluator {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator)
}

// Evaluate looks at the trace data and returns a corresponding SamplingDecision.
func (as *alwaysSample) Evaluate(context.Context, pcommon.TraceID, *samplingpolicy.TraceData) (samplingpolicy.Decision, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision), nil
}

func (*alwaysSample) IsStateful() bool { _ = "STUB: not implemented"; return false }
