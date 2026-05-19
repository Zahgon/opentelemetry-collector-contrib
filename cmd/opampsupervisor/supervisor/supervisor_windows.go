// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package supervisor

import (
	"go.uber.org/zap/zapcore"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/eventlog"
)

type windowsService struct {
	sup *Supervisor
}

func NewSvcHandler() svc.Handler { _ = "STUB: not implemented"; return *new(svc.Handler) }

func (ws *windowsService) Execute(args []string, requests <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	_ = "STUB: not implemented"
	// The first argument supplied to service.Execute is the service name. If this is
	// not provided for some reason, raise a relevant error to the system event log
	return false, 0
}

func (ws *windowsService) start(elog *eventlog.Log) error { _ = "STUB: not implemented"; return nil }

func (ws *windowsService) stop() { _ = "STUB: not implemented"; return }

func openEventLog(serviceName string) (*eventlog.Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Logger wrappings
var _ zapcore.Core = (*windowsEventLogCore)(nil)

type windowsEventLogCore struct {
	core    zapcore.Core
	elog    *eventlog.Log
	encoder zapcore.Encoder
}

func (w windowsEventLogCore) Enabled(level zapcore.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (w windowsEventLogCore) With(fields []zapcore.Field) zapcore.Core {
	_ = "STUB: not implemented"
	return *new(zapcore.Core)
}

func (w windowsEventLogCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

func (w windowsEventLogCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	_ = "STUB: not implemented"
	return nil
}

// golang.org/x/sys/windows/svc/eventlog does not support Critical level event logs

// We would not be here if debug were disabled so log as info to not drop.

func (w windowsEventLogCore) Sync() error { _ = "STUB: not implemented"; return nil }

// TODO: If supervisor logging becomes configurable, update this function to respect that config
func withWindowsCore(elog *eventlog.Log) func(zapcore.Core) zapcore.Core {
	_ = "STUB: not implemented"
	return nil
}

// Use the Windows Event Log
