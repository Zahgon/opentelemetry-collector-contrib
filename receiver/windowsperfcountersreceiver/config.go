// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package windowsperfcountersreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsperfcountersreceiver"

import (
	"go.opentelemetry.io/collector/scraper/scraperhelper"
)

// Config defines configuration for WindowsPerfCounters receiver.
type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`

	MetricMetaData map[string]MetricConfig `mapstructure:"metrics"`
	PerfCounters   []ObjectConfig          `mapstructure:"perfcounters"`
}

// MetricsConfig defines the configuration for a metric to be created.
type MetricConfig struct {
	Unit        string      `mapstructure:"unit"`
	Description string      `mapstructure:"description"`
	Gauge       GaugeMetric `mapstructure:"gauge"`
	Sum         SumMetric   `mapstructure:"sum"`
}

type GaugeMetric struct{}

type SumMetric struct {
	Aggregation string `mapstructure:"aggregation"`
	Monotonic   bool   `mapstructure:"monotonic"`
}

// ObjectConfig defines configuration for a perf counter object.
type ObjectConfig struct {
	Object    string          `mapstructure:"object"`
	Instances []string        `mapstructure:"instances"`
	Counters  []CounterConfig `mapstructure:"counters"`
}

// CounterConfig defines the individual counter in an object.
type CounterConfig struct {
	Name          string `mapstructure:"name"`
	MetricRep     `mapstructure:",squash"`
	RecreateQuery bool `mapstructure:"recreate_query"`
}

type MetricRep struct {
	Name       string            `mapstructure:"metric"`
	Attributes map[string]string `mapstructure:"attributes"`
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
