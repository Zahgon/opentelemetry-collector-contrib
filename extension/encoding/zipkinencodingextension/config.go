// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/zipkinencodingextension"

import (
	"go.opentelemetry.io/collector/confmap/xconfmap"
)

var _ xconfmap.Validator = (*Config)(nil)

type Config struct {
	Protocol string `mapstructure:"protocol"`
	Version  string `mapstructure:"version"`
	// prevent unkeyed literal initialization
	_ struct{}
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
