// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package changelist // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/changelist"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/migrate"
)

// ChangeList represents a list of changes within a section of the schema processor.  It can take in a list of different migrators for a specific section and will apply them in order, based on whether Apply or Rollback is called
type ChangeList struct {
	Migrators []migrate.Migrator
}

func (c ChangeList) Do(ss migrate.StateSelector, signal any) error {
	_ = "STUB: not implemented"
	return nil
}

// todo(ankit) in go1.23 switch to reversed iterators for this

// switch between transformer types - what do the transformers act on?

// this one acts on both spans and span events!

func (c ChangeList) Apply(signal any) error { _ = "STUB: not implemented"; return nil }

func (c ChangeList) Rollback(signal any) error { _ = "STUB: not implemented"; return nil }
