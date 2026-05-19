// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewrite // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"

import (
	"github.com/prometheus/otlptranslator"
	"github.com/prometheus/prometheus/prompb"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

const (
	sumStr        = "_sum"
	countStr      = "_count"
	bucketStr     = "_bucket"
	leStr         = "le"
	quantileStr   = "quantile"
	pInfStr       = "+Inf"
	createdSuffix = "_created"
	// maxExemplarRunes is the maximum number of UTF-8 exemplar characters
	// according to the prometheus specification
	// https://github.com/prometheus/OpenMetrics/blob/v1.0.0/specification/OpenMetrics.md#exemplars
	maxExemplarRunes = 128
	infoType         = "info"
)

type bucketBoundsData struct {
	ts    *prompb.TimeSeries
	bound float64
}

// byBucketBoundsData enables the usage of sort.Sort() with a slice of bucket bounds
type byBucketBoundsData []bucketBoundsData

func (m byBucketBoundsData) Len() int           { _ = "STUB: not implemented"; return 0 }
func (m byBucketBoundsData) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (m byBucketBoundsData) Swap(i, j int)      { _ = "STUB: not implemented"; return }

// ByLabelName enables the usage of sort.Sort() with a slice of labels
type ByLabelName []prompb.Label

func (a ByLabelName) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a ByLabelName) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (a ByLabelName) Swap(i, j int)      { _ = "STUB: not implemented"; return }

// timeSeriesSignature returns a hashed label set signature.
// The label slice should not contain duplicate label names; this method sorts the slice by label name before creating
// the signature.
// The algorithm is the same as in Prometheus' labels.StableHash function.
func timeSeriesSignature(labels []prompb.Label) uint64 { _ = "STUB: not implemented"; return 0 }

// Use xxhash.Sum64(b) for fast path as it's faster.

// If labels entry is 1KB+ do not allocate whole entry.

var seps = []byte{'\xff'}

// createAttributes creates a slice of Prometheus Labels with OTLP attributes and pairs of string values.
// Unpaired string values are ignored. String pairs overwrite OTLP labels if collisions happen and
// if logOnOverwrite is true, the overwrite is logged. Resulting label names are sanitized.
func createAttributes(resource pcommon.Resource, attributes pcommon.Map, scope pcommon.InstrumentationScope, externalLabels map[string]string,
	ignoreAttrs []string, logOnOverwrite bool, labelNamer otlptranslator.LabelNamer, disableScopeInfo bool, extras ...string,
) ([]prompb.Label, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Calculate the maximum possible number of labels we could return so we can preallocate l

// Scope info

// map ensures no duplicate label name

// Ensure attributes are sorted by key for consistent merging of keys which
// collide when sanitized.

// XXX: Should we always drop service namespace/service name/service instance ID from the labels
// (as they get mapped to other Prometheus labels)?

// Only append to existing value if the new value is different

// Map service.name + service.namespace to job

// Map service.instance.id to instance

// External labels have already been sanitized

// Skip external labels if they are overridden by metric attributes

// internal labels should be maintained

// isValidAggregationTemporality checks whether an OTel metric has a valid
// aggregation temporality for conversion to a Prometheus metric.
func isValidAggregationTemporality(metric pmetric.Metric) bool {
	_ = "STUB: not implemented"
	//exhaustive:enforce
	return false
}

func (c *prometheusConverter) addHistogramDataPoints(dataPoints pmetric.HistogramDataPointSlice,
	resource pcommon.Resource, scope pcommon.InstrumentationScope, settings Settings, baseName string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// If the sum is unset, it indicates the _sum metric point should be
// omitted

// treat sum as a sample in an individual TimeSeries

// treat count as a sample in an individual TimeSeries

// cumulative count for conversion to cumulative histogram

// process each bound, based on histograms proto definition, # of buckets = # of explicit bounds + 1

// add le=+Inf bucket

type exemplarType interface {
	pmetric.ExponentialHistogramDataPoint | pmetric.HistogramDataPoint | pmetric.NumberDataPoint
	Exemplars() pmetric.ExemplarSlice
}

func getPromExemplars[T exemplarType](pt T) []prompb.Exemplar {
	_ = "STUB: not implemented"
	return nil
}

// only append filtered attributes if it does not cause exemplar
// labels to exceed the max number of runes

// mostRecentTimestampInMetric returns the latest timestamp in a batch of metrics
func mostRecentTimestampInMetric(metric pmetric.Metric) pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *

	// handle individual metric based on type
	//exhaustive:enforce
	new(pcommon.Timestamp)
}

func (c *prometheusConverter) addSummaryDataPoints(dataPoints pmetric.SummaryDataPointSlice, resource pcommon.Resource, scope pcommon.InstrumentationScope,
	settings Settings, baseName string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// treat sum as a sample in an individual TimeSeries

// sum and count of the summary should append suffix to baseName

// treat count as a sample in an individual TimeSeries

// process each percentile/quantile

// createLabels returns a copy of baseLabels, adding to it the pair model.MetricNameLabel=name.
// If extras are provided, corresponding label pairs are also added to the returned slice.
// If extras is uneven length, the last (unpaired) extra will be ignored.
func createLabels(name string, baseLabels []prompb.Label, extras ...string) []prompb.Label {
	_ = "STUB: not implemented"
	return nil
}

// +1 for name

// getOrCreateTimeSeries returns the time series corresponding to the label set if existent, and false.
// Otherwise it creates a new one and returns that, and true.
func (c *prometheusConverter) getOrCreateTimeSeries(lbls []prompb.Label) (*prompb.TimeSeries, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// We already have this metric

// Look for a matching conflict

// We already have this metric

// New conflict

// This metric is new

// addResourceTargetInfo converts the resource to the target info metric.
func addResourceTargetInfo(resource pcommon.Resource, settings Settings, timestamp pcommon.Timestamp, converter *prometheusConverter) error {
	_ = "STUB: not implemented"
	return nil
}

// If we only have job + instance, then target_info isn't useful, so don't add it.

// We need at least one identifying label to generate target_info.

// convert ns to ms

// convertTimeStamp converts OTLP timestamp in ns to timestamp in ms
func convertTimeStamp(timestamp pcommon.Timestamp) int64 { _ = "STUB: not implemented"; return 0 }
