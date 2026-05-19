// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudlogentryencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/googlecloudlogentryencodingextension"

import (
	gojson "github.com/goccy/go-json"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func setBodyFromProto(logRecord plog.LogRecord, value gojson.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (opts fieldTranslateOptions) translateValue(dst pcommon.Value, fd protoreflect.FieldDescriptor, src gojson.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// protojson represents both of these as strings

// All the wrapper types have a single field with name
// `value` and field number 1, and are represented in
// protojson without the wrapping.

// protojson accepts either string name or enum int value; try both.

// The protojson encoding accepts either string or number for
// integer types, so try both.

func (opts fieldTranslateOptions) translateList(dst pcommon.Slice, fd protoreflect.FieldDescriptor, src gojson.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func (opts fieldTranslateOptions) translateMap(dst pcommon.Map, fd protoreflect.FieldDescriptor, src gojson.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func translateAny(dst pcommon.Map, src map[string]gojson.RawMessage) error {
	_ = "STUB: not implemented"
	// protojson represents Any as the JSON representation of the actual
	// message, plus a special @type field containing the type URL of the
	// message.
	return nil
}

// If we don't have the type, we do a best-effort JSON decode;
// some ints might be floats or strings.

func translateProtoMessage(dst pcommon.Map, desc protoreflect.MessageDescriptor, src map[string]gojson.RawMessage, opts fieldTranslateOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// Handle well-known aggregate types.

func translateInto(dst pcommon.Map, desc protoreflect.MessageDescriptor, src any, opts ...fieldTranslateFn) error {
	_ = "STUB: not implemented"
	return nil
}

func translateStr(dst pcommon.Value, src gojson.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func translateRaw(dst pcommon.Value, src gojson.RawMessage) error {
	_ = "STUB: not implemented"
	return nil
}

func getTokenType(src gojson.RawMessage) string { _ = "STUB: not implemented"; return "" }
