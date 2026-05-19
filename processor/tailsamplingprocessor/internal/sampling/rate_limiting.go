// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/sampling"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"
	"golang.org/x/time/rate"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/pkg/samplingpolicy"
)

type rateLimiting struct {
	// Rate limiter using golang.org/x/time/rate for efficient token bucket implementation.
	limiter *rate.Limiter
	logger  *zap.Logger
}

var _ samplingpolicy.Evaluator = (*rateLimiting)(nil)

// NewRateLimiting creates a policy evaluator that samples traces based on a span limit per second using a token bucket algorithm.
// The bucket capacity defaults to 2x the spans per second to allow for reasonable burst traffic.
func NewRateLimiting(settings component.TelemetrySettings, spansPerSecond int64) samplingpolicy.Evaluator {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator)
}

// NewRateLimitingWithBurstCapacity creates a rate limiting policy evaluator with a configurable
// burst capacity, using a token bucket algorithm. Tokens (spans) refill continuously at
// spansPerSecond and the bucket holds at most burstCapacity tokens. A single trace whose span
// count exceeds the burst capacity will not pass.
func NewRateLimitingWithBurstCapacity(settings component.TelemetrySettings, spansPerSecond, burstCapacity int64) samplingpolicy.Evaluator {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator)
}

// Evaluate looks at the trace data and returns a corresponding SamplingDecision based on token bucket consumption.
func (r *rateLimiting) Evaluate(_ context.Context, _ pcommon.TraceID, trace *samplingpolicy.TraceData) (samplingpolicy.Decision, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision), nil
}

func (*rateLimiting) IsStateful() bool { _ = "STUB: not implemented"; return false }
