// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package entry // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"

// ResourceField is the path to an entry resource
type ResourceField struct {
	Keys []string
	// prevent unkeyed literal initialization
	_ struct{}
}

// NewResourceField will creat a new resource field from a key
func NewResourceField(keys ...string) Field { _ = "STUB: not implemented"; return *new(Field) }

// Parent returns the parent of the current field.
// In the case that the resource field points to the root node, it is a no-op.
func (f ResourceField) Parent() ResourceField {
	_ = "STUB: not implemented"
	return *new(ResourceField)
}

// Child returns a child of the current field using the given key.
func (f ResourceField) Child(key string) ResourceField {
	_ = "STUB: not implemented"
	return *new(ResourceField)
}

// IsRoot returns a boolean indicating if this is a root level field.
func (f ResourceField) isRoot() bool { _ = "STUB: not implemented"; return false }

// String returns the string representation of this field.
func (f ResourceField) String() string { _ = "STUB: not implemented"; return "" }

// Get will return the resource value and a boolean indicating if it exists
func (f ResourceField) Get(entry *Entry) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Set will set a value on an entry's resource using the field.
// If a key already exists, it will be overwritten.
func (f ResourceField) Set(entry *Entry, value any) error { _ = "STUB: not implemented"; return nil }

// Merge will attempt to merge the contents of a map into an entry's resource.
// It will overwrite any intermediate values as necessary.
func (f ResourceField) Merge(entry *Entry, mapValues map[string]any) {
	_ = "STUB: not implemented"
	return
}

// Delete removes a value from an entry's resource using the field.
// It will return the deleted value and whether the field existed.
func (f ResourceField) Delete(entry *Entry) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

/****************
  Serialization
****************/

// UnmarshalJSON will attempt to unmarshal the field from JSON.
func (f *ResourceField) UnmarshalJSON(raw []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalYAML will attempt to unmarshal a field from YAML.
func (f *ResourceField) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalText will unmarshal a field from text
func (f *ResourceField) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }
