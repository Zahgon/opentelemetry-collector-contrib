// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package failoverconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/failoverconnector"
import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

type logsRouter struct {
	*baseFailoverRouter[consumer.Logs]
}

func newLogsRouter(provider consumerProvider[consumer.Logs], cfg *Config) (*logsRouter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Consume is the logs-specific consumption method
func (f *logsRouter) Consume(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// consumeByHealthyPipeline will consume the logs by the current healthy level
func (f *logsRouter) consumeByHealthyPipeline(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// sampleRetryConsumers iterates through all unhealthy consumers to re-establish a healthy connection
func (f *logsRouter) sampleRetryConsumers(ctx context.Context, ld plog.Logs) bool {
	_ = "STUB: not implemented"
	return false
}

type logsFailover struct {
	component.StartFunc
	component.ShutdownFunc

	config   *Config
	failover *logsRouter
	logger   *zap.Logger
}

func (*logsFailover) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

// ConsumeLogs will try to export to the current set priority level and handle failover in the case of an error
func (f *logsFailover) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *logsFailover) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func newLogsToLogs(set connector.Settings, cfg component.Config, logs consumer.Logs) (connector.Logs, error) {
	_ = "STUB: not implemented"
	return *new(connector.Logs), nil
}
