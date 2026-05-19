// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awscloudwatchlogsexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awscloudwatchlogsexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	exp "go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/cwlogs"
)

type cwlExporter struct {
	Config           *Config
	logger           *zap.Logger
	collectorID      string
	svcStructuredLog *cwlogs.Client
	pusherFactory    cwlogs.MultiStreamPusherFactory
}

type awsMetadata struct {
	LogGroupName  string `json:"logGroupName,omitempty"`
	LogStreamName string `json:"logStreamName,omitempty"`
}

type emfMetadata struct {
	AWSMetadata   *awsMetadata `json:"_aws,omitempty"`
	LogGroupName  string       `json:"log_group_name,omitempty"`
	LogStreamName string       `json:"log_stream_name,omitempty"`
}

func newCwLogsPusher(ctx context.Context, expConfig *Config, params exp.Settings) (*cwlExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create CWLogs client with aws session config

func newCwLogsExporter(ctx context.Context, config component.Config, params exp.Settings) (exp.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exp.Logs), nil
}

func (e *cwlExporter) consumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (*cwlExporter) shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func pushLogsToCWLogs(ctx context.Context, logger *zap.Logger, ld plog.Logs, config *Config, pusher cwlogs.Pusher) error {
	_ = "STUB: not implemented"
	return nil
}

type scopeCwLogBody struct {
	Name       string         `json:"name,omitempty"`
	Version    string         `json:"version,omitempty"`
	Attributes map[string]any `json:"attributes,omitempty"`
}

type cwLogBody struct {
	Body                   any             `json:"body,omitempty"`
	SeverityNumber         int32           `json:"severity_number,omitempty"`
	SeverityText           string          `json:"severity_text,omitempty"`
	DroppedAttributesCount uint32          `json:"dropped_attributes_count,omitempty"`
	Flags                  uint32          `json:"flags,omitempty"`
	TraceID                string          `json:"trace_id,omitempty"`
	SpanID                 string          `json:"span_id,omitempty"`
	Attributes             map[string]any  `json:"attributes,omitempty"`
	Scope                  *scopeCwLogBody `json:"scope,omitempty"`
	Resource               map[string]any  `json:"resource,omitempty"`
}

func logToCWLog(resourceAttrs map[string]any, scope pcommon.InstrumentationScope, log plog.LogRecord, config *Config) (*cwlogs.Event, error) {
	_ = "STUB: not implemented"
	// TODO(jbd): Benchmark and improve the allocations.
	// Evaluate go.elastic.co/fastjson as a replacement for encoding/json.
	// Replace loggroup and logstream with resource attribute
	return nil, nil
}

// Check if this is an emf log

// v1 emf json

/* v0 emf json */

// scope should have a name at least

// in milliseconds

func attrsValue(attrs pcommon.Map) map[string]any { _ = "STUB: not implemented"; return nil }
