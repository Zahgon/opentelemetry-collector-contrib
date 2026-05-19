// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlpjsonconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/otlpjsonconnector"

import (
	"context"
	"regexp"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
)

var (
	logRegex    = regexp.MustCompile(`^\{\s*"resourceLogs"\s*:\s*\[`)
	metricRegex = regexp.MustCompile(`^\{\s*"resourceMetrics"\s*:\s*\[`)
	traceRegex  = regexp.MustCompile(`^\{\s*"resourceSpans"\s*:\s*\[`)
)

// NewFactory returns a ConnectorFactory.
func NewFactory() connector.Factory { _ = "STUB: not implemented"; return *new(connector.Factory) }

// createDefaultConfig creates the default configuration.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"

	// createLogsConnector returns a connector which consume logs and export logs
	return *new(component.Config)
}

func createLogsConnector(
	_ context.Context,
	set connector.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (connector.Logs, error) {
	_ = "STUB: not implemented"
	return *new(connector.Logs), nil
}

// createTracesConnector returns a connector which consume logs and export traces
func createTracesConnector(
	_ context.Context,
	set connector.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (connector.Logs, error) {
	_ = "STUB: not implemented"
	return *new(connector.Logs), nil
}

// createMetricsConnector returns a connector which consume logs and export metrics
func createMetricsConnector(
	_ context.Context,
	set connector.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (connector.Logs, error) {
	_ = "STUB: not implemented"
	return *new(connector.Logs), nil
}
