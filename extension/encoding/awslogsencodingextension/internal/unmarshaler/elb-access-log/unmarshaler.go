// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elbaccesslogs // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler/elb-access-log"

import (
	"bufio"
	"io"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler"
)

var _ unmarshaler.StreamingLogsUnmarshaler = (*ElbAccessLogUnmarshaler)(nil)

type ElbAccessLogUnmarshaler struct {
	buildInfo component.BuildInfo
	logger    *zap.Logger
}

func NewELBAccessLogUnmarshaler(buildInfo component.BuildInfo, logger *zap.Logger) *ElbAccessLogUnmarshaler {
	_ = "STUB: not implemented"
	return nil
}

type resourceAttributes struct {
	resourceID string
}

// UnmarshalAWSLogs processes all logs from the provided reader
func (f *ElbAccessLogUnmarshaler) UnmarshalAWSLogs(reader io.Reader) (plog.Logs, error) {
	_ = "STUB: not implemented"
	// Decode as a stream but flush all at once using flush options
	return *new(plog.Logs), nil
}

// we must check for EOF with direct comparison and avoid wrapped EOF that can come from stream itself
//nolint:errorlint

// EOF indicates no logs were found, return any logs that's available

// NewLogsDecoder returns a LogsDecoder that processes ELB access logs from the provided reader.
// Auto-detects the ELB Log type (ALB, NLB, CLB, or control message) using the first log line.
// Supported sub formats:
//   - ALB/NLB/CLB access logs: Supports offset-based streaming; offset tracks bytes processed
//   - Control message: Returns empty log; offset tracks bytes processed
func (f *ElbAccessLogUnmarshaler) NewLogsDecoder(reader io.Reader, options ...encoding.DecoderOption) (encoding.LogsDecoder, error) {
	_ = "STUB: not implemented"
	return *new(encoding.LogsDecoder), nil
}

// Process line based on syntax

// Control messages are conveyed as empty logs

// return empty log with EOF

// createLogs with the expected fields for the scope logs
func (f *ElbAccessLogUnmarshaler) createLogs() (plog.Logs, plog.ResourceLogs, plog.ScopeLogs) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), *new(plog.ResourceLogs), *new(plog.ScopeLogs)
}

// setResourceAttributes based on the resourceAttributes
func (*ElbAccessLogUnmarshaler) setResourceAttributes(r *resourceAttributes, logs plog.ResourceLogs) {
	_ = "STUB: not implemented"
	return
}

// handleCLBAccessLogs handles clb access logs
func (f *ElbAccessLogUnmarshaler) handleCLBAccessLogs(fields []string, resourceAttr *resourceAttributes, scopeLogs plog.ScopeLogs) error {
	_ = "STUB: not implemented"
	return nil
}

// addToCLBAccessLogs adds clb record to provided logs based
// on the extracted logs of each resource
func (f *ElbAccessLogUnmarshaler) addToCLBAccessLogs(resourceAttr *resourceAttributes, scopeLogs plog.ScopeLogs, clbRecord CLBAccessLogRecord) {
	_ = "STUB: not implemented"
	// Convert timestamp first; if invalid, skip log creation
	return
}

// Create record log

// Set resource id

// Populate record attributes

// Set timestamp

// move recordLog to scope

// handleALBAccessLogs handles alb access logs
func (f *ElbAccessLogUnmarshaler) handleALBAccessLogs(fields []string, resourceAttr *resourceAttributes, scopeLogs plog.ScopeLogs) error {
	_ = "STUB: not implemented"
	return nil
}

// addToALBAccessLogs adds alb record to provided logs based
// on the extracted logs of each resource
func (f *ElbAccessLogUnmarshaler) addToALBAccessLogs(resourceAttr *resourceAttributes, scopeLogs plog.ScopeLogs, albRecord ALBAccessLogRecord) {
	_ = "STUB: not implemented"
	// Convert timestamp first; if invalid, skip log creation
	return
}

// Create record log

// Set resource id

// Populate record attributes

// Times are expressed in seconds with a precision of 3 decimal places in logs. Here we convert them to milliseconds.

// Set timestamp

// move recordLog to scope

// handleNLBAccessLogs handles nlb access logs
func (f *ElbAccessLogUnmarshaler) handleNLBAccessLogs(fields []string, resourceAttr *resourceAttributes, scopeLogs plog.ScopeLogs) error {
	_ = "STUB: not implemented"
	return nil
}

// addToNLBAccessLogs adds nlb record to provided logs based
// on the extracted logs of each resource
func (f *ElbAccessLogUnmarshaler) addToNLBAccessLogs(resourceAttr *resourceAttributes, scopeLogs plog.ScopeLogs, nlbRecord NLBAccessLogRecord) {
	_ = "STUB: not implemented"
	// Convert timestamp first; if invalid, skip log creation
	return
}

// Create record log

// Set resource id

// Populate record attributes

// Attributes below may be unset (set to "-") in logs

// Set timestamp

// move recordLog to scope

func peekAndGetSyntax(bufReader *bufio.Reader) (logSyntaxType, error) {
	_ = "STUB: not implemented"
	// 100 bytes should be enough to cover first sections of the log line
	return *new(logSyntaxType), nil
}

// Check for control message

// Determine syntax
