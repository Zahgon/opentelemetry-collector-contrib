// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package flinkmetricsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/flinkmetricsreceiver"

import (
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/flinkmetricsreceiver/internal/metadata"
)

const defaultEndpoint = "http://localhost:8081"

// Config defines the configuration for the various elements of the receiver agent.
type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	confighttp.ClientConfig        `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
}

// Validate validates the configuration by checking for missing or invalid fields
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }
