// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metricstarttimeprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstarttimeprocessor"

import (
	"time"

	"go.opentelemetry.io/collector/component"
)

// Config holds configuration of the metric start time processor.
type Config struct {
	Strategy string `mapstructure:"strategy"`

	// GCInterval specifies how long to wait before removing a metric from the
	// cache.
	GCInterval time.Duration `mapstructure:"gc_interval"`

	// StartTimeMetricRegex allows specifying alternate metrics to determine
	// start time using a regular expression. It only applies when the
	// `start_time_metric strategy` is used.
	StartTimeMetricRegex string `mapstructure:"start_time_metric_regex"`
}

var _ component.Config = (*Config)(nil)

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// Validate checks the configuration is valid
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }
