// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package journaldreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/journaldreceiver"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/journald"
)

// createDefaultConfig creates a config with type and version
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// receiverType implements adapter.LogReceiverType
// to create a journald receiver
type receiverType struct{}

// Type is the receiver type
func (receiverType) Type() component.Type {
	_ = "STUB: not implemented"
	return *

	// BaseConfig gets the base config from config, for now
	new(component.Type)
}

func (receiverType) BaseConfig(cfg component.Config) adapter.BaseConfig {
	_ = "STUB: not implemented"
	return *new(adapter.BaseConfig)
}

// JournaldConfig defines configuration for the journald receiver
type JournaldConfig struct {
	adapter.BaseConfig `mapstructure:",squash"`
	InputConfig        journald.Config `mapstructure:",squash"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// InputConfig unmarshals the input operator
func (receiverType) InputConfig(cfg component.Config) operator.Config {
	_ = "STUB: not implemented"
	return *new(operator.Config)
}

// CreateDefaultConfig creates a config with type and version
func (receiverType) CreateDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}
