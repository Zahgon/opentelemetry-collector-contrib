// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package remove // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/remove"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

// RootableField represents a potential field on an entry.
// It differs from a normal Field in that it allows users to
// specify `resource` or `attributes` with the intention
// of referring to "all" fields within those groups.
// It is used to get, set, and delete values at this field.
// It is deserialized from JSON dot notation.
type rootableField struct {
	entry.Field
	allResource   bool
	allAttributes bool
}

// UnmarshalJSON will unmarshal a field from JSON
func (f *rootableField) UnmarshalJSON(raw []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalYAML will unmarshal a field from YAML
func (f *rootableField) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalText will unmarshal a field from text
func (f *rootableField) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (f *rootableField) unmarshalCheckString(s string) error { _ = "STUB: not implemented"; return nil }

// Get gets the value of the field if the flags for 'allAttributes' or 'allResource' isn't set
func (f *rootableField) Get(entry *entry.Entry) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (f *rootableField) IsEmpty() bool { _ = "STUB: not implemented"; return false }
