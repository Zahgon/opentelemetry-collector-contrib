// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awslambdareceiver/internal"

import (
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
)

// NewDefaultS3LogsDecoder returns a defaultS3Unmarshaler wrapped as an encoding.LogsDecoderFactory.
func NewDefaultS3LogsDecoder() encoding.LogsDecoderFactory {
	_ = "STUB: not implemented"
	return *new(encoding.LogsDecoderFactory)
}

// defaultS3Unmarshaler defines the default S3 logs decoder for AWS Lambda receiver.
type defaultS3Unmarshaler struct{}

// UnmarshalLogs defines the built-in behavior for S3 events when no encoding extension is provided.
func (*defaultS3Unmarshaler) UnmarshalLogs(data []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// NewDefaultCWLogsDecoder returns a defaultCWLogsDecoder wrapped as an encoding.LogsDecoderFactory.
func NewDefaultCWLogsDecoder() encoding.LogsDecoderFactory {
	_ = "STUB: not implemented"
	return *new(encoding.LogsDecoderFactory)
}

// defaultCWLogsDecoder defines the default CloudWatch logs decoder for AWS Lambda receiver.
type defaultCWLogsDecoder struct{}

// UnmarshalLogs defines the built-in behavior for CloudWatch logs events when no encoding extension is provided.
func (*defaultCWLogsDecoder) UnmarshalLogs(data []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// pcommon.Timestamp is a time specified as UNIX Epoch time in nanoseconds
// but timestamp in cloudwatch logs are in milliseconds.
