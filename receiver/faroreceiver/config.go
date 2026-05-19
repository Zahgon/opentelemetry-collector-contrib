// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package faroreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/faroreceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/confmap"
)

type Config struct {
	confighttp.ServerConfig `mapstructure:",squash"`
}

var (
	_ component.Config    = (*Config)(nil)
	_ confmap.Unmarshaler = (*Config)(nil)
)

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (cfg *Config) Unmarshal(conf *confmap.Conf) error { _ = "STUB: not implemented"; return nil }
