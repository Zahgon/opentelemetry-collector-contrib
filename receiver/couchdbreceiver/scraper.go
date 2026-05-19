// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package couchdbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/couchdbreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/couchdbreceiver/internal/metadata"
)

type couchdbScraper struct {
	client   client
	config   *Config
	settings component.TelemetrySettings
	mb       *metadata.MetricsBuilder
}

func newCouchdbScraper(settings receiver.Settings, config *Config) *couchdbScraper {
	_ = "STUB: not implemented"
	return nil
}

func (c *couchdbScraper) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *couchdbScraper) scrape(context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}
