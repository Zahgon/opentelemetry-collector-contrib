// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stores // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/stores"

import (
	"sync"

	"go.uber.org/zap"
)

type nodeStats struct {
	podCnt       int
	containerCnt int
	cpuReq       uint64
	memReq       uint64
}

type nodeInfo struct {
	statsLock sync.RWMutex
	nodeStats nodeStats

	cpuLock     sync.RWMutex
	CPUCapacity uint64

	memLock     sync.RWMutex
	MemCapacity uint64

	logger *zap.Logger
}

func newNodeInfo(logger *zap.Logger) *nodeInfo { _ = "STUB: not implemented"; return nil }

func (n *nodeInfo) setCPUCapacity(cpuCapacity any) { _ = "STUB: not implemented"; return }

func (n *nodeInfo) setMemCapacity(memCapacity any) { _ = "STUB: not implemented"; return }

func (n *nodeInfo) getCPUCapacity() uint64 { _ = "STUB: not implemented"; return 0 }

func (n *nodeInfo) getMemCapacity() uint64 { _ = "STUB: not implemented"; return 0 }

func (n *nodeInfo) setNodeStats(stats nodeStats) { _ = "STUB: not implemented"; return }

func (n *nodeInfo) getNodeStats() nodeStats { _ = "STUB: not implemented"; return *new(nodeStats) }

func forceConvertToInt64(v any, logger *zap.Logger) uint64 { _ = "STUB: not implemented"; return 0 }
