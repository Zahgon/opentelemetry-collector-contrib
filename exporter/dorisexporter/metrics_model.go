// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dorisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dorisexporter"

import "go.opentelemetry.io/collector/pdata/pmetric"

type metricModel interface {
	metricType() pmetric.MetricType
	tableSuffix() string
	add(pm pmetric.Metric, dm *dMetric, e *metricsExporter) error
	size() int
	bytes() ([]byte, error)
	label() string
}

type metricModelCommon[T metric] struct {
	data []*T
	lbl  string
}

func (m *metricModelCommon[T]) size() int { _ = "STUB: not implemented"; return 0 }

func (m *metricModelCommon[T]) bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *metricModelCommon[T]) label() string {
	_ = "STUB: not implemented"

	// dMetric Basic Metric
	return ""
}

type dMetric struct {
	ServiceName        string         `json:"service_name"`
	ServiceInstanceID  string         `json:"service_instance_id"`
	MetricName         string         `json:"metric_name"`
	MetricDescription  string         `json:"metric_description"`
	MetricUnit         string         `json:"metric_unit"`
	ResourceAttributes map[string]any `json:"resource_attributes"`
	ScopeName          string         `json:"scope_name"`
	ScopeVersion       string         `json:"scope_version"`
}

// dExemplar Exemplar to Doris
type dExemplar struct {
	FilteredAttributes map[string]any `json:"filtered_attributes"`
	Timestamp          string         `json:"timestamp"`
	Value              float64        `json:"value"`
	SpanID             string         `json:"span_id"`
	TraceID            string         `json:"trace_id"`
}
