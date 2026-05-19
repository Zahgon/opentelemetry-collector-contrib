// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package batchperresourceattr // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr"

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var separator = string([]byte{0x0, 0x1})

// Option configures a batch consumer created by this package.
type Option func(*options)

type options struct {
	injectMetadata bool
}

// WithMetadataInjection enables injecting the batched resource attribute
// values as client.Metadata into the context before forwarding each
// sub-batch to the next consumer. This allows downstream components (e.g.
// an exporterhelper batcher with Partition.MetadataKeys configured) to
// partition requests by the same keys used for batching.
func WithMetadataInjection() Option { _ = "STUB: not implemented"; return *new(Option) }

// injectAttrMetadata returns a new context with client.Metadata populated
// from the given attribute map for the specified keys. Keys absent from the
// map are omitted. If no keys are found the original context is returned.
func injectAttrMetadata(ctx context.Context, attrs pcommon.Map, attrKeys []string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type batchTraces struct {
	attrKeys       []string
	injectMetadata bool
	next           consumer.Traces
}

func NewBatchPerResourceTraces(attrKey string, next consumer.Traces, opts ...Option) consumer.Traces {
	_ = "STUB: not implemented"
	return *new(consumer.Traces)
}

func NewMultiBatchPerResourceTraces(attrKeys []string, next consumer.Traces, opts ...Option) consumer.Traces {
	_ = "STUB: not implemented"
	return *new(consumer.Traces)
}

// Capabilities returns the capabilities of the next consumer because batchTraces doesn't mutate data itself.
func (bt *batchTraces) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (bt *batchTraces) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// If zero or one resource spans just call next.

// If there is a single attribute value, then call next.

// Build the resource spans for each attribute value using CopyTo and call next for each one.

type batchMetrics struct {
	attrKeys       []string
	injectMetadata bool
	next           consumer.Metrics
}

func NewBatchPerResourceMetrics(attrKey string, next consumer.Metrics, opts ...Option) consumer.Metrics {
	_ = "STUB: not implemented"
	return *new(consumer.Metrics)
}

func NewMultiBatchPerResourceMetrics(attrKeys []string, next consumer.Metrics, opts ...Option) consumer.Metrics {
	_ = "STUB: not implemented"
	return *new(consumer.Metrics)
}

// Capabilities returns the capabilities of the next consumer because batchMetrics doesn't mutate data itself.
func (bt *batchMetrics) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (bt *batchMetrics) ConsumeMetrics(ctx context.Context, td pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// If zero or one resource metrics just call next.

// If there is a single attribute value, then call next.

// Build the resource metrics for each attribute value using CopyTo and call next for each one.

type batchLogs struct {
	attrKeys       []string
	injectMetadata bool
	next           consumer.Logs
}

func NewBatchPerResourceLogs(attrKey string, next consumer.Logs, opts ...Option) consumer.Logs {
	_ = "STUB: not implemented"
	return *new(consumer.Logs)
}

func NewMultiBatchPerResourceLogs(attrKeys []string, next consumer.Logs, opts ...Option) consumer.Logs {
	_ = "STUB: not implemented"
	return *new(consumer.Logs)
}

// Capabilities returns the capabilities of the next consumer because batchLogs doesn't mutate data itself.
func (bt *batchLogs) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (bt *batchLogs) ConsumeLogs(ctx context.Context, td plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// If zero or one resource logs just call next.

// If there is a single attribute value, then call next.

// Build the resource logs for each attribute value using CopyTo and call next for each one.
