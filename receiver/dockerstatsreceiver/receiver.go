// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dockerstatsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/dockerstatsreceiver"

import (
	"context"

	ctypes "github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/docker"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/dockerstatsreceiver/internal/metadata"
)

var (
	defaultDockerAPIVersion         = docker.MustNewAPIVersion("1.44")
	minimumRequiredDockerAPIVersion = docker.MustNewAPIVersion("1.25")
)

type resultV2 struct {
	stats     *ctypes.StatsResponse
	container *docker.Container
	err       error
}

type metricsReceiver struct {
	config   *Config
	settings receiver.Settings
	client   *docker.Client
	mb       *metadata.MetricsBuilder
	cancel   context.CancelFunc
}

func newMetricsReceiver(set receiver.Settings, config *Config) *metricsReceiver {
	_ = "STUB: not implemented"
	return nil
}

func (r *metricsReceiver) clientOptions() []client.Opt { _ = "STUB: not implemented"; return nil }

func (r *metricsReceiver) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *metricsReceiver) shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *metricsReceiver) scrapeV2(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Stream is still starting up; skip until first frame arrives.

// Default (non-streaming): open a fresh connection per container each scrape.

// Don't know the number of failed stats, but one container fetch is a partial error.

func (r *metricsReceiver) recordContainerStats(now pcommon.Timestamp, containerStats *ctypes.StatsResponse, container *docker.Container) error {
	_ = "STUB: not implemented"
	return nil
}

// Always-present resource attrs + the user-configured resource attrs

func (r *metricsReceiver) recordMemoryMetrics(now pcommon.Timestamp, memoryStats *ctypes.MemoryStats) {
	_ = "STUB: not implemented"
	return
}

type blkioRecorder func(now pcommon.Timestamp, val int64, devMaj, devMin, operation string)

func (r *metricsReceiver) recordBlkioMetrics(now pcommon.Timestamp, blkioStats *ctypes.BlkioStats) {
	_ = "STUB: not implemented"
	return
}

func recordSingleBlkioStat(now pcommon.Timestamp, statEntries []ctypes.BlkioStatEntry, recorder blkioRecorder) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsReceiver) recordNetworkMetrics(now pcommon.Timestamp, networks *map[string]ctypes.NetworkStats) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsReceiver) recordCPUMetrics(now pcommon.Timestamp, v *ctypes.StatsResponse) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsReceiver) recordPidsMetrics(now pcommon.Timestamp, pidsStats *ctypes.PidsStats) {
	_ = "STUB: not implemented"
	// pidsStats are available when kernel version is >= 4.3 and pids_cgroup is supported, it is empty otherwise.
	return
}

func (r *metricsReceiver) recordBaseMetrics(now pcommon.Timestamp, base *ctypes.InspectResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// value not available or invalid

func (r *metricsReceiver) recordHostConfigMetrics(now pcommon.Timestamp, containerJSON *ctypes.InspectResponse) error {
	_ = "STUB: not implemented"
	return nil
}
