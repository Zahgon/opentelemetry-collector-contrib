// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlmetric // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"

import (
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap/zapcore"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxcache"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxcommon"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxscope"
)

var tcPool = sync.Pool{
	New: func() any {
		return &TransformContext{cache: pcommon.NewMap()}
	},
}

// ContextName is the name of the context for metrics.
// Experimental: *NOTE* this constant is subject to change or removal in the future.
const ContextName = ctxmetric.Name

var (
	_ ctxresource.Context = (*TransformContext)(nil)
	_ ctxscope.Context    = (*TransformContext)(nil)
	_ ctxmetric.Context   = (*TransformContext)(nil)
)

// TransformContext represents a metric and its associated hierarchy.
type TransformContext struct {
	resourceMetrics pmetric.ResourceMetrics
	scopeMetrics    pmetric.ScopeMetrics
	metric          pmetric.Metric
	cache           pcommon.Map
}

// MarshalLogObject serializes the metric into a zapcore.ObjectEncoder for logging.
func (tCtx *TransformContext) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// TransformContextOption represents an option for configuring a TransformContext.
type TransformContextOption func(*TransformContext)

// NewTransformContextPtr returns a new TransformContext with the provided parameters from a pool of contexts.
// Caller must call TransformContext.Close on the returned TransformContext.
func NewTransformContextPtr(resourceMetrics pmetric.ResourceMetrics, scopeMetrics pmetric.ScopeMetrics, metric pmetric.Metric, options ...TransformContextOption) *TransformContext {
	_ = "STUB: not implemented"
	return nil
}

// Close the current TransformContext.
// After this function returns this instance cannot be used.
func (tCtx *TransformContext) Close() { _ = "STUB: not implemented"; return }

// GetMetric returns the metric from the TransformContext.
func (tCtx *TransformContext) GetMetric() pmetric.Metric {
	_ = "STUB: not implemented"
	return *

	// GetMetrics returns the metric slice from the TransformContext.
	new(pmetric.Metric)
}

func (tCtx *TransformContext) GetMetrics() pmetric.MetricSlice {
	_ = "STUB: not implemented"
	return *new(pmetric.MetricSlice)
}

// GetInstrumentationScope returns the instrumentation scope from the TransformContext.
func (tCtx *TransformContext) GetInstrumentationScope() pcommon.InstrumentationScope {
	_ = "STUB: not implemented"
	return *new(pcommon.InstrumentationScope)
}

// GetResource returns the resource from the TransformContext.
func (tCtx *TransformContext) GetResource() pcommon.Resource {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource)
}

// GetScopeSchemaURLItem returns the scope schema URL item from the TransformContext.
func (tCtx *TransformContext) GetScopeSchemaURLItem() ctxcommon.SchemaURLItem {
	_ = "STUB: not implemented"
	return *

	// GetResourceSchemaURLItem returns the resource schema URL item from the TransformContext.
	new(ctxcommon.SchemaURLItem)
}

func (tCtx *TransformContext) GetResourceSchemaURLItem() ctxcommon.SchemaURLItem {
	_ = "STUB: not implemented"
	return *new(ctxcommon.SchemaURLItem)
}

// EnablePathContextNames enables the support for path's context names on statements.
// When this option is configured, all statement's paths must have a valid context prefix,
// otherwise an error is reported.
//
// Experimental: *NOTE* this option is subject to change or removal in the future.
func EnablePathContextNames() ottl.Option[*TransformContext] { _ = "STUB: not implemented"; return nil }

// StatementSequenceOption represents an option for configuring a statement sequence.
type StatementSequenceOption func(*ottl.StatementSequence[*TransformContext])

// WithStatementSequenceErrorMode sets the error mode for a statement sequence.
func WithStatementSequenceErrorMode(errorMode ottl.ErrorMode) StatementSequenceOption {
	_ = "STUB: not implemented"
	return *new(StatementSequenceOption)
}

// NewStatementSequence creates a new statement sequence with the provided statements and options.
func NewStatementSequence(statements []*ottl.Statement[*TransformContext], telemetrySettings component.TelemetrySettings, options ...StatementSequenceOption) ottl.StatementSequence[*TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

// ConditionSequenceOption represents an option for configuring a condition sequence.
type ConditionSequenceOption func(*ottl.ConditionSequence[*TransformContext])

// WithConditionSequenceErrorMode sets the error mode for a condition sequence.
func WithConditionSequenceErrorMode(errorMode ottl.ErrorMode) ConditionSequenceOption {
	_ = "STUB: not implemented"
	return *new(ConditionSequenceOption)
}

// NewConditionSequence creates a new condition sequence with the provided conditions and options.
func NewConditionSequence(conditions []*ottl.Condition[*TransformContext], telemetrySettings component.TelemetrySettings, options ...ConditionSequenceOption) ottl.ConditionSequence[*TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

// NewParser creates a new metric parser with the provided functions and options.
func NewParser(
	functions map[string]ottl.Factory[*TransformContext],
	telemetrySettings component.TelemetrySettings,
	options ...ottl.Option[*TransformContext],
) (ottl.Parser[*TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseEnum(val *ottl.EnumSymbol) (*ottl.Enum, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getCache(tCtx *TransformContext) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func pathExpressionParser(cacheGetter ctxcache.Getter[*TransformContext]) ottl.PathExpressionParser[*TransformContext] {
	_ = "STUB: not implemented"
	return nil
}
