// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logzioexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter"

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-hclog"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"google.golang.org/genproto/googleapis/rpc/status"
)

const (
	loggerName               = "logzio-exporter"
	headerRetryAfter         = "Retry-After"
	headerAuthorization      = "Authorization"
	maxHTTPResponseReadBytes = 64 * 1024
)

// logzioExporter implements an OpenTelemetry trace exporter that exports all spans to Logz.io
type logzioExporter struct {
	config   *Config
	client   *http.Client
	logger   hclog.Logger
	settings component.TelemetrySettings
}

func newLogzioExporter(cfg *Config, params exporter.Settings) (*logzioExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newLogzioTracesExporter(config *Config, set exporter.Settings) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

// disable since we rely on http.Client timeout logic.

func newLogzioLogsExporter(config *Config, set exporter.Settings) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

// disable since we rely on http.Client timeout logic.

func (exporter *logzioExporter) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (exporter *logzioExporter) pushLogData(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func mergeMapEntries(maps ...pcommon.Map) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

// Check if the key was already added

// Create a new slice and append values if the key exists:

func (exporter *logzioExporter) pushTraceData(ctx context.Context, traces ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// export is similar to otlp_http export method with changes in log messages + Permanent error for `StatusUnauthorized` and `StatusForbidden`
// https://github.com/open-telemetry/opentelemetry-collector/blob/main/exporter/otlphttpexporter/otlp.go#L127
func (exporter *logzioExporter) export(ctx context.Context, url string, request []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Discard any remaining response body when we are done reading.

// Request is successful.

// Format the error message. Use the status if it is present in the response.

// Check if the server is overwhelmed.
// See spec https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/protocol/otlp.md#throttling-1

// Fallback to 0 if the Retry-After header is not present. This will trigger the
// default backoff policy by our caller (retry handler).

// Indicate to our caller to pause for the specified number of seconds.

// All other errors are retryable, so don't wrap them in consumererror.NewPermanent().

// Read the response and decode the status.Status from the body.
// Returns nil if the response is empty or cannot be decoded.
func readResponse(resp *http.Response) *status.Status { _ = "STUB: not implemented"; return nil }

// Request failed. Read the body. OTLP spec says:
// "Response body for all HTTP 4xx and HTTP 5xx responses MUST be a
// Protobuf-encoded Status message that describes the problem."

// Decode it as Status struct. See https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/protocol/otlp.md#failures
