// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tinybirdexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tinybirdexporter"

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

const (
	headerRetryAfter  = "Retry-After"
	contentTypeNDJSON = "application/x-ndjson"
)

type tinybirdExporter struct {
	config             *Config
	client             *http.Client
	logger             *zap.Logger
	settings           component.TelemetrySettings
	userAgent          string
	maxRequestBodySize int
}

func newExporter(cfg component.Config, set exporter.Settings, opts ...option) *tinybirdExporter {
	_ = "STUB: not implemented"
	return nil
}

// Apply options

// Log the error but continue with default values

func (e *tinybirdExporter) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tinybirdExporter) pushTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tinybirdExporter) pushMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tinybirdExporter) pushLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *tinybirdExporter) exportBuffers(ctx context.Context, dataSource string, buffers []*bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

// At-most-once delivery. If we have already performed a successful export,
// we return a permanent error to indicate that the error is not retryable.

// As we have not performed any successful export, we are free to retry if needed.

func (e *tinybirdExporter) export(ctx context.Context, dataSource string, body io.Reader) error {
	_ = "STUB: not implemented"
	// Create request and add query parameters
	return nil
}

// Set headers

// Send request

// Drain the response body to avoid leaking resources.

// Check if the request was successful.

// Read error response

// If the status code is not retryable, return a permanent error.

// Check if the server is overwhelmed.

// The value of Retry-After field can be either an HTTP-date or a number of
// seconds to delay after the response is received. See https://datatracker.ietf.org/doc/html/rfc7231#section-7.1.3
//
// Tinybird Events API returns the delay-seconds in the Retry-After header.
// https://www.tinybird.co/docs/forward/get-data-in/events-api#rate-limit-headers

// Determine if the status code is retryable according to Tinybird Events API.
// See https://www.tinybird.co/docs/api-reference/events-api#return-http-status-codes
func isRetryableStatusCode(code int) bool { _ = "STUB: not implemented"; return false }
