// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelarrowexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/otelarrowexporter"

import (
	"context"

	arrowpb "github.com/open-telemetry/otel-arrow/go/api/experimental/arrow/v1"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"google.golang.org/grpc"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/otelarrowexporter/internal/arrow"
)

// NewFactory creates a factory for OTLP exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	// These defaults are taken from the experimental setup used
	// in the blog post covering Phase 1 performance results.  These
	// were the defaults used in the concurrentbatchprocessor, too.
	return *new(component.Config)
}

// The default is configured in items, this value represents
// 60-100 concurrent batches.

// This enables by default an appropriate number of consumers
// Note for this exporter the consumer's role is to take from
// the queue and call into an Arrow stream. When the exporter
// falls back to OTLP, this is the number of concurrent OTLP
// exports.

// Default to zstd compression

// We almost read 0 bytes, so no need to tune ReadBufferSize.

// The `configgrpc` default is pick_first,
// which is not great for OTel Arrow exporters
// because it concentrates load at a single
// destination.

// Note the default payload compression is

func helperOptions(e exp) []exporterhelper.Option { _ = "STUB: not implemented"; return nil }

func gRPCName(desc grpc.ServiceDesc) string { _ = "STUB: not implemented"; return "" }

var (
	arrowTracesMethod  = gRPCName(arrowpb.ArrowTracesService_ServiceDesc)
	arrowMetricsMethod = gRPCName(arrowpb.ArrowMetricsService_ServiceDesc)
	arrowLogsMethod    = gRPCName(arrowpb.ArrowLogsService_ServiceDesc)
)

func createArrowTracesStream(conn *grpc.ClientConn) arrow.StreamClientFunc {
	_ = "STUB: not implemented"
	return *new(arrow.StreamClientFunc)
}

func createTracesExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

func createArrowMetricsStream(conn *grpc.ClientConn) arrow.StreamClientFunc {
	_ = "STUB: not implemented"
	return *new(arrow.StreamClientFunc)
}

func createMetricsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

func createArrowLogsStream(conn *grpc.ClientConn) arrow.StreamClientFunc {
	_ = "STUB: not implemented"
	return *new(arrow.StreamClientFunc)
}

func createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	cfg component.Config,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}
