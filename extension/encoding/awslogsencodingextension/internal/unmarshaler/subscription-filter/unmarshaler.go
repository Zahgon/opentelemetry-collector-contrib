// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package subscriptionfilter // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler/subscription-filter"

import (
	"errors"
	"io"

	gojson "github.com/goccy/go-json"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler"
)

const ctrlMessageType = "CONTROL_MESSAGE"

var (
	errEmptyOwner     = errors.New("cloudwatch log with message type 'DATA_MESSAGE' has empty owner field")
	errEmptyLogGroup  = errors.New("cloudwatch log with message type 'DATA_MESSAGE' has empty log group field")
	errEmptyLogStream = errors.New("cloudwatch log with message type 'DATA_MESSAGE' has empty log stream field")
)

var _ unmarshaler.StreamingLogsUnmarshaler = (*SubscriptionFilterUnmarshaler)(nil)

// route is a resolved CloudWatchStream: patterns are pre-split, the inner
// encoding extension is already type-asserted, and payload is non-empty.
type route struct {
	name           string
	inner          plog.Unmarshaler
	logGroupParts  []string
	logStreamParts []string
	payload        PayloadMode
}

type SubscriptionFilterUnmarshaler struct {
	buildInfo component.BuildInfo
	routes    []route
}

func NewSubscriptionFilterUnmarshaler(buildInfo component.BuildInfo) *SubscriptionFilterUnmarshaler {
	_ = "STUB: not implemented"
	return nil
}

// ConfigureRoutes resolves the user's CloudWatchStream config against the
// host's loaded extensions and installs the resulting route table for
// per-envelope dispatch. Called from the extension's Start.
func (f *SubscriptionFilterUnmarshaler) ConfigureRoutes(streams []CloudWatchStream, host component.Host, selfID component.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalAWSLogs deserializes the given reader as CloudWatch Logs events
// into a plog.Logs, grouping logs by owner (account ID), log group, and
// log stream. When extracted fields are present (from centralized logging),
// logs are further grouped by their extracted account ID and region.
// Logs are assumed to be gzip-compressed as specified at
// https://docs.aws.amazon.com/firehose/latest/dev/writing-with-cloudwatch-logs.html.
func (f *SubscriptionFilterUnmarshaler) UnmarshalAWSLogs(reader io.Reader) (plog.Logs, error) {
	_ = "STUB: not implemented"
	// Decode as a stream but flush all at once using flush options
	return *new(plog.Logs), nil
}

// we must check for EOF with direct comparison and avoid wrapped EOF that can come from stream itself
//nolint:errorlint

// EOF indicates no logs were found, return any logs that's available

// NewLogsDecoder returns a LogsDecoder that processes CloudWatch Logs subscription filter events.
// Supported sub formats:
//   - DATA_MESSAGE: Returns logs grouped by owner, log group, and stream; offset is the number of records processed
//   - CONTROL_MESSAGE: Returns empty log; offset is the number of records processed
func (f *SubscriptionFilterUnmarshaler) NewLogsDecoder(reader io.Reader, options ...encoding.DecoderOption) (encoding.LogsDecoder, error) {
	_ = "STUB: not implemented"
	return *new(encoding.LogsDecoder), nil
}

// decodeOne reads one CloudWatch envelope from decoder. With no routes
// configured, the envelope is decoded once into cloudwatchLogsData and
// merged via appendLogs (existing behavior). With routes, the envelope is
// captured as raw bytes, the header is peeked to drive route selection, and
// dispatch happens per envelope.
func (f *SubscriptionFilterUnmarshaler) decodeOne(
	decoder *gojson.Decoder,
	logs plog.Logs,
	resourceLogsByKey map[resourceGroupKey]plog.LogRecordSlice,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Envelope dispatch hands the raw bytes directly to the inner encoding; no full parse needed.

// Both the no-route fallback and message dispatch need the parsed envelope.

func (f *SubscriptionFilterUnmarshaler) matchRoute(logGroup, logStream string) *route {
	_ = "STUB: not implemented"
	return nil
}

// dispatchEnvelope hands the raw CloudWatch envelope bytes to the inner
// encoding once. The inner is responsible for parsing and resource attrs.
func (*SubscriptionFilterUnmarshaler) dispatchEnvelope(logs plog.Logs, raw []byte, logGroup string, r route) error {
	_ = "STUB: not implemented"
	return nil
}

// dispatchMessage hands each event's message bytes to the inner encoding,
// then attaches CloudWatch resource attrs and back-fills the event timestamp.
func (*SubscriptionFilterUnmarshaler) dispatchMessage(logs plog.Logs, cwLog cloudwatchLogsData, r route) error {
	_ = "STUB: not implemented"
	return nil
}

func addCloudWatchResourceAttrs(attrs pcommon.Map, accountID, region, logGroup, logStream string) {
	_ = "STUB: not implemented"
	return
}

func backfillTimestamps(rl plog.ResourceLogs, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func resolveAccountAndRegion(event cloudwatchLogsLogEvent, owner string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// appendLogs appends log records from cwLog into the given plog.Logs, reusing
// existing ResourceLogs entries tracked by resourceLogsByKey when possible.
// Events are grouped by their extracted fields (account ID + region) and
// by log group/stream combination.
func (f *SubscriptionFilterUnmarshaler) appendLogs(logs plog.Logs, resourceLogsByKey map[resourceGroupKey]plog.LogRecordSlice, cwLog cloudwatchLogsData) {
	_ = "STUB: not implemented"
	return
}

// pcommon.Timestamp is a time specified as UNIX Epoch time in nanoseconds
// but timestamp in cloudwatch logs are in milliseconds.

// extractResourceKey extracts the resource group key from a log event.
// When extracted fields are present, uses those; otherwise falls back to owner.
func extractResourceKey(event cloudwatchLogsLogEvent, owner, logGroup, logStream string) resourceGroupKey {
	_ = "STUB: not implemented"
	return *new(resourceGroupKey)
}

func validateLog(log cloudwatchLogsData) error { _ = "STUB: not implemented"; return nil }

func validateLogFields(messageType, owner, logGroup, logStream string) error {
	_ = "STUB: not implemented"
	return nil
}
