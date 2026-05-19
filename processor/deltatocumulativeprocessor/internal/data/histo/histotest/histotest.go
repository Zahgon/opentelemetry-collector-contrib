// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package histotest // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/data/histo/histotest"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/data/histo"
)

type Histogram struct {
	Ts pcommon.Timestamp

	Bounds  histo.Bounds
	Buckets []uint64

	Count uint64
	Sum   *float64

	Min, Max *float64
}

func (hist Histogram) Into() pmetric.HistogramDataPoint {
	_ = "STUB: not implemented"
	return *new(pmetric.HistogramDataPoint)
}

type Bounds histo.Bounds

func (bs Bounds) Observe(observations ...float64) Histogram {
	_ = "STUB: not implemented"
	return *new(Histogram)
}

func ptr[T any](v T) *T { _ = "STUB: not implemented"; return nil }
