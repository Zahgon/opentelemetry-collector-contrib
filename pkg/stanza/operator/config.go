// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package operator // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
)

// Config is the configuration of an operator
type Config struct {
	Builder
}

// NewConfig wraps the builder interface in a concrete struct
func NewConfig(b Builder) Config {
	_ = "STUB: not implemented"
	return *

	// Builder is an entity that can build a single operator
	new(Config)
}

type Builder interface {
	ID() string
	Type() string
	Build(component.TelemetrySettings) (Operator, error)
	SetID(string)
}

// UnmarshalJSON will unmarshal a config from JSON.
func (c *Config) UnmarshalJSON(bytes []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalYAML will unmarshal a config from YAML.
func (c *Config) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) Unmarshal(component *confmap.Conf) error { _ = "STUB: not implemented"; return nil }
