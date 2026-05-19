// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package failoverconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/failoverconnector"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type tracesRouter struct {
	*baseFailoverRouter[consumer.Traces]
}

func newTracesRouter(provider consumerProvider[consumer.Traces], cfg *Config) (*tracesRouter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Consume is the traces-specific consumption method
func (f *tracesRouter) Consume(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// consumeByHealthyPipeline will consume the traces by the current healthy level
func (f *tracesRouter) consumeByHealthyPipeline(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// sampleRetryConsumers iterates through all unhealthy consumers to re-establish a healthy connection
func (f *tracesRouter) sampleRetryConsumers(ctx context.Context, td ptrace.Traces) bool {
	_ = "STUB: not implemented"
	return false
}

type tracesFailover struct {
	component.StartFunc
	component.ShutdownFunc

	config   *Config
	failover *tracesRouter
	logger   *zap.Logger
}

func (*tracesFailover) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

// ConsumeTraces will try to export to the current set priority level and handle failover in the case of an error
func (f *tracesFailover) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *tracesFailover) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func newTracesToTraces(set connector.Settings, cfg component.Config, traces consumer.Traces) (connector.Traces, error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}
