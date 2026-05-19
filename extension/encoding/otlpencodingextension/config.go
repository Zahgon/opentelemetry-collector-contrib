// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlpencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/otlpencodingextension"
import (
	"go.opentelemetry.io/collector/confmap/xconfmap"
)

var _ xconfmap.Validator = (*Config)(nil)

type Config struct {
	Protocol string `mapstructure:"protocol"`
	// prevent unkeyed literal initialization
	_ struct{}
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
