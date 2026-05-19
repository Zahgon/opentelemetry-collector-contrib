// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter"

import (
	"errors"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/model"
	"github.com/prometheus/otlptranslator"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

var separatorString = string([]byte{model.SeparatorByte})

type collector struct {
	accumulator accumulator
	logger      *zap.Logger

	sendTimestamps   bool
	namespace        string
	constLabels      prometheus.Labels
	metricFamilies   sync.Map
	metricExpiration time.Duration
	withoutScopeInfo bool

	metricNamer otlptranslator.MetricNamer
	labelNamer  otlptranslator.LabelNamer
}

type metricFamily struct {
	lastSeen time.Time
	mf       *dto.MetricFamily
}

func newCollector(config *Config, logger *zap.Logger) *collector {
	_ = "STUB: not implemented"
	return nil
}

// normalizeNamespace builds and returns the namespace if specified in the config
// If not specified, it returns an empty string
// If building the namespace fails, it logs the error and returns an empty string
func normalizeNamespace(configNamespace string, labelNamer otlptranslator.LabelNamer, logger *zap.Logger) string {
	_ = "STUB: not implemented"
	return ""
}

// configureMetricNamer configures the MetricNamer based on the translation strategy or legacy configuration
func configureMetricNamer(config *Config) otlptranslator.MetricNamer {
	_ = "STUB: not implemented"
	return *new(otlptranslator.MetricNamer)
}

// configureLabelNamer configures the LabelNamer based on the translation strategy or legacy configuration
func configureLabelNamer(config *Config) otlptranslator.LabelNamer {
	_ = "STUB: not implemented"
	return *new(otlptranslator.LabelNamer)
}

// getTranslationConfiguration returns the translation configuration based on the strategy or legacy settings
// Returns (withSuffixes, allowUTF8)
func getTranslationConfiguration(config *Config) (bool, bool) {
	_ = "STUB: not implemented"
	// If TranslationStrategy is explicitly set, use it (takes precedence)
	return false, false
}

// Fallback to default behavior, suffixes enabled, UTF-8 escaped to underscores.

// If feature gate is enabled, ignore AddMetricSuffixes (for deprecation)

// Default to UnderscoreEscapingWithSuffixes behavior when AddMetricSuffixes is deprecated

// Fall back to legacy AddMetricSuffixes behavior, UTF-8 escaped to underscores.

func convertExemplars(exemplars pmetric.ExemplarSlice) []prometheus.Exemplar {
	_ = "STUB: not implemented"
	return nil
}

// Describe is a no-op, because the collector dynamically allocates metrics.
// https://github.com/prometheus/client_golang/blob/v1.9.0/prometheus/collector.go#L28-L40
func (*collector) Describe(chan<- *prometheus.Desc) {
	_ = "STUB: not implemented"

	/*
	   Processing
	*/return
}

func (c *collector) processMetrics(rm pmetric.ResourceMetrics) (n int) {
	_ = "STUB: not implemented"
	return 0
}

var errUnknownMetricType = errors.New("unknown metric type")

func (c *collector) convertMetric(metric pmetric.Metric, resourceAttrs pcommon.Map, scopeName, scopeVersion, scopeSchemaURL string, scopeAttributes pcommon.Map) (prometheus.Metric, error) {
	_ = "STUB: not implemented"
	return *new(prometheus.Metric), nil
}

// defaultZeroThreshold matches the remote-write translator's default for native histograms
// when an explicit zero threshold is not provided in the datapoint.
const (
	defaultZeroThreshold = 1e-128
	cbnhScale            = -53
)

func bucketsToNativeMap(buckets pmetric.ExponentialHistogramDataPointBuckets, scaleDown int32) map[int]int64 {
	_ = "STUB: not implemented"
	return nil
}

// Effective bucket index after downscaling: ((offset + i) >> scaleDown) + 1

func (c *collector) convertExponentialHistogram(metric pmetric.Metric, resourceAttrs pcommon.Map, scopeName, scopeVersion, scopeSchemaURL string, scopeAttributes pcommon.Map) (prometheus.Metric, error) {
	_ = "STUB: not implemented"
	return *new(prometheus.Metric), nil
}

// Build metadata/labels first.

// Schema -53 (CBNH) is already supported via the classic histogram path.
// The Prometheus receiver converts NHCB to OTLP classic Histogram (not ExponentialHistogram),
// which is then handled by convertDoubleHistogram.

// Use created timestamp if start time is set (> 0), else zero value.

func (c *collector) getMetricMetadata(metric pmetric.Metric, mType *dto.MetricType, attributes, resourceAttrs pcommon.Map, scopeName, scopeVersion, scopeSchemaURL string, scopeAttributes pcommon.Map) (*prometheus.Desc, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// +2 for job and instance labels, +3 for scope name, version and schema url

func isReservedScopeAttribute(k string) bool { _ = "STUB: not implemented"; return false }

func (c *collector) convertGauge(metric pmetric.Metric, resourceAttrs pcommon.Map, scopeName, scopeVersion, scopeSchemaURL string, scopeAttributes pcommon.Map) (prometheus.Metric, error) {
	_ = "STUB: not implemented"
	return *new(prometheus.Metric), nil
}

func (c *collector) convertSum(metric pmetric.Metric, resourceAttrs pcommon.Map, scopeName, scopeVersion, scopeSchemaURL string, scopeAttributes pcommon.Map) (prometheus.Metric, error) {
	_ = "STUB: not implemented"
	return *new(prometheus.Metric), nil
}

// Prometheus currently only supports exporting counters

func (c *collector) convertSummary(metric pmetric.Metric, resourceAttrs pcommon.Map, scopeName, scopeVersion, scopeSchemaURL string, scopeAttributes pcommon.Map) (prometheus.Metric, error) {
	_ = "STUB: not implemented"
	// TODO: In the off chance that we have multiple points
	// within the same metric, how should we handle them?
	return *new(prometheus.Metric), nil
}

// There should be EXACTLY one quantile value lest it is an invalid exposition.

func (c *collector) convertDoubleHistogram(metric pmetric.Metric, resourceAttrs pcommon.Map, scopeName, scopeVersion, scopeSchemaURL string, scopeAttributes pcommon.Map) (prometheus.Metric, error) {
	_ = "STUB: not implemented"
	return *new(prometheus.Metric), nil
}

func (c *collector) createTargetInfoMetrics(resourceAttrs []pcommon.Map) ([]prometheus.Metric, error) {
	_ = "STUB: not implemented"

	// deduplicate resourceAttrs by job and instance
	return nil, nil
}

// map ensures no duplicate label name
// +2 for job and instance labels.

// Use resource attributes (other than those used for job+instance) as the
// metric labels for the target info metric

// Remove resource attributes used for job + instance

// Map service.name + service.namespace to job

// Map service.instance.id to instance

/*
Reporting
*/
func (c *collector) Collect(ch chan<- prometheus.Metric) { _ = "STUB: not implemented"; return }

func (c *collector) validateMetrics(name, description string, metricType *dto.MetricType) (help string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *collector) cleanupMetricFamilies() { _ = "STUB: not implemented"; return }
