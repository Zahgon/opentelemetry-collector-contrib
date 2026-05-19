// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package main

import (
	"golang.org/x/sys/windows"
)

var (
	kernel32API = windows.NewLazySystemDLL("kernel32.dll")

	allocConsoleProc = kernel32API.NewProc("AllocConsole")
	freeConsoleProc  = kernel32API.NewProc("FreeConsole")
)

func run() error {
	_ = "STUB: not implemented"
	// always allocate a console in case we're running as service
	return nil
}

// Per https://learn.microsoft.com/en-us/windows/console/allocconsole#remarks
// AllocConsole fails with this error when there's already a console attached, such as not being ran as service
// ignore this error and only return other errors

// No need to supply service name when startup is invoked through
// the Service Control Manager directly.

// Per https://learn.microsoft.com/en-us/windows/win32/api/winsvc/nf-winsvc-startservicectrldispatchera#return-value
// this means that the process is not running as a service, so run interactively.

// windows services don't get created with a console
// need to allocate a console in order to send CTRL_BREAK_EVENT to agent sub process
func allocConsole() error { _ = "STUB: not implemented"; return nil }

// free console once we're done with it
func freeConsole() error { _ = "STUB: not implemented"; return nil }
