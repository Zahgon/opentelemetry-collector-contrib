// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/antchfx/xmlquery"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ParseSimplifiedXMLArguments[K any] struct {
	Target ottl.StringGetter[K]
}

func NewParseSimplifiedXMLFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createParseSimplifiedXMLFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The `ParseSimplifiedXML` Converter returns a `pcommon.Map` struct that is the result of parsing the target
// string without preservation of attributes or extraneous text content.
func parseSimplifiedXML[K any](target ottl.StringGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}

func parseElement(parent *xmlquery.Node, parentMap *pcommon.Map, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

// Count the number of each element tag so we know whether it will be a member of a slice or not

// Convert the children, now knowing whether they will be a member of a slice or not

// Slice of the same element

// Get or create the slice of children

// Add the child's text content to the slice

// Parse the child to make sure there's something to add

// Child will be a map

func leafValueFromElement(node *xmlquery.Node) string {
	_ = "STUB: not implemented"
	// First check if there are any child elements. If there are, ignore any extraneous text.
	return ""
}

// No child elements, so return the first text or CDATA content
