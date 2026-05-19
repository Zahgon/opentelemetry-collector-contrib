// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package podmanreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/podmanreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/podmanreceiver/internal/metadata"
)

type metricsReceiver struct {
	config        *Config
	set           receiver.Settings
	clientFactory clientFactory
	scraper       *containerScraper
	mb            *metadata.MetricsBuilder
	cancel        context.CancelFunc
}

func newMetricsReceiver(
	set receiver.Settings,
	config *Config,
	clientFactory clientFactory,
) *metricsReceiver {
	_ = "STUB: not implemented"
	return nil
}

func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	config component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func (r *metricsReceiver) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// context for long-running operation

func (r *metricsReceiver) shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

type result struct {
	container      container
	containerStats containerStats
	err            error
}

func (r *metricsReceiver) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Don't know the number of failed metrics, but one container fetch is a partial error.

func (r *metricsReceiver) recordContainerStats(now pcommon.Timestamp, container container, stats *containerStats) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsReceiver) recordCPUMetrics(now pcommon.Timestamp, stats *containerStats) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsReceiver) recordNetworkMetrics(now pcommon.Timestamp, stats *containerStats) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsReceiver) recordMemoryMetrics(now pcommon.Timestamp, stats *containerStats) {
	_ = "STUB: not implemented"
	return
}

func (r *metricsReceiver) recordIOMetrics(now pcommon.Timestamp, stats *containerStats) {
	_ = "STUB: not implemented"
	return
}

// nanoseconds to seconds conversion truncating the fractional part
func toSecondsWithNanosecondPrecision(nanoseconds uint64) uint64 {
	_ = "STUB: not implemented"
	return 0
}
