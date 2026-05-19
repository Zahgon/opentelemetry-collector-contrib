// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filtermatcher // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filtermatcher"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterconfig"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterset"
)

type instrumentationLibraryMatcher struct {
	Name    filterset.FilterSet
	Version filterset.FilterSet
}

// PropertiesMatcher allows matching a span against various span properties.
type PropertiesMatcher struct {
	// Instrumentation libraries to compare against
	libraries []instrumentationLibraryMatcher

	// The attribute values are stored in the internal format.
	attributes AttributesMatcher

	// The attribute values are stored in the internal format.
	resources AttributesMatcher
}

// NewMatcher creates a span Matcher that matches based on the given MatchProperties.
func NewMatcher(mp *filterconfig.MatchProperties) (PropertiesMatcher, error) {
	_ = "STUB: not implemented"
	return *new(PropertiesMatcher), nil
}

// Match matches a span or log to a set of properties.
func (mp *PropertiesMatcher) Match(attributes pcommon.Map, resource pcommon.Resource, library pcommon.InstrumentationScope) bool {
	_ = "STUB: not implemented"
	return false
}
