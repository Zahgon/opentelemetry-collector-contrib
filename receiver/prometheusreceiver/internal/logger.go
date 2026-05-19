// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver/internal"

import (
	gokitLog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"go.uber.org/zap"
)

const (
	levelKey = "level"
	msgKey   = "msg"
	errKey   = "err"
)

// NewZapToGokitLogAdapter create an adapter for zap.Logger to gokitLog.Logger
func NewZapToGokitLogAdapter(logger *zap.Logger) gokitLog.Logger {
	_ = "STUB: not implemented"
	// need to skip two levels in order to get the correct caller
	// one for this method, the other for gokitLog
	return *new(gokitLog.Logger)
}

type zapToGokitLogAdapter struct {
	l *zap.SugaredLogger
}

type logData struct {
	level       level.Value
	msg         string
	otherFields []any
}

func (w *zapToGokitLogAdapter) Log(keyvals ...any) error {
	_ = "STUB: not implemented"
	// expecting key value pairs, the number of items need to be even
	return nil
}

// Extract log level and message and log them using corresponding zap function

// in case something goes wrong

func extractLogData(keyvals []any) logData { _ = "STUB: not implemented"; return *new(logData) }

// default

// check if a given key-value pair represents go-kit log message and return it
func matchLogMessage(key, val any) (string, bool) { _ = "STUB: not implemented"; return "", false }

// check if a given key-value pair represents go-kit log level and return it
func matchLogLevel(key, val any) (level.Value, bool) {
	_ = "STUB: not implemented"
	return *new(level.Value), false
}

//revive:disable:error-return

// check if a given key-value pair represents an error and return it
func matchError(key, val any) (error, bool) { _ = "STUB: not implemented"; return nil, false }

//revive:enable:error-return

// find a matching zap logging function to be used for a given level
func levelToFunc(logger *zap.SugaredLogger, lvl level.Value) func(string, ...any) {
	_ = "STUB: not implemented"
	return nil
}

// default

var _ gokitLog.Logger = (*zapToGokitLogAdapter)(nil)
