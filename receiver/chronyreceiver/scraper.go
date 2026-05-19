// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package chronyreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/chronyreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/chronyreceiver/internal/chrony"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/chronyreceiver/internal/metadata"
)

type chronyScraper struct {
	client chrony.Client
	mb     *metadata.MetricsBuilder
}

func newScraper(ctx context.Context, cfg *Config, set receiver.Settings) *chronyScraper {
	_ = "STUB: not implemented"
	return nil
}

func (cs *chronyScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}
