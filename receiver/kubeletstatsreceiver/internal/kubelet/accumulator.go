// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kubelet // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/kubelet"

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	stats "k8s.io/kubelet/pkg/apis/stats/v1alpha1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/metadata"
)

type MetricGroup string

// Values for MetricGroup enum.
const (
	ContainerMetricGroup = MetricGroup("container")
	PodMetricGroup       = MetricGroup("pod")
	NodeMetricGroup      = MetricGroup("node")
	VolumeMetricGroup    = MetricGroup("volume")
)

// ValidMetricGroups map of valid metrics.
var ValidMetricGroups = map[MetricGroup]bool{
	ContainerMetricGroup: true,
	PodMetricGroup:       true,
	NodeMetricGroup:      true,
	VolumeMetricGroup:    true,
}

type metricDataAccumulator struct {
	m                     []pmetric.Metrics
	metadata              Metadata
	logger                *zap.Logger
	metricGroupsToCollect map[MetricGroup]bool
	allNetworkInterfaces  map[MetricGroup]bool
	time                  time.Time
	mbs                   *metadata.MetricsBuilders
}

func addUptimeMetric(mb *metadata.MetricsBuilder, uptimeMetric metadata.RecordIntDataPointFunc, startTime v1.Time, currentTime pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func (a *metricDataAccumulator) nodeStats(s stats.NodeStats) { _ = "STUB: not implemented"; return }

// todo s.Runtime.ImageFs

func (a *metricDataAccumulator) podStats(s *stats.PodStats) { _ = "STUB: not implemented"; return }

func (a *metricDataAccumulator) containerStats(sPod *stats.PodStats, s *stats.ContainerStats) {
	_ = "STUB: not implemented"
	return
}

func (a *metricDataAccumulator) volumeStats(sPod *stats.PodStats, s *stats.VolumeStats) {
	_ = "STUB: not implemented"
	return
}
