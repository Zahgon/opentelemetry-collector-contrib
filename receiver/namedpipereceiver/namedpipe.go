// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package namedpipereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/namedpipereceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/namedpipe"
)

// NewFactory creates a factory for the named_pipe receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

type ReceiverType struct{}

// Type is the receiver type
func (ReceiverType) Type() component.Type {
	_ = "STUB: not implemented"
	return *

	// CreateDefaultConfig creates a config with type and version
	new(component.Type)
}

func (ReceiverType) CreateDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createDefaultConfig() *NamedPipeConfig { _ = "STUB: not implemented"; return nil }

// BaseConfig gets the base config from config, for now
func (ReceiverType) BaseConfig(cfg component.Config) adapter.BaseConfig {
	_ = "STUB: not implemented"
	return *new(adapter.BaseConfig)
}

// NamedPipeConfig defines configuration for the named_pipe receiver
type NamedPipeConfig struct {
	InputConfig        namedpipe.Config `mapstructure:",squash"`
	adapter.BaseConfig `mapstructure:",squash"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// InputConfig unmarshals the input operator
func (ReceiverType) InputConfig(cfg component.Config) operator.Config {
	_ = "STUB: not implemented"
	return *new(operator.Config)
}
