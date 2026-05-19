// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package coralogixexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configopaque"
	exp "go.opentelemetry.io/collector/exporter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type signalConfig interface {
	ToClientConn(ctx context.Context, host component.Host, settings component.TelemetrySettings, opts ...configgrpc.ToClientConnOption) (*grpc.ClientConn, error)
	ToHTTPClient(ctx context.Context, host component.Host, settings component.TelemetrySettings) (*http.Client, error)
	GetWaitForReady() bool
	GetEndpoint() string
	GetAcceptEncoding() string
}

var _ signalConfig = (*signalConfigWrapper)(nil)

type signalConfigWrapper struct {
	config *TransportConfig
}

func (w *signalConfigWrapper) ToClientConn(ctx context.Context, host component.Host, settings component.TelemetrySettings, opts ...configgrpc.ToClientConnOption) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *signalConfigWrapper) ToHTTPClient(ctx context.Context, host component.Host, settings component.TelemetrySettings) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *signalConfigWrapper) GetWaitForReady() bool { _ = "STUB: not implemented"; return false }

func (w *signalConfigWrapper) GetEndpoint() string { _ = "STUB: not implemented"; return "" }

func (w *signalConfigWrapper) GetAcceptEncoding() string { _ = "STUB: not implemented"; return "" }

func newSignalExporter(oCfg *Config, set exp.Settings, signalEndpoint string, headers configopaque.MapList) (*signalExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type signalExporter struct {
	// Input configuration.
	config *Config

	// GRPC exporter
	clientConn  *grpc.ClientConn
	callOptions []grpc.CallOption

	// HTTP exporter
	clientHTTP *http.Client

	settings component.TelemetrySettings

	// Default user-agent header.
	userAgent string

	// Cached metadata for outgoing context
	metadata metadata.MD

	rateError rateError
}

func (e *signalExporter) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *signalExporter) enhanceContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (e *signalExporter) startSignalExporter(ctx context.Context, host component.Host, signalConfig signalConfig) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Only add AcceptCompressors if a compression encoding is specified

func (e *signalExporter) EnableRateLimit() { _ = "STUB: not implemented"; return }

func (e *signalExporter) canSend() bool { _ = "STUB: not implemented"; return false }

// processError implements the common OTLP logic around request handling such as retries and throttling.
// Send a telemetry data request to the server. "perform" function is expected to make
// the actual gRPC unary call that sends the request.
func (e *signalExporter) processError(err error) error { _ = "STUB: not implemented"; return nil }
