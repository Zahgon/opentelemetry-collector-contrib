// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package windowseventlogreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowseventlogreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowseventlogreceiver/internal/discovery"
)

// getDomainControllersRemoteConfig is the function used to discover domain controllers.
// It is a variable to allow mocking in tests.
var getDomainControllersRemoteConfig = discovery.GetJoinedDomainControllersRemoteConfig

// createLogsReceiver creates a logs receiver with SID enrichment support
func createLogsReceiver(
	ctx context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

// Create SID cache if enabled

// Wrap the consumer with SID enrichment

// Create the underlying Stanza receiver with the enriched consumer

type multiLogsReceiver struct {
	receivers []receiver.Logs
}

func (m *multiLogsReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shut down already-started receivers before returning the error.

func (m *multiLogsReceiver) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// receiverType implements adapter.LogReceiverType
// to create a file tailing receiver
type receiverType struct{}

var _ adapter.LogReceiverType = (*receiverType)(nil)

// Type is the receiver type
func (receiverType) Type() component.Type {
	_ = "STUB: not implemented"
	return *

	// CreateDefaultConfig creates a config with type and version
	new(component.Type)
}

func (receiverType) CreateDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// BaseConfig gets the base config from config, for now
func (receiverType) BaseConfig(cfg component.Config) adapter.BaseConfig {
	_ = "STUB: not implemented"
	return *new(adapter.BaseConfig)
}

// InputConfig unmarshals the input operator
func (receiverType) InputConfig(cfg component.Config) operator.Config {
	_ = "STUB: not implemented"
	return *new(operator.Config)
}
