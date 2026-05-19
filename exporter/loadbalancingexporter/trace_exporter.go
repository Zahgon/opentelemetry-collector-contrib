// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loadbalancingexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter/internal/metadata"
)

const (
	pseudoAttrSpanName = "span.name"
	pseudoAttrSpanKind = "span.kind"
)

var _ exporter.Traces = (*traceExporterImp)(nil)

type exporterTraces map[*wrappedExporter]ptrace.Traces

type traceExporterImp struct {
	loadBalancer *loadBalancer
	routingKey   routingKey
	routingAttrs []string

	logger     *zap.Logger
	stopped    bool
	shutdownWg sync.WaitGroup
	telemetry  *metadata.TelemetryBuilder
}

// Create new traces exporter
func newTracesExporter(params exporter.Settings, cfg component.Config) (*traceExporterImp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*traceExporterImp) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (e *traceExporterImp) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *traceExporterImp) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *traceExporterImp) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// routingIdentifiersFromTraces reads the traces and determines an identifier that can be used to define a position on the
// ring hash. It takes the routingKey, defining what type of routing should be used, and a series of attributes
// (optionally) used if the routingKey is attrRouting.
//
// only svcRouting and attrRouting are supported. For attrRouting, any attribute, as well the "pseudo" attributes span.name
// and span.kind are supported.
//
// In practice, makes the assumption that ptrace.Traces includes only one trace of each kind, in the "trace tree".
func routingIdentifiersFromTraces(td ptrace.Traces, rType routingKey, attrs []string) (map[string]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Determine how the key should be populated.

// The simple case is the TraceID routing. In this case, we just use the string representation of the Trace ID.

// Service Name is still handled as an "attribute router", it's just expressed as a "pseudo attribute"

// By default, we'll just use the input provided.

// Composite the attributes together as a key.

// rKey will never return an error. See
// 1. https://pkg.go.dev/bytes#Buffer.Write
// 2. https://stackoverflow.com/a/70388629

// resource spans

// ils or "instrumentation library spans"

// the lowest level span (or what engineers regularly interact with)

// No matter what, there will be a key here (even if that key is "").
