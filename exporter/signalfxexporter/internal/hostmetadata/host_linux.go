// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux

// Taken from https://github.com/signalfx/golib/blob/master/metadata/hostmetadata/host-linux.go
// with minor modifications.

package hostmetadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/hostmetadata"

import (
	"context"

	"golang.org/x/sys/unix"
)

// syscallUname maps to the golib system call, but can be modified for testing
var syscallUname = unix.Uname

func fillPlatformSpecificOSData(ctx context.Context, info *hostOS) error {
	_ = "STUB: not implemented"
	return nil
}

func fillPlatformSpecificCPUData(info *hostCPU) error { _ = "STUB: not implemented"; return nil }

// according to the python doc platform.Processor usually returns the same
// value as platform.Machine
// https://docs.python.org/3/library/platform.html#platform.processor

// getLinuxVersion - adds information about the host linux version to the supplied map
func getLinuxVersion(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getStringFromFile(pattern, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
