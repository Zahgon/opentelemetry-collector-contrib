// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package traces

import (
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

// grpcExporterOptions creates the configuration options for a gRPC-based OTLP trace exporter.
// It configures the exporter with the provided endpoint, connection security settings, and headers.
func grpcExporterOptions(cfg *Config) ([]otlptracegrpc.Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// httpExporterOptions creates the configuration options for an HTTP-based OTLP trace exporter.
// It configures the exporter with the provided endpoint, URL path, connection security settings, and headers.
func httpExporterOptions(cfg *Config) ([]otlptracehttp.Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
