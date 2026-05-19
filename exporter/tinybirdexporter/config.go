// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tinybirdexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tinybirdexporter"

import (
	"regexp"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

var datasourceRegex = regexp.MustCompile(`^[\w_]+$`)

type SignalConfig struct {
	Datasource string `mapstructure:"datasource"`

	_ struct{}
}

func (cfg SignalConfig) Validate() error { _ = "STUB: not implemented"; return nil }

// Config defines configuration for the Tinybird exporter.
type Config struct {
	ClientConfig confighttp.ClientConfig                                  `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.
	RetryConfig  configretry.BackOffConfig                                `mapstructure:"retry_on_failure"`
	QueueConfig  configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`

	// Tinybird API token.
	Token   configopaque.String `mapstructure:"token"`
	Metrics metricSignalConfigs `mapstructure:"metrics"`
	Traces  SignalConfig        `mapstructure:"traces"`
	Logs    SignalConfig        `mapstructure:"logs"`
	// Wait for data to be ingested before returning a response.
	Wait bool `mapstructure:"wait"`
}

type metricSignalConfigs struct {
	MetricsGauge                SignalConfig `mapstructure:"gauge"`
	MetricsSum                  SignalConfig `mapstructure:"sum"`
	MetricsHistogram            SignalConfig `mapstructure:"histogram"`
	MetricsExponentialHistogram SignalConfig `mapstructure:"exponential_histogram"`
}

var _ component.Config = (*Config)(nil)

// Validate checks if the exporter configuration is valid
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }
