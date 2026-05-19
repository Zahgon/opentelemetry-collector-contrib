// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewrite // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"

import (
	"github.com/prometheus/otlptranslator"
	"github.com/prometheus/prometheus/prompb"
	writev2 "github.com/prometheus/prometheus/prompb/io/prometheus/write/v2"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// FromMetricsV2 converts pmetric.Metrics to Prometheus remote write format 2.0.
func FromMetricsV2(md pmetric.Metrics, settings Settings) (map[string]*writev2.TimeSeries, writev2.SymbolsTable, error) {
	_ = "STUB: not implemented"
	return nil, *new(writev2.SymbolsTable), nil
}

// prometheusConverterV2 converts from OTLP to Prometheus write 2.0 format.
type prometheusConverterV2 struct {
	unique map[uint64]*writev2.TimeSeries
	// conflicts is a map of time series signatures(an unique identifier for TS labels) to a list of TSs with the same signature.
	// this is used to handle conflicts that occur when multiple TSs have the same labels or when different labels generate the same signature.
	conflicts map[uint64][]*writev2.TimeSeries
	// conflictCount is used to track the number of conflicts that were encountered.
	conflictCount int
	symbolTable   writev2.SymbolsTable

	metricNamer otlptranslator.MetricNamer
	labelNamer  otlptranslator.LabelNamer
	unitNamer   otlptranslator.UnitNamer
}

type metadata struct {
	Type writev2.Metadata_MetricType
	Help string
	Unit string
}

func newPrometheusConverterV2(settings Settings) *prometheusConverterV2 {
	_ = "STUB: not implemented"
	return nil
}

// fromMetrics converts pmetric.Metrics to Prometheus remote write format.
func (c *prometheusConverterV2) fromMetrics(md pmetric.Metrics, settings Settings) (errs error) {
	_ = "STUB: not implemented"
	return nil
}

// use with the "target" info metric

// TODO: decide if instrumentation library information should be exported as labels

// Initialize metadata

// handle individual metrics based on type
//exhaustive:enforce

// timeSeries returns a slice of the writev2.TimeSeries that were converted from OTel format.
func (c *prometheusConverterV2) timeSeries() []writev2.TimeSeries {
	_ = "STUB: not implemented"
	return nil
}

func (c *prometheusConverterV2) addSample(sample *writev2.Sample, lbls []prompb.Label, metadata metadata) *writev2.TimeSeries {
	_ = "STUB: not implemented"
	return nil
}

// isSameMetricV2 checks if two time series are the same metric
func isSameMetricV2(ts1, ts2 *writev2.TimeSeries) bool { _ = "STUB: not implemented"; return false }

// getOrCreateTimeSeries returns the time series corresponding to the label set
func (c *prometheusConverterV2) getOrCreateTimeSeries(lbls []prompb.Label, metadata metadata) *writev2.TimeSeries {
	_ = "STUB: not implemented"
	return nil
}

// We already have this metric

// Look for a matching conflict

// We already have this metric

// New conflict

// This metric is new
