// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filestatsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filestatsreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filestatsreceiver/internal/metadata"
)

type fsScraper struct {
	include string
	logger  *zap.Logger
	mb      *metadata.MetricsBuilder
}

func (s *fsScraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func newScraper(cfg *Config, settings receiver.Settings) *fsScraper {
	_ = "STUB: not implemented"
	return nil
}
