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

// newBoolExprWithPathContextNames wraps parser in a single-context ottl.ParserCollection so
// that conditions without an explicit path context are rewritten to use contextName as their
// context. The parser must be constructed with EnablePathContextNames().
func newBoolExprWithPathContextNames[K any](
	contextName string,
	parser ottl.Parser[K],
	conditions []string,
	set component.TelemetrySettings,
	newConditionSequence func([]*ottl.Condition[K]) ottl.ConditionSequence[K],
) (*ottl.ConditionSequence[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForSpanWithPathContextNames is like NewBoolExprForSpan, but conditions may use OTTL
// path context names (e.g. `span.attributes["foo"]`). Conditions without an explicit context are
// rewritten to use the span context (e.g. `attributes["foo"]` becomes `span.attributes["foo"]`).
func NewBoolExprForSpanWithPathContextNames(conditions []string, functions map[string]ottl.Factory[*ottlspan.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (*ottl.ConditionSequence[*ottlspan.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForSpanEventWithPathContextNames is like NewBoolExprForSpanEvent, but conditions may use
// OTTL path context names (e.g. `spanevent.attributes["foo"]`). Conditions without an explicit
// context are rewritten to use the spanevent context.
func NewBoolExprForSpanEventWithPathContextNames(conditions []string, functions map[string]ottl.Factory[*ottlspanevent.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (*ottl.ConditionSequence[*ottlspanevent.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForMetricWithPathContextNames is like NewBoolExprForMetric, but conditions may use OTTL
// path context names (e.g. `metric.name`). Conditions without an explicit context are rewritten to
// use the metric context.
func NewBoolExprForMetricWithPathContextNames(conditions []string, functions map[string]ottl.Factory[*ottlmetric.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (*ottl.ConditionSequence[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForDataPointWithPathContextNames is like NewBoolExprForDataPoint, but conditions may
// use OTTL path context names (e.g. `datapoint.attributes["foo"]`). Conditions without an explicit
// context are rewritten to use the datapoint context.
func NewBoolExprForDataPointWithPathContextNames(conditions []string, functions map[string]ottl.Factory[*ottldatapoint.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (*ottl.ConditionSequence[*ottldatapoint.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForLogWithPathContextNames is like NewBoolExprForLog, but conditions may use OTTL path
// context names (e.g. `log.attributes["foo"]`). Conditions without an explicit context are
// rewritten to use the log context.
func NewBoolExprForLogWithPathContextNames(conditions []string, functions map[string]ottl.Factory[*ottllog.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (*ottl.ConditionSequence[*ottllog.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForProfileWithPathContextNames is like NewBoolExprForProfile, but conditions may use
// OTTL path context names (e.g. `profile.attributes["foo"]`). Conditions without an explicit
// context are rewritten to use the profile context.
func NewBoolExprForProfileWithPathContextNames(conditions []string, functions map[string]ottl.Factory[*ottlprofile.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (*ottl.ConditionSequence[*ottlprofile.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForResourceWithPathContextNames is like NewBoolExprForResource, but conditions may use
// OTTL path context names (e.g. `resource.attributes["foo"]`). Conditions without an explicit
// context are rewritten to use the resource context.
func NewBoolExprForResourceWithPathContextNames(conditions []string, functions map[string]ottl.Factory[*ottlresource.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (*ottl.ConditionSequence[*ottlresource.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBoolExprForScopeWithPathContextNames is like NewBoolExprForScope, but conditions may use OTTL
// path context names (e.g. `scope.name`). Conditions without an explicit context are rewritten to
// use the scope context.
func NewBoolExprForScopeWithPathContextNames(conditions []string, functions map[string]ottl.Factory[*ottlscope.TransformContext], errorMode ottl.ErrorMode, set component.TelemetrySettings) (*ottl.ConditionSequence[*ottlscope.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
