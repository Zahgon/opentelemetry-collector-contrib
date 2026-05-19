// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"
import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type IsInCIDRArguments[K any] struct {
	Target   ottl.StringGetter[K]
	Networks []ottl.StringGetter[K]
}

func NewIsInCIDRFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createIsInCIDRFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isInCIDR[K any](target ottl.StringGetter[K], networks []ottl.StringGetter[K]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	// Check if all networks are literals and pre-parse them if so.
	return nil, nil
}

// Use pre-parsed networks for literal values.

// Parse networks at runtime for dynamic values.
