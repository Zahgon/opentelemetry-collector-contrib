// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cloudfoundryreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudfoundryreceiver"

import (
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configopaque"
)

type RLPGatewayConfig struct {
	confighttp.ClientConfig `mapstructure:",squash"`
	ShardID                 string `mapstructure:"shard_id"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// LimitedTLSClientSetting is a subset of TLSClientSetting, see LimitedClientConfig for more details
type LimitedTLSClientSetting struct {
	InsecureSkipVerify bool `mapstructure:"insecure_skip_verify"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// LimitedClientConfig is a subset of ClientConfig, implemented as a separate type due to the library this
// configuration is used with not taking a preconfigured http.Client as input, but only taking these specific options
type LimitedClientConfig struct {
	Endpoint string                  `mapstructure:"endpoint"`
	TLS      LimitedTLSClientSetting `mapstructure:"tls"`

	// prevent unkeyed literal initialization
	_ struct{}
}

type UAAConfig struct {
	LimitedClientConfig `mapstructure:",squash"`
	Username            string              `mapstructure:"username"`
	Password            configopaque.String `mapstructure:"password"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// Config defines configuration for Collectd receiver.
type Config struct {
	RLPGateway RLPGatewayConfig `mapstructure:"rlp_gateway"`
	UAA        UAAConfig        `mapstructure:"uaa"`

	// prevent unkeyed literal initialization
	_ struct{}
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func validateURLOption(name, value string) error { _ = "STUB: not implemented"; return nil }
