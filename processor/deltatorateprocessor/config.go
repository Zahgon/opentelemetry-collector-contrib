// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package deltatorateprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatorateprocessor"

// Config defines the configuration for the processor.
type Config struct {
	// List of delta sum metrics to convert to rates
	Metrics []string `mapstructure:"metrics"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// Validate checks whether the input configuration has all of the required fields for the processor.
// An error is returned if there are any invalid inputs.
func (config *Config) Validate() error { _ = "STUB: not implemented"; return nil }
