// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exceptionsconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/exceptionsconnector"

import (
	"go.opentelemetry.io/collector/confmap/xconfmap"
)

// Dimension defines the dimension name and optional default value if the Dimension is missing from a span attribute.
type Dimension struct {
	Name    string  `mapstructure:"name"`
	Default *string `mapstructure:"default"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type Exemplars struct {
	Enabled bool `mapstructure:"enabled"`
	// prevent unkeyed literal initialization
	_ struct{}
}

// Config defines the configuration options for exceptionsconnector
type Config struct {
	// Dimensions defines the list of additional dimensions on top of the provided:
	// - service.name
	// - span.name
	// - span.kind
	// - status.code
	// The dimensions will be fetched from the span's attributes. Examples of some conventionally used attributes:
	// https://github.com/open-telemetry/opentelemetry-collector/blob/main/model/semconv/opentelemetry.go.
	Dimensions []Dimension `mapstructure:"dimensions"`
	// Exemplars defines the configuration for exemplars.
	Exemplars Exemplars `mapstructure:"exemplars"`
	// prevent unkeyed literal initialization
	_ struct{}
}

var _ xconfmap.Validator = (*Config)(nil)

// Validate checks if the connector configuration is valid
func (c Config) Validate() error { _ = "STUB: not implemented"; return nil }

// validateDimensions checks duplicates for reserved dimensions and additional dimensions.
func validateDimensions(dimensions []Dimension) error { _ = "STUB: not implemented"; return nil }
