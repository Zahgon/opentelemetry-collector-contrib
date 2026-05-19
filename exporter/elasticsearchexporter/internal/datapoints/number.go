// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datapoints // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/datapoints"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/elasticsearch"
)

type Number struct {
	pmetric.NumberDataPoint
	elasticsearch.MappingHintGetter
	metric pmetric.Metric
}

func NewNumber(metric pmetric.Metric, dp pmetric.NumberDataPoint) Number {
	_ = "STUB: not implemented"
	return *new(Number)
}

func (dp Number) Value() (pcommon.Value, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Value), nil
}

func (dp Number) DynamicTemplate(metric pmetric.Metric, mode DynamicTemplateMode) string {
	_ = "STUB: not implemented"
	return ""
}

// NumberDataPointValueTypeEmpty should already be discarded in numberToValue

// NumberDataPointValueTypeEmpty should already be discarded in numberToValue

func (Number) DocCount() uint64 { _ = "STUB: not implemented"; return 0 }

func (dp Number) Metric() pmetric.Metric { _ = "STUB: not implemented"; return *new(pmetric.Metric) }
