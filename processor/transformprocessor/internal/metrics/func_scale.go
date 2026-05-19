// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/metrics"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
)

type ScaleArguments struct {
	Multiplier float64
	Unit       ottl.Optional[ottl.StringGetter[*ottlmetric.TransformContext]]
}

func newScaleMetricFactory() ottl.Factory[*ottlmetric.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func createScaleFunction(_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Scale(args ScaleArguments) (ottl.ExprFunc[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func scaleExemplar(ex *pmetric.Exemplar, multiplier float64) { _ = "STUB: not implemented"; return }

func scaleSummarySlice(values pmetric.SummaryDataPointSlice, multiplier float64) {
	_ = "STUB: not implemented"
	return
}

func scaleHistogram(datapoints pmetric.HistogramDataPointSlice, multiplier float64) {
	_ = "STUB: not implemented"
	return
}

func scaleMetric(points pmetric.NumberDataPointSlice, multiplier float64) {
	_ = "STUB: not implemented"
	return
}
