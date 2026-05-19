// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver/internal/translator"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/identity"
)

type batcher struct {
	pmetric.Metrics

	resourceMetrics map[identity.Resource]pmetric.ResourceMetrics
	scopeMetrics    map[identity.Scope]pmetric.ScopeMetrics
	metrics         map[identity.Metric]pmetric.Metric
}

func newBatcher() batcher { _ = "STUB: not implemented"; return *new(batcher) }

// Dimensions stores the properties of the series that are needed in order
// to unique identify the series. This is needed in order to batch metrics by
// resource, scope, and datapoint attributes
type dimensions struct {
	name          string
	metricType    pmetric.MetricType
	resourceAttrs pcommon.Map
	scopeAttrs    pcommon.Map
	dpAttrs       pcommon.Map
	buildInfo     string
}

var metricTypeMap = map[string]pmetric.MetricType{
	"count":         pmetric.MetricTypeSum,
	"gauge":         pmetric.MetricTypeGauge,
	"rate":          pmetric.MetricTypeSum,
	"service_check": pmetric.MetricTypeGauge,
	"sketch":        pmetric.MetricTypeExponentialHistogram,
}

func parseSeriesProperties(name, metricType string, tags []string, host, version string, stringPool *StringPool) dimensions {
	_ = "STUB: not implemented"
	return *new(dimensions)
}

func (b batcher) Lookup(dim dimensions) (pmetric.Metric, identity.Metric) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric), *new(identity.Metric)
}

func (d dimensions) Resource() pcommon.Resource {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource)
}

// TODO(jesus.vazquez) review this copy

func (d dimensions) Scope() pcommon.InstrumentationScope {
	_ = "STUB: not implemented"
	return *new(pcommon.InstrumentationScope)
}

func (d dimensions) Metric() pmetric.Metric { _ = "STUB: not implemented"; return *new(pmetric.Metric) }
