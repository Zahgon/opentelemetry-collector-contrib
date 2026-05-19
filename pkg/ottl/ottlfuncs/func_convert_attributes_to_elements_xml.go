// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ConvertAttributesToElementsXMLArguments[K any] struct {
	Target ottl.StringGetter[K]
	XPath  ottl.Optional[string]
}

func NewConvertAttributesToElementsXMLFactory[K any]() ottl.Factory[K] {
	_ = "STUB: not implemented"
	return nil
}

func createConvertAttributesToElementsXMLFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// All attributes in the document

// convertAttributesToElementsXML returns a string that is a result of converting all attributes of the
// target XML into child elements. These new elements are added as the last child elements of the parent.
// e.g. <a foo="bar" hello="world"><b/></a> -> <a><hello>world</hello><foo>bar</foo><b/></a>
func convertAttributesToElementsXML[K any](target ottl.StringGetter[K], xPath string) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}
