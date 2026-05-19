// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package unmarshaler // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver/internal/unmarshaler"
import (
	"go.opentelemetry.io/collector/pdata/plog"
)

var _ plog.Unmarshaler = JSONLogsUnmarshaler{}

type JSONLogsUnmarshaler struct{}

func (JSONLogsUnmarshaler) UnmarshalLogs(buf []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	// create a new Logs struct to be populated with log data and returned
	return *new(plog.Logs), nil
}

// get json logs from the buffer

// create a new log record

// Set the unmarshaled jsonVal as the body of the log record
