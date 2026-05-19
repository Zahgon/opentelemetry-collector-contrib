// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datadogreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver"

import (
	"context"
	"io"
	"net/http"
	"net/http/httputil"
	"sync"

	lru "github.com/hashicorp/golang-lru/v2"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver/internal/translator"
)

type datadogReceiver struct {
	address            string
	config             *Config
	params             receiver.Settings
	intakeReverseProxy *httputil.ReverseProxy

	nextTracesConsumer  consumer.Traces
	nextMetricsConsumer consumer.Metrics
	nextLogsConsumer    consumer.Logs

	metricsTranslator *translator.MetricsTranslator
	statsTranslator   *translator.StatsTranslator

	server    *http.Server
	tReceiver *receiverhelper.ObsReport

	traceIDCache *lru.Cache[uint64, pcommon.TraceID]

	shutdownCh chan struct{}
	wg         sync.WaitGroup
}

// Endpoint specifies an API endpoint definition.
type endpoint struct {
	// Pattern specifies the API pattern, as registered by the HTTP handler.
	Pattern string

	// Handler specifies the http.Handler for this endpoint.
	Handler func(http.ResponseWriter, *http.Request)
}

// getEndpoints specifies the list of endpoints registered for the trace-agent API.
func (ddr *datadogReceiver) getEndpoints() []endpoint { _ = "STUB: not implemented"; return nil }

// the datadog agent is configured to use a trailing slash in some places:
// https://github.com/DataDog/datadog-agent/blob/7.64.3/comp/forwarder/defaultforwarder/endpoints/endpoints.go#L18

func newDataDogReceiver(ctx context.Context, config *Config, params receiver.Settings) (component.Component, error) {
	_ = "STUB: not implemented"
	return *new(component.Component), nil
}

func (ddr *datadogReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Starts the cleanup loop to remove idle series

func (ddr *datadogReceiver) Shutdown(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	// Signal background goroutines to stop
	return nil
}

// Wait for them to finish

func (ddr *datadogReceiver) buildInfoResponse(endpoints []endpoint) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handleInfo handles incoming /info payloads.
func (ddr *datadogReceiver) handleInfo(w http.ResponseWriter, _ *http.Request, infoResponse []byte) {
	_ = "STUB: not implemented"
	return
}

func (ddr *datadogReceiver) handleLogs(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Ping mechanism of Datadog SDK perform http request with empty body when GET /info not implemented.
// The response code should be different of 404 to be considered ok by Datadog SDK.

// try parsing as array first, then single record

// now try parsing as a single record

// handleStatsV2 handles incoming stats payloads from datadog agent
// Stats payloads are sent from the DataDog agent at /api/v0.2/stats
func (ddr *datadogReceiver) handleStatsV2(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Get content encoding from header

// Create a reader that handles decompression if needed

// Decode the StatsPayload (NOT ClientStatsPayload directly)
// The agent wraps ClientStatsPayload(s) inside a StatsPayload

// Validate the payload

// Process each ClientStatsPayload within the StatsPayload
// The agent may send multiple ClientStatsPayload entries in a single request

// Extract metadata from headers (fallback if not in payload)

// Translate each client stats payload to metrics

// Send to metrics consumer

func (ddr *datadogReceiver) handleTraces(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Ping mechanism of Datadog SDK perform http request with empty body when GET /info not implemented.
// The response code should be different of 404 to be considered ok by Datadog SDK.

// Match the response logic from dd-agent https://github.com/DataDog/datadog-agent/blob/86b2ae24f93941447a5bf0a2b6419caed77e76dd/pkg/trace/api/api.go#L511-L519

// Keep the "OK" response for these versions

// handleV1Series handles the v1 series endpoint https://docs.datadoghq.com/api/latest/metrics/#submit-metrics
func (ddr *datadogReceiver) handleV1Series(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// handleV2Series handles the v2 series endpoint https://docs.datadoghq.com/api/latest/metrics/#submit-metrics
func (ddr *datadogReceiver) handleV2Series(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// handleCheckRun handles the service checks endpoint https://docs.datadoghq.com/api/latest/service-checks/
func (ddr *datadogReceiver) handleCheckRun(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// try parsing as array first, then single service check

// now try parsing as a single service check

// handleSketches handles sketches, the underlying data structure of distributions https://docs.datadoghq.com/metrics/distributions/
func (ddr *datadogReceiver) handleSketches(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// handleIntake handles operational calls made by the agent to submit host tags and other metadata to the backend.
func (ddr *datadogReceiver) handleIntake(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// handleDistributionPoints handles the distribution points endpoint https://docs.datadoghq.com/api/latest/metrics/#submit-distribution-points
func (ddr *datadogReceiver) handleDistributionPoints(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// handleStats handles incoming stats payloads.
func (ddr *datadogReceiver) handleStats(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func createIntakeReverseProxyDirector(site, key string) func(*http.Request) {
	_ = "STUB: not implemented"
	return nil
}

// we want to use our own API key for all calls

// intake puts the API key in the query string as well

// Technically, the JSON body of the `/intake` request contains the API key as well
// (it's the top-level `apiKey` field of the payload JSON object),
// but it appears as though the value of that field does not matter,
// at least when it comes to matching the actual `DD-API-KEY` we set in the header above.
// So, to avoid a bunch of extra expensive work in the collector, we don't touch the body.

// runIdleSeriesCleanup runs the loop that checks for and removes idle series.
func (ddr *datadogReceiver) runIdleSeriesCleanup() { _ = "STUB: not implemented"; return }

// Assumes validation was done in Start(), so cleanupInterval is positive.

// createDecompressingReader creates a reader that handles decompression based on the content encoding.
// Supported encodings: gzip. Returns the original reader if encoding is empty or unsupported.
func createDecompressingReader(body io.ReadCloser, contentEncoding string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}
