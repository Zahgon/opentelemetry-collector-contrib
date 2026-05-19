// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package system // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/system"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/system/internal/metadata"
)

// Config defines user-specified configurations unique to the system detector
type Config struct {
	// The HostnameSources is a priority list of sources from which hostname will be fetched.
	// In case of the error in fetching hostname from source,
	// the next source from the list will be considered.(**default**: `["dns", "os"]`)
	HostnameSources []string `mapstructure:"hostname_sources"`

	ResourceAttributes metadata.ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

// Validate config
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func CreateDefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }
