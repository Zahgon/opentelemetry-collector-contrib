// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logdedupprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/logdedupprocessor"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

const (
	// fieldDelimiter is the delimiter used to split a field key into its parts.
	fieldDelimiter = "."

	// fieldEscapeKeyReplacement is the string used to temporarily replace escaped delimiters while splitting a field key.
	fieldEscapeKeyReplacement = "{TEMP_REPLACE}"
)

// fieldRemover handles removing excluded fields from log records
type fieldRemover struct {
	fields []*field
}

// field represents a field and it's compound key to match on
type field struct {
	keyParts []string
}

// newFieldRemover creates a new field remover based on the passed in field keys
func newFieldRemover(fieldKeys []string) *fieldRemover { _ = "STUB: not implemented"; return nil }

// RemoveFields removes any body or attribute fields that match in the log record
func (fe *fieldRemover) RemoveFields(logRecord plog.LogRecord) { _ = "STUB: not implemented"; return }

// removeField removes the field from the log record if it exists
func (f *field) removeField(logRecord plog.LogRecord) { _ = "STUB: not implemented"; return }

// If body is a map then recurse through to remove the field

// Remove all attributes

// Recurse through map and remove fields

// removeFieldFromMap recurses through the map and removes the field if it's found.
func removeFieldFromMap(valueMap pcommon.Map, keyParts []string) { _ = "STUB: not implemented"; return }

// Look for the value associated with the next key part.
// If we don't find it then return

// No more key parts that means we have found the value and remove it

// If the value is a map then recurse through with the remaining parts

// splitField splits a field key into its parts.
// It replaces escaped delimiters with the full delimiter after splitting.
func splitField(fieldKey string) []string { _ = "STUB: not implemented"; return nil }

// Replace the temporarily escaped delimiters with the actual delimiter.
