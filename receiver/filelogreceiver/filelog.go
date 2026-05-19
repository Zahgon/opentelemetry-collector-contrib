// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filelogreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/file"
)

// NewFactory creates a factory for file_log receiver
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// ReceiverType implements stanza.LogReceiverType
// to create a file tailing receiver
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

func createDefaultConfig() *FileLogConfig { _ = "STUB: not implemented"; return nil }

// BaseConfig gets the base config from config, for now
func (ReceiverType) BaseConfig(cfg component.Config) adapter.BaseConfig {
	_ = "STUB: not implemented"
	return *new(adapter.BaseConfig)
}

// FileLogConfig defines configuration for the file_log receiver
type FileLogConfig struct {
	InputConfig        file.Config `mapstructure:",squash"`
	adapter.BaseConfig `mapstructure:",squash"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// InputConfig unmarshals the input operator
func (ReceiverType) InputConfig(cfg component.Config) operator.Config {
	_ = "STUB: not implemented"
	return *new(operator.Config)
}
