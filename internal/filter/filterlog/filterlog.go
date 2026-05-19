// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filterlog // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterlog"

import (
	"context"

	"go.opentelemetry.io/collector/featuregate"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterconfig"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filtermatcher"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterset"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
)

var useOTTLBridge = featuregate.GlobalRegistry().MustRegister(
	"filter.filterlog.useOTTLBridge",
	featuregate.StageAlpha,
	featuregate.WithRegisterDescription("When enabled, filterlog will convert filterlog configuration to OTTL and use filterottl evaluation"),
	featuregate.WithRegisterReferenceURL("https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/18642"),
)

// NewSkipExpr creates a BoolExpr that on evaluation returns true if a log should NOT be processed or kept.
// The logic determining if a log should be processed is based on include and exclude settings.
// Include properties are checked before exclude settings are checked.
func NewSkipExpr(mp *filterconfig.MatchConfig) (expr.BoolExpr[*ottllog.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// propertiesMatcher allows matching a log record against various log record properties.
type propertiesMatcher struct {
	filtermatcher.PropertiesMatcher

	// log bodies to compare to.
	bodyFilters filterset.FilterSet

	// log severity texts to compare to
	severityTextFilters filterset.FilterSet

	// matcher for severity number
	severityNumberMatcher *severityNumberMatcher
}

// NewMatcher creates a LogRecord Matcher that matches based on the given MatchProperties.
func newExpr(mp *filterconfig.MatchProperties) (expr.BoolExpr[*ottllog.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Eval matches a log record to a set of properties.
// There are 3 sets of properties to match against.
// The log record names are matched, if specified.
// The log record bodies are matched, if specified.
// The attributes are then checked, if specified.
// At least one of log record names or attributes must be specified. It is
// supported to have more than one of these specified, and all specified must
// evaluate to true for a match to occur.
func (mp *propertiesMatcher) Eval(_ context.Context, tCtx *ottllog.TransformContext) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
