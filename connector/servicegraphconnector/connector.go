// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package servicegraphconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/servicegraphconnector"

import (
	"context"
	"sync"
	"time"

	"github.com/lightstep/go-expohisto/structure"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
	conventionsv125 "go.opentelemetry.io/otel/semconv/v1.25.0"
	conventionsv128 "go.opentelemetry.io/otel/semconv/v1.28.0"
	conventionsv138 "go.opentelemetry.io/otel/semconv/v1.38.0"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/servicegraphconnector/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/servicegraphconnector/internal/store"
)

const (
	metricKeySeparator = string(byte(0))
	clientKind         = "client"
	serverKind         = "server"
	virtualNodeLabel   = "virtual_node"
	millisecondsUnit   = "ms"
	secondsUnit        = "s"
)

var (
	legacyDefaultLatencyHistogramBuckets = []float64{
		2, 4, 6, 8, 10, 50, 100, 200, 400, 800, 1000, 1400, 2000, 5000, 10_000, 15_000,
	}
	defaultLatencyHistogramBuckets = []float64{
		0.002, 0.004, 0.006, 0.008, 0.01, 0.05, 0.1, 0.2, 0.4, 0.8, 1, 1.4, 2, 5, 10, 15,
	}

	defaultPeerAttributes = []string{
		string(conventionsv138.PeerServiceKey), string(conventionsv125.DBNameKey), string(conventionsv128.DBSystemKey),
	}

	defaultDatabaseNameAttributes = []string{string(conventionsv125.DBNameKey)}

	defaultMetricsFlushInterval = 60 * time.Second // 1 DPM
)

type metricSeries struct {
	dimensions  pcommon.Map
	lastUpdated int64 // Used to remove stale series
}

var _ processor.Traces = (*serviceGraphConnector)(nil)

type serviceGraphConnector struct {
	config          *Config
	logger          *zap.Logger
	metricsConsumer consumer.Metrics

	store *store.Store

	startTime time.Time

	seriesMutex                          sync.Mutex
	reqTotal                             map[string]int64
	reqFailedTotal                       map[string]int64
	reqClientDurationSecondsCount        map[string]uint64
	reqClientDurationSecondsSum          map[string]float64
	reqClientDurationSecondsBucketCounts map[string][]uint64
	reqClientDurationExpHistogram        map[string]*structure.Histogram[float64]
	reqServerDurationSecondsCount        map[string]uint64
	reqServerDurationSecondsSum          map[string]float64
	reqServerDurationSecondsBucketCounts map[string][]uint64
	reqServerDurationExpHistogram        map[string]*structure.Histogram[float64]
	reqDurationBounds                    []float64

	metricMutex sync.RWMutex
	keyToMetric map[string]metricSeries

	telemetryBuilder *metadata.TelemetryBuilder

	shutdownCh chan any
}

func newConnector(set component.TelemetrySettings, config component.Config, next consumer.Metrics) (*serviceGraphConnector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *serviceGraphConnector) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *serviceGraphConnector) metricFlushLoop(ctx context.Context, flushInterval time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (p *serviceGraphConnector) flushMetrics(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip empty metrics.

// Firstly, export md to avoid being impacted by downstream trace serviceGraphConnector errors/latency.

func (p *serviceGraphConnector) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (*serviceGraphConnector) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (p *serviceGraphConnector) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// If metricsFlushInterval is not set, flush metrics immediately.

// Not return error here to avoid impacting traces.

func (p *serviceGraphConnector) aggregateMetrics(ctx context.Context, td ptrace.Traces) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// If service.name doesn't exist, skip processing this trace

// override connection type and continue processing as span kind client

// A database request will only have one span, we don't wait for the server
// span but just copy details from the client span

// override connection type and continue processing as span kind server

// this span is not part of an edge

// UpsertEdge will only return ErrTooManyItems

func (p *serviceGraphConnector) upsertDimensions(kind string, m map[string]string, resourceAttr, spanAttr pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func (*serviceGraphConnector) upsertPeerAttributes(m []string, peers map[string]string, spanAttr pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func (p *serviceGraphConnector) onComplete(e *store.Edge) { _ = "STUB: not implemented"; return }

func (p *serviceGraphConnector) onExpire(e *store.Edge) { _ = "STUB: not implemented"; return }

func (p *serviceGraphConnector) aggregateMetricsForEdge(e *store.Edge) {
	_ = "STUB: not implemented"
	return
}

func (p *serviceGraphConnector) updateSeries(key string, dimensions pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

// Overwrite the series if it already exists

func (p *serviceGraphConnector) dimensionsForSeries(key string) (pcommon.Map, bool) {
	_ = "STUB: not implemented"
	return *new(pcommon.Map), false
}

func (p *serviceGraphConnector) updateCountMetrics(key string) { _ = "STUB: not implemented"; return }

func (p *serviceGraphConnector) updateErrorMetrics(key string) { _ = "STUB: not implemented"; return }

func (p *serviceGraphConnector) updateDurationMetrics(key string, serverDuration, clientDuration float64) {
	_ = "STUB: not implemented"
	return
}

func (p *serviceGraphConnector) updateServerDurationMetrics(key string, duration float64) {
	_ = "STUB: not implemented"
	return
}

// Search bucket index

func (p *serviceGraphConnector) updateClientDurationMetrics(key string, duration float64) {
	_ = "STUB: not implemented"
	return
}

// Search bucket index

func buildDimensions(e *store.Edge) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func addExtraLabel(dimensions pcommon.Map, label, value string) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

// nowWithOffset returns the current time minus the configured offset
func (p *serviceGraphConnector) nowWithOffset() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (p *serviceGraphConnector) buildMetrics() (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Obtain write lock to reset data

func (p *serviceGraphConnector) collectCountMetrics(ilm pmetric.ScopeMetrics) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Support other aggregation temporalities

// TODO: Support other aggregation temporalities

func (p *serviceGraphConnector) collectLatencyMetrics(ilm pmetric.ScopeMetrics) error {
	_ = "STUB: not implemented"
	// TODO: Remove this once legacy metric names are removed
	return nil
}

func (p *serviceGraphConnector) collectClientLatencyMetrics(ilm pmetric.ScopeMetrics) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Support other aggregation temporalities

// TODO: Support exemplars

func (p *serviceGraphConnector) collectServerLatencyMetrics(ilm pmetric.ScopeMetrics, mName string) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Support other aggregation temporalities

// TODO: Support exemplars

func (p *serviceGraphConnector) buildMetricKey(clientName, serverName, connectionType, failed string, edgeDimensions map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

// storeExpirationLoop periodically expires old entries from the store.
func (p *serviceGraphConnector) storeExpirationLoop(d time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (*serviceGraphConnector) getPeerHost(m []string, peers map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

// cacheLoop periodically cleans the cache
func (p *serviceGraphConnector) cacheLoop(d time.Duration) { _ = "STUB: not implemented"; return }

// cleanCache removes series that have not been updated in 15 minutes
func (p *serviceGraphConnector) cleanCache() { _ = "STUB: not implemented"; return }

// spanDuration returns the duration of the given span in seconds (legacy ms).
func spanDuration(span ptrace.Span) float64 { _ = "STUB: not implemented"; return 0 }

// durationToFloat converts the given duration to the number of seconds (legacy ms) it represents.
func durationToFloat(d time.Duration) float64 { _ = "STUB: not implemented"; return 0 }

func mapDurationsToFloat(vs []time.Duration) []float64 { _ = "STUB: not implemented"; return nil }
