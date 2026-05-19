// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package unmarshaler // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver/internal/unmarshaler"
import (
	"go.opentelemetry.io/collector/pdata/plog"
	"golang.org/x/text/encoding"
)

var _ plog.Unmarshaler = (*TextLogsUnmarshaler)(nil)

type TextLogsUnmarshaler struct {
	decoder *encoding.Decoder
}

func NewTextLogsUnmarshaler(encodingName string) (*TextLogsUnmarshaler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *TextLogsUnmarshaler) UnmarshalLogs(buf []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}
