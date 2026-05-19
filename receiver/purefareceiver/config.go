// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package purefareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefareceiver"

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefareceiver/internal"
)

var _ component.Config = (*Config)(nil)

// Config relating to Array Metric Scraper.
type Config struct {
	confighttp.ClientConfig `mapstructure:",squash"`

	// Settings contains settings for the individual scrapers
	Settings *Settings `mapstructure:"settings"`

	// Array represents the list of arrays to query
	Array []internal.ScraperConfig `mapstructure:"array"`

	// Hosts represents the list of hosts to query
	Hosts []internal.ScraperConfig `mapstructure:"hosts"`

	// Directories represents the list of directories to query
	Directories []internal.ScraperConfig `mapstructure:"directories"`

	// Pods represents the list of pods to query
	Pods []internal.ScraperConfig `mapstructure:"pods"`

	// Volumes represents the list of volumes to query
	Volumes []internal.ScraperConfig `mapstructure:"volumes"`

	// Env represents the respective environment value valid to scrape
	Env string `mapstructure:"env"`

	// ArrayName represents the display name that is appended to the received metrics, as the `host` label if not provided by OpenMetrics output, and to the `fa_array_name` label always
	ArrayName string `mapstructure:"fa_array_name"`

	// Namespace selects the OpenMetrics namespace requested
	Namespace string `mapstructure:"namespace"`
}

type Settings struct {
	ReloadIntervals *ReloadIntervals `mapstructure:"reload_intervals"`

	// prevent unkeyed literal initialization
	_ struct{}
}

type ReloadIntervals struct {
	Array       time.Duration `mapstructure:"array"`
	Hosts       time.Duration `mapstructure:"hosts"`
	Directories time.Duration `mapstructure:"directories"`
	Pods        time.Duration `mapstructure:"pods"`
	Volumes     time.Duration `mapstructure:"volumes"`

	// prevent unkeyed literal initialization
	_ struct{}
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
