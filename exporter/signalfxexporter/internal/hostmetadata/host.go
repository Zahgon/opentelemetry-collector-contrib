// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Taken from https://github.com/signalfx/golib/blob/master/metadata/hostmetadata/host.go
// with minor modifications.

package hostmetadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/hostmetadata"

import (
	"context"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

const cpuStatsTimeout = 10 * time.Second

// Map library functions to unexported package variables for testing purposes.
var (
	cpuInfo          = cpu.InfoWithContext
	cpuCounts        = cpu.CountsWithContext
	memVirtualMemory = mem.VirtualMemoryWithContext
	hostInfo         = host.InfoWithContext
)

// hostCPU information about the host
type hostCPU struct {
	HostPhysicalCPUs int
	HostLogicalCPUs  int
	HostCPUCores     int64
	HostCPUModel     string
	HostMachine      string
	HostProcessor    string
}

// toStringMap returns the hostCPU as a string map
func (c *hostCPU) toStringMap() map[string]string { _ = "STUB: not implemented"; return nil }

// getCPU - adds information about the host cpu to the supplied map
func getCPU(ctx context.Context) (info *hostCPU, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// On Windows this can sometimes take longer than the default timeout (10 seconds).

// get cpu infoStats

// get physical cpu stats

// get logical cpu stats

// Count physical CPU cores by tracking unique {PhysicalID, CoreID} pairs.

// TODO: This is not ideal... if there are different processors
// we will only report one of the models... This is unlikely to happen,
// but it could

// hostOS is a struct containing information about the host os
type hostOS struct {
	HostOSName        string
	HostKernelName    string
	HostKernelRelease string
	HostKernelVersion string
	HostLinuxVersion  string
}

// toStringMap returns a map of key/value metadata about the host os
func (o *hostOS) toStringMap() map[string]string { _ = "STUB: not implemented"; return nil }

// getOS returns a struct with information about the host os
func getOS(ctx context.Context) (info *hostOS, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// in gopsutil KernelVersion returns what we would expect for Kernel Release

// Memory stores memory collected from the host
type Memory struct {
	Total int
}

// toStringMap returns a map of key/value metadata about the host memory
// where memory sizes are reported in Kb
func (m *Memory) toStringMap() map[string]string { _ = "STUB: not implemented"; return nil }

// getMemory returns the amount of memory on the host as datatype.USize
func getMemory(ctx context.Context) (*Memory, error) { _ = "STUB: not implemented"; return nil, nil }

func bytesToKilobytes(b int) int { _ = "STUB: not implemented"; return 0 }
