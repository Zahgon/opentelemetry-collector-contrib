// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/sampling"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/pkg/samplingpolicy"
)

type Drop struct {
	// the subpolicy evaluators
	subpolicies []samplingpolicy.Evaluator
	logger      *zap.Logger
}

func NewDrop(
	logger *zap.Logger,
	subpolicies []samplingpolicy.Evaluator,
) samplingpolicy.Evaluator {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator)
}

// Evaluate looks at the trace data and returns a corresponding SamplingDecision.
func (c *Drop) Evaluate(ctx context.Context, traceID pcommon.TraceID, trace *samplingpolicy.TraceData) (samplingpolicy.Decision, error) {
	_ = "STUB: not implemented"
	// The policy iterates over all sub-policies and returns Dropped if all
	// sub-policies returned a Sampled Decision. If any subpolicy returns
	// NotSampled, it returns NotSampled Decision.
	return *new(samplingpolicy.Decision), nil
}

//nolint:staticcheck // SA1019: Use of inverted decisions until they are fully removed.

func (c *Drop) IsStateful() bool { _ = "STUB: not implemented"; return false }
