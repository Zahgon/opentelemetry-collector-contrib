// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package hostmetricsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver"

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/confmap/xconfmap"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
)

// Config defines configuration for HostMetrics receiver.
type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	Scrapers                       map[component.Type]component.Config `mapstructure:"-"`
	// RootPath is the host's root directory (linux only).
	RootPath string `mapstructure:"root_path"`

	// Collection interval for metadata.
	// Metadata of the particular entity is collected when the entity changes.
	// In addition metadata of all entities is collected periodically even if no changes happen.
	// Setting the duration to 0 will disable periodic collection (however will not impact
	// metadata collection on changes).
	MetadataCollectionInterval time.Duration `mapstructure:"metadata_collection_interval"`
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

// dynamically load the individual collector configs based on the key name
