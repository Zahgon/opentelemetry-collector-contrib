// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureeventhubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pipeline"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent"
)

var errUnexpectedConfigurationType = errors.New("failed to cast configuration to azure event hub config")

type eventhubReceiverFactory struct {
	receivers *sharedcomponent.SharedComponents
}

// NewFactory creates a factory for the Azure Event Hub receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func (f *eventhubReceiverFactory) createLogsReceiver(
	_ context.Context,
	settings receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

func (f *eventhubReceiverFactory) createMetricsReceiver(
	_ context.Context,
	settings receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func (f *eventhubReceiverFactory) createTracesReceiver(
	_ context.Context,
	settings receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Traces), nil
}

func (f *eventhubReceiverFactory) getReceiver(
	signal pipeline.Signal,
	cfg component.Config,
	settings receiver.Settings,
) (component.Component, error) {
	_ = "STUB: not implemented"
	return *new(component.Component), nil
}
