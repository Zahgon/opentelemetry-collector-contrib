// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

const spanIDFuncName = "SpanID"

type SpanIDArguments[K any] struct {
	Target ottl.ByteSliceLikeGetter[K]
}

func NewSpanIDFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createSpanIDFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func spanID[K any](target ottl.ByteSliceLikeGetter[K]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodeHexToSpanID(b []byte) (pcommon.SpanID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID), nil
}
