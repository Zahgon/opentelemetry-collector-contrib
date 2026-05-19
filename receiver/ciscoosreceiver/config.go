// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ciscoosreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciscoosreceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/xconfmap"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciscoosreceiver/internal/connection"
)

// DeviceConfig represents configuration for a single Cisco device in the devices list.
type DeviceConfig struct {
	// DO NOT USE unkeyed struct initialization
	_ struct{} `mapstructure:"-"`

	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`

	Auth connection.AuthConfig `mapstructure:"auth"`
}

// Config defines configuration for Cisco OS receiver.
type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`

	// Devices is the list of Cisco devices to monitor.
	Devices []DeviceConfig `mapstructure:"devices"`

	Scrapers map[component.Type]component.Config `mapstructure:"-"`
}

var (
	_ xconfmap.Validator  = (*Config)(nil)
	_ confmap.Unmarshaler = (*Config)(nil)
)

// Validate checks the receiver configuration is valid
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// Unmarshal a config.Parser into the config struct.
func (cfg *Config) Unmarshal(componentParser *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil
}

// load the non-dynamic config normally

// dynamically load the individual scraper configs based on the key name
