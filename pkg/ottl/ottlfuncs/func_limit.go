// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type LimitArguments[K any] struct {
	Target       ottl.PMapGetSetter[K]
	Limit        int64
	PriorityKeys []string
}

func NewLimitFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createLimitFunction[K any](fCtx ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func limit[K any](target ottl.PMapGetSetter[K], limit int64, priorityKeys []string, logger *zap.Logger) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
