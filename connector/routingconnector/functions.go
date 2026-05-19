// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package routingconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
)

func createRouteFunction[K any](ottl.FunctionContext, ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func standardFunctions[K any]() map[string]ottl.Factory[K] {
	_ = "STUB: not implemented"
	// standard converters do not transform data, so we can safely use them
	return nil
}

func spanFunctions() map[string]ottl.Factory[*ottlspan.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}
