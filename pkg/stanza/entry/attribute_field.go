// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package entry // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"

// AttributeField is the path to an entry attribute
type AttributeField struct {
	Keys []string
	// prevent unkeyed literal initialization
	_ struct{}
}

// NewAttributeField will creat a new attribute field from a key
func NewAttributeField(keys ...string) Field { _ = "STUB: not implemented"; return *new(Field) }

// Parent returns the parent of the current field.
// In the case that the attribute field points to the root node, it is a no-op.
func (f AttributeField) Parent() AttributeField {
	_ = "STUB: not implemented"
	return *new(AttributeField)
}

// Child returns a child of the current field using the given key.
func (f AttributeField) Child(key string) AttributeField {
	_ = "STUB: not implemented"
	return *new(AttributeField)
}

// IsRoot returns a boolean indicating if this is a root level field.
func (f AttributeField) isRoot() bool { _ = "STUB: not implemented"; return false }

// String returns the string representation of this field.
func (f AttributeField) String() string { _ = "STUB: not implemented"; return "" }

// Get will return the attribute value and a boolean indicating if it exists
func (f AttributeField) Get(entry *Entry) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Set will set a value on an entry's attributes using the field.
// If a key already exists, it will be overwritten.
func (f AttributeField) Set(entry *Entry, value any) error { _ = "STUB: not implemented"; return nil }

// Merge will attempt to merge the contents of a map into an entry's attributes.
// It will overwrite any intermediate values as necessary.
func (f AttributeField) Merge(entry *Entry, mapValues map[string]any) {
	_ = "STUB: not implemented"
	return
}

// Delete removes a value from an entry's attributes using the field.
// It will return the deleted value and whether the field existed.
func (f AttributeField) Delete(entry *Entry) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

/****************
  Serialization
****************/

// UnmarshalJSON will attempt to unmarshal the field from JSON.
func (f *AttributeField) UnmarshalJSON(raw []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalYAML will attempt to unmarshal a field from YAML.
func (f *AttributeField) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalText will unmarshal a field from text
func (f *AttributeField) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }
