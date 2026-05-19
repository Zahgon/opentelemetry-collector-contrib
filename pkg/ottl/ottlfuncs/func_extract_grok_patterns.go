// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ExtractGrokPatternsArguments[K any] struct {
	Target             ottl.StringGetter[K]
	Pattern            ottl.StringGetter[K]
	NamedCapturesOnly  ottl.Optional[bool]
	PatternDefinitions ottl.Optional[[]string]
}

func NewExtractGrokPatternsFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createExtractGrokPatternsFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractGrokPatterns[K any](target, pattern ottl.StringGetter[K], nco ottl.Optional[bool], patternDefinitions ottl.Optional[[]string]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// split pattern in format key=val

// keep whole string 20 characters long including ...
