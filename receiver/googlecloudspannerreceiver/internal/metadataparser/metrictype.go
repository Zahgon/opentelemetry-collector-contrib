// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadataparser // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadataparser"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"
)

type MetricDataType string

const (
	UnknownMetricDataType MetricDataType = "unknown"
	GaugeMetricDataType   MetricDataType = "gauge"
	SumMetricDataType     MetricDataType = "sum"
)

type AggregationType string

const (
	UnknownAggregationType    AggregationType = "unknown"
	DeltaAggregationType      AggregationType = "delta"
	CumulativeAggregationType AggregationType = "cumulative"
)

type MetricType struct {
	DataType    MetricDataType  `yaml:"type"`
	Aggregation AggregationType `yaml:"aggregation"`
	Monotonic   bool            `yaml:"monotonic"`
}

func (metricType MetricType) dataType() (pmetric.MetricType, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.MetricType), nil
}

func (metricType MetricType) aggregationTemporality() (pmetric.AggregationTemporality, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.AggregationTemporality), nil
}

func (metricType MetricType) toMetricType() (metadata.MetricType, error) {
	_ = "STUB: not implemented"
	return *new(metadata.MetricType), nil
}
