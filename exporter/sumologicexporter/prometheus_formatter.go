// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumologicexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sumologicexporter"

import (
	"regexp"
	"strings"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type dataPoint interface {
	Timestamp() pcommon.Timestamp
	Attributes() pcommon.Map
}

type prometheusFormatter struct {
	sanitNameRegex *regexp.Regexp
	replacer       *strings.Replacer
}

type prometheusTags string

const (
	prometheusLeTag       string = "le"
	prometheusQuantileTag string = "quantile"
	prometheusInfValue    string = "+Inf"
)

func newPrometheusFormatter() prometheusFormatter {
	_ = "STUB: not implemented"
	return *new(prometheusFormatter)
}

// `\`, `"` and `\n` should be escaped, everything else should be left as-is
// see: https://github.com/prometheus/docs/blob/main/content/docs/instrumenting/exposition_formats.md#line-format

// PrometheusLabels returns all attributes as sanitized prometheus labels string
func (f *prometheusFormatter) tags2String(attr, labels pcommon.Map) prometheusTags {
	_ = "STUB: not implemented"
	return *new(prometheusTags)
}

func formatKeyValuePair(key []byte, value string) string { _ = "STUB: not implemented"; return "" }

// Use strings.Builder and not fmt.Sprintf as it uses significantly less
// allocations.

// We preallocate space for key, value, equal sign and quotes.

// stringsJoinAndSurround joins the strings in s slice using the separator adds front
// to the front of the resulting string and back at the end.
//
// This has a benefit over using the strings.Join() of using just one strings.Builder
// instance and hence using less allocations to produce the final string.
func stringsJoinAndSurround(s []string, separator, front, back string) string {
	_ = "STUB: not implemented"
	return ""
}

// Count the total strings summarized length for the preallocation.

// We preallocate space for all the entires in the provided slice together with
// the separator as well as the surrounding characters.

// sanitizeKeyBytes returns sanitized key byte slice by replacing
// all non-allowed chars with `_`
func (f *prometheusFormatter) sanitizeKeyBytes(s []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// sanitizeValue returns sanitized value string performing the following substitutions:
// `/` -> `//`
// `"` -> `\"`
// "\n" -> `\n`
func (f *prometheusFormatter) sanitizeValue(s string) string { _ = "STUB: not implemented"; return "" }

// doubleLine builds metric based on the given arguments where value is float64
func (f *prometheusFormatter) doubleLine(name string, attributes prometheusTags, value float64, timestamp pcommon.Timestamp) string {
	_ = "STUB: not implemented"
	return ""
}

// intLine builds metric based on the given arguments where value is int64
func (f *prometheusFormatter) intLine(name string, attributes prometheusTags, value int64, timestamp pcommon.Timestamp) string {
	_ = "STUB: not implemented"
	return ""
}

// uintLine builds metric based on the given arguments where value is uint64
func (f *prometheusFormatter) uintLine(name string, attributes prometheusTags, value uint64, timestamp pcommon.Timestamp) string {
	_ = "STUB: not implemented"
	return ""
}

// doubleValueLine returns prometheus line with given value
func (f *prometheusFormatter) doubleValueLine(name string, value float64, dp dataPoint, attributes pcommon.Map) string {
	_ = "STUB: not implemented"
	return ""
}

// uintValueLine returns prometheus line with given value
func (f *prometheusFormatter) uintValueLine(name string, value uint64, dp dataPoint, attributes pcommon.Map) string {
	_ = "STUB: not implemented"
	return ""
}

// numberDataPointValueLine returns prometheus line with value from pmetric.NumberDataPoint
func (f *prometheusFormatter) numberDataPointValueLine(name string, dp pmetric.NumberDataPoint, attributes pcommon.Map) string {
	_ = "STUB: not implemented"
	return ""
}

// sumMetric returns _sum suffixed metric name
func (*prometheusFormatter) sumMetric(name string) string { _ = "STUB: not implemented"; return "" }

// countMetric returns _count suffixed metric name
func (*prometheusFormatter) countMetric(name string) string { _ = "STUB: not implemented"; return "" }

// bucketMetric returns _bucket suffixed metric name
func (*prometheusFormatter) bucketMetric(name string) string { _ = "STUB: not implemented"; return "" }

// mergeAttributes gets two pcommon.Maps and returns new which contains values from both of them
func (*prometheusFormatter) mergeAttributes(attributes, additionalAttributes pcommon.Map) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

// doubleGauge2Strings converts DoubleGauge record to a list of strings (one per dataPoint)
func (f *prometheusFormatter) gauge2Strings(metric pmetric.Metric, attributes pcommon.Map) []string {
	_ = "STUB: not implemented"
	return nil
}

// doubleSum2Strings converts Sum record to a list of strings (one per dataPoint)
func (f *prometheusFormatter) sum2Strings(metric pmetric.Metric, attributes pcommon.Map) []string {
	_ = "STUB: not implemented"
	return nil
}

// summary2Strings converts Summary record to a list of strings
// n+2 where n is number of quantiles and 2 stands for sum and count metrics per each data point
func (f *prometheusFormatter) summary2Strings(metric pmetric.Metric, attributes pcommon.Map) []string {
	_ = "STUB: not implemented"
	return nil
}

// histogram2Strings converts Histogram record to a list of strings,
// (n+1) where n is number of bounds plus two for sum and count per each data point
func (f *prometheusFormatter) histogram2Strings(metric pmetric.Metric, attributes pcommon.Map) []string {
	_ = "STUB: not implemented"
	return nil
}

// according to the spec, it's valid to have no buckets at all

// metric2String returns stringified metricPair
func (f *prometheusFormatter) metric2String(metric pmetric.Metric, attributes pcommon.Map) string {
	_ = "STUB: not implemented"
	return ""
}
