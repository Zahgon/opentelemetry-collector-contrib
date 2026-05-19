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

const (
	defaultHashSalt = "default-hash-seed"
)

type probabilisticSampler struct {
	logger    *zap.Logger
	threshold uint64
	hashSalt  string
}

var _ samplingpolicy.Evaluator = (*probabilisticSampler)(nil)

// NewProbabilisticSampler creates a policy evaluator that samples a percentage of
// traces.
func NewProbabilisticSampler(settings component.TelemetrySettings, hashSalt string, samplingPercentage float64) samplingpolicy.Evaluator {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator)
}

// calculate threshold once

// Evaluate looks at the trace data and returns a corresponding SamplingDecision.
func (s *probabilisticSampler) Evaluate(_ context.Context, traceID pcommon.TraceID, _ *samplingpolicy.TraceData) (samplingpolicy.Decision, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision), nil
}

func (*probabilisticSampler) IsStateful() bool {
	_ = "STUB: not implemented"

	// calculateThreshold converts a ratio into a value between 0 and MaxUint64
	return false
}

func calculateThreshold(ratio float64) uint64 {
	_ = "STUB: not implemented"
	// Use big.Float and big.Int to calculate threshold because directly convert
	// math.MaxUint64 to float64 will cause digits/bits to be cut off if the converted value
	// doesn't fit into bits that are used to store digits for float64 in Golang
	return 0
}

// hashTraceID creates a hash using the FNV-1a algorithm.
func hashTraceID(salt string, b []byte) uint64 { _ = "STUB: not implemented"; return 0 }

// the implementation fnv.Write() never returns an error, see hash/fnv/fnv.go
