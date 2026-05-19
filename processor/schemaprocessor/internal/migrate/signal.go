// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package migrate // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/migrate"

import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/alias"

// SignalType allows for type constraints in order
// to apply to potential type defined strings.
type SignalType interface {
	~string
}

// SignalNameChange allows for migrating types that
// implement the `alias.NamedSignal` interface.
type SignalNameChange struct {
	updates  map[string]string
	rollback map[string]string
}

// NewSignalNameChange will create a `Signal` that will check the provided mappings if it can update a `alias.NamedSignal`
// and if no values are provided for `matches`, then all values will be updated.
func NewSignalNameChange[Key, Value SignalType](mappings map[Key]Value) SignalNameChange {
	_ = "STUB: not implemented"
	return *new(SignalNameChange)
}

func (SignalNameChange) IsMigrator() { _ = "STUB: not implemented"; return }

func (s *SignalNameChange) Do(ss StateSelector, signal alias.NamedSignal) {
	_ = "STUB: not implemented"
	return
}
