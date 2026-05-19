// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package spanmetricsconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector"

import (
	"time"

	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/confmap/xconfmap"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector/internal/metrics"
)

const (
	delta      = "AGGREGATION_TEMPORALITY_DELTA"
	cumulative = "AGGREGATION_TEMPORALITY_CUMULATIVE"
)

var defaultHistogramBucketsMs = []float64{
	2, 4, 6, 8, 10, 50, 100, 200, 400, 800, 1000, 1400, 2000, 5000, 10_000, 15_000,
}

var defaultDeltaTimestampCacheSize = 1000

// Dimension defines the dimension name and optional default value if the Dimension is missing from a span attribute.
type Dimension struct {
	Name    string  `mapstructure:"name"`
	Default *string `mapstructure:"default"`
	// prevent unkeyed literal initialization
	_ struct{}
}

// Config defines the configuration options for spanmetricsconnector.
type Config struct {
	// Dimensions defines the list of additional dimensions on top of the provided:
	// - service.name
	// - span.name
	// - span.kind
	// - status.code
	// - collector.instance.id
	// The dimensions will be fetched from the span's attributes. Examples of some conventionally used attributes:
	// https://github.com/open-telemetry/opentelemetry-collector/blob/main/model/semconv/opentelemetry.go.
	Dimensions        []Dimension `mapstructure:"dimensions"`
	CallsDimensions   []Dimension `mapstructure:"calls_dimensions"`
	ExcludeDimensions []string    `mapstructure:"exclude_dimensions"`

	// DimensionsCacheSize defines the size of cache for storing Dimensions, which helps to avoid cache memory growing
	// indefinitely over the lifetime of the collector.
	// Optional. See defaultDimensionsCacheSize in connector.go for the default value.
	// Deprecated [v0.130.0]:  Please use AggregationCardinalityLimit instead
	DimensionsCacheSize int `mapstructure:"dimensions_cache_size"`

	// ResourceMetricsCacheSize defines the size of the cache holding metrics for a service. This is mostly relevant for
	// cumulative temporality to avoid memory leaks and correct metric timestamp resets.
	// Optional. See defaultResourceMetricsCacheSize in connector.go for the default value.
	ResourceMetricsCacheSize int `mapstructure:"resource_metrics_cache_size"`

	// ResourceMetricsKeyAttributes filters the resource attributes used to create the resource metrics key hash.
	// This can be used to avoid situations where resource attributes may change across service restarts, causing
	// metric counters to break (and duplicate). A resource does not need to have all of the attributes. The list
	// must include enough attributes to properly identify unique resources or risk aggregating data from more
	// than one service and span.
	// e.g. ["service.name", "telemetry.sdk.language", "telemetry.sdk.name"]
	// See https://opentelemetry.io/docs/specs/semconv/resource/ for possible attributes.
	ResourceMetricsKeyAttributes []string `mapstructure:"resource_metrics_key_attributes"`

	AggregationTemporality string `mapstructure:"aggregation_temporality"`

	Histogram HistogramConfig `mapstructure:"histogram"`

	// MetricsEmitInterval is the time period between when metrics are flushed or emitted to the configured MetricsExporter.
	MetricsFlushInterval time.Duration `mapstructure:"metrics_flush_interval"`

	// MetricsExpiration is the time period after which, if no new spans are received, metrics are considered stale and will no longer be exported.
	// Default value (0) means that the metrics will never expire.
	MetricsExpiration time.Duration `mapstructure:"metrics_expiration"`

	// SeriesExpiration is the time period after which individual metric series are considered stale and will no longer be exported.
	// Default value (0) means that individual metric series will never expire.
	SeriesExpiration time.Duration `mapstructure:"series_expiration"`

	// TimestampCacheSize controls the size of the cache used to keep track of delta metrics' TimestampUnixNano the last time it was flushed
	TimestampCacheSize *int `mapstructure:"metric_timestamp_cache_size"`

	// Namespace is the namespace of the metrics emitted by the connector.
	Namespace string `mapstructure:"namespace"`

	// Exemplars defines the configuration for exemplars.
	Exemplars ExemplarsConfig `mapstructure:"exemplars"`

	// Events defines the configuration for events section of spans.
	Events EventsConfig `mapstructure:"events"`

	IncludeInstrumentationScope []string `mapstructure:"include_instrumentation_scope"`

	AggregationCardinalityLimit int `mapstructure:"aggregation_cardinality_limit"`

	// Add the resource attributes to the resulting metrics (disabled by default)
	// This option enables the old behavior
	// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/42103
	AddResourceAttributes bool `mapstructure:"add_resource_attributes"`

	// EnableMetricsSamplingMethod adds the sampling.method attribute ("extrapolated" or "counted") to metrics.
	// When false (default), the attribute is not added.
	EnableMetricsSamplingMethod bool `mapstructure:"enable_metrics_sampling_method"`
}

type HistogramConfig struct {
	Disable     bool                                                `mapstructure:"disable"`
	Unit        metrics.Unit                                        `mapstructure:"unit"`
	Exponential configoptional.Optional[ExponentialHistogramConfig] `mapstructure:"exponential"`
	Explicit    configoptional.Optional[ExplicitHistogramConfig]    `mapstructure:"explicit"`
	Dimensions  []Dimension                                         `mapstructure:"dimensions"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type ExemplarsConfig struct {
	Enabled         bool `mapstructure:"enabled"`
	MaxPerDataPoint int  `mapstructure:"max_per_data_point"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type ExponentialHistogramConfig struct {
	MaxSize int32 `mapstructure:"max_size"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type ExplicitHistogramConfig struct {
	// Buckets is the list of durations representing explicit histogram buckets.
	Buckets []time.Duration `mapstructure:"buckets"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type EventsConfig struct {
	// Enabled is a flag to enable events.
	Enabled bool `mapstructure:"enabled"`
	// Dimensions defines the list of dimensions to add to the events metric.
	Dimensions []Dimension `mapstructure:"dimensions"`
	// prevent unkeyed literal initialization
	_ struct{}
}

var _ xconfmap.Validator = (*Config)(nil)

// Validate checks if the processor configuration is valid
func (c Config) Validate() error { _ = "STUB: not implemented"; return nil }

// GetAggregationTemporality converts the string value given in the config into a AggregationTemporality.
// Returns cumulative, unless delta is correctly specified.
func (c Config) GetAggregationTemporality() pmetric.AggregationTemporality {
	_ = "STUB: not implemented"
	return *new(pmetric.AggregationTemporality)
}

func (c Config) GetDeltaTimestampCacheSize() int { _ = "STUB: not implemented"; return 0 }

// validateDimensions checks duplicates for reserved dimensions and additional dimensions.
func validateDimensions(dimensions []Dimension) error { _ = "STUB: not implemented"; return nil }

// validateEventDimensions checks for empty and duplicates for the dimensions configured.
func validateEventDimensions(enabled bool, dimensions []Dimension) error {
	_ = "STUB: not implemented"
	return nil
}
