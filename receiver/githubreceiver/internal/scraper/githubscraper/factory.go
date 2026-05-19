// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package githubscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/githubreceiver/internal/scraper/githubscraper"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/githubreceiver/internal"
)

// This file implements factory for the GitHub Scraper as part of the GitHub Receiver

const (
	TypeStr                     = "scraper"
	defaultConcurrencyLimit     = 50
	defaultHTTPTimeout          = 15 * time.Second
	defaultMergedPRLookbackDays = 30
	defaultMaxRetries           = 10
)

type Factory struct{}

func (*Factory) CreateDefaultConfig() internal.Config {
	_ = "STUB: not implemented"
	return *new(internal.Config)
}

// Default to 50 concurrent goroutines

func (*Factory) CreateMetricsScraper(
	_ context.Context,
	params receiver.Settings,
	cfg internal.Config,
) (scraper.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(scraper.Metrics), nil
}
