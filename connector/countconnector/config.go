// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package countconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/countconnector"

import (
	"go.opentelemetry.io/collector/confmap"
)

// Default metrics are emitted if no conditions are specified.
const (
	defaultMetricNameSpans      = "trace.span.count"
	defaultMetricDescSpans      = "The number of spans observed."
	defaultMetricNameSpanEvents = "trace.span.event.count"
	defaultMetricDescSpanEvents = "The number of span events observed."

	defaultMetricNameMetrics    = "metric.count"
	defaultMetricDescMetrics    = "The number of metrics observed."
	defaultMetricNameDataPoints = "metric.datapoint.count"
	defaultMetricDescDataPoints = "The number of data points observed."

	defaultMetricNameLogs = "log.record.count"
	defaultMetricDescLogs = "The number of log records observed."

	defaultMetricNameProfiles = "profile.count"
	defaultMetricDescProfiles = "The number of profiles observed."
)

// Config for the connector
type Config struct {
	Spans      map[string]MetricInfo `mapstructure:"spans"`
	SpanEvents map[string]MetricInfo `mapstructure:"spanevents"`
	Metrics    map[string]MetricInfo `mapstructure:"metrics"`
	DataPoints map[string]MetricInfo `mapstructure:"datapoints"`
	Logs       map[string]MetricInfo `mapstructure:"logs"`
	Profiles   map[string]MetricInfo `mapstructure:"profiles"`
	// prevent unkeyed literal initialization
	_ struct{}
}

// MetricInfo for a data type
type MetricInfo struct {
	Description string            `mapstructure:"description"`
	Conditions  []string          `mapstructure:"conditions"`
	Attributes  []AttributeConfig `mapstructure:"attributes"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type AttributeConfig struct {
	Key          string `mapstructure:"key"`
	DefaultValue any    `mapstructure:"default_value"`
	// prevent unkeyed literal initialization
	_ struct{}
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (i *MetricInfo) validateAttributes() error { _ = "STUB: not implemented"; return nil }

var _ confmap.Unmarshaler = (*Config)(nil)

// Unmarshal with custom logic to override default values if user has specified any custom metrics.
func (c *Config) Unmarshal(componentParser *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil

	// Nothing to do if there is no config given.
}

// Start from defaults provided by createDefaultConfig.
// Unmarshal into a temporary struct and override only sections that are provided and non-empty.

// Spans

// Span events

// Metrics

// Data points

// Logs

// Profiles

func defaultSpansConfig() map[string]MetricInfo { _ = "STUB: not implemented"; return nil }

func defaultSpanEventsConfig() map[string]MetricInfo { _ = "STUB: not implemented"; return nil }

func defaultMetricsConfig() map[string]MetricInfo { _ = "STUB: not implemented"; return nil }

func defaultDataPointsConfig() map[string]MetricInfo { _ = "STUB: not implemented"; return nil }

func defaultLogsConfig() map[string]MetricInfo { _ = "STUB: not implemented"; return nil }

func defaultProfilesConfig() map[string]MetricInfo { _ = "STUB: not implemented"; return nil }
