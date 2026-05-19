// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package alibabacloudlogserviceexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

// newTracesExporter return a new LogService trace exporter.
func newTracesExporter(set exporter.Settings, cfg component.Config) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

type logServiceTraceSender struct {
	logger *zap.Logger
	client logServiceClient
}

func (s *logServiceTraceSender) pushTraceData(
	_ context.Context,
	td ptrace.Traces,
) error {
	_ = "STUB: not implemented"
	return nil
}
