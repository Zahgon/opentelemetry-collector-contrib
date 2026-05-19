// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/antchfx/xmlquery"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ConvertTextToElementsXMLArguments[K any] struct {
	Target      ottl.StringGetter[K]
	XPath       ottl.Optional[string]
	ElementName ottl.Optional[string]
}

func NewConvertTextToElementsXMLFactory[K any]() ottl.Factory[K] {
	_ = "STUB: not implemented"
	return nil
}

func createConvertTextToElementsXMLFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convertTextToElementsXML returns a string that is a result of wrapping any extraneous text nodes with a dedicated element.
func convertTextToElementsXML[K any](target ottl.StringGetter[K], xPath, elementName string) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}

func convertTextToElementsForNode(parent *xmlquery.Node, elementName string, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

// ok
// ok

// Convert any child nodes and count text and element nodes.

// If there are no values to wrap, or if there is exactly one value OR one element, this node is all set.

// At this point, we either have multiple values, or a mix of values and elements.
// Either way, we need to wrap the values.

// Change this node into an element
