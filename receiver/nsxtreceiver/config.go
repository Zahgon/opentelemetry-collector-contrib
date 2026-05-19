// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package nsxtreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver"

import (
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver/internal/metadata"
)

// Config is the configuration for the NSX receiver
type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	confighttp.ClientConfig        `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	Username                       string              `mapstructure:"username"`
	Password                       configopaque.String `mapstructure:"password"`
}

// Validate returns if the NSX configuration is valid
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
