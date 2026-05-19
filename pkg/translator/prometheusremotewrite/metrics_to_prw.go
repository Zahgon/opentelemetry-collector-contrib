// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewrite // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"

import (
	"github.com/prometheus/otlptranslator"
	"github.com/prometheus/prometheus/prompb"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type Settings struct {
	Namespace           string
	ExternalLabels      map[string]string
	DisableTargetInfo   bool
	DisableScopeInfo    bool
	AddMetricSuffixes   bool
	TranslationStrategy string
	SendMetadata        bool
}

// FromMetrics converts pmetric.Metrics to Prometheus remote write format.
func FromMetrics(md pmetric.Metrics, settings Settings) (map[string]*prompb.TimeSeries, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// prometheusConverter converts from OTel write format to Prometheus write format.
type prometheusConverter struct {
	unique    map[uint64]*prompb.TimeSeries
	conflicts map[uint64][]*prompb.TimeSeries

	metricNamer otlptranslator.MetricNamer
	labelNamer  otlptranslator.LabelNamer
	unitNamer   otlptranslator.UnitNamer
}

func getTranslationConfiguration(settings Settings) (withSuffixes, utf8Allowed bool) {
	_ = "STUB: not implemented"
	return false, false
}

func newPrometheusConverter(settings Settings) *prometheusConverter {
	_ = "STUB: not implemented"
	return nil
}

// fromMetrics converts pmetric.Metrics to Prometheus remote write format.
func (c *prometheusConverter) fromMetrics(md pmetric.Metrics, settings Settings) (errs error) {
	_ = "STUB: not implemented"
	return nil
}

// keep track of the most recent timestamp in the ResourceMetrics for
// use with the "target" info metric

// TODO: decide if instrumentation library information should be exported as labels

// handle individual metrics based on type
//exhaustive:enforce

// timeSeries returns a slice of the prompb.TimeSeries that were converted from OTel format.
func (c *prometheusConverter) timeSeries() []prompb.TimeSeries {
	_ = "STUB: not implemented"
	return nil
}

func isSameMetric(ts *prompb.TimeSeries, lbls []prompb.Label) bool {
	_ = "STUB: not implemented"
	return false
}

// addExemplars adds exemplars for the dataPoint. For each exemplar, if it can find a bucket bound corresponding to its value,
// the exemplar is added to the bucket bound's time series, provided that the time series' has samples.
func (*prometheusConverter) addExemplars(dataPoint pmetric.HistogramDataPoint, bucketBounds []bucketBoundsData) {
	_ = "STUB: not implemented"
	return
}

// addSample finds a TimeSeries that corresponds to lbls, and adds sample to it.
// If there is no corresponding TimeSeries already, it's created.
// The corresponding TimeSeries is returned.
// If either lbls is nil/empty or sample is nil, nothing is done.
func (c *prometheusConverter) addSample(sample *prompb.Sample, lbls []prompb.Label) *prompb.TimeSeries {
	_ = "STUB: not implemented"
	return nil
}

// This shouldn't happen
