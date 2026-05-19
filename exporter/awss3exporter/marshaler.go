// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awss3exporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter"

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type marshaler interface {
	MarshalTraces(td ptrace.Traces) ([]byte, error)
	MarshalLogs(ld plog.Logs) ([]byte, error)
	MarshalMetrics(md pmetric.Metrics) ([]byte, error)
	format() string
	compressed() bool
}

var ErrUnknownMarshaler = errors.New("unknown marshaler")

func newMarshalerFromEncoding(encoding *component.ID, fileFormat string, host component.Host, logger *zap.Logger) (marshaler, error) {
	_ = "STUB: not implemented"
	return *new(marshaler), nil
}

// cast with ok to avoid panics.

func newMarshaler(mType MarshalerType, logger *zap.Logger) (marshaler, error) {
	_ = "STUB: not implemented"
	return *new(marshaler), nil
}
