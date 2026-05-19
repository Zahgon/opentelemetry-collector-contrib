// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tinybirdexporter/internal"

import (
	"go.opentelemetry.io/collector/pdata/plog"
)

type logSignal struct {
	ResourceSchemaURL  string            `json:"resource_schema_url"`
	ResourceAttributes map[string]string `json:"resource_attributes"`
	ServiceName        string            `json:"service_name"`
	ScopeSchemaURL     string            `json:"scope_schema_url"`
	ScopeAttributes    map[string]string `json:"scope_attributes"`
	ScopeName          string            `json:"scope_name"`
	ScopeVersion       string            `json:"scope_version"`
	Timestamp          string            `json:"timestamp"`
	TraceID            string            `json:"trace_id"`
	SpanID             string            `json:"span_id"`
	Flags              uint32            `json:"flags"`
	SeverityText       string            `json:"severity_text"`
	SeverityNumber     int32             `json:"severity_number"`
	LogAttributes      map[string]string `json:"log_attributes"`
	Body               string            `json:"body"`
}

func ConvertLogs(ld plog.Logs, encoder Encoder) error { _ = "STUB: not implemented"; return nil }
