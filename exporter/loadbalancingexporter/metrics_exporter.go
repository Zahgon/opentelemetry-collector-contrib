// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loadbalancingexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter/internal/metadata"
)

var _ exporter.Metrics = (*metricExporterImp)(nil)

type metricExporterImp struct {
	loadBalancer *loadBalancer
	routingKey   routingKey
	routingAttrs []string

	logger     *zap.Logger
	stopped    bool
	shutdownWg sync.WaitGroup
	telemetry  *metadata.TelemetryBuilder
}

func newMetricsExporter(params exporter.Settings, cfg component.Config) (*metricExporterImp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default case for empty routing key

func (*metricExporterImp) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (e *metricExporterImp) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *metricExporterImp) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *metricExporterImp) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// Now assign each batch to an exporter, and merge as we go

func splitMetricsByResourceServiceName(md pmetric.Metrics) (map[string]pmetric.Metrics, []error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func splitMetricsByResourceID(md pmetric.Metrics) map[string]pmetric.Metrics {
	_ = "STUB: not implemented"
	return nil
}

func splitMetricsByMetricName(md pmetric.Metrics) map[string]pmetric.Metrics {
	_ = "STUB: not implemented"
	return nil
}

func splitMetricsByStreamID(md pmetric.Metrics) map[string]pmetric.Metrics {
	_ = "STUB: not implemented"
	return nil
}

func splitMetricsByAttributes(md pmetric.Metrics, attrs []string) map[string]pmetric.Metrics {
	_ = "STUB: not implemented"
	return nil
}

// All split attributes are on resource, so no per-scope/datapoint keying.

// All split attributes are on resource/scope, so no per-datapoint keying.

func forEachMetricDataPoint(rm pmetric.ResourceMetrics, sm pmetric.ScopeMetrics, m pmetric.Metric, fn func(dp attrPoint, md pmetric.Metrics)) {
	_ = "STUB: not implemented"
	return
}

type attrPoint interface {
	Attributes() pcommon.Map
}

func appendMetricsByKey(results map[string]pmetric.Metrics, key string, mds pmetric.Metrics) {
	_ = "STUB: not implemented"
	return
}

func cloneMetricWithoutType(rm pmetric.ResourceMetrics, sm pmetric.ScopeMetrics, m pmetric.Metric) (md pmetric.Metrics, mClone pmetric.Metric) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), *new(pmetric.Metric)
}
