// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/metrics"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
)

type copyMetricArguments struct {
	Name        ottl.Optional[ottl.StringGetter[*ottlmetric.TransformContext]]
	Description ottl.Optional[ottl.StringGetter[*ottlmetric.TransformContext]]
	Unit        ottl.Optional[ottl.StringGetter[*ottlmetric.TransformContext]]
}

func newCopyMetricFactory() ottl.Factory[*ottlmetric.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func createCopyMetricFunction(_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func copyMetric(name, desc, unit ottl.Optional[ottl.StringGetter[*ottlmetric.TransformContext]]) (ottl.ExprFunc[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
