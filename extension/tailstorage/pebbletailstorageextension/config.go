// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pebbletailstorageextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/tailstorage/pebbletailstorageextension"

type Config struct {
	// Directory is where the extension stores Pebble DB files.
	Directory string `mapstructure:"directory"`
	// prevent unkeyed literal initialization
	_ struct{}
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
