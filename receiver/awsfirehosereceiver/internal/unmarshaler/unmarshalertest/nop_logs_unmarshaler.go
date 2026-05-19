// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package unmarshalertest // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/unmarshalertest"

import (
	"go.opentelemetry.io/collector/pdata/plog"
)

// NopLogsUnmarshaler is a LogsUnmarshaler that doesn't do anything
// with the inputs and just returns the logs and error passed in.
type NopLogsUnmarshaler struct {
	logs plog.Logs
	err  error
}

var _ plog.Unmarshaler = (*NopLogsUnmarshaler)(nil)

// NewNopLogs provides a nop logs unmarshaler with the default
// plog.Logs and no error.
func NewNopLogs() *NopLogsUnmarshaler { _ = "STUB: not implemented"; return nil }

// NewWithLogs provides a nop logs unmarshaler with the passed
// in logs as the result of the Unmarshal and no error.
func NewWithLogs(logs plog.Logs) *NopLogsUnmarshaler { _ = "STUB: not implemented"; return nil }

// NewErrLogs provides a nop logs unmarshaler with the passed
// in error as the Unmarshal error.
func NewErrLogs(err error) *NopLogsUnmarshaler { _ = "STUB: not implemented"; return nil }

// Unmarshal deserializes the records into logs.
func (u *NopLogsUnmarshaler) UnmarshalLogs([]byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *

	// Type of the serialized messages.
	new(plog.Logs), nil
}

func (*NopLogsUnmarshaler) Type() string { _ = "STUB: not implemented"; return "" }
