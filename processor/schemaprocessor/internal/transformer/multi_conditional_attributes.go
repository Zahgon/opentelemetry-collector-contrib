// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package transformer // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/transformer"

import (
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/migrate"
)

// SpanEventConditionalAttributes is an transformer that powers the [Span Event's rename_attributes] change.
// [Span Event's rename_attributes]: https://opentelemetry.io/docs/specs/otel/schemas/file_format_v1.1.0/#rename_attributes-transformation-1
type SpanEventConditionalAttributes struct {
	MultiConditionalAttributeSet migrate.MultiConditionalAttributeSet
}

func (SpanEventConditionalAttributes) IsMigrator() { _ = "STUB: not implemented"; return }

func (o SpanEventConditionalAttributes) Do(ss migrate.StateSelector, span ptrace.Span) error {
	_ = "STUB: not implemented"
	return nil
}
