// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mezmoexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/mezmoexporter"

import (
	"time"

	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

const (
	// defaultTimeout
	defaultTimeout time.Duration = 5 * time.Second

	// defaultIngestURL
	defaultIngestURL = "https://logs.mezmo.com/otel/ingest/rest"

	// See https://docs.mezmo.com/docs/Mezmo-ingestion-service-limits for details

	// Maximum payload in bytes that can be POST'd to the REST endpoint
	maxBodySize     = 10 * 1024 * 1024
	maxMessageSize  = 16 * 1024
	maxMetaDataSize = 32 * 1024
	maxAppnameLen   = 512
	maxLogLevelLen  = 80
)

// Config defines configuration for Mezmo exporter.
type Config struct {
	confighttp.ClientConfig   `mapstructure:",squash"`                                 // squash ensures fields are correctly decoded in embedded struct.
	QueueSettings             configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`
	configretry.BackOffConfig `mapstructure:"retry_on_failure"`

	// IngestURL is the URL to send telemetry to.
	IngestURL string `mapstructure:"ingest_url"`

	// Token is the authentication token provided by Mezmo.
	IngestKey configopaque.String `mapstructure:"ingest_key"`
}

// returns default http client settings
func createDefaultClientConfig() confighttp.ClientConfig {
	_ = "STUB: not implemented"
	return *new(confighttp.ClientConfig)
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
