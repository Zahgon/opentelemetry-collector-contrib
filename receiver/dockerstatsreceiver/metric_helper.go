// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dockerstatsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/dockerstatsreceiver"

import (
	ctypes "github.com/moby/moby/api/types/container"
)

const nanosInASecond = 1e9

func calculateMemoryPercent(limit, usedNoCache uint64) float64 {
	_ = "STUB: not implemented"
	// MemoryStats.Limit will never be 0 unless the container is not running and we haven't
	// got any data from cgroup
	return 0
}

// calculateCPULimit calculate the number of cpus assigned to a container.
//
// Calculation is based on 3 alternatives by the following order:
// - nanocpus:   if set by i.e docker run -cpus=2
// - cpusetCpus: if set by i.e docker run -docker run -cpuset-cpus="0,2"
// - cpuquota:   if set by i.e docker run -cpu-quota=50000
//
// See https://docs.docker.com/config/containers/resource_constraints/#configure-the-default-cfs-scheduler for background.
func calculateCPULimit(hostConfig *ctypes.HostConfig) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Default CFS Period

// parseCPUSet helper function to decompose -cpuset-cpus value into number os cpus.
func parseCPUSet(line string) (float64, error) { _ = "STUB: not implemented"; return 0, nil }
