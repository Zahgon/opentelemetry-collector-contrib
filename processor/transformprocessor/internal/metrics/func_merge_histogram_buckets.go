// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/metrics"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
)

type mergeHistogramBucketsArguments struct {
	Bound float64
}

func newMergeHistogramBucketsFactory() ottl.Factory[*ottldatapoint.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func createMergeHistogramBucketsFunction(_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[*ottldatapoint.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeHistogramBuckets(bound float64) (ottl.ExprFunc[*ottldatapoint.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mergeHistogramBucketsFromDataPoint(dp pmetric.HistogramDataPoint, bound float64) {
	_ = "STUB: not implemented"
	return
}

// findBoundIndex finds the index of a target bound in the bounds slice with epsilon tolerance
func findBoundIndex(bounds *[]float64, target float64) int { _ = "STUB: not implemented"; return 0 }
