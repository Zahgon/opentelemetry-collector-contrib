// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package expo implements various operations on exponential histograms and their bucket counts
package expo // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/data/expo"

import "go.opentelemetry.io/collector/pdata/pmetric"

type (
	DataPoint = pmetric.ExponentialHistogramDataPoint
	Buckets   = pmetric.ExponentialHistogramDataPointBuckets
)

// Abs returns a view into the buckets using an absolute scale
func Abs(bs Buckets) Absolute { _ = "STUB: not implemented"; return *new(Absolute) }

type buckets = Buckets

// Absolute addresses bucket counts using an absolute scale, such that it is
// interoperable with [Scale].
//
// It spans from [[Absolute.Lower]:[Absolute.Upper]]
//
// NOTE: The zero-value is unusable, use [Abs] to construct
type Absolute struct {
	buckets
}

// Abs returns the value at absolute index 'at'
func (a Absolute) Abs(at int) uint64 { _ = "STUB: not implemented"; return 0 }

// Upper returns the minimal index outside the set, such that every i < Upper
func (a Absolute) Upper() int { _ = "STUB: not implemented"; return 0 }

// Lower returns the minimal index inside the set, such that every i >= Lower
func (a Absolute) Lower() int { _ = "STUB: not implemented"; return 0 }

func (a Absolute) idx(at int) (int, bool) { _ = "STUB: not implemented"; return 0, false }
