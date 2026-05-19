// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type InsertXMLArguments[K any] struct {
	Target      ottl.StringGetter[K]
	XPath       string
	SubDocument ottl.StringGetter[K]
}

func NewInsertXMLFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createInsertXMLFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// insertXML returns a XML formatted string that is a result of inserting another XML document into
// the content of each selected target element.
func insertXML[K any](target ottl.StringGetter[K], xPath string, subGetter ottl.StringGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}

// AddChild updates c.NextSibling but not subDoc.FirstChild
// so we need to get the handle to it prior to the update.
