// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlspan // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"

import (
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap/zapcore"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxcache"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxcommon"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxspan"
)

var tcPool = sync.Pool{
	New: func() any {
		return &TransformContext{cache: pcommon.NewMap()}
	},
}

// ContextName is the name of the context for spans.
// Experimental: *NOTE* this constant is subject to change or removal in the future.
const ContextName = ctxspan.Name

var _ zapcore.ObjectMarshaler = (*TransformContext)(nil)

// TransformContext represents a span and its associated hierarchy.
type TransformContext struct {
	resourceSpans ptrace.ResourceSpans
	scopeSpans    ptrace.ScopeSpans
	span          ptrace.Span
	cache         pcommon.Map
}

// MarshalLogObject serializes the TransformContext into a zapcore.ObjectEncoder for logging.
func (tCtx *TransformContext) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// TransformContextOption represents an option for configuring a TransformContext.
type TransformContextOption func(*TransformContext)

// NewTransformContextPtr returns a new TransformContext with the provided parameters from a pool of contexts.
// Caller must call TransformContext.Close on the returned TransformContext.
func NewTransformContextPtr(resourceSpans ptrace.ResourceSpans, scopeSpans ptrace.ScopeSpans, span ptrace.Span, options ...TransformContextOption) *TransformContext {
	_ = "STUB: not implemented"
	return nil
}

// Close the current TransformContext.
// After this function returns this instance cannot be used.
func (tCtx *TransformContext) Close() { _ = "STUB: not implemented"; return }

// GetSpan returns the span from the TransformContext.
func (tCtx *TransformContext) GetSpan() ptrace.Span {
	_ = "STUB: not implemented"

	// GetInstrumentationScope returns the instrumentation scope from the TransformContext.
	return *new(ptrace.Span)
}

func (tCtx *TransformContext) GetInstrumentationScope() pcommon.InstrumentationScope {
	_ = "STUB: not implemented"
	return *new(pcommon.InstrumentationScope)
}

// GetResource returns the resource from the TransformContext.
func (tCtx *TransformContext) GetResource() pcommon.Resource {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource)
}

// GetResourceSchemaURLItem returns the schema URL item for the resource from the TransformContext.
func (tCtx *TransformContext) GetResourceSchemaURLItem() ctxcommon.SchemaURLItem {
	_ = "STUB: not implemented"
	return *new(ctxcommon.SchemaURLItem)
}

// GetScopeSchemaURLItem returns the schema URL item for the scope from the TransformContext.
func (tCtx *TransformContext) GetScopeSchemaURLItem() ctxcommon.SchemaURLItem {
	_ = "STUB: not implemented"
	return *

	// EnablePathContextNames enables the support for path's context names on statements.
	// When this option is configured, all statement's paths must have a valid context prefix,
	// otherwise an error is reported.
	//
	// Experimental: *NOTE* this option is subject to change or removal in the future.
	new(ctxcommon.SchemaURLItem)
}

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

// NewParser creates a new span parser with the provided functions and options.
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
