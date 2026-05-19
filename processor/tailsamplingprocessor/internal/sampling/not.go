// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/sampling"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/pkg/samplingpolicy"
)

type not struct {
	logger             *zap.Logger
	subPolicyEvaluator samplingpolicy.Evaluator
}

var _ samplingpolicy.Evaluator = (*not)(nil)

// NewNot creates a new not policy evaluator returns the opposite of the decision of the wrapped policy.
func NewNot(logger *zap.Logger, subPolicyEvaluator samplingpolicy.Evaluator) samplingpolicy.Evaluator {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator)
}

// Evaluate looks at the trace data and returns a corresponding SamplingResult.
// The not policy return the opposite of the decision of the wrapped policy.
func (n *not) Evaluate(ctx context.Context, traceID pcommon.TraceID, trace *samplingpolicy.TraceData) (samplingpolicy.Decision, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision), nil
}

// Return the opposite of the decision

// For any other decision types (Unspecified, Pending, Error), just let them bubble up.

func (n *not) IsStateful() bool { _ = "STUB: not implemented"; return false }
