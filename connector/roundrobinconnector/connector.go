// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package roundrobinconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/roundrobinconnector"

import (
	"context"
	"sync/atomic"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pipeline"
)

func allConsumers[T any](r router[T]) ([]T, error) { _ = "STUB: not implemented"; return nil, nil }

type router[T any] interface {
	PipelineIDs() []pipeline.ID
	Consumer(pipelineIDs ...pipeline.ID) (T, error)
}

func newLogs(nextConsumer consumer.Logs) (connector.Logs, error) {
	_ = "STUB: not implemented"
	return *new(connector.Logs), nil
}

func newMetrics(nextConsumer consumer.Metrics) (connector.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(connector.Metrics), nil
}

func newTraces(nextConsumer consumer.Traces) (connector.Traces, error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}

// roundRobin is used to pass signals directly from one pipeline to one of the configured once in a round-robin mode.
// This is useful when there is a need to scale (shard) data processing and downstream components do not
// handle concurrent requests very well.
type roundRobin struct {
	component.StartFunc
	component.ShutdownFunc
	nextConsumer atomic.Uint64
	nextMetrics  []consumer.Metrics
	nextLogs     []consumer.Logs
	nextTraces   []consumer.Traces
}

func (*roundRobin) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (rr *roundRobin) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (rr *roundRobin) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (rr *roundRobin) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}
