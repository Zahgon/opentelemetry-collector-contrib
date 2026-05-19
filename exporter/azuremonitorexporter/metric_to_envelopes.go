// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"

import (
	"github.com/microsoft/ApplicationInsights-Go/appinsights/contracts"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type metricPacker struct {
	logger *zap.Logger
}

type timedMetricDataPoint struct {
	dataPoint  *contracts.DataPoint
	timestamp  pcommon.Timestamp
	attributes pcommon.Map
}

type metricTimedData interface {
	getTimedDataPoints() []*timedMetricDataPoint
}

// MetricToEnvelopes packages metrics into a slice of Application Insight envelopes.
func (packer *metricPacker) MetricToEnvelopes(metric pmetric.Metric, resource pcommon.Resource, instrumentationScope pcommon.InstrumentationScope) []*contracts.Envelope {
	_ = "STUB: not implemented"
	return nil
}

func (packer *metricPacker) sanitize(sanitizeFunc func() []string) {
	_ = "STUB: not implemented"
	return
}

func newMetricPacker(logger *zap.Logger) *metricPacker { _ = "STUB: not implemented"; return nil }

func (packer metricPacker) getMetricTimedData(metric pmetric.Metric) metricTimedData {
	_ = "STUB: not implemented"
	//exhaustive:enforce
	return *new(metricTimedData)
}

type scalarMetric struct {
	name           string
	dataPointSlice pmetric.NumberDataPointSlice
}

func newScalarMetric(name string, dataPointSlice pmetric.NumberDataPointSlice) *scalarMetric {
	_ = "STUB: not implemented"
	return nil
}

func (m scalarMetric) getTimedDataPoints() []*timedMetricDataPoint {
	_ = "STUB: not implemented"
	return nil
}

type histogramMetric struct {
	name           string
	dataPointSlice pmetric.HistogramDataPointSlice
}

func newHistogramMetric(name string, dataPointSlice pmetric.HistogramDataPointSlice) *histogramMetric {
	_ = "STUB: not implemented"
	return nil
}

func (m histogramMetric) getTimedDataPoints() []*timedMetricDataPoint {
	_ = "STUB: not implemented"
	return nil
}

type exponentialHistogramMetric struct {
	name           string
	dataPointSlice pmetric.ExponentialHistogramDataPointSlice
}

func newExponentialHistogramMetric(name string, dataPointSlice pmetric.ExponentialHistogramDataPointSlice) *exponentialHistogramMetric {
	_ = "STUB: not implemented"
	return nil
}

func (m exponentialHistogramMetric) getTimedDataPoints() []*timedMetricDataPoint {
	_ = "STUB: not implemented"
	return nil
}

type summaryMetric struct {
	name           string
	dataPointSlice pmetric.SummaryDataPointSlice
}

func newSummaryMetric(name string, dataPointSlice pmetric.SummaryDataPointSlice) *summaryMetric {
	_ = "STUB: not implemented"
	return nil
}

func (m summaryMetric) getTimedDataPoints() []*timedMetricDataPoint {
	_ = "STUB: not implemented"
	return nil
}
