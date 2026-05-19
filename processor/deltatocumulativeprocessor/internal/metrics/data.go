// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/metrics"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type Sum Metric

func (s Sum) Len() int { _ = "STUB: not implemented"; return 0 }

func (s Sum) Ident() Ident { _ = "STUB: not implemented"; return *new(Ident) }

func (s Sum) SetAggregationTemporality(at pmetric.AggregationTemporality) {
	_ = "STUB: not implemented"
	return
}

type Histogram Metric

func (s Histogram) Len() int { _ = "STUB: not implemented"; return 0 }

func (s Histogram) Ident() Ident { _ = "STUB: not implemented"; return *new(Ident) }

func (s Histogram) SetAggregationTemporality(at pmetric.AggregationTemporality) {
	_ = "STUB: not implemented"
	return
}

type ExpHistogram Metric

func (s ExpHistogram) Len() int { _ = "STUB: not implemented"; return 0 }

func (s ExpHistogram) Ident() Ident { _ = "STUB: not implemented"; return *new(Ident) }

func (s ExpHistogram) SetAggregationTemporality(at pmetric.AggregationTemporality) {
	_ = "STUB: not implemented"
	return
}

type Gauge Metric

func (s Gauge) Len() int { _ = "STUB: not implemented"; return 0 }

func (s Gauge) Ident() Ident { _ = "STUB: not implemented"; return *new(Ident) }

func (Gauge) SetAggregationTemporality(pmetric.AggregationTemporality) {
	_ = "STUB: not implemented"
	return
}

type Summary Metric

func (s Summary) Len() int { _ = "STUB: not implemented"; return 0 }

func (s Summary) Ident() Ident { _ = "STUB: not implemented"; return *new(Ident) }

func (Summary) SetAggregationTemporality(pmetric.AggregationTemporality) {
	_ = "STUB: not implemented"
	return
}
