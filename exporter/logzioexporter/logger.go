// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logzioexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter"

import (
	"io"
	"log"

	"github.com/hashicorp/go-hclog"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// hclog2ZapLogger implements Hashicorp's hclog.Logger interface using Uber's zap.Logger. It's a workaround for plugin
// system. go-plugin doesn't support other logger than hclog. This logger implements only methods used by the go-plugin.
type hclog2ZapLogger struct {
	Zap  *zap.Logger
	name string
}

func (*hclog2ZapLogger) Log(hclog.Level, string, ...any) { _ = "STUB: not implemented"; return }

func (*hclog2ZapLogger) ImpliedArgs() []any { _ = "STUB: not implemented"; return nil }

func (l *hclog2ZapLogger) Name() string { _ = "STUB: not implemented"; return "" }

func (*hclog2ZapLogger) StandardWriter(*hclog.StandardLoggerOptions) io.Writer {
	_ = "STUB: not implemented"

	// Trace implementation.
	return *new(io.Writer)
}

func (*hclog2ZapLogger) Trace(string, ...any) {
	_ = "STUB: not implemented"

	// Debug implementation.
	return
}

func (l *hclog2ZapLogger) Debug(msg string, args ...any) { _ = "STUB: not implemented"; return }

// Info implementation.
func (l *hclog2ZapLogger) Info(msg string, args ...any) { _ = "STUB: not implemented"; return }

// Warn implementation.
func (l *hclog2ZapLogger) Warn(msg string, args ...any) { _ = "STUB: not implemented"; return }

// Error implementation.
func (l *hclog2ZapLogger) Error(msg string, args ...any) { _ = "STUB: not implemented"; return }

// IsTrace implementation.
func (*hclog2ZapLogger) IsTrace() bool {
	_ = "STUB: not implemented"

	// IsDebug implementation.
	return false
}

func (*hclog2ZapLogger) IsDebug() bool {
	_ = "STUB: not implemented"

	// IsInfo implementation.
	return false
}

func (*hclog2ZapLogger) IsInfo() bool {
	_ = "STUB: not implemented"

	// IsWarn implementation.
	return false
}

func (*hclog2ZapLogger) IsWarn() bool {
	_ = "STUB: not implemented"

	// IsError implementation.
	return false
}

func (*hclog2ZapLogger) IsError() bool {
	_ = "STUB: not implemented"

	// With implementation.
	return false
}

func (l *hclog2ZapLogger) With(args ...any) hclog.Logger {
	_ = "STUB: not implemented"
	return *new(hclog.Logger)
}

// Named implementation.
func (l *hclog2ZapLogger) Named(name string) hclog.Logger {
	_ = "STUB: not implemented"
	return *new(hclog.Logger)
}

// ResetNamed implementation.
func (*hclog2ZapLogger) ResetNamed(string) hclog.Logger {
	_ = "STUB: not implemented"
	// no need to implement that as go-plugin doesn't use this method.
	return *new(hclog.Logger)
}

// SetLevel implementation.
func (*hclog2ZapLogger) SetLevel(hclog.Level) {
	_ = "STUB: not implemented"
	// no need to implement that as go-plugin doesn't use this method.
	return
}

// GetLevel implementation.
func (*hclog2ZapLogger) GetLevel() hclog.Level {
	_ = "STUB: not implemented"
	// no need to implement that as go-plugin doesn't use this method.
	return *new(hclog.Level)
}

// StandardLogger implementation.
func (*hclog2ZapLogger) StandardLogger(*hclog.StandardLoggerOptions) *log.Logger {
	_ = "STUB: not implemented"
	// no need to implement that as go-plugin doesn't use this method.
	return nil
}

func argsToFields(args ...any) []zapcore.Field { _ = "STUB: not implemented"; return nil }
