// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func uuidV7[K any]() (ottl.ExprFunc[K], error) { _ = "STUB: not implemented"; return nil, nil }

func createUUIDv7Function[K any](_ ottl.FunctionContext, _ ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewUUIDv7Factory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }
