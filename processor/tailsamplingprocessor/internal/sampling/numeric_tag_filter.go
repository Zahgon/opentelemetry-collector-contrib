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

type numericAttributeFilter struct {
	key         string
	minValue    *int64
	maxValue    *int64
	logger      *zap.Logger
	invertMatch bool
}

var _ samplingpolicy.Evaluator = (*numericAttributeFilter)(nil)

// NewNumericAttributeFilter creates a policy evaluator that samples all traces with
// the given attribute in the given numeric range. If minValue is nil, it will use math.MinInt64.
// If maxValue is nil, it will use math.MaxInt64. At least one of minValue or maxValue must be set.
func NewNumericAttributeFilter(settings component.TelemetrySettings, key string, minValue, maxValue *int64, invertMatch bool) samplingpolicy.Evaluator {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator)
}

// Evaluate looks at the trace data and returns a corresponding SamplingDecision.
func (naf *numericAttributeFilter) Evaluate(_ context.Context, _ pcommon.TraceID, trace *samplingpolicy.TraceData) (samplingpolicy.Decision, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision), nil
}

// Get the effective min/max values

func (*numericAttributeFilter) IsStateful() bool { _ = "STUB: not implemented"; return false }
