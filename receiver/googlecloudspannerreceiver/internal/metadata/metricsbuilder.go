// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/filter"
)

type MetricsBuilder interface {
	Build(dataPoints []*MetricsDataPoint) (pmetric.Metrics, error)
	Shutdown() error
}

type metricsFromDataPointBuilder struct {
	filterResolver filter.ItemFilterResolver
}

func NewMetricsFromDataPointBuilder(filterResolver filter.ItemFilterResolver) MetricsBuilder {
	_ = "STUB: not implemented"
	return *new(MetricsBuilder)
}

func (b *metricsFromDataPointBuilder) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (b *metricsFromDataPointBuilder) Build(dataPoints []*MetricsDataPoint) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (b *metricsFromDataPointBuilder) groupAndFilter(dataPoints []*MetricsDataPoint) (map[MetricsDataPointKey][]*MetricsDataPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Cardinality filtering

func (b *metricsFromDataPointBuilder) filter(metricName string, dataPoints []*MetricsDataPoint) ([]*MetricsDataPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Creating new slice instead of removing elements from source slice because removing by value is not efficient operation.
// Need to use such approach for preserving data points order.
