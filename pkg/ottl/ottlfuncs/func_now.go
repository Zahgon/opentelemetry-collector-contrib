// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func now[K any]() (ottl.ExprFunc[K], error) { _ = "STUB: not implemented"; return nil, nil }

func createNowFunction[K any](_ ottl.FunctionContext, _ ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewNowFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }
