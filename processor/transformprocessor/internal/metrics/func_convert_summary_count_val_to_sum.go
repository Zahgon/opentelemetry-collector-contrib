// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/metrics"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
)

type convertSummaryCountValToSumArguments struct {
	StringAggTemp string
	Monotonic     bool
	Suffix        ottl.Optional[string]
}

func newConvertSummaryCountValToSumFactory() ottl.Factory[*ottldatapoint.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func createConvertSummaryCountValToSumFunction(_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[*ottldatapoint.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertSummaryCountValToSum(stringAggTemp string, monotonic bool, suffix ottl.Optional[string]) (ottl.ExprFunc[*ottldatapoint.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
