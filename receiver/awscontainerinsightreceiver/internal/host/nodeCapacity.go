// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package host // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/host"

import (
	"context"
	"os"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
	"go.uber.org/zap"
)

type nodeCapacityProvider interface {
	getMemoryCapacity() int64
	getNumCores() int64
}

type nodeCapacity struct {
	memCapacity int64
	cpuCapacity int64
	logger      *zap.Logger

	// osLstat returns a FileInfo describing the named file.
	osLstat       func(name string) (os.FileInfo, error)
	virtualMemory func(ctx context.Context) (*mem.VirtualMemoryStat, error)
	cpuInfo       func(ctx context.Context) ([]cpu.InfoStat, error)
}

type nodeCapacityOption func(*nodeCapacity)

func newNodeCapacity(logger *zap.Logger, options ...nodeCapacityOption) (nodeCapacityProvider, error) {
	_ = "STUB: not implemented"
	return *new(nodeCapacityProvider), nil
}

func (nc *nodeCapacity) parseMemory(ctx context.Context) { _ = "STUB: not implemented"; return }

// If any error happen, then there will be no mem utilization metrics

func (nc *nodeCapacity) parseCPU(ctx context.Context) { _ = "STUB: not implemented"; return }

// If any error happen, then there will be no cpu utilization metrics

func (nc *nodeCapacity) getNumCores() int64 { _ = "STUB: not implemented"; return 0 }

func (nc *nodeCapacity) getMemoryCapacity() int64 { _ = "STUB: not implemented"; return 0 }
