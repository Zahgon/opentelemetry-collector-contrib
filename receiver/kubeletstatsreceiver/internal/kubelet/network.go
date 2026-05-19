// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kubelet // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/kubelet"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	stats "k8s.io/kubelet/pkg/apis/stats/v1alpha1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/metadata"
)

type getNetworkDataFunc func(s *stats.NetworkStats) (rx, tx *uint64)

type getInterfaceDataFunc func(s *stats.InterfaceStats) (rx, tx *uint64)

func addNetworkMetrics(mb *metadata.MetricsBuilder, networkMetrics metadata.NetworkMetrics, s *stats.NetworkStats, currentTime pcommon.Timestamp, allInterfaces bool) {
	_ = "STUB: not implemented"
	return
}

// Because stats.NetworkStats.Interfaces contains metrics for all interfaces, including default,
// we don't need to iterate over stats.NetworkStats.InterfaceStats for it, hence we return here

func recordNetworkDataPoint(mb *metadata.MetricsBuilder, recordDataPoint metadata.RecordIntDataPointWithDirectionFunc, s *stats.NetworkStats, getData getNetworkDataFunc, currentTime pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func getNetworkIO(s *stats.NetworkStats) (*uint64, *uint64) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getNetworkErrors(s *stats.NetworkStats) (*uint64, *uint64) {
	_ = "STUB: not implemented"
	return nil, nil
}

func recordInterfaceDataPoint(mb *metadata.MetricsBuilder, recordDataPoint metadata.RecordIntDataPointWithDirectionFunc, s *stats.InterfaceStats, getData getInterfaceDataFunc, currentTime pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func getInterfaceIO(s *stats.InterfaceStats) (*uint64, *uint64) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getInterfaceErrors(s *stats.InterfaceStats) (*uint64, *uint64) {
	_ = "STUB: not implemented"
	return nil, nil
}
