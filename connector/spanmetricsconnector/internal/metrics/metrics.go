// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector/internal/metrics"

import (
	"time"

	"github.com/lightstep/go-expohisto/structure"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// https://github.com/open-telemetry/opentelemetry-go/blob/3ae002c3caf3e44387f0554dfcbbde2c5aab7909/sdk/metric/internal/aggregate/limit.go#L11C36-L11C50
const overflowKey = "otel.metric.overflow"

type Key string

type HistogramMetrics interface {
	GetOrCreate(key Key, attributesFun BuildAttributesFun, startTimestamp pcommon.Timestamp, lastSeen time.Time) (Histogram, bool)
	BuildMetrics(pmetric.Metric, pcommon.Timestamp, func(Key, pcommon.Timestamp) pcommon.Timestamp, pmetric.AggregationTemporality)
	ClearExemplars()
	ExpireSeries(expiration time.Duration, now time.Time)
}

type Histogram interface {
	Observe(value float64)
	ObserveN(value float64, n uint64)
	AddExemplar(traceID pcommon.TraceID, spanID pcommon.SpanID, value float64)
}

type explicitHistogramMetrics struct {
	metrics          map[Key]*explicitHistogram
	bounds           []float64
	maxExemplarCount int
	cardinalityLimit int
}

type exponentialHistogramMetrics struct {
	metrics          map[Key]*exponentialHistogram
	maxSize          int32
	maxExemplarCount int
	cardinalityLimit int
}

type explicitHistogram struct {
	attributes pcommon.Map
	exemplars  pmetric.ExemplarSlice

	bucketCounts []uint64
	count        uint64
	sum          float64

	bounds []float64

	maxExemplarCount int

	startTimestamp pcommon.Timestamp
	lastSeen       time.Time
}

type exponentialHistogram struct {
	attributes pcommon.Map
	exemplars  pmetric.ExemplarSlice

	histogram *structure.Histogram[float64]

	maxExemplarCount int

	startTimestamp pcommon.Timestamp
	lastSeen       time.Time
}

type BuildAttributesFun func() pcommon.Map

func NewExponentialHistogramMetrics(maxSize int32, maxExemplarCount, cardinalityLimit int) HistogramMetrics {
	_ = "STUB: not implemented"
	return *new(HistogramMetrics)
}

func NewExplicitHistogramMetrics(bounds []float64, maxExemplarCount, cardinalityLimit int) HistogramMetrics {
	_ = "STUB: not implemented"
	return *new(HistogramMetrics)
}

func (m *explicitHistogramMetrics) IsCardinalityLimitReached() bool {
	_ = "STUB: not implemented"
	return false
}

func (m *explicitHistogramMetrics) GetOrCreate(key Key, attributesFun BuildAttributesFun, startTimestamp pcommon.Timestamp, lastSeen time.Time) (Histogram, bool) {
	_ = "STUB: not implemented"
	return *new(Histogram), false
}

// check if overflowKey already exists

func (m *explicitHistogramMetrics) BuildMetrics(
	metric pmetric.Metric,
	timestamp pcommon.Timestamp,
	startTimeStampGenerator func(Key, pcommon.Timestamp) pcommon.Timestamp,
	temporality pmetric.AggregationTemporality,
) {
	_ = "STUB: not implemented"
	return
}

func (m *explicitHistogramMetrics) ClearExemplars() { _ = "STUB: not implemented"; return }

func (m *explicitHistogramMetrics) ExpireSeries(expiration time.Duration, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (m *exponentialHistogramMetrics) IsCardinalityLimitReached() bool {
	_ = "STUB: not implemented"
	return false
}

func (m *exponentialHistogramMetrics) GetOrCreate(key Key, attributesFun BuildAttributesFun, startTimeStamp pcommon.Timestamp, lastSeen time.Time) (Histogram, bool) {
	_ = "STUB: not implemented"
	return *new(Histogram), false
}

// check if overflowKey already exists

func (m *exponentialHistogramMetrics) BuildMetrics(
	metric pmetric.Metric,
	timestamp pcommon.Timestamp,
	startTimeStampGenerator func(Key, pcommon.Timestamp) pcommon.Timestamp,
	temporality pmetric.AggregationTemporality,
) {
	_ = "STUB: not implemented"
	return
}

// expoHistToExponentialDataPoint copies `lightstep/go-expohisto` structure.Histogram to
// pmetric.ExponentialHistogramDataPoint
func expoHistToExponentialDataPoint(agg *structure.Histogram[float64], dp pmetric.ExponentialHistogramDataPoint) {
	_ = "STUB: not implemented"
	return
}

func (m *exponentialHistogramMetrics) ClearExemplars() { _ = "STUB: not implemented"; return }

func (m *exponentialHistogramMetrics) ExpireSeries(expiration time.Duration, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (h *explicitHistogram) Observe(value float64) { _ = "STUB: not implemented"; return }

// Binary search to find the value bucket index.

func (h *explicitHistogram) ObserveN(value float64, n uint64) { _ = "STUB: not implemented"; return }

// Binary search to find the value bucket index.

func (h *explicitHistogram) AddExemplar(traceID pcommon.TraceID, spanID pcommon.SpanID, value float64) {
	_ = "STUB: not implemented"
	return
}

func (h *exponentialHistogram) Observe(value float64) { _ = "STUB: not implemented"; return }

func (h *exponentialHistogram) ObserveN(value float64, n uint64) { _ = "STUB: not implemented"; return }

func (h *exponentialHistogram) AddExemplar(traceID pcommon.TraceID, spanID pcommon.SpanID, value float64) {
	_ = "STUB: not implemented"
	return
}

type Sum struct {
	attributes pcommon.Map
	count      uint64

	exemplars        pmetric.ExemplarSlice
	maxExemplarCount int

	startTimestamp pcommon.Timestamp
	// isFirst is used to track if this datapoint is new to the Sum. This
	// is used to ensure that new Sum metrics being with 0, and then are incremented
	// to the desired value.  This avoids Prometheus throwing away the first
	// value in the series, due to the transition from null -> x.
	isFirst  bool
	lastSeen time.Time
}

func (s *Sum) Add(value uint64) { _ = "STUB: not implemented"; return }

func NewSumMetrics(maxExemplarCount, cardinalityLimit int) SumMetrics {
	_ = "STUB: not implemented"
	return *new(SumMetrics)
}

type SumMetrics struct {
	metrics          map[Key]*Sum
	maxExemplarCount int
	cardinalityLimit int
}

func (m *SumMetrics) IsCardinalityLimitReached() bool { _ = "STUB: not implemented"; return false }

func (m *SumMetrics) GetOrCreate(key Key, attributesFun BuildAttributesFun, startTimestamp pcommon.Timestamp, lastSeen time.Time) (*Sum, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// check when new key coming

// check if overflowKey already exists

func (s *Sum) AddExemplar(traceID pcommon.TraceID, spanID pcommon.SpanID, value float64) {
	_ = "STUB: not implemented"
	return
}

func (m *SumMetrics) BuildMetrics(
	metric pmetric.Metric,
	timestamp pcommon.Timestamp,
	startTimeStampGenerator func(Key, pcommon.Timestamp) pcommon.Timestamp,
	temporality pmetric.AggregationTemporality,
) {
	_ = "STUB: not implemented"
	return
}

func (m *SumMetrics) ClearExemplars() { _ = "STUB: not implemented"; return }

func (m *SumMetrics) ExpireSeries(expiration time.Duration, now time.Time) {
	_ = "STUB: not implemented"
	return
}
