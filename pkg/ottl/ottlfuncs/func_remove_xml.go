// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/antchfx/xmlquery"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type RemoveXMLArguments[K any] struct {
	Target ottl.StringGetter[K]
	XPath  string
}

func NewRemoveXMLFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createRemoveXMLFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// removeXML returns a XML formatted string that is a result of removing all matching nodes from the target XML.
// This currently supports removal of elements, attributes, text values, comments, and CharData.
func removeXML[K any](target ottl.StringGetter[K], xPath string) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}

func validateXPath(xPath string) error { _ = "STUB: not implemented"; return nil }

// Aside from parsing the XML document, this function also ensures that
// the XML declaration is included in the result only if it was present in
// the original document.
func parseNodesXML(targetVal string) (*xmlquery.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
