// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cloudflarereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudflarereceiver"

import (
	"errors"

	"go.opentelemetry.io/collector/config/configtls"
)

// Config holds all the parameters to start an HTTP server that can be sent logs from CloudFlare
type Config struct {
	Logs LogsConfig `mapstructure:"logs"`

	// prevent unkeyed literal initialization
	_ struct{}
}

type LogsConfig struct {
	Secret          string                  `mapstructure:"secret"`
	Endpoint        string                  `mapstructure:"endpoint"`
	TLS             *configtls.ServerConfig `mapstructure:"tls"`
	Attributes      map[string]string       `mapstructure:"attributes"`
	TimestampField  string                  `mapstructure:"timestamp_field"`
	TimestampFormat string                  `mapstructure:"timestamp_format"`
	Separator       string                  `mapstructure:"separator"`
	// MaxRequestBodySize sets the maximum request body size in bytes. 0 means no limit.
	// Default: 20MB
	MaxRequestBodySize int64 `mapstructure:"max_request_body_size,omitempty"`

	// prevent unkeyed literal initialization
	_ struct{}
}

var (
	errNoEndpoint = errors.New("an endpoint must be specified")
	errNoCert     = errors.New("tls was configured, but no cert file was specified")
	errNoKey      = errors.New("tls was configured, but no key file was specified")

	defaultTimestampField  = "EdgeStartTimestamp"
	defaultTimestampFormat = "rfc3339"
	defaultSeparator       = "."
)

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// 20MB default

// Validate timestamp_format if provided

// Missing key

// Missing cert
