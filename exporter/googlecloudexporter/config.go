// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudexporter"

import (
	"github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/collector"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

// Config defines configuration for Google Cloud exporter.
type Config struct {
	collector.Config `mapstructure:",squash"`

	// Timeout for all API calls. If not set, defaults to 12 seconds.
	TimeoutSettings exporterhelper.TimeoutConfig                             `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.
	QueueSettings   configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`
}

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }
