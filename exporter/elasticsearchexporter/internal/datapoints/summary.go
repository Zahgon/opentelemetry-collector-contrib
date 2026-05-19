// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datapoints // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/datapoints"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/elasticsearch"
)

type Summary struct {
	pmetric.SummaryDataPoint
	elasticsearch.MappingHintGetter
	metric pmetric.Metric
}

func NewSummary(metric pmetric.Metric, dp pmetric.SummaryDataPoint) Summary {
	_ = "STUB: not implemented"
	return *new(Summary)
}

func (dp Summary) Value() (pcommon.Value, error) {
	_ = "STUB: not implemented"
	// TODO: Add support for quantiles
	// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/34561
	return *new(pcommon.Value), nil
}

func (Summary) DynamicTemplate(_ pmetric.Metric, mode DynamicTemplateMode) string {
	_ = "STUB: not implemented"
	return ""
}

func (dp Summary) DocCount() uint64 { _ = "STUB: not implemented"; return 0 }

func (dp Summary) Metric() pmetric.Metric { _ = "STUB: not implemented"; return *new(pmetric.Metric) }
