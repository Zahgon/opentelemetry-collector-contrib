// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/metrics"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
)

type convertGaugeToSumArguments struct {
	StringAggTemp string
	Monotonic     bool
}

func newConvertGaugeToSumFactory() ottl.Factory[*ottlmetric.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func createConvertGaugeToSumFunction(_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertGaugeToSum(stringAggTemp string, monotonic bool) (ottl.ExprFunc[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Setting the data type removed all the data points, so we must move them back to the metric.
