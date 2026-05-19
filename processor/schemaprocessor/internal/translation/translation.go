// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translation // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/translation"

import (
	"sort"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	ast11 "go.opentelemetry.io/otel/schema/v1.1/ast"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/alias"
)

// Translation defines the complete abstraction of schema translation file
// that is defined as part of the https://opentelemetry.io/docs/reference/specification/schemas/file_format_v1.0.0/
// Each instance of Translation is "Target Aware", meaning that given a schemaURL as an input
// it will convert from the given input, to the configured target.
//
// Note: as an optimisation, once a Translation is returned from the manager,
//
//	there is no checking the incoming signals if the schema family is a match.
type Translation interface {
	// SupportedVersion checks to see if the provided version is defined as part
	// of this translation since it is useful to know if the translation is missing
	// updates.
	SupportedVersion(v *Version) bool

	// TargetSchemaURL returns the target schema URL for this translation.
	TargetSchemaURL() string

	// ApplyAllResourceChanges will modify the resource part of the incoming signals
	// This applies to all telemetry types and should be applied there
	ApplyAllResourceChanges(in alias.Resource, inSchemaURL string) error

	// ApplyScopeSpanChanges will modify all spans and span events within the incoming signals
	ApplyScopeSpanChanges(in ptrace.ScopeSpans, inSchemaURL string) error

	// ApplyScopeLogChanges will modify all logs within the incoming signal
	ApplyScopeLogChanges(in plog.ScopeLogs, inSchemaURL string) error

	// ApplyScopeMetricChanges will update all metrics including
	// histograms, exponential histograms, summaries, sum and gauges
	ApplyScopeMetricChanges(in pmetric.ScopeMetrics, inSchemaURL string) error
}

type translator struct {
	targetSchemaURL string
	target          *Version
	indexes         map[Version]int // map from version to index in revisions containing the pertinent Version
	revisions       []RevisionV1
	copyFromVersion *Version

	log *zap.Logger
}

type iterator func() (r RevisionV1, more bool)

var (
	_ sort.Interface = (*translator)(nil)
	_ Translation    = (*translator)(nil)
)

func (t *translator) loadTranslation(content *ast11.Schema) error {
	_ = "STUB: not implemented"
	return nil
}

// When copyFromVersion is set, attribute renames preserve both old
// and new names for revisions between copyFromVersion and the target
// (in either direction).

func newTranslatorFromSchema(log *zap.Logger, targetSchemaURL string, schemaFileSchema *ast11.Schema, copyFromVersion *Version) (*translator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTranslator(log *zap.Logger, targetSchemaURL, schema string, copyFromVersion *Version) (*translator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *translator) Len() int { _ = "STUB: not implemented"; return 0 }

func (t *translator) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (t *translator) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (t *translator) TargetSchemaURL() string { _ = "STUB: not implemented"; return "" }

func (t *translator) SupportedVersion(v *Version) bool { _ = "STUB: not implemented"; return false }

func (t *translator) ApplyAllResourceChanges(resource alias.Resource, inSchemaURL string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *translator) ApplyScopeLogChanges(scopeLogs plog.ScopeLogs, inSchemaURL string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *translator) ApplyScopeSpanChanges(scopeSpans ptrace.ScopeSpans, inSchemaURL string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *translator) ApplyScopeMetricChanges(scopeMetrics pmetric.ScopeMetrics, inSchemaURL string) error {
	_ = "STUB: not implemented"
	return nil
}

// iterator abstracts the logic to perform the migrations of "From Version to Version".
// The return values an iterator type and translation status that
// should be compared against Revert, Update, NoChange
// to determine what should be applied.
// In the event that the ChangeSet has not yet been created (it is possible on a cold start)
// then the iterator will wait til the update has been made
//
// Note: Once an iterator has been made, the passed context MUST cancel or run to completion
//
//	in order for the read lock to be released if either Revert or Upgrade has been returned.
func (t *translator) iterator(from *Version) (iterator, int) {
	_ = "STUB: not implemented"
	return *new(iterator), 0
}

// In the event of an update, the iterator needs to also run that version
// for the signal to be the correct version.

// we need to not run the starting version to start with, that's already been done!

// Performs a bounds check and if it has reached stop

// The iterator value needs to move the opposite direction of what
// status is defined as so subtracting it to progress the iterator.
