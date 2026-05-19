// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package histo // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/data/histo"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type DataPoint = pmetric.HistogramDataPoint

type Bounds []float64

// Default boundaries, as defined per SDK spec:
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/metrics/sdk.md#explicit-bucket-histogram-aggregation
var DefaultBounds = Bounds{0, 5, 10, 25, 50, 75, 100, 250, 500, 750, 1000, 2500, 5000, 7500, 10000}

func (bs Bounds) Observe(observations ...float64) DataPoint {
	_ = "STUB: not implemented"
	return *new(DataPoint)
}
