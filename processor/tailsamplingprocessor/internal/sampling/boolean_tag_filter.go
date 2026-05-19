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

type booleanAttributeFilter struct {
	key         string
	value       bool
	logger      *zap.Logger
	invertMatch bool
}

var _ samplingpolicy.Evaluator = (*booleanAttributeFilter)(nil)

// NewBooleanAttributeFilter creates a policy evaluator that samples all traces with
// the given attribute that match the supplied boolean value.
func NewBooleanAttributeFilter(settings component.TelemetrySettings, key string, value, invertMatch bool) samplingpolicy.Evaluator {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator)
}

// Evaluate looks at the trace data and returns a corresponding SamplingDecision.
func (baf *booleanAttributeFilter) Evaluate(_ context.Context, _ pcommon.TraceID, trace *samplingpolicy.TraceData) (samplingpolicy.Decision, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision), nil
}

func (*booleanAttributeFilter) IsStateful() bool { _ = "STUB: not implemented"; return false }
