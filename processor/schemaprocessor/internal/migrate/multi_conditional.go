// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package migrate // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/migrate"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

type set = map[string]struct{}

// MultiConditionalAttributeSet maps from string keys to possible values for each of those keys.  The Do function then checks passed in values for each key against the list provided here in the constructor.  If there is a matching value for each key, the attribute changes are applied.
type MultiConditionalAttributeSet struct {
	// map from string keys (in the intended case "event.name" and "span.name") to a set of acceptable values.
	keysToPossibleValues map[string]set
	attrs                AttributeChangeSet
}

type MultiConditionalAttributeSetSlice []*MultiConditionalAttributeSet

func NewMultiConditionalAttributeSet[Match ValueMatch](mappings map[string]string, copyAttributes bool, matches map[string][]Match) MultiConditionalAttributeSet {
	_ = "STUB: not implemented"
	return *new(MultiConditionalAttributeSet)
}

func (MultiConditionalAttributeSet) IsMigrator() {
	_ = "STUB: not implemented"

	// Do function applies the attribute changes if the passed in values match the expected values provided in the constructor.  Uses the Do method of the embedded AttributeChangeSet
	return
}

func (ca *MultiConditionalAttributeSet) Do(ss StateSelector, attrs pcommon.Map, keyToCheckVals map[string]string) (errs error) {
	_ = "STUB: not implemented"
	return nil
}

func (ca *MultiConditionalAttributeSet) check(keyToCheckVals map[string]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// We must already have a key matching the input key!  If not, return an error
// indicates a programming error, should be impossible if using the class correctly

// if there's nothing in here, match all values

// if we've gone through every one of the keys, and they've all generated matches, return true
