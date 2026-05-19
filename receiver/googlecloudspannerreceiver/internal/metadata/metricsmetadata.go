// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"

import (
	"time"

	"cloud.google.com/go/spanner"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/datasource"
)

type MetricsMetadataType int32

const (
	MetricsMetadataTypeCurrentStats MetricsMetadataType = iota
	MetricsMetadataTypeIntervalStats
)

type MetricsMetadata struct {
	Name                string
	Query               string
	MetricNamePrefix    string
	TimestampColumnName string
	HighCardinality     bool
	// In addition to common metric labels
	QueryLabelValuesMetadata  []LabelValueMetadata
	QueryMetricValuesMetadata []MetricValueMetadata
}

func (metadata *MetricsMetadata) timestamp(row *spanner.Row) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (metadata *MetricsMetadata) toLabelValues(row *spanner.Row) ([]LabelValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toLabelValue(labelValueMetadata LabelValueMetadata, row *spanner.Row) (LabelValue, error) {
	_ = "STUB: not implemented"
	return *new(LabelValue), nil
}

func (metadata *MetricsMetadata) toMetricValues(row *spanner.Row) ([]MetricValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toMetricValue(metricValueMetadata MetricValueMetadata, row *spanner.Row) (MetricValue, error) {
	_ = "STUB: not implemented"
	return *new(MetricValue), nil
}

func (metadata *MetricsMetadata) RowToMetricsDataPoints(databaseID *datasource.DatabaseID, row *spanner.Row) ([]*MetricsDataPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reading labels

// Reading metrics

func (metadata *MetricsMetadata) toMetricsDataPoints(databaseID *datasource.DatabaseID, timestamp time.Time,
	labelValues []LabelValue, metricValues []MetricValue,
) []*MetricsDataPoint {
	_ = "STUB: not implemented"
	return nil
}

func (metadata *MetricsMetadata) MetadataType() MetricsMetadataType {
	_ = "STUB: not implemented"
	return *new(MetricsMetadataType)
}
