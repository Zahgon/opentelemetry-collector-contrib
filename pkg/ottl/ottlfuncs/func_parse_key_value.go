// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ParseKeyValueArguments[K any] struct {
	Target        ottl.StringGetter[K]
	Delimiter     ottl.Optional[string]
	PairDelimiter ottl.Optional[string]
}

func NewParseKeyValueFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createParseKeyValueFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseKeyValue[K any](target ottl.StringGetter[K], d, p ottl.Optional[string]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
