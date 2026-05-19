// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package collectdreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/collectdreceiver"

import (
	"time"

	"go.opentelemetry.io/collector/config/confighttp"
)

// Config defines configuration for Collectd receiver.
type Config struct {
	confighttp.ServerConfig `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct
	Timeout                 time.Duration            `mapstructure:"timeout"`
	Encoding                string                   `mapstructure:"encoding"`
	AttributesPrefix        string                   `mapstructure:"attributes_prefix"`
}

func (c *Config) Validate() error {
	_ = "STUB: not implemented"
	// CollectD receiver only supports JSON encoding. We expose a config option
	// to make it explicit and obvious to the users.
	return nil
}
