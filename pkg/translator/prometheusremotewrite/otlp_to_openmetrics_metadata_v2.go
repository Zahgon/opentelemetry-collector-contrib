// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewrite // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheusremotewrite"

import (
	writev2 "github.com/prometheus/prometheus/prompb/io/prometheus/write/v2"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func otelMetricTypeToPromMetricTypeV2(otelMetric pmetric.Metric) writev2.Metadata_MetricType {
	_ = "STUB: not implemented"
	// metric metadata can be used to support Prometheus types that don't exist
	// in OpenTelemetry.
	return *new(writev2.Metadata_MetricType)
}
