// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureblobreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureblobreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent"
)

const (
	logsContainerName   = "logs"
	tracesContainerName = "traces"
	defaultCloud        = AzureCloudType
)

var errUnexpectedConfigurationType = errors.New("failed to cast configuration to Azure Blob Config")

type blobReceiverFactory struct {
	receivers *sharedcomponent.SharedComponents
}

// NewFactory returns a factory for Azure Blob receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func (*blobReceiverFactory) createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func (f *blobReceiverFactory) createLogsReceiver(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

func (f *blobReceiverFactory) createTracesReceiver(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Traces), nil
}

func (f *blobReceiverFactory) getReceiver(
	set receiver.Settings,
	cfg component.Config,
) (component.Component, error) {
	_ = "STUB: not implemented"
	return *new(component.Component), nil
}

func (*blobReceiverFactory) getEventHandler(cfg *Config, logger *zap.Logger) (eventHandler, error) {
	_ = "STUB: not implemented"
	return *new(eventHandler), nil
}

// If Event Hub is not configured, use the Blob Event Handler
