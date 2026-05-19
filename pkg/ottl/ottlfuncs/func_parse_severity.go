// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

const (
	// http2xx is a special key that is represents a range from 200 to 299
	http2xx = "2xx"

	// http3xx is a special key that is represents a range from 300 to 399
	http3xx = "3xx"

	// http4xx is a special key that is represents a range from 400 to 499
	http4xx = "4xx"

	// http5xx is a special key that is represents a range from 500 to 599
	http5xx = "5xx"

	minKey = "min"
	maxKey = "max"

	rangeKey  = "range"
	equalsKey = "equals"
)

type ParseSeverityArguments[K any] struct {
	Target  ottl.Getter[K]
	Mapping ottl.PMapGetter[K]
}

func NewParseSeverityFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createParseSeverityFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseSeverity[K any](target ottl.Getter[K], mapping ottl.PMapGetter[K]) ottl.ExprFunc[K] {
	_ = "STUB: not implemented"
	// retrieve the mapping as a literal PMap and use it for all evaluations
	return nil
}

// convert the mapping to criteria objects to validate its structure

func evaluateSeverity(value any, severities map[string]criteriaSet) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type criteriaSet []criteria

func (cs criteriaSet) evaluate(value any) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type criteria struct {
	Equals []string
	Range  *valueRange
}

type valueRange struct {
	Min int64
	Max int64
}

func (c *criteria) evaluate(value any) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func newCriteriaFromMap(m map[string]any) (*criteria, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
