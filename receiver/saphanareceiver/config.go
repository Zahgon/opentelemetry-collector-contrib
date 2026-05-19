// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package saphanareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/saphanareceiver"

import (
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/saphanareceiver/internal/metadata"
)

// Errors for missing required config parameters.
const (
	ErrNoUsername = "invalid config: missing username"
	ErrNoPassword = "invalid config: missing password" // #nosec G101 - not hardcoded credentials
)

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	confignet.TCPAddrConfig        `mapstructure:",squash"`
	configtls.ClientConfig         `mapstructure:"tls,omitempty"`
	// MetricsBuilderConfig defines which metrics/attributes to enable for the scraper
	metadata.MetricsBuilderConfig `mapstructure:",squash"`

	Username string              `mapstructure:"username"`
	Password configopaque.String `mapstructure:"password"`
}

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }
