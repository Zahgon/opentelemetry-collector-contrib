// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadataparser // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadataparser"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"
)

type Metric struct {
	Label    `yaml:",inline"`
	DataType MetricType `yaml:"data"`
	Unit     string     `yaml:"unit"`
}

func (metric Metric) toMetricValueMetadata() (metadata.MetricValueMetadata, error) {
	_ = "STUB: not implemented"
	return *new(metadata.MetricValueMetadata), nil
}
