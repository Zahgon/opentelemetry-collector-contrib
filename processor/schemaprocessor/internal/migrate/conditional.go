// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package migrate // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/migrate"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ValueMatch defines the expected match type
// that is used on creation of `ConditionalAttributeSet`
type ValueMatch interface {
	~string
}

// ConditionalAttributeSet represents a rename_attribute that will happen only if the passed in value matches the `on` set.
type ConditionalAttributeSet struct {
	on    map[string]struct{}
	attrs AttributeChangeSet
}

type ConditionalAttributeSetSlice []*ConditionalAttributeSet

func NewConditionalAttributeSet[Match ValueMatch](mappings map[string]string, copyAttributes bool, matches ...Match) ConditionalAttributeSet {
	_ = "STUB: not implemented"
	return *new(ConditionalAttributeSet)
}

func (ConditionalAttributeSet) IsMigrator() {
	_ = "STUB: not implemented"

	// Do applies the attribute changes specified in the constructor if any of the values in values matches the matches specified in the constructor.
	return
}

func (ca *ConditionalAttributeSet) Do(ss StateSelector, attrs pcommon.Map, values ...string) (errs error) {
	_ = "STUB: not implemented"
	return nil
}

func (ca *ConditionalAttributeSet) check(values ...string) bool {
	_ = "STUB: not implemented"
	return false
}
