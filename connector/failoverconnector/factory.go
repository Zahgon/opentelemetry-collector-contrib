// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package failoverconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/failoverconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
)

func NewFactory() connector.Factory { _ = "STUB: not implemented"; return *new(connector.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesToTraces(
	ctx context.Context,
	set connector.Settings,
	cfg component.Config,
	traces consumer.Traces,
) (connector.Traces, error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}

// If queue is disabled, return the raw failover connector

// If queue is enabled, wrap with exporterhelper

// Return testable wrapper that exposes internal failover router

func createMetricsToMetrics(
	ctx context.Context,
	set connector.Settings,
	cfg component.Config,
	metrics consumer.Metrics,
) (connector.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(connector.Metrics), nil
}

// If queue is disabled, return the raw failover connector directly (original behavior)

// If queue is enabled, wrap with exporterhelper

// Return testable wrapper that exposes internal failover router

func createLogsToLogs(
	ctx context.Context,
	set connector.Settings,
	cfg component.Config,
	logs consumer.Logs,
) (connector.Logs, error) {
	_ = "STUB: not implemented"
	return *new(connector.Logs), nil
}

// If queue is disabled, return the raw failover connector directly (original behavior)

// If queue is enabled, wrap with exporterhelper

// Return testable wrapper that exposes internal failover router
