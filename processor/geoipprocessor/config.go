// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package geoipprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/otel/attribute"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor/internal/provider"
)

const (
	providersKey = "providers"
)

type ContextID string

const (
	resource ContextID = "resource"
	record   ContextID = "record"
)

func (c *ContextID) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// Config holds the configuration for the GeoIP processor.
type Config struct {
	// Providers specifies the sources to extract geographical information about a given IP.
	Providers map[string]provider.Config `mapstructure:"-"`

	// Context section allows specifying the source type to look for the IP. Available options: resource or record.
	Context ContextID `mapstructure:"context"`

	// An array of attribute names, which are used for the IP address lookup
	Attributes []attribute.Key `mapstructure:"attributes"`
}

var (
	_ component.Config    = (*Config)(nil)
	_ confmap.Unmarshaler = (*Config)(nil)
)

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// validate all provider's configuration

// Unmarshal a config.Parser into the config struct.
func (cfg *Config) Unmarshal(componentParser *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil
}

// load the non-dynamic config normally

// dynamically load the individual providers configs based on the key name

// retrieve `providers` configuration section

// loop through all defined providers and load their configuration
