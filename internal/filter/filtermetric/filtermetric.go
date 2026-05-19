// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filtermetric // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filtermetric"

import (
	"go.opentelemetry.io/collector/featuregate"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterconfig"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
)

var UseOTTLBridge = featuregate.GlobalRegistry().MustRegister(
	"filter.filtermetric.useOTTLBridge",
	featuregate.StageAlpha,
	featuregate.WithRegisterDescription("When enabled, filtermetric will convert filtermetric configuration to OTTL and use filterottl evaluation"),
	featuregate.WithRegisterReferenceURL("https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/18642"),
)

// NewSkipExpr creates a BoolExpr that on evaluation returns true if a metric should NOT be processed or kept.
// The logic determining if a metric should be processed is based on include and exclude settings.
// Include properties are checked before exclude settings are checked.
func NewSkipExpr(include, exclude *filterconfig.MetricMatchProperties) (expr.BoolExpr[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewMatcher constructs a metric Matcher. If an 'expr' match type is specified,
// returns an expr matcher, otherwise a name matcher.
func newExpr(mp *filterconfig.MetricMatchProperties) (expr.BoolExpr[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
