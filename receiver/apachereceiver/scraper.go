// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package apachereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachereceiver"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachereceiver/internal/metadata"
)

type apacheScraper struct {
	settings   component.TelemetrySettings
	cfg        *Config
	httpClient *http.Client
	mb         *metadata.MetricsBuilder
	serverName string
	port       string
}

func newApacheScraper(
	settings receiver.Settings,
	cfg *Config,
	serverName string,
	port string,
) *apacheScraper {
	_ = "STUB: not implemented"
	return nil
}

func (r *apacheScraper) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *apacheScraper) scrape(context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func addPartialIfError(errs *scrapererror.ScrapeErrors, err error) {
	_ = "STUB: not implemented"
	return
}

// GetStats collects metric stats by making a get request at an endpoint.
func (r *apacheScraper) GetStats() (string, error) { _ = "STUB: not implemented"; return "", nil }

// parseStats converts a response body key:values into a map.
func parseStats(resp string) map[string]string { _ = "STUB: not implemented"; return nil }

type scoreboardCountsByLabel map[metadata.AttributeScoreboardState]int64

// parseScoreboard quantifies the symbolic mapping of the scoreboard.
func parseScoreboard(values string) scoreboardCountsByLabel {
	_ = "STUB: not implemented"
	return *new(scoreboardCountsByLabel)
}

// kbytesToBytes converts 1 Kibibyte to 1024 bytes.
func kbytesToBytes(i int64) int64 { _ = "STUB: not implemented"; return 0 }
