// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package systemscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciscoosreceiver/internal/scraper/systemscraper"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciscoosreceiver/internal/connection"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciscoosreceiver/internal/scraper/systemscraper/internal/metadata"
)

// systemScraper collects system-level metrics for the Cisco device
type systemScraper struct {
	logger          *zap.Logger
	config          *Config
	mb              *metadata.MetricsBuilder
	collectionCount int
	deviceTarget    string
	rpcClient       *connection.RPCClient
}

func (s *systemScraper) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Log authentication method

func (s *systemScraper) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *systemScraper) ScrapeMetrics(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Set resource attributes

// Set resource attributes

// collectCPUUtilization collects CPU utilization metric from the device
func (s *systemScraper) collectCPUUtilization(_ context.Context) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Parse based on OS type

// IOS or IOS XE

// collectMemoryUtilization collects memory utilization metric from the device
func (s *systemScraper) collectMemoryUtilization(_ context.Context) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Parse memory utilization
