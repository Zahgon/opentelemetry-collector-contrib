// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zookeeperscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/scraper/zookeeperscraper"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/scraper/zookeeperscraper/internal/metadata"
)

// Constants to define entries in the output of "mntr" command.
const (
	avgLatencyMetricKey              = "zk_avg_latency"
	maxLatencyMetricKey              = "zk_max_latency"
	minLatencyMetricKey              = "zk_min_latency"
	packetsReceivedMetricKey         = "zk_packets_received"
	packetsSentMetricKey             = "zk_packets_sent"
	numAliveConnectionsMetricKey     = "zk_num_alive_connections"
	outstandingRequestsMetricKey     = "zk_outstanding_requests"
	zNodeCountMetricKey              = "zk_znode_count"
	watchCountMetricKey              = "zk_watch_count"
	ephemeralsCountMetricKey         = "zk_ephemerals_count"
	approximateDataSizeMetricKey     = "zk_approximate_data_size"
	openFileDescriptorCountMetricKey = "zk_open_file_descriptor_count"
	maxFileDescriptorCountMetricKey  = "zk_max_file_descriptor_count"

	fSyncThresholdExceedCountMetricKey = "zk_fsync_threshold_exceed_count"

	followersMetricKey       = "zk_followers"
	syncedFollowersMetricKey = "zk_synced_followers"
	pendingSyncsMetricKey    = "zk_pending_syncs"

	serverStateKey = "zk_server_state"
	zkVersionKey   = "zk_version"

	ruokKey = "ruok"
)

// metricCreator handles generation of metric and metric recording
type metricCreator struct {
	computedMetricStore map[string]int64
	mb                  *metadata.MetricsBuilder
	// Temporary feature gates while transitioning to metrics without a direction attribute
}

func newMetricCreator(mb *metadata.MetricsBuilder) *metricCreator {
	_ = "STUB: not implemented"
	return nil
}

func (m *metricCreator) recordDataPointsFunc(metric string) func(ts pcommon.Timestamp, val int64) {
	_ = "STUB: not implemented"
	return nil
}

func (m *metricCreator) generateComputedMetrics(logger *zap.Logger, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	// not_synced Followers Count
	return
}

func (m *metricCreator) computeNotSyncedFollowersMetric(ts pcommon.Timestamp) error {
	_ = "STUB: not implemented"
	return nil
}
