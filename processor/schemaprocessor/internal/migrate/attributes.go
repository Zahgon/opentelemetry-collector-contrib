// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package migrate // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/migrate"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// AttributeChangeSet represents a rename_attributes type operation.
// The listed changes are duplicated twice
// to allow for simplified means of transition to or from a revision.
type AttributeChangeSet struct {
	// The keys are the old attribute name used in the previous version, the values are the
	// new attribute name starting from this version (comment from ast.AttributeMap)
	updates map[string]string
	// the inverse of the updates map
	rollback       map[string]string
	copyAttributes bool
}

// NewAttributeChangeSet allows for typed strings to be used as part
// of the invocation that will be converted into the default string type.
func NewAttributeChangeSet(mappings map[string]string, copyAttributes bool) AttributeChangeSet {
	_ = "STUB: not implemented"
	// for ambiguous rollbacks (if updates contains entries with multiple keys that have the same value), rollback contains the last key iterated over in mappings
	return *new(AttributeChangeSet)
}

func (AttributeChangeSet) IsMigrator() { _ = "STUB: not implemented"; return }

func (a *AttributeChangeSet) Do(ss StateSelector, attrs pcommon.Map) (errs error) {
	_ = "STUB: not implemented"
	return nil
}

// In copy mode, preserve the original attribute and also
// write the renamed copy (unless the target already exists).

// Original key k is preserved — falls through to CopyTo below.

// The schema file format spec does not define behavior for name
// conflicts (when the target attribute already exists). We treat
// identical values as a no-op and only report an error when they differ.
