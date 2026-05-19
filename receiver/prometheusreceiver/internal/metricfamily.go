// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver/internal"

import (
	"github.com/prometheus/prometheus/model/exemplar"
	"github.com/prometheus/prometheus/model/histogram"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/scrape"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type metricFamily struct {
	mtype pmetric.MetricType
	// isMonotonic only applies to sums
	isMonotonic bool
	groups      map[uint64]*metricGroup
	name        string
	metadata    *scrape.MetricMetadata
	groupOrders []*metricGroup
}

// metricGroup, represents a single metric of a metric family. for example a histogram metric is usually represent by
// a couple data complexValue (buckets and count/sum), a group of a metric family always share a same set of tags. for
// simple types like counter and gauge, each data point is a group of itself
type metricGroup struct {
	mtype    pmetric.MetricType
	ts       int64
	ls       labels.Labels
	count    float64
	hasCount bool
	sum      float64
	hasSum   bool
	// This corresponds to the `_created` sample found from the metric parsing.
	// - https://github.com/prometheus/OpenMetrics/blob/main/specification/OpenMetrics.md#timestamps
	// - https://github.com/prometheus/OpenMetrics/blob/main/specification/OpenMetrics.md#counter-1
	createdSeconds float64
	value          float64
	hasValue       bool
	hValue         *histogram.Histogram
	fhValue        *histogram.FloatHistogram
	complexValue   []*dataPoint
	exemplars      pmetric.ExemplarSlice
	isNHCB         bool // true if this is a Native Histogram Custom Buckets (schema -53)
}

func newMetricFamily(metricName string, mc scrape.MetricMetadataStore, logger *zap.Logger, isNativeHistogram, isNHCB bool) *metricFamily {
	_ = "STUB: not implemented"
	return nil
}

// Native histograms have intrinsic metric type, use it,
// regardless of what metadata says.

// includesMetric returns true if the metric is part of the family
func (mf *metricFamily) includesMetric(metricName string) bool {
	_ = "STUB: not implemented"
	return false
}

// If it is a merged family type, then it should match the
// family name when suffixes are trimmed.

// If it isn't a merged type, the metricName and family name should match

func (mg *metricGroup) sortPoints() { _ = "STUB: not implemented"; return }

func (mg *metricGroup) toDistributionPoint(dest pmetric.HistogramDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

// if the final bucket is +Inf, we ignore it

// for OTLP the bounds won't include +inf

// Buckets still need to be sent to know to set them as stale,
// but a staleness NaN converted to uint64 would be an extremely large number.
// Setting to 0 instead.

// Add the final bucket based on the total count

// The timestamp MUST be in retrieved from milliseconds and converted to nanoseconds.

// toExponentialHistogramDataPoints is based on
// https://opentelemetry.io/docs/specs/otel/compatibility/prometheus_and_openmetrics/#exponential-histograms
func (mg *metricGroup) toExponentialHistogramDataPoints(dest pmetric.ExponentialHistogramDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

// We do not set Min or Max as native histograms don't have that information.

// The count and sum are initialized to 0, so we don't need to set them.

// Input is a float native histogram. This conversion will lose
// precision,but we don't actually expect float histograms in scrape,
// since these are typically the result of operations on integer
// native histograms in the database.

// -1 because OTEL offset are for the lower bound, not the upper bound

// -1 because OTEL offset are for the lower bound, not the upper bound

// The count and sum are initialized to 0, so we don't need to set them.

// -1 because OTEL offset are for the lower bound, not the upper bound

// -1 because OTEL offset are for the lower bound, not the upper bound

// This should never happen.

func convertDeltaBuckets(spans []histogram.Span, deltas []int64, buckets pcommon.UInt64Slice) {
	_ = "STUB: not implemented"
	return
}

func convertAbsoluteBuckets(spans []histogram.Span, counts []float64, buckets pcommon.UInt64Slice) {
	_ = "STUB: not implemented"
	return
}

// convertNHCBBDeltBuckets converts NHCB delta buckets to otel bucket counts.
func convertNHCBBDeltBuckets(histogram *histogram.Histogram) []uint64 {
	_ = "STUB: not implemented"
	return nil
}

// convertNHCBAbsoluteBuckets converts NHCB absolute buckets to otel bucket counts.
func convertNHCBAbsoluteBuckets(histogram *histogram.FloatHistogram) []uint64 {
	_ = "STUB: not implemented"
	return nil
}

// This intentionally truncates the float value to an integer (e.g. 5.7 becomes 5).

func (mg *metricGroup) setExemplars(exemplars pmetric.ExemplarSlice) {
	_ = "STUB: not implemented"
	return
}

func (mg *metricGroup) toSummaryPoint(dest pmetric.SummaryDataPointSlice) {
	_ = "STUB: not implemented"
	// expecting count to be provided, however, in the following two cases, they can be missed.
	// 1. data is corrupted
	// 2. ignored by startValue evaluation
	return
}

// Quantiles still need to be sent to know to set them as stale,
// but a staleness NaN converted to uint64 would be an extremely large number.
// By not setting the quantile value, it will default to 0.

// Based on the summary description from https://prometheus.io/docs/concepts/metric_types/#summary
// the quantiles are calculated over a sliding time window, however, the count is the total count of
// observations and the corresponding sum is a sum of all observed values, thus the sum and count used
// at the global level of the metricspb.SummaryValue
// The timestamp MUST be in retrieved from milliseconds and converted to nanoseconds.

func (mg *metricGroup) toNumberDataPoint(dest pmetric.NumberDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

// gauge/undefined types have no start time.

func populateAttributes(mType pmetric.MetricType, ls labels.Labels, dest pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

// empty label values should be omitted

func (mf *metricFamily) loadMetricGroupOrCreate(groupKey uint64, ls labels.Labels, ts int64) *metricGroup {
	_ = "STUB: not implemented"
	return nil
}

// maintaining data insertion order is helpful to generate stable/reproducible metric output

func (mf *metricFamily) addSeries(seriesRef uint64, metricName string, ls labels.Labels, t int64, v float64) error {
	_ = "STUB: not implemented"
	return nil
}

// always use the timestamp from count, because is the only required field for histograms and summaries.

// addCreationTimestamp updates the metric group cache with the created timestamp for the group.
// The parser gets the created time in ms however and we must convert it to seconds here.
// - https://github.com/prometheus/prometheus/blob/2bf6f4c9dcbb1ad2e8fef70c6a48d8fc44a7f57c/model/textparse/interface.go#L77-L80
func (mf *metricFamily) addCreationTimestamp(seriesRef uint64, ls labels.Labels, atMs, ctMs int64) {
	_ = "STUB: not implemented"
	return
}

func (mf *metricFamily) addExponentialHistogramSeries(seriesRef uint64, metricName string, ls labels.Labels, t int64, h *histogram.Histogram, fh *histogram.FloatHistogram) error {
	_ = "STUB: not implemented"
	return nil
}

// addNHCBSeries adds a Native Histogram Custom Buckets (NHCB) series to the metric family.
func (mf *metricFamily) addNHCBSeries(seriesRef uint64, metricName string, ls labels.Labels, t int64, h *histogram.Histogram, fh *histogram.FloatHistogram) error {
	_ = "STUB: not implemented"
	return nil
}

func (mf *metricFamily) appendMetric(metrics pmetric.MetricSlice, trimSuffixes bool) {
	_ = "STUB: not implemented"
	return
}

// Trims type and unit suffixes from metric name

// Everything else should be set to a Gauge.

func (mf *metricFamily) addExemplar(seriesRef uint64, e exemplar.Exemplar) {
	_ = "STUB: not implemented"
	return
}

func convertExemplar(pe exemplar.Exemplar, e pmetric.Exemplar) { _ = "STUB: not implemented"; return }

/*
	decodeAndCopyToLowerBytes copies src to dst on lower bytes instead of higher

1. If len(src) > len(dst) -> copy first len(dst) bytes as it is. Example -> src = []byte{0xab,0xcd,0xef,0xgh,0xij}, dst = [2]byte, result dst = [2]byte{0xab, 0xcd}
2. If len(src) = len(dst) -> copy src to dst as it is
3. If len(src) < len(dst) -> prepend required 0s and then add src to dst. Example -> src = []byte{0xab, 0xcd}, dst = [8]byte, result dst = [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xab, 0xcd}
*/
func decodeAndCopyToLowerBytes(dst, src []byte) error { _ = "STUB: not implemented"; return nil }
