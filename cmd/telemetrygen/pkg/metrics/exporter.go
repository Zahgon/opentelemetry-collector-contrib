// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics

import (
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
)

// grpcExporterOptions creates the configuration options for a gRPC-based OTLP metric exporter.
// It configures the exporter with the provided endpoint, connection security settings, and headers.
func grpcExporterOptions(cfg *Config) ([]otlpmetricgrpc.Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// httpExporterOptions creates the configuration options for an HTTP-based OTLP metric exporter.
// It configures the exporter with the provided endpoint, URL path, connection security settings, and headers.
func httpExporterOptions(cfg *Config) ([]otlpmetrichttp.Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
