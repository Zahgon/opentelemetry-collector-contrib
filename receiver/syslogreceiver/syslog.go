// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package syslogreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/syslogreceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/syslog"
)

// NewFactory creates a factory for syslog receiver
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// ReceiverType implements adapter.LogReceiverType
// to create a syslog receiver
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

// BaseConfig gets the base config from config, for now
func (ReceiverType) BaseConfig(cfg component.Config) adapter.BaseConfig {
	_ = "STUB: not implemented"
	return *new(adapter.BaseConfig)
}

// SysLogConfig defines configuration for the syslog receiver
type SysLogConfig struct {
	InputConfig        syslog.Config `mapstructure:",squash"`
	adapter.BaseConfig `mapstructure:",squash"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// InputConfig unmarshals the input operator
func (ReceiverType) InputConfig(cfg component.Config) operator.Config {
	_ = "STUB: not implemented"
	return *new(operator.Config)
}

func (cfg *SysLogConfig) Unmarshal(componentParser *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil

	// Nothing to do if there is no config given.
}
