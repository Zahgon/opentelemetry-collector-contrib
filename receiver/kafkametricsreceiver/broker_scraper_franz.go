// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkametricsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkametricsreceiver"

import (
	"context"

	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkametricsreceiver/internal/metadata"
)

const logRetentionHours = "log.retention.hours"

type brokerScraperFranz struct {
	// franz-go handles (lazy created on first scrape)
	adm *kadm.Client
	cl  *kgo.Client

	settings receiver.Settings
	config   Config
	mb       *metadata.MetricsBuilder
	host     component.Host
}

func (s *brokerScraperFranz) start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *brokerScraperFranz) shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *brokerScraperFranz) ensureClients(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *brokerScraperFranz) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// ---- brokers count ----

// If we cannot list brokers, emit what we have (resource attrs) and return the error

// If log retention metric is disabled, we are done.

// Iterate the result and record the metric for each broker entry we can parse.

// Look up this broker's config set by resource name (broker id as string).
// fn can be nil to just return the entry

// kadm.Config has Key and MaybeValue() for the string value.
// We only care about log.retention.hours here.

// factory for franz-go scraper (internal; selected via gate at the call site later)
func createBrokerScraperFranz(_ context.Context, cfg Config, settings receiver.Settings) (scraper.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(scraper.Metrics), nil
}
