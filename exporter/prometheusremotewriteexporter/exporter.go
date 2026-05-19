// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewriteexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter"

import (
	"context"
	"net/http"
	"net/url"
	"sync"

	remoteapi "github.com/prometheus/client_golang/exp/api/remote"
	"github.com/prometheus/prometheus/prompb"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/otel/attribute"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"
)

type prwTelemetry interface {
	recordTranslationFailure(ctx context.Context)
	recordTranslatedTimeSeries(ctx context.Context, numTS int)
	recordRemoteWriteSentBatch(ctx context.Context)
	setNumberConsumer(ctx context.Context, n int64)
	recordWrittenSamples(ctx context.Context, numSamples int64)
	recordWrittenHistograms(ctx context.Context, numHistograms int64)
	recordWrittenExemplars(ctx context.Context, numExemplars int64)
}

type prwTelemetryOtel struct {
	telemetryBuilder *metadata.TelemetryBuilder
	otelAttrs        []attribute.KeyValue
}

func (p *prwTelemetryOtel) setNumberConsumer(ctx context.Context, n int64) {
	_ = "STUB: not implemented"
	return
}

func (p *prwTelemetryOtel) recordRemoteWriteSentBatch(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (p *prwTelemetryOtel) recordTranslationFailure(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (p *prwTelemetryOtel) recordTranslatedTimeSeries(ctx context.Context, numTS int) {
	_ = "STUB: not implemented"
	return
}

func (p *prwTelemetryOtel) recordWrittenSamples(ctx context.Context, numSamples int64) {
	_ = "STUB: not implemented"
	return
}

func (p *prwTelemetryOtel) recordWrittenHistograms(ctx context.Context, numHistograms int64) {
	_ = "STUB: not implemented"
	return
}

func (p *prwTelemetryOtel) recordWrittenExemplars(ctx context.Context, numExemplars int64) {
	_ = "STUB: not implemented"
	return
}

type gogoProto interface {
	Size() int
	MarshalToSizedBuffer([]byte) (int, error)
}

type buffer struct {
	protobuf []byte
	snappy   []byte
}

func (b *buffer) MarshalAndEncode(req gogoProto) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If we don't pass a buffer large enough, Snappy Encode function will not use it and instead will allocate a new buffer.
// Manually grow the buffer to make sure Snappy uses it and we can re-use it afterwards.

// A reusable buffer pool for serializing protobufs and compressing them with Snappy.
var bufferPool = sync.Pool{
	New: func() any {
		return &buffer{
			protobuf: nil,
			snappy:   nil,
		}
	},
}

// prwExporter converts OTLP metrics to Prometheus remote write TimeSeries and sends them to a remote endpoint.
type prwExporter struct {
	endpointURL         *url.URL
	client              *http.Client
	wg                  *sync.WaitGroup
	closeChan           chan struct{}
	concurrency         int
	userAgentHeader     string
	maxBatchSizeBytes   int
	clientSettings      *confighttp.ClientConfig
	settings            component.TelemetrySettings
	retrySettings       configretry.BackOffConfig
	retryOnHTTP429      bool
	wal                 *prweWAL
	exporterSettings    prometheusremotewrite.Settings
	telemetry           prwTelemetry
	RemoteWriteProtoMsg remoteapi.WriteMessageType

	// When concurrency is enabled, concurrent goroutines would potentially
	// fight over the same batchState object. To avoid this, we use a pool
	// to provide each goroutine with its own state.
	batchStatePool sync.Pool
}

func newPRWTelemetry(set exporter.Settings, endpointURL *url.URL) (prwTelemetry, error) {
	_ = "STUB: not implemented"
	return *new(prwTelemetry), nil
}

// newPRWExporter initializes a new prwExporter instance and sets fields accordingly.
func newPRWExporter(cfg *Config, set exporter.Settings) (*prwExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set the desired number of consumers as a metric for the exporter.

// Start creates the prometheus client
func (prwe *prwExporter) Start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (prwe *prwExporter) shutdownWALIfEnabled() error { _ = "STUB: not implemented"; return nil }

// Shutdown stops the exporter from accepting incoming calls(and return error), and wait for current export operations
// to finish before returning
func (prwe *prwExporter) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (prwe *prwExporter) pushMetricsV1(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// Call export even if a conversion error, since there may be points that were successfully converted.

// PushMetrics converts metrics to Prometheus remote write TimeSeries and send to remote endpoint. It maintain a map of
// TimeSeries, validates and handles each individual metric, adding the converted TimeSeries to the map, and finally
// exports the map.
func (prwe *prwExporter) PushMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// If feature flag not enabled support only RW1.

// If feature flag was enabled check if we want to send RW1 or RW2.

func validateAndSanitizeExternalLabels(cfg *Config) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (prwe *prwExporter) handleExport(ctx context.Context, tsMap map[string]*prompb.TimeSeries, m []*prompb.MetricMetadata) error {
	_ = "STUB: not implemented"
	// There are no metrics to export, so return.
	return nil
}

// Calls the helper function to convert and batch the TsMap to the desired format

// Perform a direct export otherwise.

// Otherwise the WAL is enabled, and just persist the requests to the WAL

// export sends a Snappy-compressed WriteRequest containing TimeSeries to a remote write endpoint in order
func (prwe *prwExporter) export(ctx context.Context, requests []*prompb.WriteRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// used to wait for workers to be finished

// Run concurrencyLimit of workers until there
// is no more requests to execute in the input channel.

func (prwe *prwExporter) handleRequests(ctx context.Context, input chan *prompb.WriteRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Check firstly to ensure that the context wasn't cancelled.

func (prwe *prwExporter) execute(ctx context.Context, buf []byte) error {
	_ = "STUB: not implemented"

	// executeFunc can be used for backoff and non backoff scenarios.
	return nil
}

// check there was no timeout in the component level to avoid retries
// to continue to run after a timeout

// continue

// Create the HTTP POST request to send to the endpoint

// Add necessary headers specified by:
// https://cortexmetrics.io/docs/apis/#remote-api

// If feature flag not enabled support only RW1

// Per the Prometheus remote write 2.0 specification, the response should contain
// X-Prometheus-Remote-Write-Samples-Written header.
// If the header is missing, it suggests that the endpoint does not support RW2 or the
// implementation is not compliant with the specification. Reference:
// https://prometheus.io/docs/specs/prw/remote_write_spec_2_0/#required-written-response-headers

// 2xx status code is considered a success
// 5xx errors are recoverable and the exporter should retry
// Reference for different behavior according to status code:
// https://github.com/prometheus/prometheus/pull/2552/files#diff-ae8db9d16d8057358e49d694522e7186

// 429 errors are recoverable and the exporter should retry if RetryOnHTTP429 enabled
// Reference: https://github.com/prometheus/prometheus/pull/12677

// Use the BackOff instance to retry the func with exponential backoff.

// A permanent error is being returned here so we don't retry on context deadline exceeded.

func (prwe *prwExporter) walEnabled() bool { _ = "STUB: not implemented"; return false }

func (prwe *prwExporter) turnOnWALIfEnabled(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
