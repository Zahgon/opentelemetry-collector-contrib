// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewrite // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"

import (
	"github.com/prometheus/prometheus/prompb"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func otelMetricTypeToPromMetricType(otelMetric pmetric.Metric) prompb.MetricMetadata_MetricType {
	_ = "STUB: not implemented"
	// metric metadata can be used to support Prometheus types that don't exist
	// in OpenTelemetry.
	return *new(prompb.MetricMetadata_MetricType)
}

func OtelMetricsToMetadata(md pmetric.Metrics, settings Settings) ([]*prompb.MetricMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
