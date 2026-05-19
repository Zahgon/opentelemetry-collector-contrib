// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package vcenterreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver"

import (
	"net/url"

	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver/internal/metadata"
)

// Config is the configuration of the receiver
type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	configtls.ClientConfig         `mapstructure:"tls,omitempty"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	Endpoint                       string              `mapstructure:"endpoint"`
	Username                       string              `mapstructure:"username"`
	Password                       configopaque.String `mapstructure:"password"`
	MaxQueryMetrics                int                 `mapstructure:"max_query_metrics"`
}

// Validate checks to see if the supplied config will work for the receiver
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// SDKUrl returns the url for the vCenter SDK
func (c *Config) SDKUrl() (*url.URL, error) { _ = "STUB: not implemented"; return nil, nil }
