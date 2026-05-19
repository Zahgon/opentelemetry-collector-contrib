// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package apmstats // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/apmstats"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type traceToTraceConnector struct {
	logger         *zap.Logger
	tracesConsumer consumer.Traces // the next component in the pipeline to ingest traces after connector
}

func newTraceToTraceConnector(logger *zap.Logger, nextConsumer consumer.Traces) *traceToTraceConnector {
	_ = "STUB: not implemented"
	return nil
}

// Start implements the component interface.
func (*traceToTraceConnector) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"

	// Shutdown implements the component interface.
	return nil
}

func (*traceToTraceConnector) Shutdown(context.Context) error {
	_ = "STUB: not implemented"

	// Capabilities implements the consumer interface.
	// tells use whether the component(connector) will mutate the data passed into it. if set to true the connector does modify the data
	return nil
}

func (*traceToTraceConnector) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

// ConsumeTraces implements the consumer interface.
func (c *traceToTraceConnector) ConsumeTraces(ctx context.Context, traces ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}
