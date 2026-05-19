// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxotelcol // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxotelcol"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

var errOTelColContextDisabled = errors.New("OTTL `otelcol` context requires the `ottl.contexts.enableOTelColContext` feature gate to be enabled")

func PathGetSetter[K any](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertStringArrToValueSlice(vals []string) pcommon.Value {
	_ = "STUB: not implemented"
	return *new(pcommon.Value)
}

func getIndexableValueFromStringArr[K any](ctx context.Context, tCtx K, keys []ottl.Key[K], strSlice []string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
