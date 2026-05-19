// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ntpreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ntpreceiver"

import (
	"context"
	"time"

	"github.com/beevik/ntp"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ntpreceiver/internal/metadata"
)

var queryWithOptions = ntp.QueryWithOptions

type ntpScraper struct {
	logger   *zap.Logger
	mb       *metadata.MetricsBuilder
	version  int
	timeout  time.Duration
	endpoint string
}

func (s *ntpScraper) scrape(context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func newScraper(cfg *Config, settings receiver.Settings) *ntpScraper {
	_ = "STUB: not implemented"
	return nil
}
