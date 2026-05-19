// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package textencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/textencodingextension"

type Config struct {
	Encoding              string `mapstructure:"encoding"`
	MarshalingSeparator   string `mapstructure:"marshaling_separator"`
	UnmarshalingSeparator string `mapstructure:"unmarshaling_separator"`
	// prevent unkeyed literal initialization
	_ struct{}
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
