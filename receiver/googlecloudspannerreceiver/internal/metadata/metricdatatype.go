// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"

import "go.opentelemetry.io/collector/pdata/pmetric"

type MetricType interface {
	MetricType() pmetric.MetricType
	AggregationTemporality() pmetric.AggregationTemporality
	IsMonotonic() bool
}

type metricValueDataType struct {
	dataType               pmetric.MetricType
	aggregationTemporality pmetric.AggregationTemporality
	isMonotonic            bool
}

func NewMetricType(dataType pmetric.MetricType, aggregationTemporality pmetric.AggregationTemporality,
	isMonotonic bool,
) MetricType {
	_ = "STUB: not implemented"
	return *new(MetricType)
}

func (metricValueDataType metricValueDataType) MetricType() pmetric.MetricType {
	_ = "STUB: not implemented"
	return *new(pmetric.MetricType)
}

func (metricValueDataType metricValueDataType) AggregationTemporality() pmetric.AggregationTemporality {
	_ = "STUB: not implemented"
	return *new(pmetric.AggregationTemporality)
}

func (metricValueDataType metricValueDataType) IsMonotonic() bool {
	_ = "STUB: not implemented"
	return false
}
