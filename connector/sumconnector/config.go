// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/sumconnector"

import (
	"go.opentelemetry.io/collector/confmap/xconfmap"
)

// Config for the connector
type Config struct {
	Spans      map[string]MetricInfo `mapstructure:"spans"`
	SpanEvents map[string]MetricInfo `mapstructure:"spanevents"`
	Metrics    map[string]MetricInfo `mapstructure:"metrics"`
	DataPoints map[string]MetricInfo `mapstructure:"datapoints"`
	Logs       map[string]MetricInfo `mapstructure:"logs"`
	// prevent unkeyed literal initialization
	_ struct{}
}

// MetricInfo for a data type
type MetricInfo struct {
	Description     string            `mapstructure:"description"`
	Conditions      []string          `mapstructure:"conditions"`
	Attributes      []AttributeConfig `mapstructure:"attributes"`
	SourceAttribute string            `mapstructure:"source_attribute"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type AttributeConfig struct {
	Key          string `mapstructure:"key"`
	DefaultValue any    `mapstructure:"default_value"`
	// prevent unkeyed literal initialization
	_ struct{}
}

func (c *Config) Validate() (combinedErrors error) { _ = "STUB: not implemented"; return nil }

func (i *MetricInfo) validateAttributes() error { _ = "STUB: not implemented"; return nil }

var _ xconfmap.Validator = (*Config)(nil)
