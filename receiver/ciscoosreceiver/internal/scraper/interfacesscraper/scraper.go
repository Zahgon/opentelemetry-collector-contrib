// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package interfacesscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciscoosreceiver/internal/scraper/interfacesscraper"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciscoosreceiver/internal/connection"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciscoosreceiver/internal/scraper/interfacesscraper/internal/metadata"
)

// interfacesScraper collects interface metrics from Cisco devices
type interfacesScraper struct {
	logger       *zap.Logger
	config       *Config
	mb           *metadata.MetricsBuilder
	deviceTarget string
	rpcClient    *connection.RPCClient
}

func (s *interfacesScraper) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *interfacesScraper) ScrapeMetrics(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (s *interfacesScraper) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *interfacesScraper) parseInterfaceData(ctx context.Context) ([]*Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
