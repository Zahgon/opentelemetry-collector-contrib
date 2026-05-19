// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlresource // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"

import (
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap/zapcore"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxcache"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxcommon"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxresource"
)

var tcPool = sync.Pool{
	New: func() any {
		return &TransformContext{cache: pcommon.NewMap()}
	},
}

// ContextName is the name of the context for resources.
// Experimental: *NOTE* this constant is subject to change or removal in the future.
const ContextName = ctxresource.Name

var (
	_ ctxresource.Context     = (*TransformContext)(nil)
	_ zapcore.ObjectMarshaler = (*TransformContext)(nil)
)

// TransformContext represents a resource and its associated hierarchy.
type TransformContext struct {
	resource      pcommon.Resource
	cache         pcommon.Map
	schemaURLItem ctxcommon.SchemaURLItem
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
func NewTransformContextPtr(resource pcommon.Resource, schemaURLItem ctxcommon.SchemaURLItem, options ...TransformContextOption) *TransformContext {
	_ = "STUB: not implemented"
	return nil
}

// Close the current TransformContext.
// After this function returns this instance cannot be used.
func (tCtx *TransformContext) Close() { _ = "STUB: not implemented"; return }

// GetResource returns the resource from the TransformContext.
func (tCtx *TransformContext) GetResource() pcommon.Resource {
	_ = "STUB: not implemented"
	return *

	// GetResourceSchemaURLItem returns the schema URL item for the resource from the TransformContext.
	new(pcommon.Resource)
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

// NewParser creates a new resource parser with the provided functions and options.
func NewParser(
	functions map[string]ottl.Factory[*TransformContext],
	telemetrySettings component.TelemetrySettings,
	options ...ottl.Option[*TransformContext],
) (ottl.Parser[*TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseEnum(_ *ottl.EnumSymbol) (*ottl.Enum, error) { _ = "STUB: not implemented"; return nil, nil }

func getCache(tCtx *TransformContext) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func pathExpressionParser(cacheGetter ctxcache.Getter[*TransformContext]) ottl.PathExpressionParser[*TransformContext] {
	_ = "STUB: not implemented"
	return nil
}
