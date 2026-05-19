// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlprofile // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlprofile"

import (
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.uber.org/zap/zapcore"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxcache"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxcommon"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxprofile"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxscope"
)

var tcPool = sync.Pool{
	New: func() any {
		return &TransformContext{cache: pcommon.NewMap()}
	},
}

// ContextName is the name of the context for profiles.
// Experimental: *NOTE* this constant is subject to change or removal in the future.
const ContextName = ctxprofile.Name

var (
	_ ctxresource.Context     = (*TransformContext)(nil)
	_ ctxscope.Context        = (*TransformContext)(nil)
	_ ctxprofile.Context      = (*TransformContext)(nil)
	_ zapcore.ObjectMarshaler = (*TransformContext)(nil)
)

// TransformContext represents a profile and its associated hierarchy.
type TransformContext struct {
	profile              pprofile.Profile
	dictionary           pprofile.ProfilesDictionary
	instrumentationScope pcommon.InstrumentationScope
	resource             pcommon.Resource
	cache                pcommon.Map
	scopeProfiles        pprofile.ScopeProfiles
	resourceProfiles     pprofile.ResourceProfiles
}

// MarshalLogObject serializes the profile into a zapcore.ObjectEncoder for logging.
func (tCtx *TransformContext) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// TransformContextOption represents an option for configuring a TransformContext.
type TransformContextOption func(*TransformContext)

// NewTransformContext creates a new TransformContext with the provided parameters.
//
// Deprecated: [v0.145.0] Use NewTransformContextPtr.
func NewTransformContext(profile pprofile.Profile, dictionary pprofile.ProfilesDictionary, instrumentationScope pcommon.InstrumentationScope, resource pcommon.Resource, scopeProfiles pprofile.ScopeProfiles, resourceProfiles pprofile.ResourceProfiles, options ...TransformContextOption) TransformContext {
	_ = "STUB: not implemented"
	return *new(TransformContext)
}

// NewTransformContextPtr returns a new TransformContext with the provided parameters from a pool of contexts.
// Caller must call TransformContext.Close on the returned TransformContext.
func NewTransformContextPtr(resourceProfiles pprofile.ResourceProfiles, scopeProfiles pprofile.ScopeProfiles, profile pprofile.Profile, dictionary pprofile.ProfilesDictionary, options ...TransformContextOption) *TransformContext {
	_ = "STUB: not implemented"
	return nil
}

// Close the current TransformContext.
// After this function returns this instance cannot be used.
func (tCtx *TransformContext) Close() { _ = "STUB: not implemented"; return }

// GetProfile returns the profile from the TransformContext.
func (tCtx *TransformContext) GetProfile() pprofile.Profile {
	_ = "STUB: not implemented"
	return *

	// GetProfilesDictionary returns the profiles dictionary from the TransformContext.
	new(pprofile.Profile)
}

func (tCtx *TransformContext) GetProfilesDictionary() pprofile.ProfilesDictionary {
	_ = "STUB: not implemented"
	return *

	// GetInstrumentationScope returns the instrumentation scope from the TransformContext.
	new(pprofile.ProfilesDictionary)
}

func (tCtx *TransformContext) GetInstrumentationScope() pcommon.InstrumentationScope {
	_ = "STUB: not implemented"
	return *new(pcommon.InstrumentationScope)
}

// GetResource returns the resource from the TransformContext.
func (tCtx *TransformContext) GetResource() pcommon.Resource {
	_ = "STUB: not implemented"
	return *

	// GetScopeSchemaURLItem returns the scope schema URL item from the TransformContext.
	new(pcommon.Resource)
}

func (tCtx *TransformContext) GetScopeSchemaURLItem() ctxcommon.SchemaURLItem {
	_ = "STUB: not implemented"
	return *new(ctxcommon.SchemaURLItem)
}

// GetResourceSchemaURLItem returns the resource schema URL item from the TransformContext.
func (tCtx *TransformContext) GetResourceSchemaURLItem() ctxcommon.SchemaURLItem {
	_ = "STUB: not implemented"
	return *new(ctxcommon.SchemaURLItem)
}

// NewParser creates a new profile parser with the provided functions and options.
func NewParser(functions map[string]ottl.Factory[*TransformContext], telemetrySettings component.TelemetrySettings, options ...ottl.Option[*TransformContext]) (ottl.Parser[*TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
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
