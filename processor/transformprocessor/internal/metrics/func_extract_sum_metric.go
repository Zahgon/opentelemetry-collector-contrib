// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/metrics"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
)

const sumFuncName = "extract_sum_metric"

type extractSumMetricArguments struct {
	Monotonic bool
	Suffix    ottl.Optional[string]
}

func newExtractSumMetricFactory() ottl.Factory[*ottlmetric.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func createExtractSumMetricFunction(_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractSumMetric(monotonic bool, suffix ottl.Optional[string]) (ottl.ExprFunc[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// note that unlike Histograms, the Sum field is required for Summaries

func addSumDataPoint(dataPoint sumCountDataPoint, destination pmetric.NumberDataPointSlice) {
	_ = "STUB: not implemented"
	return
}

func getAggregationTemporality(metric pmetric.Metric) pmetric.AggregationTemporality {
	_ = "STUB: not implemented"
	return *new(pmetric.AggregationTemporality)
}

// Summaries don't have an aggregation temporality, but they *should* be cumulative based on the Openmetrics spec.
// This should become an optional argument once those are available in OTTL.

func invalidMetricTypeError(name string, metric pmetric.Metric) error {
	_ = "STUB: not implemented"
	return nil
}
