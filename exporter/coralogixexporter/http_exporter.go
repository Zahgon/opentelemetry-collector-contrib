// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package coralogixexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/pdata/plog/plogotlp"
	"go.opentelemetry.io/collector/pdata/pmetric/pmetricotlp"
	"go.opentelemetry.io/collector/pdata/ptrace/ptraceotlp"
)

type httpExporter struct {
	Client   *http.Client
	Endpoint string
}

type httpLogsExporter struct {
	httpExporter
}

type httpMetricsExporter struct {
	httpExporter
}

type httpTracesExporter struct {
	httpExporter
}

// httpError wraps HTTP status codes, headers, and response body for error handling
type httpError struct {
	StatusCode int
	Header     http.Header
	Body       []byte
	Message    string
}

func (e *httpError) Error() string { _ = "STUB: not implemented"; return "" }

func newHTTPLogsExporter(client *http.Client, config *Config) httpLogsExporter {
	_ = "STUB: not implemented"
	return *new(httpLogsExporter)
}

func newHTTPMetricsExporter(client *http.Client, config *Config) httpMetricsExporter {
	_ = "STUB: not implemented"
	return *new(httpMetricsExporter)
}

func newHTTPTracesExporter(client *http.Client, config *Config) httpTracesExporter {
	_ = "STUB: not implemented"
	return *new(httpTracesExporter)
}

// doRequest performs the HTTP request with the given body and path
func (c *httpExporter) doRequest(ctx context.Context, body []byte, path string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *httpLogsExporter) Export(ctx context.Context, request plogotlp.ExportRequest) (plogotlp.ExportResponse, error) {
	_ = "STUB: not implemented"
	return *new(plogotlp.ExportResponse), nil
}

func (e *httpMetricsExporter) Export(ctx context.Context, request pmetricotlp.ExportRequest) (pmetricotlp.ExportResponse, error) {
	_ = "STUB: not implemented"
	return *new(pmetricotlp.ExportResponse), nil
}

func (e *httpTracesExporter) Export(ctx context.Context, request ptraceotlp.ExportRequest) (ptraceotlp.ExportResponse, error) {
	_ = "STUB: not implemented"
	return *new(ptraceotlp.ExportResponse), nil
}
