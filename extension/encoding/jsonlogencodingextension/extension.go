// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jsonlogencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/jsonlogencodingextension"

import (
	"context"
	"io"

	"github.com/goccy/go-json"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
)

var (
	_ encoding.LogsMarshalerExtension   = (*jsonLogExtension)(nil)
	_ encoding.LogsUnmarshalerExtension = (*jsonLogExtension)(nil)
)

type jsonLogExtension struct {
	config *Config
}

func (e *jsonLogExtension) MarshalLogs(ld plog.Logs) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// special handling for inline attributes Mode

// check for processing mode so we can return the best format

// if multiple logs, then consider exporting as ndjson

// default mode

func (e *jsonLogExtension) UnmarshalLogs(buf []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// Default mode to handle arrays having backward compatibility

func (*jsonLogExtension) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (*jsonLogExtension) Shutdown(context.Context) error {
	_ = "STUB: not implemented"

	// streamReader is a wrapper to process input stream and return processed JSON records one by one
	return nil
}

type streamReader struct {
	decoder *json.Decoder
	current map[string]any
	err     error
	done    bool
}

func newStreamReader(r io.Reader) *streamReader { _ = "STUB: not implemented"; return nil }

func (r *streamReader) next() bool { _ = "STUB: not implemented"; return false }

// EOF signals the end

// Record error and let caller handles the result

func (r *streamReader) value() (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }
