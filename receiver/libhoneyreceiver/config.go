// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package libhoneyreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver"

import (
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/confmap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver/internal/libhoneyevent"
)

// Config represents the receiver config settings within the collector's config.yaml
type Config struct {
	HTTP           configoptional.Optional[HTTPConfig] `mapstructure:"http"`
	AuthAPI        string                              `mapstructure:"auth_api"`
	Wrapper        string                              `mapstructure:"wrapper"`
	FieldMapConfig libhoneyevent.FieldMapConfig        `mapstructure:"fields"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// HTTPConfig defines the configuration for the HTTP server receiving traces.
type HTTPConfig struct {
	confighttp.ServerConfig `mapstructure:",squash"`

	// The URL path to receive traces on. If omitted "/" will be used.
	TracesURLPaths []string `mapstructure:"traces_url_paths,omitempty"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// Validate ensures the HTTP configuration is set.
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// Unmarshal unmarshals the configuration from the given configuration and then checks for errors.
func (cfg *Config) Unmarshal(conf *confmap.Conf) error {
	_ = "STUB: not implemented"
	// first load the config normally
	return nil
}

func sanitizeURLPath(urlPath string) (string, error) { _ = "STUB: not implemented"; return "", nil }
