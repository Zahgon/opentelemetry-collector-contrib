// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunkhecexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	translator "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/splunk"
)

const (
	metricsPrefix              = "otelcol_exporter_splunkhec_"
	defaultHBSentMetricsName   = metricsPrefix + "heartbeats_sent"
	defaultHBFailedMetricsName = metricsPrefix + "heartbeats_failed"
)

type heartbeater struct {
	hbDoneChan chan struct{}
}

func getMetricsName(overrides map[string]string, metricName string) string {
	_ = "STUB: not implemented"
	return ""
}

func newHeartbeater(config *Config, buildInfo component.BuildInfo, pushLogFn func(ctx context.Context, ld plog.Logs) error, meter metric.Meter) *heartbeater {
	_ = "STUB: not implemented"
	return nil
}

func (h *heartbeater) shutdown() { _ = "STUB: not implemented"; return }

func (*heartbeater) sendHeartbeat(config *Config, buildInfo component.BuildInfo, pushLogFn func(ctx context.Context, ld plog.Logs) error) error {
	_ = "STUB: not implemented"
	return nil
}

// there is only use case for open census metrics recording for now. Extend to use open telemetry in the future.
func observe(heartbeatsSent, heartbeatsFailed metric.Int64Counter, attrs attribute.Set, err error) {
	_ = "STUB: not implemented"
	return
}

func generateHeartbeatLog(hecToOtelAttrs translator.HecToOtelAttrs, buildInfo component.BuildInfo) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}
