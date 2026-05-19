// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Inspired from "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/serializer/map.go"

package serializer // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opensearchexporter/internal/serializer"

import (
	"bytes"

	"github.com/elastic/go-structform/json"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func Map(m pcommon.Map, buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

// Enable ExplicitRadixPoint such that 1.0 is encoded as 1.0 instead of 1.
// This is required to generate the correct dynamic mapping in ES.

func writeMap(v *json.Visitor, m pcommon.Map, stringifyMapValues bool) {
	_ = "STUB: not implemented"
	return
}

func WriteValue(v *json.Visitor, val pcommon.Value, stringifyMaps bool) {
	_ = "STUB: not implemented"
	return
}
