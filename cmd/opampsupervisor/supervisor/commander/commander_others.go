// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package commander

import (
	"os"
	"syscall"
)

func sendShutdownSignal(process *os.Process) error { _ = "STUB: not implemented"; return nil }

func sysProcAttrs() *syscall.SysProcAttr {
	_ = "STUB: not implemented"
	// On non-windows systems, no extra attributes are needed.
	return nil
}
