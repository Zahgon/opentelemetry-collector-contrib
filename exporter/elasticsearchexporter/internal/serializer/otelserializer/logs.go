// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelserializer // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/serializer/otelserializer"

import (
	"bytes"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/elasticsearch"
)

func (*Serializer) SerializeLog(resource pcommon.Resource, resourceSchemaURL string, scope pcommon.InstrumentationScope, scopeSchemaURL string, record plog.LogRecord, idx elasticsearch.Index, buf *bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func writeLogBody(w *jsonWriter, record plog.LogRecord, first bool) bool {
	_ = "STUB: not implemented"
	return false
}

// output must be an array of objects due to ES limitations
// otherwise, wrap the array in an object
