// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package entry // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"

// BodyField is a field found on an entry body.
type BodyField struct {
	Keys []string
	// prevent unkeyed literal initialization
	_ struct{}
}

// NewBodyField creates a new field from an ordered array of keys.
func NewBodyField(keys ...string) Field { _ = "STUB: not implemented"; return *new(Field) }

// Parent returns the parent of the current field.
// In the case that the body field points to the root node, it is a no-op.
func (f BodyField) Parent() BodyField { _ = "STUB: not implemented"; return *new(BodyField) }

// Child returns a child of the current field using the given key.
func (f BodyField) Child(key string) BodyField { _ = "STUB: not implemented"; return *new(BodyField) }

// IsRoot returns a boolean indicating if this is a root level field.
func (f BodyField) isRoot() bool { _ = "STUB: not implemented"; return false }

// String returns the string representation of this field.
func (f BodyField) String() string { _ = "STUB: not implemented"; return "" }

// Get will retrieve a value from an entry's body using the field.
// It will return the value and whether the field existed.
func (f BodyField) Get(entry *Entry) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Set will set a value on an entry's body using the field.
// If a key already exists, it will be overwritten.
func (f BodyField) Set(entry *Entry, value any) error { _ = "STUB: not implemented"; return nil }

// Merge will attempt to merge the contents of a map into an entry's body.
// It will overwrite any intermediate values as necessary.
func (f BodyField) Merge(entry *Entry, mapValues map[string]any) { _ = "STUB: not implemented"; return }

// Delete removes a value from an entry's body using the field.
// It will return the deleted value and whether the field existed.
func (f BodyField) Delete(entry *Entry) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

/****************
  Serialization
****************/

// UnmarshalJSON will attempt to unmarshal the field from JSON.
func (f *BodyField) UnmarshalJSON(raw []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalYAML will attempt to unmarshal a field from YAML.
func (f *BodyField) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalText will unmarshal a field from text
func (f *BodyField) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }
