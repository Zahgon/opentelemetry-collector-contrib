// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pprofreceiver/internal"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/scraper/xscraper"
)

var _ xscraper.Profiles = &HTTPClientScraper{}

type HTTPClientScraper struct {
	ClientConfig confighttp.ClientConfig
	Settings     component.TelemetrySettings
	client       *http.Client
}

func (hcs *HTTPClientScraper) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (hcs *HTTPClientScraper) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (hcs *HTTPClientScraper) ScrapeProfiles(_ context.Context) (pprofile.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles), nil
}
