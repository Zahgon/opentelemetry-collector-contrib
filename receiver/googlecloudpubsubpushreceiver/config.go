// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudpubsubpushreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubpushreceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
)

var _ component.Config = (*Config)(nil)

// Config defines configuration for the receiver.
type Config struct {
	confighttp.ServerConfig `mapstructure:",squash"`

	// Endpoint identifies the expected encoding of messages
	// received from Pub/Sub.
	Encoding *component.ID `mapstructure:"encoding"`
}

const defaultEndpoint = "0.0.0.0:8080"

// createDefaultConfig creates the default configuration for the receiver.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// Validate checks if the receiver configuration is valid.
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
