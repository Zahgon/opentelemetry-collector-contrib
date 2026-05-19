// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stefexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/stefexporter"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

// Config defines configuration for STEF exporter.
type Config struct {
	exporterhelper.TimeoutConfig `mapstructure:",squash"`
	QueueConfig                  configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`
	RetryConfig                  configretry.BackOffConfig                                `mapstructure:"retry_on_failure"`
	configgrpc.ClientConfig      `mapstructure:",squash"`
}

var _ component.Config = (*Config)(nil)

// Validate checks if the exporter configuration is valid
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// Validate that the port is in the address

// TODO: move this to configgrpc.ClientConfig to avoid this code duplication (copied from OTLP exporter).
func (c *Config) sanitizedEndpoint() string { _ = "STUB: not implemented"; return "" }
