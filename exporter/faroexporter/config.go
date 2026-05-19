// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package faroexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/faroexporter"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

// Config defines configuration settings for the Faro exporter.
type Config struct {
	confighttp.ClientConfig `mapstructure:",squash"`
	QueueConfig             configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`
	RetryConfig             configretry.BackOffConfig                                `mapstructure:"retry_on_failure"`
}

var (
	_ component.Config    = (*Config)(nil)
	_ confmap.Unmarshaler = (*Config)(nil)
)

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (c *Config) Unmarshal(component *confmap.Conf) error { _ = "STUB: not implemented"; return nil }
