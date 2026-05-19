// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelarrowexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/otelarrowexporter"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/plog/plogotlp"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pmetric/pmetricotlp"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/pdata/ptrace/ptraceotlp"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/otelarrowexporter/internal/arrow"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/netstats"
)

type exp interface {
	getSettings() exporter.Settings
	getConfig() component.Config

	start(context.Context, component.Host) error
	shutdown(context.Context) error

	pushTraces(context.Context, ptrace.Traces) error
	pushMetrics(context.Context, pmetric.Metrics) error
	pushLogs(context.Context, plog.Logs) error
}

type baseExporter struct {
	// Input configuration.
	config *Config

	// gRPC clients and connection.
	traceExporter  ptraceotlp.GRPCClient
	metricExporter pmetricotlp.GRPCClient
	logExporter    plogotlp.GRPCClient
	clientConn     *grpc.ClientConn
	metadata       metadata.MD
	callOptions    []grpc.CallOption
	settings       exporter.Settings
	netReporter    *netstats.NetworkReporter

	// Default user-agent header.
	userAgent string

	// OTel-Arrow optional state
	arrow *arrow.Exporter
	// streamClientFunc is the stream constructor
	streamClientFactory streamClientFactory
}

var _ exp = (*baseExporter)(nil)

type streamClientFactory func(conn *grpc.ClientConn) arrow.StreamClientFunc

// Crete new exporter and start it. The exporter will begin connecting but
// this function may return before the connection is established.
func newExporter(cfg component.Config, set exporter.Settings, streamClientFactory streamClientFactory, userAgent string, netReporter *netstats.NetworkReporter) (exp, error) {
	_ = "STUB: not implemented"
	return *new(exp), nil
}

func (e *baseExporter) getSettings() exporter.Settings {
	_ = "STUB: not implemented"
	return *new(exporter.Settings)
}

func (e *baseExporter) getConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func (e *baseExporter) setMetadata(md metadata.MD) { _ = "STUB: not implemented"; return }

// start actually creates the gRPC connection. The client construction is deferred till this point as this
// is the only place we get hold of Extensions which are required to construct auth round tripper.
func (e *baseExporter) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Note this sets static outgoing context for all future stream requests.

// Get the auth extension, we'll use it to enrich the request context.

// ignore the error below b/c Validate() was called

// use the configured compressor.

func (e *baseExporter) shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// arrowSendAndWait gets an available stream and tries to send using
// Arrow if it is configured.  A (false, nil) result indicates for the
// caller to fall back to ordinary OTLP.
//
// Note that ctx is has not had enhanceContext() called, meaning it
// will have outgoing gRPC metadata only when an upstream processor or
// receiver placed it there.
func (e *baseExporter) arrowSendAndWait(ctx context.Context, data any) (sent bool, _ error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *baseExporter) pushTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: These should be counted, similar to dropped items.

func (e *baseExporter) pushMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: These should be counted, similar to dropped items.

func (e *baseExporter) pushLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: These should be counted, similar to dropped items.

func (e *baseExporter) enhanceContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func processError(err error) error {
	_ = "STUB: not implemented"

	// Request is successful, we are done.
	return nil
}

// We have an error, check gRPC status code.

// Not really an error, still success.

// Now, this is this a real error.

// It is not a retryable error, we should not retry.

// Check if server returned throttling information.

// We are throttled. Wait before retrying as requested by the server.

// Need to retry.

func shouldRetry(code codes.Code, retryInfo *errdetails.RetryInfo) bool {
	_ = "STUB: not implemented"
	return false
}

// These are retryable errors.

// Retry only if RetryInfo was supplied by the server.
// This indicates that the server can still recover from resource exhaustion.

// Don't retry on any other code.

func getRetryInfo(status *status.Status) *errdetails.RetryInfo {
	_ = "STUB: not implemented"
	return nil
}

func getThrottleDuration(t *errdetails.RetryInfo) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
