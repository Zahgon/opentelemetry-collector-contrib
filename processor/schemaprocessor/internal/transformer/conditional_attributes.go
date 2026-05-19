// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package transformer // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/transformer"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/migrate"
)

// MetricDataPointAttributes is a conditional Transformer that acts on [pmetric.Metric]'s DataPoint's attributes.  It powers the [Metric's rename_attributes] transformation.
// [Metric's rename_attributes]: https://opentelemetry.io/docs/specs/otel/schemas/file_format_v1.1.0/#rename_attributes-transformation-2
type MetricDataPointAttributes struct {
	ConditionalAttributeChange migrate.ConditionalAttributeSet
}

func (MetricDataPointAttributes) IsMigrator() { _ = "STUB: not implemented"; return }

func (o MetricDataPointAttributes) Do(ss migrate.StateSelector, metric pmetric.Metric) error {
	_ = "STUB: not implemented"
	return nil
}

// SpanConditionalAttributes is a conditional Transformer that acts on [ptrace.Span]'s name.  It powers the [Span's rename_attributes] transformation.
// [Span's rename_attributes]: https://opentelemetry.io/docs/specs/otel/schemas/file_format_v1.1.0/#rename_attributes-transformation
type SpanConditionalAttributes struct {
	Migrator migrate.ConditionalAttributeSet
}

func (SpanConditionalAttributes) IsMigrator() { _ = "STUB: not implemented"; return }

func (o SpanConditionalAttributes) Do(ss migrate.StateSelector, span ptrace.Span) error {
	_ = "STUB: not implemented"
	return nil
}
