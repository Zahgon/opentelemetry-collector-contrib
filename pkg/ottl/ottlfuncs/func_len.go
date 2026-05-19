// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

const (
	typeError = `target arg must be of type string, []any, map[string]any, pcommon.Map, pcommon.Slice, pcommon.Value (of type String, Map, Slice) or a supported slice type from the plog, pmetric or ptrace packages`
)

type LenArguments[K any] struct {
	Target ottl.Getter[K]
}

func NewLenFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createLenFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func computeLen[K any](target ottl.Getter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
