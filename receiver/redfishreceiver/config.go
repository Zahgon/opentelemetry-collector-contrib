// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package redfishreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redfishreceiver"

import (
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redfishreceiver/internal/metadata"
)

type redfishConfig struct {
	Version string `mapstructure:"version"`
}

type Server struct {
	BaseURL          string              `mapstructure:"base_url"`
	User             string              `mapstructure:"username"`
	Pwd              configopaque.String `mapstructure:"password"`
	Insecure         bool                `mapstructure:"insecure"`
	Timeout          string              `mapstructure:"timeout"`
	Redfish          redfishConfig       `mapstructure:"redfish"`
	ComputerSystemID string              `mapstructure:"computer_system_id"`
	Resources        []Resource          `mapstructure:"resources"`
}

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	Servers                        []Server `mapstructure:"servers"`

	// prevent unkeyed literal initialization
	_ struct{}
}

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }
