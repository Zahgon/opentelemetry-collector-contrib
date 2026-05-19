// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/sampling"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/pkg/samplingpolicy"
)

type subpolicy struct {
	// the subpolicy evaluator
	evaluator samplingpolicy.Evaluator

	// spans per second allocated to each subpolicy
	allocatedSPS int64

	// spans per second that each subpolicy sampled in this period
	sampledSPS int64

	name string
}

// Composite evaluator and its internal data
type Composite struct {
	// the subpolicy evaluators
	subpolicies []*subpolicy

	// maximum total spans per second that must be sampled
	maxTotalSPS int64

	// current unix timestamp second
	currentSecond int64

	// The time provider (can be different from clock for testing purposes)
	timeProvider TimeProvider

	logger          *zap.Logger
	recordSubPolicy bool
}

var _ samplingpolicy.Evaluator = (*Composite)(nil)

// SubPolicyEvalParams defines the evaluator and max rate for a sub-policy
type SubPolicyEvalParams struct {
	Evaluator         samplingpolicy.Evaluator
	MaxSpansPerSecond int64
	Name              string
}

// NewComposite creates a policy evaluator that samples all subpolicies.
func NewComposite(
	logger *zap.Logger,
	maxTotalSpansPerSecond int64,
	subPolicyParams []SubPolicyEvalParams,
	timeProvider TimeProvider,
	recordSubPolicy bool,
) samplingpolicy.Evaluator {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator)
}

// We are just starting, so there is no previous input, set it to 0

// Evaluate looks at the trace data and returns a corresponding SamplingDecision.
func (c *Composite) Evaluate(ctx context.Context, traceID pcommon.TraceID, trace *samplingpolicy.TraceData) (samplingpolicy.Decision, error) {
	_ = "STUB: not implemented"
	// Rate limiting works by counting spans that are sampled during each 1 second
	// time period. Until the total number of spans during a particular second
	// exceeds the allocated number of spans-per-second the traces are sampled,
	// once the limit is exceeded the traces are no longer sampled. The counter
	// restarts at the beginning of each second.
	// Current counters and rate limits are kept separately for each subpolicy.
	return *new(samplingpolicy.Decision), nil
}

// This is a new second

// Reset counters

//nolint:staticcheck // SA1019: Use of inverted decisions until they are fully removed.

// The subpolicy made a decision to Sample. Now we need to make our decision.

// Calculate resulting SPS counter if we decide to sample this trace

// Check if the rate will be within the allocated bandwidth.

// Let the sampling happen

// We exceeded the rate limit. Don't sample this trace.
// Note that we will continue evaluating new incoming traces against
// allocated SPS, we do not update sub.sampledSPS here in order to give
// chance to another smaller trace to be accepted later.

func (c *Composite) IsStateful() bool { _ = "STUB: not implemented"; return false }
