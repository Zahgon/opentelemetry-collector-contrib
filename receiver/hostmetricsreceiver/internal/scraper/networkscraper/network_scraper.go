// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package networkscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/networkscraper"

import (
	"context"

	"github.com/shirou/gopsutil/v4/net"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterset"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/networkscraper/internal/metadata"
)

const (
	networkMetricsLen     = 4
	connectionsMetricsLen = 1
)

// scraper for Network Metrics
type networkScraper struct {
	settings  scraper.Settings
	config    *Config
	mb        *metadata.MetricsBuilder
	startTime pcommon.Timestamp
	includeFS filterset.FilterSet
	excludeFS filterset.FilterSet

	// for mocking
	bootTime    func(context.Context) (uint64, error)
	ioCounters  func(context.Context, bool) ([]net.IOCountersStat, error)
	connections func(context.Context, string) ([]net.ConnectionStat, error)
	conntrack   func(context.Context) ([]net.FilterStat, error)
}

// newNetworkScraper creates a set of Network related metrics
func newNetworkScraper(_ context.Context, settings scraper.Settings, cfg *Config) (*networkScraper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *networkScraper) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *networkScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (s *networkScraper) recordNetworkCounterMetrics(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// get total stats only
/*perNetworkInterfaceController=*/

// filter network interfaces by name

func (s *networkScraper) recordNetworkPacketsMetric(now pcommon.Timestamp, ioCountersSlice []net.IOCountersStat) {
	_ = "STUB: not implemented"
	return
}

func (s *networkScraper) recordNetworkDroppedPacketsMetric(now pcommon.Timestamp, ioCountersSlice []net.IOCountersStat) {
	_ = "STUB: not implemented"
	return
}

func (s *networkScraper) recordNetworkErrorPacketsMetric(now pcommon.Timestamp, ioCountersSlice []net.IOCountersStat) {
	_ = "STUB: not implemented"
	return
}

func (s *networkScraper) recordNetworkIOMetric(now pcommon.Timestamp, ioCountersSlice []net.IOCountersStat) {
	_ = "STUB: not implemented"
	return
}

func (s *networkScraper) recordNetworkConnectionsMetrics(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func getTCPConnectionStatusCounts(connections []net.ConnectionStat) map[string]int64 {
	_ = "STUB: not implemented"
	return nil
}

func (s *networkScraper) recordNetworkConnectionsMetric(now pcommon.Timestamp, connectionStateCounts map[string]int64) {
	_ = "STUB: not implemented"
	return
}

func (s *networkScraper) filterByInterface(ioCounters []net.IOCountersStat) []net.IOCountersStat {
	_ = "STUB: not implemented"
	return nil
}

func (s *networkScraper) includeInterface(interfaceName string) bool {
	_ = "STUB: not implemented"
	return false
}
