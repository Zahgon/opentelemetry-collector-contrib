// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal

type SchemaElement interface {
	setIsPointer(value bool)
	setDescription(description string)
	setOptional(value bool)
}

type SchemaObject interface {
	AddProperty(name string, property SchemaElement)
	AddEmbedded(element SchemaElement)
}

type BaseSchemaElement struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	IsPointer   bool   `json:"x-pointer,omitempty" yaml:"x-pointer,omitempty"`
	IsOptional  bool   `json:"x-optional,omitempty" yaml:"x-optional,omitempty"`
}

func (b *BaseSchemaElement) setIsPointer(value bool) { _ = "STUB: not implemented"; return }

func (b *BaseSchemaElement) setDescription(value string) { _ = "STUB: not implemented"; return }

func (b *BaseSchemaElement) setOptional(value bool) { _ = "STUB: not implemented"; return }

type RefSchemaElement struct {
	BaseSchemaElement `json:",inline" yaml:",inline"`
	Ref               string `json:"$ref" yaml:"$ref"`
}

type FieldSchemaElement struct {
	BaseSchemaElement `json:",inline" yaml:",inline"`
	ElementType       SchemaType `json:"type,omitempty" yaml:"type,omitempty"`
	CustomElementType string     `json:"x-customType,omitempty" yaml:"x-customType,omitempty"`
	Format            string     `json:"format,omitempty" yaml:"format,omitempty"`
}

type ArraySchemaElement struct {
	FieldSchemaElement `json:",inline" yaml:",inline"`
	Items              SchemaElement `json:"items" yaml:"items"`
}
type ObjectSchemaElement struct {
	SchemaObject         `json:"-" yaml:"-"`
	FieldSchemaElement   `json:",inline" yaml:",inline"`
	Properties           map[string]SchemaElement `json:"properties,omitempty" yaml:"properties,omitempty"`
	AdditionalProperties SchemaElement            `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
	AllOf                []SchemaElement          `json:"allOf,omitempty" yaml:"allOf,omitempty"`
}

func (s *ObjectSchemaElement) AddProperty(name string, property SchemaElement) {
	_ = "STUB: not implemented"
	return
}

func (s *ObjectSchemaElement) AddEmbedded(element SchemaElement) {
	_ = "STUB: not implemented"
	// prevent duplicates
	return
}

type DefsSchemaElement map[string]SchemaElement

func (d DefsSchemaElement) AddDef(name string, property SchemaElement) {
	_ = "STUB: not implemented"
	return
}

type Schema struct {
	Defs                DefsSchemaElement `json:"$defs,omitempty" yaml:"$defs,omitempty"`
	ObjectSchemaElement `json:",inline" yaml:",inline"`
}

func (s *Schema) ToJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Schema) ToYAML() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func CreateSchema() *Schema { _ = "STUB: not implemented"; return nil }

func CreateSimpleField(fieldType SchemaType, description string) *FieldSchemaElement {
	_ = "STUB: not implemented"
	return nil
}

func CreateArrayField(itemType SchemaElement, description string) *ArraySchemaElement {
	_ = "STUB: not implemented"
	return nil
}

func CreateRefField(ref, description string) *RefSchemaElement {
	_ = "STUB: not implemented"
	return nil
}

func CreateObjectField(description string) *ObjectSchemaElement {
	_ = "STUB: not implemented"
	return nil
}

func CreateMapField(valueType SchemaElement, description string) *ObjectSchemaElement {
	_ = "STUB: not implemented"
	return nil
}

type SchemaType string

const (
	SchemaTypeObject  SchemaType = "object"
	SchemaTypeArray   SchemaType = "array"
	SchemaTypeString  SchemaType = "string"
	SchemaTypeInteger SchemaType = "integer"
	SchemaTypeNumber  SchemaType = "number"
	SchemaTypeBoolean SchemaType = "boolean"
	SchemaTypeAny     SchemaType = ""
	SchemaTypeUnknown SchemaType = "-"
)

func mergeSchemas(base SchemaObject, additional SchemaElement) error {
	_ = "STUB: not implemented"
	return nil
}
