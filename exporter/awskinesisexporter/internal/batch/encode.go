// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package batch // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter/internal/batch"

import (
	"errors"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var (
	// ErrUnsupportedEncoding is used when the encoder type does not support the type of encoding
	ErrUnsupportedEncoding = errors.New("unsupported type to encode")
	// ErrUnknownExportEncoder is used when a named encoding doesn't not exist
	ErrUnknownExportEncoder = errors.New("unknown encoding export format")
)

// Encoder transforms the internal pipeline format into a configurable
// format that is then used to export to kinesis.
type Encoder interface {
	Metrics(md pmetric.Metrics) (*Batch, error)

	Traces(td ptrace.Traces) (*Batch, error)

	Logs(ld plog.Logs) (*Batch, error)
}

func NewEncoder(named string, batchOptions ...Option) (Encoder, error) {
	_ = "STUB: not implemented"
	return *new(Encoder), nil
}

// Jaeger encoding is a special case
// since the internal libraries offer no means of ptrace.TraceMarshaller.
// In order to preserve historical behavior, a custom type
// is used until it can be replaced.
