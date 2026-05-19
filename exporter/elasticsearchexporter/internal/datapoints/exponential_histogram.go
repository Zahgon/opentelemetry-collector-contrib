// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datapoints // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/datapoints"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/elasticsearch"
)

type ExponentialHistogram struct {
	pmetric.ExponentialHistogramDataPoint
	elasticsearch.MappingHintGetter
	metric pmetric.Metric
}

func NewExponentialHistogram(metric pmetric.Metric, dp pmetric.ExponentialHistogramDataPoint) ExponentialHistogram {
	_ = "STUB: not implemented"
	return *new(ExponentialHistogram)
}

func (dp ExponentialHistogram) Value() (pcommon.Value, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Value), nil
}

func (dp ExponentialHistogram) DynamicTemplate(_ pmetric.Metric, mode DynamicTemplateMode) string {
	_ = "STUB: not implemented"
	return ""
}

func (dp ExponentialHistogram) DocCount() uint64 { _ = "STUB: not implemented"; return 0 }

func (dp ExponentialHistogram) Metric() pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}
