// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filtermatcher // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filtermatcher"

import (
	"errors"

	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterconfig"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterset"
)

type AttributesMatcher []AttributeMatcher

type valueIdentifier struct {
	value     pcommon.Value
	valueHash [16]byte
}

// AttributeMatcher is a attribute key/value pair to match to.
type AttributeMatcher struct {
	Key string
	// If both AttributeValue and StringFilter are nil only check for key existence.
	AttributeValue *valueIdentifier
	// StringFilter is needed to match against a regular expression
	StringFilter filterset.FilterSet
}

var errUnexpectedAttributeType = errors.New("unexpected attribute type")

func NewAttributesMatcher(config filterset.Config, attributes []filterconfig.Attribute) (AttributesMatcher, error) {
	_ = "STUB: not implemented"
	// Convert attribute values from mp representation to in-memory representation.
	return *new(AttributesMatcher), nil
}

// Match attributes specification against a span/log.
func (ma AttributesMatcher) Match(attrs pcommon.Map) bool {
	_ = "STUB: not implemented"
	// If there are no attributes to match against, the span/log matches.
	return false
}

// At this point, it is expected of the span/log to have attributes because of
// len(ma) != 0. This means for spans/logs with no attributes, it does not match.

// Check that all expected properties are set.

func attributeStringValue(attr pcommon.Value) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func attributeValueMatch(vi *valueIdentifier, val pcommon.Value) bool {
	_ = "STUB: not implemented"
	return false
}

// Use hash for other complex data types.
