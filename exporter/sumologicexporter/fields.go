// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumologicexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sumologicexporter"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// fields represents metadata
type fields struct {
	orig        pcommon.Map
	initialized bool
}

func newFields(attrMap pcommon.Map) fields { _ = "STUB: not implemented"; return *new(fields) }

func (f fields) isInitialized() bool { _ = "STUB: not implemented"; return false }

// string returns fields as ordered key=value string with `, ` as separator
func (f fields) string() string { _ = "STUB: not implemented"; return "" }

// Don't add source related attributes to fields as they are handled separately
// and are added to the payload either as special HTTP headers or as resources
// attributes.

// Skip empty field

// sanitizeFields sanitize field (key or value) to be correctly parsed by sumologic receiver
// It modifies the field in place.
func (fields) sanitizeField(fld []byte) { _ = "STUB: not implemented"; return }
