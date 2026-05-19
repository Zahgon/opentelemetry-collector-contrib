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

type connectorLogs struct {
	config       Config
	logsConsumer consumer.Logs
	logger       *zap.Logger

	component.StartFunc
	component.ShutdownFunc
}

// newLogsConnector is a function to create a new connector for logs extraction
func newLogsConnector(set connector.Settings, config component.Config, logsConsumer consumer.Logs) *connectorLogs {
	_ = "STUB: not implemented"
	return nil
}

// Capabilities implements the consumer interface
func (*connectorLogs) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

// ConsumeLogs method is called for each instance of a log sent to the connector
func (c *connectorLogs) ConsumeLogs(ctx context.Context, pl plog.Logs) error {
	_ = "STUB: not implemented"
	// loop through the levels of logs
	return nil
}

// Check if the "resourceLogs" key exists in the JSON data

// If it's a metric or trace payload, simply continue

// If no regex matches, log the invalid payload
