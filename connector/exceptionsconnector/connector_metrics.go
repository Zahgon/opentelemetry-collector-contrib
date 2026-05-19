// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exceptionsconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/exceptionsconnector"

import (
	"bytes"
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/pdatautil"
)

const (
	metricKeySeparator = string(byte(0))
)

type metricsConnector struct {
	lock   sync.Mutex
	config Config

	// Additional dimensions to add to metrics.
	dimensions []pdatautil.Dimension

	keyBuf *bytes.Buffer

	metricsConsumer consumer.Metrics
	component.StartFunc
	component.ShutdownFunc

	exceptions map[string]*exception

	logger *zap.Logger

	// The starting time of the data points.
	startTimestamp pcommon.Timestamp
}

type exception struct {
	count     int
	attrs     pcommon.Map
	exemplars pmetric.ExemplarSlice
}

func newMetricsConnector(logger *zap.Logger, config component.Config) *metricsConnector {
	_ = "STUB: not implemented"
	return nil
}

// Capabilities implements the consumer interface.
func (*metricsConnector) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

// ConsumeTraces implements the consumer.Traces interface.
// It aggregates the trace data to generate metrics.
func (c *metricsConnector) ConsumeTraces(ctx context.Context, traces ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *metricsConnector) exportMetrics(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// collectExceptions collects the exception metrics data and writes it into the metrics object.
func (c *metricsConnector) collectExceptions(ilm pmetric.ScopeMetrics) error {
	_ = "STUB: not implemented"
	return nil
}

// Reset the exemplars for the next batch of spans.

func (c *metricsConnector) addException(excKey string, attrs pcommon.Map) *exception {
	_ = "STUB: not implemented"
	return nil
}

func (c *metricsConnector) addExemplar(exc *exception, traceID pcommon.TraceID, spanID pcommon.SpanID) {
	_ = "STUB: not implemented"
	return
}

func buildDimensionKVs(dimensions []pdatautil.Dimension, serviceName string, span ptrace.Span, eventAttrs, resourceAttrs pcommon.Map) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

// buildKey builds the metric key from the service name and span metadata such as kind, status_code and
// will attempt to add any additional dimensions the user has configured that match the span's attributes
// or resource attributes. If the dimension exists in both, the span's attributes, being the most specific, takes precedence.
//
// The metric key is a simple concatenation of dimension values, delimited by a null character.
func buildKey(dest *bytes.Buffer, serviceName string, span ptrace.Span, optionalDims []pdatautil.Dimension, eventAttrs, resourceAttrs pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func concatDimensionValue(dest *bytes.Buffer, value string, prefixSep bool) {
	_ = "STUB: not implemented"
	return
}
