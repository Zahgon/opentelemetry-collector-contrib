// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package agentcomponents // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/agentcomponents"

import (
	tracelog "github.com/DataDog/datadog-agent/pkg/trace/log"
	"go.uber.org/zap"
)

var _ tracelog.Logger = &ZapLogger{}

// ZapLogger implements the tracelog.Logger interface on top of a zap.Logger
type ZapLogger struct {
	// Logger is the internal zap logger
	Logger *zap.Logger
	// prevent unkeyed literal initialization
	_ struct{}
}

// Trace implements Logger.
func (*ZapLogger) Trace(...any) {
	_ = "STUB: not implemented" /* N/A */

	// Tracef implements Logger.
	return
}

func (*ZapLogger) Tracef(string, ...any) {
	_ = "STUB: not implemented" /* N/A */

	// Debug implements Logger.
	return
}

func (z *ZapLogger) Debug(v ...any) { _ = "STUB: not implemented"; return }

// Debugf implements Logger.
func (z *ZapLogger) Debugf(format string, params ...any) { _ = "STUB: not implemented"; return }

// Info implements Logger.
func (z *ZapLogger) Info(v ...any) { _ = "STUB: not implemented"; return }

// Infof implements Logger.
func (z *ZapLogger) Infof(format string, params ...any) { _ = "STUB: not implemented"; return }

// Warn implements Logger.
func (z *ZapLogger) Warn(v ...any) error { _ = "STUB: not implemented"; return nil }

// Warnf implements Logger.
func (z *ZapLogger) Warnf(format string, params ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// Error implements Logger.
func (z *ZapLogger) Error(v ...any) error { _ = "STUB: not implemented"; return nil }

// Errorf implements Logger.
func (z *ZapLogger) Errorf(format string, params ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// Critical implements Logger.
func (z *ZapLogger) Critical(v ...any) error { _ = "STUB: not implemented"; return nil }

// Criticalf implements Logger.
func (z *ZapLogger) Criticalf(format string, params ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// Flush implements Logger.
func (z *ZapLogger) Flush() { _ = "STUB: not implemented"; return }
