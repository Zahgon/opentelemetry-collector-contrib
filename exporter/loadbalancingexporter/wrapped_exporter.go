// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loadbalancingexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/otel/attribute"
)

// wrappedExporter is an exporter that waits for the data processing to complete before shutting down.
// consumeWG has to be incremented explicitly by the consumer of the wrapped exporter.
type wrappedExporter struct {
	component.Component
	consumeWG sync.WaitGroup

	// we store the attributes here for both cases, to avoid new allocations on the hot path
	endpointAttr attribute.Set
	successAttr  attribute.Set
	failureAttr  attribute.Set
}

func newWrappedExporter(exp component.Component, identifier string) *wrappedExporter {
	_ = "STUB: not implemented"
	return nil
}

func (we *wrappedExporter) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (we *wrappedExporter) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (we *wrappedExporter) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (we *wrappedExporter) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}
