// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package routingconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

type logsConnector struct {
	component.StartFunc
	component.ShutdownFunc

	logger *zap.Logger
	config *Config
	router *router[consumer.Logs]
}

func newLogsConnector(
	set connector.Settings,
	config component.Config,
	logs consumer.Logs,
) (*logsConnector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*logsConnector) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (c *logsConnector) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// all logs are routed

// If error during statement evaluation consider it as not a match.

// If error during statement evaluation consider it as not a match.

// If error during statement evaluation consider it as not a match.

// If error during statement evaluation consider it as not a match.

// anything left wasn't matched by any route. Send to default consumer

func groupAllLogs(
	groups map[consumer.Logs]plog.Logs,
	cons consumer.Logs,
	logs plog.Logs,
) {
	_ = "STUB: not implemented"
	return
}
