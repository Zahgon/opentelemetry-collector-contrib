// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs

import (
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
)

// grpcExporterOptions creates the configuration options for a gRPC-based OTLP log exporter.
// It configures the exporter with the provided endpoint, connection security settings, and headers.
func grpcExporterOptions(cfg *Config) ([]otlploggrpc.Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// httpExporterOptions creates the configuration options for an HTTP-based OTLP log exporter.
// It configures the exporter with the provided endpoint, URL path, connection security settings, and headers.
func httpExporterOptions(cfg *Config) ([]otlploghttp.Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
