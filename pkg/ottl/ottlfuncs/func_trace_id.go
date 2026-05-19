// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

const traceIDFuncName = "TraceID"

type TraceIDArguments[K any] struct {
	Target ottl.ByteSliceLikeGetter[K]
}

func NewTraceIDFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createTraceIDFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func traceID[K any](target ottl.ByteSliceLikeGetter[K]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodeHexToTraceID(b []byte) (pcommon.TraceID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceID), nil
}
