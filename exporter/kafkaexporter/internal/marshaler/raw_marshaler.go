// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package marshaler // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter/internal/marshaler"

import (
	"errors"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

var (
	errUnsupported               = errors.New("unsupported serialization")
	_              LogsMarshaler = RawLogsMarshaler{}
)

type RawLogsMarshaler struct{}

func (r RawLogsMarshaler) MarshalLogs(logs plog.Logs, yield func(key, value []byte)) error {
	_ = "STUB: not implemented"
	return nil
}

func (r RawLogsMarshaler) logBodyAsBytes(value pcommon.Value) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (RawLogsMarshaler) interfaceAsBytes(value any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
