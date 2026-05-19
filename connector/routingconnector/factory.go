// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package routingconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector"

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
	return *new(component.Config)
}

// createTracesToTraces creates a traces to traces connector based on provided config.
func createTracesToTraces(
	_ context.Context,
	set connector.Settings,
	cfg component.Config,
	traces consumer.Traces,
) (connector.Traces, error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}

// createMetricsToMetrics creates a metrics to metrics connector based on provided config.
func createMetricsToMetrics(
	_ context.Context,
	set connector.Settings,
	cfg component.Config,
	metrics consumer.Metrics,
) (connector.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(connector.Metrics), nil
}

// createLogsToLogs creates a logs to logs connector based on provided config.
func createLogsToLogs(
	_ context.Context,
	set connector.Settings,
	cfg component.Config,
	logs consumer.Logs,
) (connector.Logs, error) {
	_ = "STUB: not implemented"
	return *new(connector.Logs), nil
}
