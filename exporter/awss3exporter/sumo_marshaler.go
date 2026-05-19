// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awss3exporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter"

import (
	"bytes"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

const (
	logBodyKey = "log"
)

type sumoMarshaler struct{}

func (*sumoMarshaler) format() string { _ = "STUB: not implemented"; return "" }

func newSumoICMarshaler() sumoMarshaler { _ = "STUB: not implemented"; return *new(sumoMarshaler) }

func logEntry(buf *bytes.Buffer, format string, a ...any) { _ = "STUB: not implemented"; return }

func attributeValueToString(v pcommon.Value) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func valueToJSON(m any) (string, error) { _ = "STUB: not implemented"; return "", nil }

const (
	SourceCategoryKey = "_sourceCategory"
	SourceHostKey     = "_sourceHost"
	SourceNameKey     = "_sourceName"
)

func (sumoMarshaler) MarshalLogs(ld plog.Logs) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// continue to next attribute

func getMessageJSON(lr plog.LogRecord) (string, error) {
	_ = "STUB: not implemented"
	// The "message" fields is a JSON created from combining the actual log body and log-level attributes,
	// where the log body is stored under "log" key.
	// More info:
	// https://help.sumologic.com/docs/send-data/opentelemetry-collector/data-source-configurations/additional-configurations-reference/#mapping-opentelemetry-concepts-to-sumo-logic
	// Create a new map to avoid mutating the original log record
	return "", nil
}

func (s sumoMarshaler) MarshalTraces(_ ptrace.Traces) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s sumoMarshaler) MarshalMetrics(_ pmetric.Metrics) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
