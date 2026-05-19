// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/sampling"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/pkg/samplingpolicy"
)

type And struct {
	// the subpolicy evaluators
	subpolicies []samplingpolicy.Evaluator
	logger      *zap.Logger
}

func NewAnd(
	logger *zap.Logger,
	subpolicies []samplingpolicy.Evaluator,
) samplingpolicy.Evaluator {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator)
}

// Evaluate looks at the trace data and returns a corresponding SamplingDecision.
func (c *And) Evaluate(ctx context.Context, traceID pcommon.TraceID, trace *samplingpolicy.TraceData) (samplingpolicy.Decision, error) {
	_ = "STUB: not implemented"
	// The policy iterates over all sub-policies and returns Sampled if all sub-policies returned a Sampled Decision.
	// If any subpolicy returns NotSampled or InvertNotSampled, it returns NotSampled Decision.
	return *new(samplingpolicy.Decision), nil
}

//nolint:staticcheck // SA1019: Use of inverted decisions until they are fully removed.

func (c *And) IsStateful() bool { _ = "STUB: not implemented"; return false }
