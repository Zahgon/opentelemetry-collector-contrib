// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
// Originally copied from https://github.com/signalfx/signalfx-agent/blob/fbc24b0fdd3884bd0bbfbd69fe3c83f49d4c0b77/pkg/apm/log/log.go

package log // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/apm/log"

// Fields is a map that is used to populated logging context.
type Fields map[string]any

type nilLogger struct{}

func (nilLogger) Debug(string) { _ = "STUB: not implemented"; return }

func (nilLogger) Warn(string) { _ = "STUB: not implemented"; return }

func (nilLogger) Error(string) { _ = "STUB: not implemented"; return }

func (nilLogger) Info(string) { _ = "STUB: not implemented"; return }

func (nilLogger) Panic(string) { _ = "STUB: not implemented"; return }

func (nilLogger) WithFields(Fields) Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (nilLogger) WithError(error) Logger {
	_ = "STUB: not implemented"

	// Nil logger is a silent logger interface.
	return *new(Logger)
}

var Nil = nilLogger{}

var _ Logger = (*nilLogger)(nil)

// Logger is generic logging interface.
type Logger interface {
	Debug(msg string)
	Warn(msg string)
	Error(msg string)
	Info(msg string)
	Panic(msg string)
	WithFields(fields Fields) Logger
	WithError(err error) Logger
}
