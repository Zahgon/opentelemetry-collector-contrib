// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package alibabacloudlogserviceexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

// newMetricsExporter return a new LogService metrics exporter.
func newMetricsExporter(set exporter.Settings, cfg component.Config) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

type logServiceMetricsSender struct {
	logger *zap.Logger
	client logServiceClient
}

func (s *logServiceMetricsSender) pushMetricsData(
	_ context.Context,
	md pmetric.Metrics,
) error {
	_ = "STUB: not implemented"
	return nil
}
