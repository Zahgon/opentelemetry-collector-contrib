// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package spanmetricsconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector"

import (
	"bytes"
	"context"
	"sync"
	"time"

	"github.com/hashicorp/golang-lru/v2/simplelru"
	"github.com/jonboulle/clockwork"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	conventions "go.opentelemetry.io/otel/semconv/v1.40.0"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector/internal/cache"
	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector/internal/metrics"
	utilattri "github.com/open-telemetry/opentelemetry-collector-contrib/internal/pdatautil"
)

const (
	serviceNameKey                 = string(conventions.ServiceNameKey)
	spanNameKey                    = "span.name"                          // OpenTelemetry non-standard constant.
	spanKindKey                    = "span.kind"                          // OpenTelemetry non-standard constant.
	statusCodeKey                  = "status.code"                        // OpenTelemetry non-standard constant.
	collectorInstanceKey           = "collector.instance.id"              // OpenTelemetry non-standard constant.
	otelStatusCodeKey              = "otel.status_code"                   // OpenTelemetry non-standard constant.
	instrumentationScopeNameKey    = "span.instrumentation.scope.name"    // OpenTelemetry non-standard constant.
	instrumentationScopeVersionKey = "span.instrumentation.scope.version" // OpenTelemetry non-standard constant.
	metricKeySeparator             = string(byte(0))

	defaultResourceMetricsCacheSize = 1000

	metricNameDuration = "duration"
	metricNameCalls    = "calls"
	metricNameEvents   = "events"

	metricAttrSamplingMethod = "sampling.method"

	defaultUnit = metrics.Seconds

	// https://github.com/open-telemetry/opentelemetry-go/blob/3ae002c3caf3e44387f0554dfcbbde2c5aab7909/sdk/metric/internal/aggregate/limit.go#L11C36-L11C50
	overflowKey = "otel.metric.overflow"

	defaultMaxPerDatapoint = 5
)

type connectorImp struct {
	lock   sync.Mutex
	logger *zap.Logger
	config Config

	metricsConsumer consumer.Metrics

	// Additional dimensions to add to metrics.
	dimensions []utilattri.Dimension

	resourceMetrics *cache.Cache[resourceKey, *resourceMetrics]

	resourceMetricsKeyAttributes map[string]struct{}

	keyBuf *bytes.Buffer

	clock   clockwork.Clock
	ticker  clockwork.Ticker
	done    chan struct{}
	started bool

	shutdownOnce sync.Once

	// Event dimensions to add to the events metric.
	eDimensions []utilattri.Dimension

	// Calls dimensions to add to the events metric.
	callsDimensions []utilattri.Dimension

	// duration dimensions to add to the events metric.
	durationDimensions []utilattri.Dimension

	events EventsConfig

	// Tracks the last TimestampUnixNano for delta metrics so that they represent an uninterrupted series. Unused for cumulative span metrics.
	lastDeltaTimestamps *simplelru.LRU[metrics.Key, pcommon.Timestamp]
	instanceID          string
}

type resourceMetrics struct {
	histograms metrics.HistogramMetrics
	sums       metrics.SumMetrics
	events     metrics.SumMetrics
	attributes pcommon.Map
	// lastSeen captures when the last data points for this resource were recorded.
	lastSeen time.Time
}

func newDimensions(cfgDims []Dimension) []utilattri.Dimension {
	_ = "STUB: not implemented"
	return nil
}

func newConnector(logger *zap.Logger, config component.Config, clock clockwork.Clock, instanceID string) (*connectorImp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func initHistogramMetrics(cfg Config) metrics.HistogramMetrics {
	_ = "STUB: not implemented"
	return *new(metrics.HistogramMetrics)
}

// unitDivider returns a unit divider to convert nanoseconds to milliseconds or seconds.
func unitDivider(u metrics.Unit) int64 { _ = "STUB: not implemented"; return 0 }

func durationsToUnits(vs []time.Duration, unitDivider int64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

// Start implements the component.Component interface.
func (p *connectorImp) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown implements the component.Component interface.
func (p *connectorImp) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

// Capabilities implements the consumer interface.
func (*connectorImp) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

// ConsumeTraces implements the consumer.Traces interface.
// It aggregates the trace data to generate metrics.
func (p *connectorImp) ConsumeTraces(_ context.Context, traces ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *connectorImp) exportMetrics(ctx context.Context) { _ = "STUB: not implemented"; return }

// This component no longer needs to read the metrics once built, so it is safe to unlock.

// buildMetrics collects the computed raw metrics data and builds OTLP metrics.
func (p *connectorImp) buildMetrics() pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

/**
 * To represent an uninterrupted stream of metrics as per the spec, the (StartTimestamp, Timestamp)'s of successive data points should be:
 * - For cumulative metrics: (T1, T2), (T1, T3), (T1, T4) ...
 * - For delta metrics: (T1, T2), (T2, T3), (T3, T4) ...
 */

// Collect lastDeltaTimestamps keys that need to be updated. Metrics can share the same key, so defer the update.

// For delta metrics, cache the current data point's timestamp, which will be the start timestamp for the next data points in the series

func (p *connectorImp) resetState() {
	_ = "STUB: not implemented"
	// If delta metrics, reset accumulated data
	return
}

// If none of these features are enabled then we can skip the remaining operations.
// Enabling either of these features requires to go over resource metrics and do operation on each.

// Exemplars are only relevant to this batch of traces, so must be cleared within the lock

// If metrics expiration is configured, remove metrics that haven't been seen for longer than the expiration period.

// aggregateMetrics aggregates the raw metrics from the input trace data.
//
// Metrics are grouped by resource attributes.
// Each metric is identified by a key that is built from the service name
// and span metadata such as name, kind, status_code and any additional
// dimensions the user has configured.
func (p *connectorImp) aggregateMetrics(traces ptrace.Traces) { _ = "STUB: not implemented"; return }

// Local cache for adjusted count - no synchronization needed.
// Consecutive spans from the same trace share identical tracestates.

// Protect against end timestamps before start timestamps. Assume 0 duration.

// aggregate sums metrics

// aggregate histogram metrics

// aggregate events metrics

// We cannot use event.Attributes().CopyTo(rscAdnEventAttrs) because it overrides the existing keys.

func (p *connectorImp) addExemplar(span ptrace.Span, duration float64, h metrics.Histogram) {
	_ = "STUB: not implemented"
	return
}

type resourceKey [16]byte

func (p *connectorImp) createResourceKey(attr pcommon.Map) resourceKey {
	_ = "STUB: not implemented"
	return *new(resourceKey)
}

func (p *connectorImp) getOrCreateResourceMetrics(attr pcommon.Map) *resourceMetrics {
	_ = "STUB: not implemented"
	return nil
}

// If expiration is enabled, track the last seen time.

func (p *connectorImp) buildAttributes(
	serviceName string,
	span ptrace.Span,
	resourceAttrs pcommon.Map,
	dimensions []utilattri.Dimension,
	instrumentationScope pcommon.InstrumentationScope,
	isAdjustedCount bool,
) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func addResourceAttributes(attrs *pcommon.Map, dimensions []utilattri.Dimension, span ptrace.Span, resourceAttrs pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func concatDimensionValue(dest *bytes.Buffer, value string, prefixSep bool) {
	_ = "STUB: not implemented"
	return
}

// buildKey builds the metric key from the service name and span metadata such as name, kind, status_code and
// will attempt to add any additional dimensions the user has configured that match the span's attributes
// or resource/event attributes. If the dimension exists in both, the span's attributes, being the most specific, takes precedence.
//
// When enable_metrics_sampling_method is true, the sampling method (extrapolated or counted) is included in the key
// so that spans with different tracestate produce distinct metrics.
//
// The metric key is a simple concatenation of dimension values, delimited by a null character.
func (p *connectorImp) buildKey(serviceName string, span ptrace.Span, optionalDims []utilattri.Dimension, resourceOrEventAttrs pcommon.Map, isAdjusted bool) metrics.Key {
	_ = "STUB: not implemented"
	return *new(metrics.Key)
}

// buildMetricName builds the namespace prefix for the metric name.
func buildMetricName(namespace, name string) string { _ = "STUB: not implemented"; return "" }
