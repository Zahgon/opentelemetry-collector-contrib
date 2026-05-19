// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package textencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/textencodingextension"

import (
	"io"
	"regexp"

	"go.opentelemetry.io/collector/pdata/plog"
	txt "golang.org/x/text/encoding"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
)

type textLogCodec struct {
	decoder               *txt.Decoder
	marshalingSeparator   string
	unmarshalingSeparator *regexp.Regexp
}

func (r *textLogCodec) UnmarshalLogs(buf []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	// Decode as a stream but flush all at once using flush options
	return *new(plog.Logs), nil
}

// NewLogsDecoder implements the encoding.LogsCodec interface. Tracks offset by bytes read from the stream.
func (r *textLogCodec) NewLogsDecoder(reader io.Reader, options ...encoding.DecoderOption) (encoding.LogsDecoder, error) {
	_ = "STUB: not implemented"
	return *new(encoding.LogsDecoder), nil
}

// Discard non-zero offset from the reader before scanning for log records

// Request more data until EOF

// check for stream EOF which results in empty log batch

func (r *textLogCodec) MarshalLogs(ld plog.Logs) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
