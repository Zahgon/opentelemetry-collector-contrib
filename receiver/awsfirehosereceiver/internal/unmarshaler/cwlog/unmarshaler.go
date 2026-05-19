// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cwlog // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwlog"

import (
	"errors"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

const (
	TypeStr = "cwlogs"

	attributeAWSCloudWatchLogGroupName  = "aws.cloudwatch.log_group_name"
	attributeAWSCloudWatchLogStreamName = "aws.cloudwatch.log_stream_name"
)

var (
	errInvalidRecords   = errors.New("record format invalid")
	errMissingOwner     = errors.New("cloudwatch log record is missing owner field")
	errMissingLogGroup  = errors.New("cloudwatch log record is missing logGroup field")
	errMissingLogStream = errors.New("cloudwatch log record is missing logStream field")
)

// Unmarshaler for the CloudWatch Log JSON record format.
type Unmarshaler struct {
	logger    *zap.Logger
	buildInfo component.BuildInfo
	gzipPool  sync.Pool
}

var _ plog.Unmarshaler = (*Unmarshaler)(nil)

// NewUnmarshaler creates a new instance of the Unmarshaler.
func NewUnmarshaler(logger *zap.Logger, buildInfo component.BuildInfo) *Unmarshaler {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalLogs deserializes the given record as CloudWatch Logs events
// into a plog.Logs, grouping logs by owner (account ID), log group, and
// log stream. Logs are assumed to be gzip-compressed as specified at
// https://docs.aws.amazon.com/firehose/latest/dev/writing-with-cloudwatch-logs.html.
func (u *Unmarshaler) UnmarshalLogs(compressedRecord []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// Deprecated: [v0.121.0] Use `string(conventions.AWSLogGroupNamesKey)` instead

// Deprecated: [v0.121.0] Use `string(conventions.AWSLogStreamNamesKey)` instead

// pcommon.Timestamp is a time specified as UNIX Epoch time in nanoseconds
// but timestamp in cloudwatch logs are in milliseconds.

func parseLog(data []byte) (log cWLog, control bool, _ error) {
	_ = "STUB: not implemented"
	return *new(cWLog), false, nil
}

// Type of the serialized messages.
func (*Unmarshaler) Type() string { _ = "STUB: not implemented"; return "" }
