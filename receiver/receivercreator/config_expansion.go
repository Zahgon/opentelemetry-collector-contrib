// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package receivercreator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/receivercreator"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
)

// evalBackticksInConfigValue expands any expressions within backticks inside configValue
// using variables from env.
//
// Note that when evaluating multiple expressions that the expanded result is always
// a string. For instance:
//
//	`true``false` -> "truefalse"
//
// However if there is only one expansion then the expanded result will keep the type
// of the expression. For instance:
//
//	`"secure" in pod.labels` -> true (boolean)
func evalBackticksInConfigValue(configValue string, env observer.EndpointEnv) (any, error) {
	_ = "STUB: not implemented"
	// Tracks index into configValue where an expression (backtick) begins. -1 is unset.
	return *new(any), nil
}

// Accumulate expanded string.

// Accumulate results of calls to eval for use at the end to return well-typed
// results if possible.

// Loop through configValue one rune at a time using exprStartIndex to keep track of
// inside or outside of expressions.

// Opening backtick encountered, expression starts one after current index.

// Closing backtick encountered, evaluate the expression.

// If expression has no text inside it return an error.

// Reset start index since this expression just closed.

// Should always be a closing backtick if it's balanced.

// If there was only one expansion and it is equal to the full output string return the expansion
// itself so that it retains the type returned by the expression. Might be a bool, int, etc.
// instead of a string.

// expandConfig will walk the provided user config and expand any `backticked` content
// with associated observer.EndpointEnv values.
func expandConfig(cfg userConfigMap, env observer.EndpointEnv) (userConfigMap, error) {
	_ = "STUB: not implemented"
	return *new(userConfigMap), nil
}

// expandAny recursively expands any expressions in backticks inside values of input using
// env as variables available within the expression, returning a copy of input
func expandAny(input any, env observer.EndpointEnv) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// expanded strings aren't guaranteed to remain them, so we
// coerce to any for shared []any expansion path
