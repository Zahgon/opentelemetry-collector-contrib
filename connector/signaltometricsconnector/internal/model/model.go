// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package model // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector/internal/model"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/signaltometricsconnector/config"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type MetricKey struct {
	Name        string
	Type        pmetric.MetricType
	Unit        string
	Description string
}

type ExplicitHistogram[K any] struct {
	Buckets []float64
	Count   *ottl.ValueExpression[K]
	Value   *ottl.ValueExpression[K]
}

func (h *ExplicitHistogram[K]) fromConfig(
	mi *config.Histogram,
	pc *ottl.ParserCollection[*ottl.ValueExpression[K]],
	contextName string,
) error {
	_ = "STUB: not implemented"
	return nil
}

type ExponentialHistogram[K any] struct {
	MaxSize int32
	Count   *ottl.ValueExpression[K]
	Value   *ottl.ValueExpression[K]
}

func (h *ExponentialHistogram[K]) fromConfig(
	mi *config.ExponentialHistogram,
	pc *ottl.ParserCollection[*ottl.ValueExpression[K]],
	contextName string,
) error {
	_ = "STUB: not implemented"
	return nil
}

type Sum[K any] struct {
	Value       *ottl.ValueExpression[K]
	IsMonotonic bool
}

func (s *Sum[K]) fromConfig(
	mi *config.Sum,
	pc *ottl.ParserCollection[*ottl.ValueExpression[K]],
	contextName string,
) error {
	_ = "STUB: not implemented"
	return nil
}

type Gauge[K any] struct {
	Value *ottl.ValueExpression[K]
}

func (s *Gauge[K]) fromConfig(
	mi *config.Gauge,
	pc *ottl.ParserCollection[*ottl.ValueExpression[K]],
	contextName string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// attributeEntry represents a single entry in include_resource_attributes
// or attributes. Exactly one of Key or Expression must be set.
type attributeEntry[K any] struct {
	Key          string
	Expression   *ottl.ValueExpression[K]
	Optional     bool
	DefaultValue pcommon.Value
}

// AttributeKeyValue represents a resolved attribute entry with a static
// key. This is the output of ResolveAttributes and
// ResolveIncludeResourceAttributes after OTTL expressions have been
// evaluated and flattened into the final ordered list.
type AttributeKeyValue struct {
	Key          string
	Optional     bool
	DefaultValue pcommon.Value
}

type MetricDef[K any] struct {
	Key                       MetricKey
	includeResourceAttributes []attributeEntry[K]
	attributes                []attributeEntry[K]
	// hasExprResAttrs is true if any includeResourceAttributes entry
	// has a keys_expression. When false, the static-only fast path
	// skips OTTL evaluation and dedup overhead in
	// ResolveIncludeResourceAttributes.
	hasExprResAttrs bool
	// hasExprAttrs is true if any attributes entry has a
	// keys_expression. When false, the static-only fast path skips
	// OTTL evaluation and dedup overhead in ResolveAttributes, and
	// sortedStaticAttrs is used directly by ComputeAttributesHash
	// to avoid per-call sorting.
	hasExprAttrs bool
	// sortedStaticAttrs is a pre-sorted copy of the static-key
	// attributes, built at construction time. Used by
	// ResolveAttributes (static fast path) and ComputeAttributesHash
	// to avoid per-call allocation and sorting when there are no
	// keys_expression entries.
	sortedStaticAttrs []AttributeKeyValue
	// sortedStaticResAttrs is the same for includeResourceAttributes.
	sortedStaticResAttrs []AttributeKeyValue
	Conditions           *ottl.ConditionSequence[K]
	ExponentialHistogram *ExponentialHistogram[K]
	ExplicitHistogram    *ExplicitHistogram[K]
	Sum                  *Sum[K]
	Gauge                *Gauge[K]
}

func (md *MetricDef[K]) FromMetricInfo(
	mi config.MetricInfo,
	pc *ottl.ParserCollection[*ottl.ValueExpression[K]],
	contextName string,
	conditions *ottl.ConditionSequence[K],
) error {
	_ = "STUB: not implemented"
	return nil
}

// Detect whether any entries use keys_expression and pre-build
// sorted static attribute lists for the fast path.

// ResolveIncludeResourceAttributes evaluates OTTL expressions in
// include_resource_attributes and returns a flat, ordered list of
// AttributeKeyValue entries. When no keys_expression entries are
// configured, the pre-built sortedStaticResAttrs is returned directly
// without allocation or OTTL evaluation.
func (md *MetricDef[K]) ResolveIncludeResourceAttributes(ctx context.Context, tCtx K) ([]AttributeKeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ResolveAttributes evaluates OTTL expressions in attributes and
// returns a flat, ordered list of AttributeKeyValue entries. When no
// keys_expression entries are configured, the pre-built
// sortedStaticAttrs is returned directly without allocation or OTTL
// evaluation.
func (md *MetricDef[K]) ResolveAttributes(ctx context.Context, tCtx K) ([]AttributeKeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MatchAttributes checks if all required static key attributes are
// present in the given map. This is a cheap pre-check that does not
// require a transform context and can be used to skip expensive
// transform context creation when the entity should not be processed.
// OTTL expression entries and entries with default values or optional
// flag are not checked — the presence of an expression entry means
// there is a possibility of attributes resolving, so MatchAttributes
// will not reject the entity on that basis.
func (md *MetricDef[K]) MatchAttributes(attrs pcommon.Map) bool {
	_ = "STUB: not implemented"
	return false
}

// ComputeAttributesHash returns a 128-bit hash that identifies the
// filtered attribute set. The resolved list is expected to be sorted
// alphabetically by key (guaranteed by both resolveEntries and the
// pre-sorted sortedStaticAttrs).
func (*MetricDef[K]) ComputeAttributesHash(attrs pcommon.Map, resolved []AttributeKeyValue) [16]byte {
	_ = "STUB: not implemented"
	return nil
}

// FilterResourceAttributes builds a pcommon.Map from the resolved
// include_resource_attributes list. If the list is empty, all resource
// attributes are copied. CollectorInstanceInfo is always appended.
func (md *MetricDef[K]) FilterResourceAttributes(
	attrs pcommon.Map,
	resolved []AttributeKeyValue,
	collectorInfo CollectorInstanceInfo,
) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

// FilterAttributes builds a pcommon.Map from the resolved attributes
// list. Entries are applied in order so later entries override earlier
// ones for the same key.
func (*MetricDef[K]) FilterAttributes(attrs pcommon.Map, resolved []AttributeKeyValue) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

// filterByResolved creates a pcommon.Map by iterating the resolved
// AttributeKeyValue list in order. Later entries with the same key
// override earlier ones.
func filterByResolved(attrs pcommon.Map, resolved []AttributeKeyValue, extraCapacity int) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

// resolveEntries evaluates OTTL expressions in entries and returns a
// flat, deduplicated, sorted list of AttributeKeyValue. Entries are
// processed in reverse so the last occurrence of each key (highest
// priority) is kept and earlier duplicates are discarded. The result
// is sorted alphabetically by key for deterministic hashing — this is
// safe because dedup guarantees unique keys and pcommon.Map is
// order-independent.
func resolveEntries[K any](ctx context.Context, tCtx K, entries []attributeEntry[K]) ([]AttributeKeyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildStaticAttrs checks if any entries have OTTL expressions and
// builds a pre-sorted list of static-key AttributeKeyValue entries.
// Returns (hasExpr, sortedStaticAttrs).
func buildStaticAttrs[K any](entries []attributeEntry[K]) (bool, []AttributeKeyValue) {
	_ = "STUB: not implemented"
	return false, nil
}

func parseAttributeEntries[K any](
	cfgs []config.Attribute,
	pc *ottl.ParserCollection[*ottl.ValueExpression[K]],
	contextName string,
) ([]attributeEntry[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// evalKeysExpression evaluates the OTTL expression and returns the
// resolved attribute keys. Returns nil (no error) if the expression
// evaluates to nil.
func evalKeysExpression[K any](
	ctx context.Context,
	tCtx K,
	entry attributeEntry[K],
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func copyAttribute(key string, defaultValue pcommon.Value, src, dst pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func extractStringSlice(val any) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
