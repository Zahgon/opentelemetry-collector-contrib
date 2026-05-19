// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package entry // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"

const (
	AttributesPrefix = "attributes"
	ResourcePrefix   = "resource"
	BodyPrefix       = "body"
)

// Field represents a potential field on an entry.
// It is used to get, set, and delete values at this field.
// It is deserialized from JSON dot notation.
type Field struct {
	FieldInterface
	// prevent unkeyed literal initialization
	_ struct{}
}

// RootableField is a Field that may refer directly to "attributes" or "resource"
type RootableField struct {
	Field
	// prevent unkeyed literal initialization
	_ struct{}
}

// FieldInterface is a field on an entry.
type FieldInterface interface {
	Get(*Entry) (any, bool)
	Set(entry *Entry, value any) error
	Delete(entry *Entry) (any, bool)
	String() string
}

// UnmarshalJSON will unmarshal a field from JSON
func (f *Field) UnmarshalJSON(raw []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSON will unmarshal a field from JSON
func (r *RootableField) UnmarshalJSON(raw []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalYAML will unmarshal a field from YAML
func (f *Field) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalYAML will unmarshal a field from YAML
func (r *RootableField) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalText will unmarshal a field from text
func (f *Field) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalText will unmarshal a field from text
func (r *RootableField) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (f *Field) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func NewField(s string) (Field, error) { _ = "STUB: not implemented"; return *new(Field), nil }

func newField(s string, rootable bool) (Field, error) {
	_ = "STUB: not implemented"
	return *new(Field), nil
}

type splitState uint

const (
	// Begin is the beginning state of a field split
	Begin splitState = iota
	// InBracket is the state of a field split inside a bracket
	InBracket
	// InQuote is the state of a field split inside a quote
	InQuote
	// OutQuote is the state of a field split outside a quote
	OutQuote
	// OutBracket is the state of a field split outside a bracket
	OutBracket
	// InUnbracketedToken is the state field split on any token outside brackets
	InUnbracketedToken
)

func fromJSONDot(s string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// shouldn't be possible

// toJSONDot returns the JSON dot notation for a field.
func toJSONDot(prefix string, keys []string) string { _ = "STUB: not implemented"; return "" }

// getNestedMap will get a nested map assigned to a key.
// If the map does not exist, it will create and return it.
func getNestedMap(currentMap map[string]any, key string) map[string]any {
	_ = "STUB: not implemented"
	return nil
}
