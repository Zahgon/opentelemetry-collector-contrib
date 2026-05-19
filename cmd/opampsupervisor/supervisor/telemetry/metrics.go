// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package telemetry

import (
	"context"

	"go.opentelemetry.io/otel/metric"
)

const (
	CollectorHealthStatusMetric = "supervisor.agent.health_status"
)

type Metrics struct {
	collectorHealthStatusMetric metric.Int64UpDownCounter
	healthStatus                bool
}

func NewMetrics(meterProvider metric.MeterProvider) (*Metrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initialize metrics to 0 to ensure they are exported

func (m *Metrics) SetCollectorHealthStatus(ctx context.Context, healthy bool) {
	_ = "STUB: not implemented"
	// Only update the metric if the health status has changed
	return
}
