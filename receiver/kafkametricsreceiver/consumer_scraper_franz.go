// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkametricsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkametricsreceiver"

import (
	"context"
	"regexp"

	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kgo"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkametricsreceiver/internal/metadata"
)

type consumerScraperFranz struct {
	adm *kadm.Client
	cl  *kgo.Client

	settings    receiver.Settings
	groupFilter *regexp.Regexp
	topicFilter *regexp.Regexp
	config      Config
	mb          *metadata.MetricsBuilder
	host        component.Host
}

func (s *consumerScraperFranz) start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *consumerScraperFranz) shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *consumerScraperFranz) ensureClients(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *consumerScraperFranz) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// 1) list & filter groups

// 2) list & filter topics
// non-internal only, same as sarama default

// 3) compute partition list + end offsets for matched topics

// Build helpers equivalent to Sarama path

// lo.Topic, lo.Partition, lo.Offset

// 4) describe groups for member counts

// 5) per group: fetch committed offsets for matched topics and compute metrics

// Factory helper for franz-go path (selected under the feature gate later).
func createConsumerScraperFranz(_ context.Context, cfg Config, settings receiver.Settings) (scraper.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(scraper.Metrics), nil
}
