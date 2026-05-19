// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filterspan // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterspan"

import (
	"context"

	"go.opentelemetry.io/collector/featuregate"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterconfig"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filtermatcher"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterset"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
)

var useOTTLBridge = featuregate.GlobalRegistry().MustRegister(
	"filter.filterspan.useOTTLBridge",
	featuregate.StageAlpha,
	featuregate.WithRegisterDescription("When enabled, filterspan will convert filterspan configuration to OTTL and use filterottl evaluation"),
	featuregate.WithRegisterReferenceURL("https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/18642"),
)

// NewSkipExpr creates a BoolExpr that on evaluation returns true if a span should NOT be processed or kept.
// The logic determining if a span should be processed is based on include and exclude settings.
// Include properties are checked before exclude settings are checked.
func NewSkipExpr(mp *filterconfig.MatchConfig) (expr.BoolExpr[*ottlspan.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// propertiesMatcher allows matching a span against various span properties.
type propertiesMatcher struct {
	filtermatcher.PropertiesMatcher

	// Service names to compare to.
	serviceFilters filterset.FilterSet

	// Span names to compare to.
	nameFilters filterset.FilterSet

	// Span kinds to compare to
	kindFilters filterset.FilterSet
}

// newExpr creates a BoolExpr that matches based on the given MatchProperties.
func newExpr(mp *filterconfig.MatchProperties) (expr.BoolExpr[*ottlspan.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Eval matches a span and service to a set of properties.
// see filterconfig.MatchProperties for more details
func (mp *propertiesMatcher) Eval(_ context.Context, tCtx *ottlspan.TransformContext) (bool, error) {
	_ = "STUB: not implemented"
	// If a set of properties was not in the mp, all spans are considered to match on that property
	return false, nil
}

// Check resource and spans for service.name

// serviceNameForResource gets the service name for a specified Resource.
func serviceNameForResource(resource pcommon.Resource) string { _ = "STUB: not implemented"; return "" }
