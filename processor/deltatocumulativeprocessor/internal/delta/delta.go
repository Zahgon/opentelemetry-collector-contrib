// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package delta // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/delta"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/data"
)

type ErrOlderStart struct {
	Start  pcommon.Timestamp
	Sample pcommon.Timestamp
}

func (e ErrOlderStart) Error() string { _ = "STUB: not implemented"; return "" }

type ErrOutOfOrder struct {
	Last   pcommon.Timestamp
	Sample pcommon.Timestamp
}

func (e ErrOutOfOrder) Error() string { _ = "STUB: not implemented"; return "" }

type Type[Self any] interface {
	pmetric.NumberDataPoint | pmetric.HistogramDataPoint | pmetric.ExponentialHistogramDataPoint

	StartTimestamp() pcommon.Timestamp
	Timestamp() pcommon.Timestamp
	SetTimestamp(pcommon.Timestamp)
	CopyTo(Self)
}

type Aggregator struct {
	data.Aggregator
}

func Aggregate[T Type[T]](state, dp T, aggregate func(state, dp T) error) error {
	_ = "STUB: not implemented"
	return nil
}

// first sample of series, no state to aggregate with

// belongs to older series

// out of order

func (aggr Aggregator) Numbers(state, dp pmetric.NumberDataPoint) error {
	_ = "STUB: not implemented"
	return nil
}

func (aggr Aggregator) Histograms(state, dp pmetric.HistogramDataPoint) error {
	_ = "STUB: not implemented"
	return nil
}

func (aggr Aggregator) Exponential(state, dp pmetric.ExponentialHistogramDataPoint) error {
	_ = "STUB: not implemented"
	return nil
}
