// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/metrics"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
)

const sumCountName = "extract_count_metric"

type extractCountMetricArguments struct {
	Monotonic bool
	Suffix    ottl.Optional[string]
}

func newExtractCountMetricFactory() ottl.Factory[*ottlmetric.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func createExtractCountMetricFunction(_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractCountMetric(monotonic bool, suffix ottl.Optional[string]) (ottl.ExprFunc[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use the default unit as the original metric unit does not apply to the 'count' field

func addCountDataPoint(dataPoint sumCountDataPoint, destination pmetric.NumberDataPointSlice) {
	_ = "STUB: not implemented"
	return
}
