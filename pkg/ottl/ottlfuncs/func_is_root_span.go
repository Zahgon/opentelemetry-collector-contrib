// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
)

// Deprecated: [v0.142.0] use NewIsRootSpanFactoryNew.
func NewIsRootSpanFactory() ottl.Factory[ottlspan.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func createIsRootSpanFunctionLegacy(_ ottl.FunctionContext, _ ottl.Arguments) (ottl.ExprFunc[ottlspan.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewIsRootSpanFactoryNew() ottl.Factory[*ottlspan.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func createIsRootSpanFunction(_ ottl.FunctionContext, _ ottl.Arguments) (ottl.ExprFunc[*ottlspan.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isRootSpan() (ottl.ExprFunc[*ottlspan.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
