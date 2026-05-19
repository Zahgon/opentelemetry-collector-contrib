// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunkhecexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter"

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"

	translator "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/splunk"
)

// allow monkey patching for injecting pushLogData function in test
var getPushLogFn = func(c *client) func(ctx context.Context, ld plog.Logs) error {
	return c.pushLogData
}

// iterState captures a state of iteration over the pdata Logs/Metrics/Traces instances.
type iterState struct {
	resource int // index in ResourceLogs/ResourceMetrics/ResourceSpans list
	library  int // index in ScopeLogs/ScopeMetrics/ScopeSpans list
	record   int // index in Logs/Metrics/Spans list
	done     bool
}

func (s iterState) empty() bool { _ = "STUB: not implemented"; return false }

// client sends the data to the splunk backend.
type client struct {
	config            *Config
	logger            *zap.Logger
	wg                sync.WaitGroup
	telemetrySettings component.TelemetrySettings
	hecWorker         hecWorker
	buildInfo         component.BuildInfo
	heartbeater       *heartbeater
	bufferPool        bufferPool
	exporterName      string
	meter             metric.Meter
}

func newClient(set exporter.Settings, cfg *Config, maxContentLength uint) *client {
	_ = "STUB: not implemented"
	return nil
}

func newLogsClient(set exporter.Settings, cfg *Config) *client {
	_ = "STUB: not implemented"
	return nil
}

func newTracesClient(set exporter.Settings, cfg *Config) *client {
	_ = "STUB: not implemented"
	return nil
}

func newMetricsClient(set exporter.Settings, cfg *Config) *client {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) pushMetricsData(
	ctx context.Context,
	md pmetric.Metrics,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) pushTraceData(
	ctx context.Context,
	td ptrace.Traces,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) pushLogData(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// All logs in a batch have the same access token after batchperresourceattr, so we can just check the first one.

// All logs in a batch have only one type (regular or profiling logs) after perScopeBatcher,
// so we can just check the first one.

// A guesstimated value > length of bytes of a single event.
// Added to buffer capacity so that buffer is likely to grow by reslicing when buf.Len() > bufCap.
const (
	bufCapPadding        = uint(4096)
	libraryHeaderName    = "X-Splunk-Instrumentation-Library"
	profilingLibraryName = "otel.profiling"
)

func isProfilingData(sl plog.ScopeLogs) bool { _ = "STUB: not implemented"; return false }

// pushLogDataInBatches sends batches of Splunk events in JSON format.
// The batch content length is restricted to MaxContentLengthLogs.
// ld log records are parsed to Splunk events.
func (c *client) pushLogDataInBatches(ctx context.Context, ld plog.Logs, headers map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// fillLogsBuffer fills the buffer with Splunk events until the buffer is full or all logs are processed.
func (c *client) fillLogsBuffer(logs plog.Logs, buf buffer, is iterState) (iterState, []error) {
	_ = "STUB: not implemented"
	return *new(iterState), nil
}

// Reset library index for next resource.

// Reset record index for next library.

// Parsing log record to Splunk event.

// TODO record this drop as a metric

// JSON encoding event and writing to buffer.

// Continue adding events to buffer up to capacity.

func (c *client) fillMetricsBuffer(metrics pmetric.Metrics, buf buffer, is iterState) (iterState, []error) {
	_ = "STUB: not implemented"
	return *new(iterState), nil
}

// Reset library index for next resource.

// Reset record index for next library.

// Parsing metric record to Splunk event.

// JSON encoding event and writing to buffer.

// Continue adding events to buffer up to capacity.

func (c *client) fillMetricsBufferMultiMetrics(events []*translator.Event, buf buffer, is iterState) (iterState, []error) {
	_ = "STUB: not implemented"
	return *new(iterState), nil
}

// JSON encoding event and writing to buffer.

func (c *client) fillTracesBuffer(traces ptrace.Traces, buf buffer, is iterState) (iterState, []error) {
	_ = "STUB: not implemented"
	return *new(iterState), nil
}

// Reset library index for next resource.

// Reset record index for next library.

// Parsing span record to Splunk event.

// JSON encoding event and writing to buffer.

// Continue adding events to buffer up to capacity.

// pushMultiMetricsDataInBatches sends batches of Splunk multi-metric events in JSON format.
// The batch content length is restricted to MaxContentLengthMetrics.
// md metrics are parsed to Splunk events.
func (c *client) pushMultiMetricsDataInBatches(ctx context.Context, md pmetric.Metrics, headers map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Parsing metric record to Splunk event.

// pushMetricsDataInBatches sends batches of Splunk events in JSON format.
// The batch content length is restricted to MaxContentLengthMetrics.
// md metrics are parsed to Splunk events.
func (c *client) pushMetricsDataInBatches(ctx context.Context, md pmetric.Metrics, headers map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// pushTracesDataInBatches sends batches of Splunk events in JSON format.
// The batch content length is restricted to MaxContentLengthMetrics.
// td traces are parsed to Splunk events.
func (c *client) pushTracesDataInBatches(ctx context.Context, td ptrace.Traces, headers map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) postEvents(ctx context.Context, buf buffer, headers map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// subLogs returns a subset of logs starting from the state.
func subLogs(src plog.Logs, state iterState) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

// subMetrics returns a subset of metrics starting from the state.
func subMetrics(src pmetric.Metrics, state iterState) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

func subTraces(src ptrace.Traces, state iterState) ptrace.Traces {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces)
}

func (c *client) stop(context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *client) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func checkHecHealth(ctx context.Context, client *http.Client, healthCheckURL *url.URL) error {
	_ = "STUB: not implemented"
	return nil
}

func buildHTTPClient(ctx context.Context, config *Config, host component.Host, telemetrySettings component.TelemetrySettings) (*http.Client, error) {
	_ = "STUB: not implemented"
	// we handle compression explicitly.
	return nil, nil
}

func buildHTTPHeaders(config *Config, buildInfo component.BuildInfo) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

var jsonBufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

// marshalEvent marshals an event to JSON
func marshalEvent(event *translator.Event, sizeLimit uint, writer io.Writer) (error, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
