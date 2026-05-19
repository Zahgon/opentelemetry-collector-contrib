// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ExtractPatternsArguments[K any] struct {
	Target  ottl.StringGetter[K]
	Pattern ottl.StringGetter[K]
}

func NewExtractPatternsFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createExtractPatternsFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractPatterns[K any](target, pattern ottl.StringGetter[K]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip whole match
