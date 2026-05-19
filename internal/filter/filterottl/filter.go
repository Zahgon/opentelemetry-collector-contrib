// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filterottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterottl"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlprofile"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlscope"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
)

// NewBoolExprForSpan creates a BoolExpr[*ottlspan.TransformContext] that will return true if any of the given OTTL conditions evaluate to true.
// The passed in functions should use the ottlspan.TransformContext.
// If a function named `match` is not present in the function map it will be added automatically so that parsing works as expected
func NewBoolExprForSpan(conditions []string, functions map[string]ottl.Factory[*ottlspan.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (*ottl.ConditionSequence[*ottlspan.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForSpanWithOptions is like NewBoolExprForSpan, but with additional options.
func NewBoolExprForSpanWithOptions(conditions []string, functions map[string]ottl.Factory[*ottlspan.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings, parserOptions []ottl.Option[*ottlspan.TransformContext]) (*ottl.ConditionSequence[*ottlspan.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForSpanEvent creates a BoolExpr[*ottlspanevent.TransformContext] that will return true if any of the given OTTL conditions evaluate to true.
// The passed in functions should use the ottlspanevent.TransformContext.
// If a function named `match` is not present in the function map it will be added automatically so that parsing works as expected
func NewBoolExprForSpanEvent(conditions []string, functions map[string]ottl.Factory[*ottlspanevent.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (*ottl.ConditionSequence[*ottlspanevent.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForSpanEventWithOptions is like NewBoolExprForSpanEvent, but with additional options.
func NewBoolExprForSpanEventWithOptions(conditions []string, functions map[string]ottl.Factory[*ottlspanevent.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings, parserOptions []ottl.Option[*ottlspanevent.TransformContext]) (*ottl.ConditionSequence[*ottlspanevent.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForMetric creates a BoolExpr[*ottlmetric.TransformContext] that will return true if any of the given OTTL conditions evaluate to true.
// The passed in functions should use the ottlmetric.TransformContext.
// If a function named `match` is not present in the function map it will be added automatically so that parsing works as expected
func NewBoolExprForMetric(conditions []string, functions map[string]ottl.Factory[*ottlmetric.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (*ottl.ConditionSequence[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForMetricWithOptions is like NewBoolExprForMetric, but with additional options.
func NewBoolExprForMetricWithOptions(conditions []string, functions map[string]ottl.Factory[*ottlmetric.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings, parserOptions []ottl.Option[*ottlmetric.TransformContext]) (*ottl.ConditionSequence[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForDataPoint creates a BoolExpr[*ottldatapoint.TransformContext] that will return true if any of the given OTTL conditions evaluate to true.
// The passed in functions should use the ottldatapoint.TransformContext.
// If a function named `match` is not present in the function map it will be added automatically so that parsing works as expected
func NewBoolExprForDataPoint(conditions []string, functions map[string]ottl.Factory[*ottldatapoint.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (*ottl.ConditionSequence[*ottldatapoint.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForDataPointWithOptions is like NewBoolExprForDataPoint, but with additional options.
func NewBoolExprForDataPointWithOptions(conditions []string, functions map[string]ottl.Factory[*ottldatapoint.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings, parserOptions []ottl.Option[*ottldatapoint.TransformContext]) (*ottl.ConditionSequence[*ottldatapoint.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForLog creates a BoolExpr[*ottllog.TransformContext] that will return true if any of the given OTTL conditions evaluate to true.
// The passed in functions should use the ottllog.TransformContext.
// If a function named `match` is not present in the function map it will be added automatically so that parsing works as expected
func NewBoolExprForLog(conditions []string, functions map[string]ottl.Factory[*ottllog.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (*ottl.ConditionSequence[*ottllog.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForLogWithOptions is like NewBoolExprForLog, but with additional options.
func NewBoolExprForLogWithOptions(conditions []string, functions map[string]ottl.Factory[*ottllog.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings, parserOptions []ottl.Option[*ottllog.TransformContext]) (*ottl.ConditionSequence[*ottllog.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForProfile creates a BoolExpr[*ottlprofile.TransformContext] that will return true if any of the given OTTL conditions evaluate to true.
// The passed in functions should use the ottlprofile.TransformContext.
// If a function named `match` is not present in the function map it will be added automatically so that parsing works as expected
func NewBoolExprForProfile(conditions []string, functions map[string]ottl.Factory[*ottlprofile.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (*ottl.ConditionSequence[*ottlprofile.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForProfileWithOptions is like NewBoolExprForProfile, but with additional options.
func NewBoolExprForProfileWithOptions(conditions []string, functions map[string]ottl.Factory[*ottlprofile.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings, parserOptions []ottl.Option[*ottlprofile.TransformContext]) (*ottl.ConditionSequence[*ottlprofile.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForResource creates a BoolExpr[*ottlresource.TransformContext] that will return true if any of the given OTTL conditions evaluate to true.
// The passed in functions should use the ottlresource.TransformContext.
// If a function named `match` is not present in the function map it will be added automatically so that parsing works as expected
func NewBoolExprForResource(conditions []string, functions map[string]ottl.Factory[*ottlresource.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (*ottl.ConditionSequence[*ottlresource.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForResourceWithOptions is like NewBoolExprForResource, but with additional options.
func NewBoolExprForResourceWithOptions(conditions []string, functions map[string]ottl.Factory[*ottlresource.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings, parserOptions []ottl.Option[*ottlresource.TransformContext]) (*ottl.ConditionSequence[*ottlresource.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForScope creates a BoolExpr[*ottlscope.TransformContext] that will return true if any of the given OTTL conditions evaluate to true.
// The passed in functions should use the ottlresource.TransformContext.
// If a function named `match` is not present in the function map it will be added automatically so that parsing works as expected
func NewBoolExprForScope(conditions []string, functions map[string]ottl.Factory[*ottlscope.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (*ottl.ConditionSequence[*ottlscope.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForScopeWithOptions is like NewBoolExprForScope, but with additional options.
func NewBoolExprForScopeWithOptions(conditions []string, functions map[string]ottl.Factory[*ottlscope.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings, parserOptions []ottl.Option[*ottlscope.TransformContext]) (*ottl.ConditionSequence[*ottlscope.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
