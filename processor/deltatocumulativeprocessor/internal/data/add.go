// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package data // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/data"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// Aggregator performs an operation on two datapoints.
// Given [pmetric] types are mutable by nature, this logically works as follows:
//
//	*state = op(state, dp)
//
// See [Adder] for an implementation.
type Aggregator interface {
	Numbers(state, dp pmetric.NumberDataPoint) error
	Histograms(state, dp pmetric.HistogramDataPoint) error
	Exponential(state, dp pmetric.ExponentialHistogramDataPoint) error
}

var _ Aggregator = (*Adder)(nil)

// Adder adds (+) datapoints.
type Adder struct{}

var maxBuckets = 160

func (Adder) Numbers(state, dp pmetric.NumberDataPoint) error {
	_ = "STUB: not implemented"
	return nil
}

func (Adder) Histograms(state, dp pmetric.HistogramDataPoint) error {
	_ = "STUB: not implemented"
	// bounds different: no way to merge, so reset observation to new boundaries
	return nil
}

// spec requires len(BucketCounts) == len(ExplicitBounds)+1.
// given we have limited error handling at this stage (and already verified boundaries are correct),
// doing a best-effort add of whatever we have appears reasonable.

func (Adder) Exponential(state, dp pmetric.ExponentialHistogramDataPoint) error {
	_ = "STUB: not implemented"
	return nil
}

// Downscale if an expected number of buckets after the merge is too large.
