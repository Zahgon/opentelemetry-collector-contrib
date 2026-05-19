// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package roundrobinconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/roundrobinconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
)

// NewFactory returns a ConnectorFactory.
func NewFactory() connector.Factory { _ = "STUB: not implemented"; return *new(connector.Factory) }

// createDefaultConfig creates the default configuration.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"

	// createLogsToLogs creates a log receiver based on provided config.
	return *new(component.Config)
}

func createLogsToLogs(
	_ context.Context,
	_ connector.Settings,
	_ component.Config,
	nextConsumer consumer.Logs,
) (connector.Logs, error) {
	_ = "STUB: not implemented"
	return *new(connector.Logs), nil
}

// createMetricsToMetrics creates a metrics receiver based on provided config.
func createMetricsToMetrics(
	_ context.Context,
	_ connector.Settings,
	_ component.Config,
	nextConsumer consumer.Metrics,
) (connector.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(connector.Metrics), nil
}

// createTracesToTraces creates a trace receiver based on provided config.
func createTracesToTraces(
	_ context.Context,
	_ connector.Settings,
	_ component.Config,
	nextConsumer consumer.Traces,
) (connector.Traces, error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}
