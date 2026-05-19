// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package alibabacloudlogserviceexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

// newLogsExporter return a new LogService logs exporter.
func newLogsExporter(set exporter.Settings, cfg component.Config) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

type logServiceLogsSender struct {
	logger *zap.Logger
	client logServiceClient
}

func (s *logServiceLogsSender) pushLogsData(
	_ context.Context,
	md plog.Logs,
) error {
	_ = "STUB: not implemented"
	return nil
}
