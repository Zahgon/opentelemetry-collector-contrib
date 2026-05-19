// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package s3accesslog // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler/s3-access-log"

import (
	"io"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler"
)

const (
	// any field can be set to - to indicate that the data was unknown
	// or unavailable, or that the field was not applicable to this request.
	//
	// See https://docs.aws.amazon.com/AmazonS3/latest/userguide/LogFormat.html.
	unknownField = "-"
)

var _ unmarshaler.StreamingLogsUnmarshaler = (*S3AccessLogUnmarshaler)(nil)

type S3AccessLogUnmarshaler struct {
	buildInfo component.BuildInfo
}

func NewS3AccessLogUnmarshaler(buildInfo component.BuildInfo) *S3AccessLogUnmarshaler {
	_ = "STUB: not implemented"
	return nil
}

type resourceAttributes struct {
	bucketOwner string
	bucketName  string
}

func (s *S3AccessLogUnmarshaler) UnmarshalAWSLogs(reader io.Reader) (plog.Logs, error) {
	_ = "STUB: not implemented"
	// Decode as a stream but flush all at once using flush options
	return *new(plog.Logs), nil
}

// we must check for EOF with direct comparison and avoid wrapped EOF that can come from stream itself
//nolint:errorlint

// EOF indicates no logs were found, return any logs that's available

// NewLogsDecoder returns a LogsDecoder that processes S3 access logs from the provided reader.
// Parses space-delimited log lines following the S3 server access log format.
// Supports offset-based streaming; offset tracks bytes processed
func (s *S3AccessLogUnmarshaler) NewLogsDecoder(reader io.Reader, options ...encoding.DecoderOption) (encoding.LogsDecoder, error) {
	_ = "STUB: not implemented"
	return *new(encoding.LogsDecoder), nil
}

// createLogs with the expected fields for the scope logs
func (s *S3AccessLogUnmarshaler) createLogs() (plog.Logs, plog.ResourceLogs, plog.ScopeLogs) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), *new(plog.ResourceLogs), *new(plog.ScopeLogs)
}

// setResourceAttributes based on the resourceAttributes
func (*S3AccessLogUnmarshaler) setResourceAttributes(r *resourceAttributes, logs plog.ResourceLogs) {
	_ = "STUB: not implemented"
	return
}

// scanField gets the next value in the log line by moving
// one space. If the value starts with a quote, it moves
// forward until it finds the ending quote, and considers
// that just 1 value. Otherwise, it returns the value as
// it is.
func scanField(logLine string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// if there is a quote, we need to get the rest of the value
// remove first quote

// Remove space after closing quote if present

func handleLog(resourceAttr *resourceAttributes, scopeLogs plog.ScopeLogs, log string) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip unknown fields for forward compatibility with future AWS S3
// access log format additions. AWS may append new fields without a
// format version bump.
// See https://docs.aws.amazon.com/AmazonS3/latest/userguide/LogFormat.html.

// acl required field can be '-' to indicate that no ACL was required

// This is the timestamp that follows a strict format
// "[DD/MM/YYYY:HH:mm:ss zone]". Since zone is after a
// space, we cut the string again to remove the zone.

func addField(field int, value string, resourceAttr *resourceAttributes, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// If the request required an ACL for authorization, the string is Yes.
// If no ACLs were required, the string is -.

// The value is one of following: TLSv1.1, TLSv1.2, TLSv1.3.

// netProtocol returns protocol name and version based on proto value
func netProtocol(proto string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
