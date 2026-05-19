// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"encoding/xml"

	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ParseXMLArguments[K any] struct {
	Target ottl.StringGetter[K]
}

func NewParseXMLFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createParseXMLFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseXML returns a `pcommon.Map` struct that is a result of parsing the target string as XML
func parseXML[K any](target ottl.StringGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	return nil
}

// maxXMLElementDepth mirrors encoding/xml maxUnmarshalDepth to bound the
// recursion depth of XML parsing input.
const maxXMLElementDepth = 10_000

type xmlElement struct {
	depth      int
	tag        string
	attributes []xml.Attr
	text       string
	children   []xmlElement
}

// UnmarshalXML implements xml.Unmarshaler for xmlElement
func (a *xmlElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil
}

// End element means we've reached the end of parsing

// Strip leading/trailing spaces to ignore newlines and
// indentation in formatted XML

// ignore comments
// ignore processing instructions
// ignore directives

// intoMap converts and adds the xmlElement into the provided pcommon.Map.
func (a xmlElement) intoMap(m pcommon.Map) { _ = "STUB: not implemented"; return }
