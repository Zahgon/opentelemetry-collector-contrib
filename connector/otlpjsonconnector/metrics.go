// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlpjsonconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/otlpjsonconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

type connectorMetrics struct {
	config          Config
	metricsConsumer consumer.Metrics
	logger          *zap.Logger

	component.StartFunc
	component.ShutdownFunc
}

// newMetricsConnector is a function to create a new connector for metrics extraction
func newMetricsConnector(set connector.Settings, config component.Config, metricsConsumer consumer.Metrics) *connectorMetrics {
	_ = "STUB: not implemented"
	return nil
}

// Capabilities implements the consumer interface.
func (*connectorMetrics) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

// ConsumeLogs method is called for each instance of a log sent to the connector
func (c *connectorMetrics) ConsumeLogs(ctx context.Context, pl plog.Logs) error {
	_ = "STUB: not implemented"
	// loop through the levels of logs
	return nil
}

// If it's a log or trace payload, simply continue

// If no regex matches, log the invalid payload
