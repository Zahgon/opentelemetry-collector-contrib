// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package solarwindsapmsettingsextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/solarwindsapmsettingsextension"

import (
	"time"

	"go.opentelemetry.io/collector/component"
)

type Config struct {
	Endpoint string        `mapstructure:"endpoint"`
	Key      string        `mapstructure:"key"`
	Interval time.Duration `mapstructure:"interval"`
}

const (
	DefaultEndpoint = "apm.collector.na-01.cloud.solarwinds.com:443"
	DefaultInterval = time.Duration(10) * time.Second
	MinimumInterval = time.Duration(5) * time.Second
	MaximumInterval = time.Duration(60) * time.Second
)

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func (cfg *Config) Validate() error {
	_ = "STUB: not implemented"
	// Endpoint
	return nil
}

// Replaced by the default

// Key

/**
 * Service name is empty. We are trying our best effort to resolve the service name
 */

// Interval

func resolveServiceNameBestEffort() string { _ = "STUB: not implemented"; return "" }
