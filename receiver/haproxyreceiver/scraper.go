// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package haproxyreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/haproxyreceiver"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/haproxyreceiver/internal/metadata"
)

var showStatsCommand = []byte("show stat\n")

type haproxyScraper struct {
	cfg               *Config
	httpClient        *http.Client
	logger            *zap.Logger
	mb                *metadata.MetricsBuilder
	telemetrySettings component.TelemetrySettings
}

func (s *haproxyScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (*haproxyScraper) readStats(buf []byte) ([]map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The CSV output starts with `# `, removing it to be able to read headers.

func (s *haproxyScraper) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func newScraper(cfg *Config, settings receiver.Settings) *haproxyScraper {
	_ = "STUB: not implemented"
	return nil
}
