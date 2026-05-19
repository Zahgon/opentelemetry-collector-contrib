// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics

import (
	"time"

	"github.com/spf13/pflag"

	"github.com/open-telemetry/opentelemetry-collector-contrib/cmd/telemetrygen/internal/config"
)

// Config describes the test scenario.
type Config struct {
	config.Config
	NumMetrics              int
	MetricName              string
	MetricType              MetricType
	AggregationTemporality  AggregationTemporality
	SpanID                  string
	TraceID                 string
	EnforceUniqueTimeseries bool
	UniqueTimelimit         time.Duration
}

// NewConfig creates a new Config with default values.
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// Flags registers config flags.
func (c *Config) Flags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// SetDefaults sets the default values for the configuration
// This is called before parsing the command line flags and when
// calling NewConfig()
func (c *Config) SetDefaults() { _ = "STUB: not implemented"; return }

// Use Gauge as default metric type.

// Use cumulative temporality as default.

// Validate validates the test scenario parameters.
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
