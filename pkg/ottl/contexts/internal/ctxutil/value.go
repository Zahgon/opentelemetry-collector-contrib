// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxutil"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func SetValue(value pcommon.Value, val any) error { _ = "STUB: not implemented"; return nil }

func getIndexableValue[K any](ctx context.Context, tCtx K, value pcommon.Value, keys []ottl.Key[K]) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func SetIndexableValue[K any](ctx context.Context, tCtx K, currentValue pcommon.Value, val any, keys []ottl.Key[K]) error {
	_ = "STUB: not implemented"
	return nil
}
