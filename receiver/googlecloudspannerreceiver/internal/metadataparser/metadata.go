// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadataparser // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadataparser"

import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"

type Metadata struct {
	Name                string   `yaml:"name"`
	Query               string   `yaml:"query"`
	MetricNamePrefix    string   `yaml:"metric_name_prefix"`
	TimestampColumnName string   `yaml:"timestamp_column_name"`
	HighCardinality     bool     `yaml:"high_cardinality"`
	Labels              []Label  `yaml:"labels"`
	Metrics             []Metric `yaml:"metrics"`
}

func (m Metadata) MetricsMetadata() (*metadata.MetricsMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m Metadata) toLabelValuesMetadata() ([]metadata.LabelValueMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m Metadata) toMetricValuesMetadata() ([]metadata.MetricValueMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
