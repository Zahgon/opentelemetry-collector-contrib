// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package golden // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func normalizeTimestamps(metrics pmetric.Metrics) { _ = "STUB: not implemented"; return }

//exhaustive:enforce

// returns a map of the original timestamps with their corresponding normalized values.
// normalization entails setting nonunique subsequent timestamps to the same value while incrementing unique timestamps by a set value of 1,000,000 ns
func normalizeTimeSeries(timeSeries []pcommon.Timestamp) map[pcommon.Timestamp]pcommon.Timestamp {
	_ = "STUB: not implemented"
	return nil
}

// normalize values

func normalTime(timeSeriesIndex int) pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

type dataPointSlice[T dataPoint] interface {
	Len() int
	At(i int) T
}

type dataPoint interface {
	pmetric.NumberDataPoint | pmetric.HistogramDataPoint | pmetric.ExponentialHistogramDataPoint | pmetric.SummaryDataPoint
	Attributes() pcommon.Map
	StartTimestamp() pcommon.Timestamp
	SetStartTimestamp(pcommon.Timestamp)
	Timestamp() pcommon.Timestamp
	SetTimestamp(pcommon.Timestamp)
}

func normalizeDataPointSlice[T dataPoint](dps dataPointSlice[T]) { _ = "STUB: not implemented"; return }

// Find any other data points in the time series
