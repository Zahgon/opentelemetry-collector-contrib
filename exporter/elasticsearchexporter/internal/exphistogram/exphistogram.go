// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package exphistogram contains utility functions for exponential histogram conversions.
package exphistogram // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/exphistogram"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// LowerBoundary calculates the lower boundary given index and scale.
// Adopted from https://opentelemetry.io/docs/specs/otel/metrics/data-model/#producer-expectations
func LowerBoundary(index, scale int) float64 { _ = "STUB: not implemented"; return 0 }

// Use this form in case the equation above computes +Inf
// as the lower boundary of a valid bucket.

// LowerBoundaryNegativeScale calculates the lower boundary for scale <= 0.
// Adopted from https://opentelemetry.io/docs/specs/otel/metrics/data-model/#producer-expectations
func LowerBoundaryNegativeScale(index, scale int) float64 { _ = "STUB: not implemented"; return 0 }

// bucketValueFunc computes the representative value for a bucket given its
// lower and upper boundaries.
type bucketValueFunc func(lb, ub float64) float64

// midpointBucketValue returns the midpoint (centroid) of the bucket boundaries.
func midpointBucketValue(lb, ub float64) float64 { _ = "STUB: not implemented"; return 0 }

// rawBucketValue returns the upper boundary of the bucket directly.
func rawBucketValue(_, ub float64) float64 {
	_ = "STUB: not implemented"

	// ToRaw converts an OTLP exponential histogram data point to counts and
	// boundary values without any midpoint approximation. Each bucket's
	// representative value is the upper boundary of the bucket.
	return 0
}

func ToRaw(dp pmetric.ExponentialHistogramDataPoint) (counts []int64, values []float64) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToTDigest converts an OTLP exponential histogram data point to T-Digest counts and mean centroid values.
func ToTDigest(dp pmetric.ExponentialHistogramDataPoint) (counts []int64, values []float64) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toHistogram(dp pmetric.ExponentialHistogramDataPoint, valueFn bucketValueFunc) (counts []int64, values []float64) {
	_ = "STUB: not implemented"
	return nil, nil
}

func safeUint64ToInt64(v uint64) int64 { _ = "STUB: not implemented"; return 0 }

//nolint:goset // overflow checked
