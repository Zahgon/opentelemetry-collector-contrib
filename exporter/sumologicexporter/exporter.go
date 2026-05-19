// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumologicexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sumologicexporter"

import (
	"context"
	"net/http"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pipeline"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sumologicexporter/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/sumologicextension"
)

const (
	logsDataURL    = "/api/v1/collector/logs"
	metricsDataURL = "/api/v1/collector/metrics"
	tracesDataURL  = "/api/v1/collector/traces"
)

type sumologicexporter struct {
	config *Config
	host   component.Host
	logger *zap.Logger

	clientLock sync.RWMutex
	client     *http.Client

	prometheusFormatter prometheusFormatter

	// Lock around data URLs is needed because the reconfiguration of the exporter
	// can happen asynchronously whenever the exporter is re registering.
	dataURLsLock   sync.RWMutex
	dataURLMetrics string
	dataURLLogs    string
	dataURLTraces  string

	foundSumologicExtension bool
	sumologicExtension      *sumologicextension.SumologicExtension

	stickySessionCookieLock sync.RWMutex
	stickySessionCookie     string

	id               component.ID
	sender           *sender
	telemetryBuilder *metadata.TelemetryBuilder
}

func initExporter(cfg *Config, set exporter.Settings) (*sumologicexporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE: client is now set in start()

func newLogsExporter(
	ctx context.Context,
	params exporter.Settings,
	cfg *Config,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

// Disable exporterhelper Timeout, since we are using a custom mechanism
// within exporter itself

func newMetricsExporter(
	ctx context.Context,
	params exporter.Settings,
	cfg *Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

// Disable exporterhelper Timeout, since we are using a custom mechanism
// within exporter itself

func newTracesExporter(
	ctx context.Context,
	params exporter.Settings,
	cfg *Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

// Disable exporterhelper Timeout, since we are using a custom mechanism
// within exporter itself

// start starts the exporter
func (se *sumologicexporter) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (se *sumologicexporter) configure(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// If user specified using sumologicextension as auth but none was
// found then return an error.

// If we're using sumologicextension as authentication extension and
// endpoint was not set then send data on a collector generic ingest URL
// with authentication set by sumologicextension.

// Clean authenticator if set to sumologic.
// Setting to null in configuration doesn't work, so we have to force it that way.

func (se *sumologicexporter) setHTTPClient(client *http.Client) { _ = "STUB: not implemented"; return }

func (se *sumologicexporter) getHTTPClient() *http.Client { _ = "STUB: not implemented"; return nil }

func (se *sumologicexporter) setDataURLs(logs, metrics, traces string) {
	_ = "STUB: not implemented"
	return
}

func (se *sumologicexporter) getDataURLs() (logs, metrics, traces string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

func (*sumologicexporter) shutdown(context.Context) error {
	_ = "STUB: not implemented"

	// pushLogsData groups data with common metadata and sends them as separate batched requests.
	// It returns the number of unsent logs and an error which contains a list of dropped records
	// so they can be handled by OTC retry mechanism
	return nil
}

func (se *sumologicexporter) pushLogsData(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	// Follow different execution path for OTLP format
	return nil
}

// Iterate over ResourceLogs

// Copy all dropped records to Logs
// NOTE: we only copy resource and log records here.
// Scope is not handled properly but it never was.

// pushMetricsData groups data with common metadata and send them as separate batched requests
// it returns number of unsent metrics and error which contains list of dropped records
// so they can be handle by the OTC retry mechanism
func (se *sumologicexporter) pushMetricsData(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// handleUnauthorizedErrors checks if any of the provided errors is an unauthorized error.
// In which case it triggers exporter reconfiguration which in turn takes the credentials
// from sumologicextension which at this point should already detect the problem with
// authorization (via heartbeats) and prepare new collector credentials to be available.
func (se *sumologicexporter) handleUnauthorizedErrors(ctx context.Context, errs ...error) {
	_ = "STUB: not implemented"
	return
}

// It's enough to successfully reconfigure the exporter just once.

func (se *sumologicexporter) pushTracesData(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (se *sumologicexporter) StickySessionCookie() string { _ = "STUB: not implemented"; return "" }

func (se *sumologicexporter) SetStickySessionCookie(stickySessionCookie string) {
	_ = "STUB: not implemented"
	return
}

// get the destination url for a given signal type
// this mostly adds signal-specific suffixes if the format is otlp
func getSignalURL(oCfg *Config, endpointURL string, signal pipeline.Signal) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func sanitizeURL(urlString string) string { _ = "STUB: not implemented"; return "" }

func nchars(b byte, n int) string { _ = "STUB: not implemented"; return "" }
