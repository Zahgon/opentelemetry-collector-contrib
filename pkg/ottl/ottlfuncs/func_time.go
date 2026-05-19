// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type TimeArguments[K any] struct {
	Time     ottl.StringGetter[K]
	Format   string
	Location ottl.Optional[string]
	Locale   ottl.Optional[string]
}

func NewTimeFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createTimeFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Time[K any](inputTime ottl.StringGetter[K], format string, location, locale ottl.Optional[string]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
