// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package nfsscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/nfsscraper"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/nfsscraper/internal/metadata"
)

var (
	// 3 net metrics + 3 rpc metrics + len(nfs_scraper_linux.go:nfsV3Procedures) + len(nfs_scraper_linux.go:nfsV4Procedures) = 97 metrics
	nfsMetricsLen = 97
	// 3 repcache metrics + 1 fh metric + 2 io metrics + 1 thread metric + 3 net metrics + 5 rpc metrics + len(nfs_scraper_linux.go:nfsdV3Procedures) + len(nfs_scraper_linux.go:nfsdV4Procedures) + len(nfs_scraper_linux.go:nfsdV4Operations) = 115 metrics
	nfsdMetricsLen = 115
)

// nfsScraper for NFS Metrics
type nfsScraper struct {
	settings scraper.Settings
	config   *Config
	mb       *metadata.MetricsBuilder

	getNfsStats  func() (*NfsStats, error)
	getNfsdStats func() (*nfsdStats, error)

	nfsStats  *NfsStats
	nfsdStats *nfsdStats
}

// newNfsScraper creates a metric scraper for NFS metrics
func newNfsScraper(settings scraper.Settings, cfg *Config) *nfsScraper {
	_ = "STUB: not implemented"
	return nil
}

func (s *nfsScraper) start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *nfsScraper) scrape(context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (s *nfsScraper) recordNfsMetrics(now pcommon.Timestamp) { _ = "STUB: not implemented"; return }

func (s *nfsScraper) recordNfsdMetrics(now pcommon.Timestamp) { _ = "STUB: not implemented"; return }
