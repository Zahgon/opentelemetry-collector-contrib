// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// The objmodel package provides tools for converting OpenTelemetry Log records into
// JSON documents.
//
// The JSON parsing in Elasticsearch does not support parsing JSON documents
// with duplicate fields. The fields in the document can be sort and duplicate entries
// can be removed before serializing. Deduplication ensures that ambiguous
// events can still be indexed.
//
// With attributes map encoded as a list of key value
// pairs, we might find some structured loggers that create log records with
// duplicate fields. Although the AttributeMap wrapper tries to give a
// dictionary like view into the list, it is not 'complete'. When iterating the map
// for encoding, we still will encounter the duplicates.
// The AttributeMap helpers treat the first occurrence as the actual field.
// For high-performance structured loggers (e.g. zap) the AttributeMap
// semantics are not necessarily correct. Most often the last occurrence will be
// what we want to export, as the last occurrence represents the last overwrite
// within a context/dictionary (the leaf-logger its context).
// Some Loggers might even allow users to create a mix of dotted and dedotted fields.
// The Document type also tries to combine these into a proper structure, such that these mixed
// representations have a unique encoding only, which allows us to properly remove duplicates.
//
// The `.` is special to Elasticsearch. In order to handle common prefixes and attributes
// being a mix of key value pairs with dots and complex objects, we flatten the document first
// before we deduplicate. Final dedotting is optional and only required when
// Ingest Node is used. But either way, we try to present only well formed
// document to Elasticsearch.

package objmodel // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/objmodel"

import (
	"io"
	"time"

	"github.com/elastic/go-structform/json"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// Document is an intermediate representation for converting open telemetry records with arbitrary attributes
// into a JSON document that can be processed by Elasticsearch.
type Document struct {
	fields           []field
	dynamicTemplates map[string]string
}

type field struct {
	key   string
	value Value
}

// Value type that can be added to a Document.
type Value struct {
	kind Kind
	ui   uint64
	i    int64
	dbl  float64
	str  string
	arr  []Value
	doc  Document
	ts   time.Time
}

// Kind represent the internal kind of a value stored in a Document.
type Kind uint8

// Enum values for Kind.
const (
	KindNil Kind = iota
	KindBool
	KindInt
	KindUInt
	KindDouble
	KindString
	KindArr
	KindObject
	KindTimestamp
	KindIgnore
	KindUnflattenableObject // Unflattenable object is an object that should not be flattened at serialization time
)

const tsLayout = "2006-01-02T15:04:05.000000000Z"

var (
	nilValue    = Value{kind: KindNil}
	ignoreValue = Value{kind: KindIgnore}
)

// DocumentFromAttributes creates a document from a OpenTelemetry attribute
// map. All nested maps will be flattened, with keys being joined using a `.` symbol.
func DocumentFromAttributes(am pcommon.Map) Document {
	_ = "STUB: not implemented"
	return *new(Document)
}

// DocumentFromAttributesWithPath creates a document from a OpenTelemetry attribute
// map. All nested maps will be flattened, with keys being joined using a `.` symbol.
//
// All keys in the map will be prefixed with path.
func DocumentFromAttributesWithPath(path string, am pcommon.Map) Document {
	_ = "STUB: not implemented"
	return *new(Document)
}

func (doc *Document) Clone() *Document { _ = "STUB: not implemented"; return nil }

func (doc *Document) AddDynamicTemplate(path, template string) { _ = "STUB: not implemented"; return }

func (doc *Document) DynamicTemplates() map[string]string { _ = "STUB: not implemented"; return nil }

// AddTimestamp adds a raw timestamp value to the Document.
func (doc *Document) AddTimestamp(key string, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

// Add adds a converted value to the document.
func (doc *Document) Add(key string, v Value) { _ = "STUB: not implemented"; return }

// AddString adds a string to the document.
func (doc *Document) AddString(key, v string) { _ = "STUB: not implemented"; return }

// AddSpanID adds the hex presentation of a SpanID to the document. If the SpanID
// is empty, no value will be added.
func (doc *Document) AddSpanID(key string, id pcommon.SpanID) { _ = "STUB: not implemented"; return }

// AddTraceID adds the hex presentation of a TraceID value to the document. If the TraceID
// is empty, no value will be added.
func (doc *Document) AddTraceID(key string, id pcommon.TraceID) { _ = "STUB: not implemented"; return }

// AddInt adds an integer value to the document.
func (doc *Document) AddInt(key string, value int64) { _ = "STUB: not implemented"; return }

// AddUInt adds an unsigned integer value to the document.
func (doc *Document) AddUInt(key string, value uint64) { _ = "STUB: not implemented"; return }

// AddAttributes expands and flattens all key-value pairs from the input attribute map into
// the document.
func (doc *Document) AddAttributes(key string, attributes pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

// AddAttribute converts and adds a AttributeValue to the document. If the attribute represents a map,
// the fields will be flattened.
func (doc *Document) AddAttribute(key string, attribute pcommon.Value) {
	_ = "STUB: not implemented"
	return
}

// do not add 'null'

// AddEvents converts and adds span events to the document.
func (doc *Document) AddEvents(key string, events ptrace.SpanEventSlice) {
	_ = "STUB: not implemented"
	return
}

// AddLinks adds a slice of span links to the document.
func (doc *Document) AddLinks(key string, links ptrace.SpanLinkSlice) {
	_ = "STUB: not implemented"
	return
}

func (doc *Document) sort() { _ = "STUB: not implemented"; return }

// Dedup removes fields from the document, that have duplicate keys.
// The filtering only keeps the last value for a key.
// protectedSet is an optional map of field paths that should never get .value suffix.
// Dedup ensure that keys are sorted.
func (doc *Document) Dedup(protectedSet map[string]struct{}) {
	_ = "STUB: not implemented"
	// 1. Always ensure the fields are sorted, Dedup support requires
	// Fields to be sorted.
	return
}

// 2. rename fields if a primitive value is overwritten by an object,
//    EXCEPT for protected fields (well-defined schema fields like ECS).
//    For example the pair (path.x=1, path.x.a="test") becomes:
//    (path.x.value=1, path.x.a="test").
//
//    However, if path.x is a protected field, we skip the renaming and instead
//    remove the conflicting nested field (path.x.a) to preserve the protected field.
//
//    NOTE: We do the renaming, in order to preserve the original value
//    in case of conflicts after dedotting, which would lead to the removal of the field.
//    For example docker/k8s labels tend to use `.`, which need to be handled in case
//    The collector does pass us these kind of labels as an AttributeMap.
//
//    NOTE: If the embedded document already has a field name `value`, we will remove the renamed
//    field in favor of the `value` field in the document.
//
//    This step removes potential conflicts when dedotting and serializing fields.

// This is a protected field - mark all nested fields under it as ignore.

// Normal case: rename to .value

// 3. mark duplicates as 'ignore'
//
//    This step ensures that we do not have duplicate fields names when serializing.
//    Elasticsearch JSON parser will fail otherwise.

// 4. fix objects that might be stored in arrays

func newJSONVisitor(w io.Writer) *json.Visitor { _ = "STUB: not implemented"; return nil }

// Enable ExplicitRadixPoint such that 1.0 is encoded as 1.0 instead of 1.
// This is required to generate the correct dynamic mapping in ES.

// Serialize writes the document to the given writer. The document fields will be
// deduplicated and, if dedot is true, turned into nested objects prior to
// serialization.
func (doc *Document) Serialize(w io.Writer, dedot bool, protectedSet map[string]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (doc *Document) iterJSON(v *json.Visitor, dedot bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (doc *Document) iterJSONFlat(w *json.Visitor) error { _ = "STUB: not implemented"; return nil }

func (doc *Document) iterJSONDedot(w *json.Visitor) error { _ = "STUB: not implemented"; return nil }

// decrease object level until last reported and current key have the same path prefix

// remove levels and append write list of outstanding '}' into the writer

// no common prefix, close all objects we reported so far.

// increase object level up to current field

// report value

// close all pending object levels

// StringValue create a new value from a string.
func StringValue(str string) Value { _ = "STUB: not implemented"; return *new(Value) }

// IntValue creates a new value from an integer.
func IntValue(i int64) Value { _ = "STUB: not implemented"; return *new(Value) }

// UIntValue creates a new value from an unsigned integer.
func UIntValue(i uint64) Value { _ = "STUB: not implemented"; return *new(Value) }

// DoubleValue creates a new value from a double value..
func DoubleValue(d float64) Value { _ = "STUB: not implemented"; return *new(Value) }

// BoolValue creates a new value from a double value..
func BoolValue(b bool) Value { _ = "STUB: not implemented"; return *new(Value) }

// ArrValue combines multiple values into an array value.
func ArrValue(values ...Value) Value { _ = "STUB: not implemented"; return *new(Value) }

// TimestampValue create a new value from a time.Time.
func TimestampValue(ts time.Time) Value { _ = "STUB: not implemented"; return *new(Value) }

// UnflattenableObjectValue creates a unflattenable object from a map
func UnflattenableObjectValue(m pcommon.Map) Value { _ = "STUB: not implemented"; return *new(Value) }

// ValueFromAttribute converts a AttributeValue into a value.
func ValueFromAttribute(attr pcommon.Value) Value { _ = "STUB: not implemented"; return *new(Value) }

func (v *Value) sort() { _ = "STUB: not implemented"; return }

// Dedup recursively dedups keys in stored documents.
//
// NOTE: The value MUST be sorted.
func (v *Value) Dedup(protectedSet map[string]struct{}) { _ = "STUB: not implemented"; return }

func (v *Value) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (v *Value) iterJSON(w *json.Visitor, dedot bool) error { _ = "STUB: not implemented"; return nil }

// NaN and Inf are undefined for JSON. Let's serialize to "null"

func arrFromAttributes(aa pcommon.Slice) []Value { _ = "STUB: not implemented"; return nil }

func appendAttributeFields(fields []field, path string, am pcommon.Map) []field {
	_ = "STUB: not implemented"
	return nil
}

func appendAttributeValue(fields []field, path, key string, attr pcommon.Value) []field {
	_ = "STUB: not implemented"
	return nil
}

func flattenKey(path, key string) string { _ = "STUB: not implemented"; return "" }

func commonObjPrefix(a, b string) int { _ = "STUB: not implemented"; return 0 }
