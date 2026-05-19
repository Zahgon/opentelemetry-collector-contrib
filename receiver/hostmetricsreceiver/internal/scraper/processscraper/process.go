// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package processscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper"

import (
	"context"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/process"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper/internal/metadata"
)

// processMetadata stores process related metadata along
// with the process handle, and provides a function to
// initialize a pcommon.Resource with the metadata

type processMetadata struct {
	pid        int32
	parentPid  int32
	executable *executableMetadata
	command    *commandMetadata
	username   string
	handle     processHandle
	createTime int64
}

type executableMetadata struct {
	name   string
	path   string
	cgroup string
}

type commandMetadata struct {
	command          string
	commandLine      string
	commandLineSlice []string
}

func (m *processMetadata) buildResource(rb *metadata.ResourceBuilder) pcommon.Resource {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource)
}

// TODO insert slice here once this is supported by the data model
// (see https://github.com/open-telemetry/opentelemetry-collector/pull/1142)

// processHandles provides a wrapper around []*process.Process
// to support testing

type processHandles interface {
	Pid(index int) int32
	At(index int) processHandle
	Len() int
}

type processHandle interface {
	NameWithContext(context.Context) (string, error)
	ExeWithContext(context.Context) (string, error)
	UsernameWithContext(context.Context) (string, error)
	CmdlineWithContext(context.Context) (string, error)
	CmdlineSliceWithContext(context.Context) ([]string, error)
	TimesWithContext(context.Context) (*cpu.TimesStat, error)
	MemoryInfoWithContext(context.Context) (*process.MemoryInfoStat, error)
	MemoryPercentWithContext(context.Context) (float32, error)
	IOCountersWithContext(context.Context) (*process.IOCountersStat, error)
	NumThreadsWithContext(context.Context) (int32, error)
	ThreadsWithContext(context.Context) (map[int32]*cpu.TimesStat, error)
	CreateTimeWithContext(context.Context) (int64, error)
	PpidWithContext(context.Context) (int32, error)
	PageFaultsWithContext(context.Context) (*process.PageFaultsStat, error)
	NumCtxSwitchesWithContext(context.Context) (*process.NumCtxSwitchesStat, error)
	NumFDsWithContext(context.Context) (int32, error)
	GetProcessHandleCountWithContext(context.Context) (int64, error)
	// If gatherUsed is true, the currently used value will be gathered and added to the resulting RlimitStat.
	RlimitUsageWithContext(ctx context.Context, gatherUsed bool) ([]process.RlimitStat, error)
	CgroupWithContext(ctx context.Context) (string, error)
}

type gopsProcessHandles struct {
	handles []wrappedProcessHandle
}

func (p *gopsProcessHandles) Pid(index int) int32 { _ = "STUB: not implemented"; return 0 }

func (p *gopsProcessHandles) At(index int) processHandle {
	_ = "STUB: not implemented"
	return *new(processHandle)
}

func (p *gopsProcessHandles) Len() int { _ = "STUB: not implemented"; return 0 }

const (
	flagParentPidSet             = 1 << 0
	flagUseInitialNumThreadsOnce = 1 << 1
)

type wrappedProcessHandle struct {
	*process.Process
	parentPid         int32
	initialNumThreads int32
	flags             uint8 // bitfield to track if fields are set
}

func (p *wrappedProcessHandle) CgroupWithContext(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *wrappedProcessHandle) PpidWithContext(ctx context.Context) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *wrappedProcessHandle) NumThreadsWithContext(ctx context.Context) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// The number of threads can fluctuate so use the initially cached value only the first time.

func parentPid(ctx context.Context, handle processHandle, pid int32) (int32, error) {
	_ = "STUB: not implemented"
	// special case for pid 0 and pid 1 in darwin
	return 0, nil
}
