// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package skywalkingreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/skywalkingreceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/confmap"
)

const (
	// The config field id to load the protocol map from
	protocolsFieldName = "protocols"
)

// Protocols is the configuration for the supported protocols.
type Protocols struct {
	GRPC *configgrpc.ServerConfig `mapstructure:"grpc"`
	HTTP *confighttp.ServerConfig `mapstructure:"http"`
}

// Config defines configuration for skywalking receiver.
type Config struct {
	Protocols `mapstructure:"protocols"`
	// prevent unkeyed literal initialization
	_ struct{}
}

var (
	_ component.Config    = (*Config)(nil)
	_ confmap.Unmarshaler = (*Config)(nil)
)

// Validate checks the receiver configuration is valid
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// Unmarshal a config.Parser into the config struct.
func (cfg *Config) Unmarshal(componentParser *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalExact will not set struct properties to nil even if no key is provided,
// so set the protocol structs to nil where the keys were omitted.
