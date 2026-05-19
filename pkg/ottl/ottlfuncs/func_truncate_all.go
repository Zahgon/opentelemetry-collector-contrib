// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type TruncateAllArguments[K any] struct {
	Target   ottl.PMapGetSetter[K]
	Limit    int64
	Utf8Safe ottl.Optional[bool]
}

func NewTruncateAllFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createTruncateAllFunction[K any](fCtx ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func TruncateAll[K any](target ottl.PMapGetSetter[K], limit int64, utf8Safe ottl.Optional[bool], logger *zap.Logger) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Back up to a valid UTF-8 boundary if we're in the middle of a rune
