// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ToKeyValueStringArguments[K any] struct {
	Target        ottl.PMapGetter[K]
	Delimiter     ottl.Optional[string]
	PairDelimiter ottl.Optional[string]
	SortOutput    ottl.Optional[bool]
}

func NewToKeyValueStringFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createToKeyValueStringFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toKeyValueString[K any](target ottl.PMapGetter[K], d, p ottl.Optional[string], s ottl.Optional[bool]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convertMapToKV converts a pcommon.Map to a key value string
func convertMapToKV(target pcommon.Map, delimiter, pairDelimiter string, sortOutput bool) string {
	_ = "STUB: not implemented"
	return ""
}

// Sort by keys

// Convert KV pairs

func buildKVString(k string, v pcommon.Value, delimiter, pairDelimiter string) string {
	_ = "STUB: not implemented"
	return ""
}

func escapeAndQuoteKV(s, delimiter, pairDelimiter string) string {
	_ = "STUB: not implemented"
	return ""
}
