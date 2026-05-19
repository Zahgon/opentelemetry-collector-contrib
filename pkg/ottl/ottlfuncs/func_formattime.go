// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type FormatTimeArguments[K any] struct {
	Time   ottl.TimeGetter[K]
	Format string
}

func NewFormatTimeFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createFormatTimeFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FormatTime[K any](timeValue ottl.TimeGetter[K], format string) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
