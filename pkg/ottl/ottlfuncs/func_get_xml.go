// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type GetXMLArguments[K any] struct {
	Target ottl.StringGetter[K]
	XPath  string
}

func NewGetXMLFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createGetXMLFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getXML returns a XML formatted string that is a result of matching elements from the target XML.
func getXML[K any](target ottl.StringGetter[K], xPath string) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}

// get the value
