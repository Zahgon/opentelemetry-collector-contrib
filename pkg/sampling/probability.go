// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/sampling"

import (
	"errors"
)

// ErrProbabilityRange is returned when a value should be in the range [1/MaxAdjustedCount, 1].
var ErrProbabilityRange = errors.New("sampling probability out of the range [1/MaxAdjustedCount, 1]")

// MinSamplingProbability is the smallest representable probability
// and is the inverse of MaxAdjustedCount.
const MinSamplingProbability = 1.0 / float64(MaxAdjustedCount)

// probabilityInRange tests MinSamplingProb <= prob <= 1.
func probabilityInRange(prob float64) bool { _ = "STUB: not implemented"; return false }

// ProbabilityToThreshold converts a probability to a Threshold.  It
// returns an error when the probability is out-of-range.
func ProbabilityToThreshold(prob float64) (Threshold, error) {
	_ = "STUB: not implemented"
	return *new(Threshold), nil
}

// ProbabilityToThresholdWithPrecision is like ProbabilityToThreshold
// with support for reduced precision.  The `precision` argument determines
// how many significant hex digits will be used to encode the exact
// probability.
func ProbabilityToThresholdWithPrecision(fraction float64, precision int) (Threshold, error) {
	_ = "STUB: not implemented"
	// Assume full precision at 0.
	return *new(Threshold), nil
}

// Special case for prob == 1.

// Calculate the amount of precision needed to encode the
// threshold with reasonable precision.  Here, we count the
// number of leading `0` or `f` characters and automatically
// add precision to preserve relative error near the extremes.
//
// Frexp() normalizes both the fraction and one-minus the
// fraction, because more digits of precision are needed if
// either value is near zero.  Frexp returns an exponent <= 0.
//
// If `exp <= -4`, there will be a leading hex `0` or `f`.
// For every multiple of -4, another leading `0` or `f`
// appears, so this raises precision accordingly.

// Compute the threshold

// Round to the specified precision, if less than the maximum.

// Probability is the sampling ratio in the range [MinSamplingProb, 1].
func (th Threshold) Probability() float64 { _ = "STUB: not implemented"; return 0 }
