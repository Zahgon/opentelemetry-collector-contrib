// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsxrayexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray/telemetry"
)

const (
	maxSegmentsPerPut = int(50) // limit imposed by PutTraceSegments API
)

// newTracesExporter creates an exporter.Traces that converts to an X-Ray PutTraceSegments
// request and then posts the request to the configured region's X-Ray endpoint.
func newTracesExporter(ctx context.Context, cfg *Config, set exporter.Settings, registry telemetry.Registry) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

// record error

func extractResourceSpans(config component.Config, logger *zap.Logger, td ptrace.Traces) []string {
	_ = "STUB: not implemented"
	return nil
}

func wrapErrorIfBadRequest(err error) error { _ = "STUB: not implemented"; return nil }
