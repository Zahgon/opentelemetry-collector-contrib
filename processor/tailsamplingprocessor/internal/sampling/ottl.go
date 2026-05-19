// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/sampling"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/pkg/samplingpolicy"
)

type ottlConditionFilter struct {
	sampleSpanExpr      *ottl.ConditionSequence[*ottlspan.TransformContext]
	sampleSpanEventExpr *ottl.ConditionSequence[*ottlspanevent.TransformContext]
	errorMode           ottl.ErrorMode
	logger              *zap.Logger
}

var _ samplingpolicy.Evaluator = (*ottlConditionFilter)(nil)

// NewOTTLConditionFilter looks at the trace data and returns a corresponding SamplingDecision.
func NewOTTLConditionFilter(settings component.TelemetrySettings, spanConditions, spanEventConditions []string, errMode ottl.ErrorMode) (samplingpolicy.Evaluator, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator), nil
}

func (ocf *ottlConditionFilter) Evaluate(ctx context.Context, traceID pcommon.TraceID, trace *samplingpolicy.TraceData) (samplingpolicy.Decision, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision), nil
}

// Now we reach span level and begin evaluation with parsed expr.
// The evaluation will break when:
// 1. error happened.
// 2. "Sampled" decision made.
// Otherwise, it will keep evaluating and finally exit with "NotSampled" decision.

// Span evaluation

// Span event evaluation

func (*ottlConditionFilter) IsStateful() bool { _ = "STUB: not implemented"; return false }
