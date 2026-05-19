// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package commander

import (
	"os"
	"syscall"

	"golang.org/x/sys/windows"
)

var (
	kernel32API = windows.NewLazySystemDLL("kernel32.dll")

	ctrlEventProc = kernel32API.NewProc("GenerateConsoleCtrlEvent")
)

func sendShutdownSignal(process *os.Process) error {
	_ = "STUB: not implemented"
	// signaling with os.Interrupt is not supported on windows systems,
	// so we need to use the windows API to properly send a graceful shutdown signal.
	// See: https://learn.microsoft.com/en-us/windows/console/generateconsolectrlevent
	return nil
}

func sysProcAttrs() *syscall.SysProcAttr {
	_ = "STUB: not implemented"
	// By default, a ctrl-break event applies to a whole process group, which ends up
	// shutting down the supervisor. Instead, we spawn the collector in its own process
	// group, so that sending a ctrl-break event shuts down just the collector.
	return nil
}
