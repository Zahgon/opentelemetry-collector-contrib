// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"
import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type UserAgentArguments[K any] struct {
	UserAgent ottl.StringGetter[K]
}

func NewUserAgentFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createUserAgentFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func userAgent[K any](userAgentSource ottl.StringGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented" //revive:disable-line:var-naming
	return nil
}
