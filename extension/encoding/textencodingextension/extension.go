// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package textencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/textencodingextension"

import (
	"context"
	"io"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
)

var (
	_ encoding.LogsMarshalerExtension   = (*textExtension)(nil)
	_ encoding.LogsUnmarshalerExtension = (*textExtension)(nil)
	_ encoding.LogsDecoderExtension     = (*textExtension)(nil)
)

type textExtension struct {
	config      *Config
	textEncoder *textLogCodec
}

func (e *textExtension) UnmarshalLogs(buf []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func (e *textExtension) MarshalLogs(ld plog.Logs) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *textExtension) NewLogsDecoder(reader io.Reader, options ...encoding.DecoderOption) (encoding.LogsDecoder, error) {
	_ = "STUB: not implemented"
	return *new(encoding.LogsDecoder), nil
}

func (e *textExtension) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (*textExtension) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }
