// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"context"
	"regexp"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

const invalidRegexErrMsg = "the regex pattern supplied to %s '%q' is not a valid pattern: %w"

type dynamicRegex[K any] struct {
	funcName string
	getter   ottl.StringGetter[K]
	value    *regexp.Regexp
}

// newDynamicRegex creates a new dynamicRegex instance that handles both literal and dynamic regex patterns.
// If the pattern is a literal value, it compiles the regex immediately and caches it.
// If the pattern is dynamic, it defers compilation until runtime.
func newDynamicRegex[K any](funcName string, getter ottl.StringGetter[K]) (*dynamicRegex[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// compile returns a compiled regex pattern. If the pattern was pre-compiled (literal), it returns the cached version.
// Otherwise, it retrieves the pattern value at runtime and compiles it.
func (l *dynamicRegex[K]) compile(ctx context.Context, tCtx K) (*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
