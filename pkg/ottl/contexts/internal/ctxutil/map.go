// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxutil"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func GetMapValue[K any](ctx context.Context, tCtx K, m pcommon.Map, keys []ottl.Key[K]) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func SetMapValue[K any](ctx context.Context, tCtx K, m pcommon.Map, keys []ottl.Key[K], val any) error {
	_ = "STUB: not implemented"
	return nil
}

func GetMapKeyName[K any](ctx context.Context, tCtx K, key ottl.Key[K]) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func FetchValueFromExpression[K any, T int64 | string](ctx context.Context, tCtx K, key ottl.Key[K]) (*T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SetMap(target pcommon.Map, val any) error { _ = "STUB: not implemented"; return nil }

func GetMap(val any) (pcommon.Map, error) { _ = "STUB: not implemented"; return *new(pcommon.Map), nil }
