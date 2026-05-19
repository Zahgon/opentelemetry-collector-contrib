// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/metrics"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
)

func newConvertSumToGaugeFactory() ottl.Factory[*ottlmetric.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func createConvertSumToGaugeFunction(_ ottl.FunctionContext, _ ottl.Arguments) (ottl.ExprFunc[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertSumToGauge() (ottl.ExprFunc[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Setting the data type removed all the data points, so we must copy them back to the metric.
