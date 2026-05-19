// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package condition // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor/internal/condition"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlscope"
)

// resourceConditionBuilder creates signal-specific conditions (parsedLogConditions, parsedMetricConditions, parsedTraceConditions, parsedProfileConditions) from resource conditions
type resourceConditionBuilder[R any] func([]*ottl.Condition[*ottlresource.TransformContext], component.TelemetrySettings, ottl.ErrorMode) R

// scopeConditionBuilder creates signal-specific conditions (parsedLogConditions, parsedMetricConditions, parsedTraceConditions, parsedProfileConditions) from scope conditions
type scopeConditionBuilder[R any] func([]*ottl.Condition[*ottlscope.TransformContext], component.TelemetrySettings, ottl.ErrorMode) R

func withCommonParsers[R any](resourceFunctions map[string]ottl.Factory[*ottlresource.TransformContext], resourceBuilder resourceConditionBuilder[R], scopeBuilder scopeConditionBuilder[R]) ottl.ParserCollectionOption[R] {
	_ = "STUB: not implemented"
	return nil
}

func resourceConditionsConverter[R any](builder resourceConditionBuilder[R]) ottl.ParsedConditionsConverter[*ottlresource.TransformContext, R] {
	_ = "STUB: not implemented"
	return nil
}

func scopeConditionsConverter[R any](builder scopeConditionBuilder[R]) ottl.ParsedConditionsConverter[*ottlscope.TransformContext, R] {
	_ = "STUB: not implemented"
	return nil
}
