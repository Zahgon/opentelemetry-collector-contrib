// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type TruncateTimeArguments[K any] struct {
	Time     ottl.TimeGetter[K]
	Duration ottl.DurationGetter[K]
}

func NewTruncateTimeFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createTruncateTimeFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func TruncateTime[K any](inputTime ottl.TimeGetter[K], inputDuration ottl.DurationGetter[K]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
