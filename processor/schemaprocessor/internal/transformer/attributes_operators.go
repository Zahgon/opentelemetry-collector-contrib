// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package transformer contains various Transformers that represent a high level operation - typically a single "change" block from the schema change file.  They rely on Migrators to do the actual work of applying the change to the data.  Transformers accept and operate on a specific type of pdata (logs, metrics, etc)
package transformer // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/transformer"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/migrate"
)

// MetricAttributes is a Transformer that acts on [pmetric.Metric]'s DataPoint's attributes.  It is part of the [AllAttributes].
type MetricAttributes struct {
	AttributeChange migrate.AttributeChangeSet
}

func (MetricAttributes) IsMigrator() { _ = "STUB: not implemented"; return }

func (o MetricAttributes) Do(ss migrate.StateSelector, metric pmetric.Metric) error {
	_ = "STUB: not implemented"
	return nil
}

// LogAttributes is a Transformer that acts on [plog.LogRecord] attributes.  It powers the [Log's rename_attributes] transformation.  It also powers the [AllAttributes].
// [Log's rename_attributes]: https://opentelemetry.io/docs/specs/otel/schemas/file_format_v1.1.0/#rename_attributes-transformation-3
type LogAttributes struct {
	AttributeChange migrate.AttributeChangeSet
}

func (LogAttributes) IsMigrator() { _ = "STUB: not implemented"; return }

func (o LogAttributes) Do(ss migrate.StateSelector, log plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// SpanAttributes is a Transformer that acts on [ptrace.Span] attributes.  It powers the [Span's rename_attributes] transformation.  It also powers the [AllAttributes].
// [Span's rename_attributes]: https://opentelemetry.io/docs/specs/otel/schemas/file_format_v1.1.0/#rename_attributes-transformation
type SpanAttributes struct {
	AttributeChange migrate.AttributeChangeSet
}

func (SpanAttributes) IsMigrator() { _ = "STUB: not implemented"; return }

func (o SpanAttributes) Do(ss migrate.StateSelector, span ptrace.Span) error {
	_ = "STUB: not implemented"
	return nil
}

// SpanEventAttributes is a Transformer that acts on [ptrace.SpanEvent] attributes.  It is part of the [AllAttributes].
type SpanEventAttributes struct {
	AttributeChange migrate.AttributeChangeSet
}

func (SpanEventAttributes) IsMigrator() { _ = "STUB: not implemented"; return }

func (o SpanEventAttributes) Do(ss migrate.StateSelector, spanEvent ptrace.SpanEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// ResourceAttributes is a Transformer that acts on [pcommon.Resource] attributes.  It powers the [Resource's rename_attributes] transformation.  It also powers the [AllAttributes].
// [Resource's rename_attributes]: https://opentelemetry.io/docs/specs/otel/schemas/file_format_v1.1.0/#resources-section
type ResourceAttributes struct {
	AttributeChange migrate.AttributeChangeSet
}

func (ResourceAttributes) IsMigrator() { _ = "STUB: not implemented"; return }

func (o ResourceAttributes) Do(ss migrate.StateSelector, resource pcommon.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

// AllAttributes is a Transformer that acts on .  It is a wrapper around the other attribute transformers.  It powers the [All rename_attributes] transformation.
// [All rename_attributes]: https://opentelemetry.io/docs/specs/otel/schemas/file_format_v1.1.0/#all-section
type AllAttributes struct {
	MetricAttributes    MetricAttributes
	LogAttributes       LogAttributes
	SpanAttributes      SpanAttributes
	SpanEventAttributes SpanEventAttributes
	ResourceAttributes  ResourceAttributes
}

func NewAllAttributesTransformer(set migrate.AttributeChangeSet) AllAttributes {
	_ = "STUB: not implemented"
	return *new(AllAttributes)
}

func (AllAttributes) IsMigrator() { _ = "STUB: not implemented"; return }

func (o AllAttributes) Do(ss migrate.StateSelector, data any) error {
	_ = "STUB: not implemented"
	return nil
}
