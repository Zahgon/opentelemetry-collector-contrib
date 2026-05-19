// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package avrologencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/avrologencodingextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
)

var _ encoding.LogsUnmarshalerExtension = (*avroLogExtension)(nil)

type avroLogExtension struct {
	deserializer avroDeserializer
}

func newExtension(config *Config) (*avroLogExtension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *avroLogExtension) UnmarshalLogs(buf []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// removes time.Time values as FromRaw does not support it

// Set the unmarshaled avro as the body of the log record

func replaceLogicalTypes(m map[string]any) { _ = "STUB: not implemented"; return }

func transformValue(value any) any { _ = "STUB: not implemented"; return *new(any) }

func (*avroLogExtension) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (*avroLogExtension) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }
