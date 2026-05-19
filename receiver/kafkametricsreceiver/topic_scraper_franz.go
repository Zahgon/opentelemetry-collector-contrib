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

const (
	minInsyncReplicas = "min.insync.replicas"
	retentionMs       = "retention.ms"
	retentionBytes    = "retention.bytes"
)

type topicScraperFranz struct {
	adm *kadm.Client
	cl  *kgo.Client

	settings    receiver.Settings
	topicFilter *regexp.Regexp
	config      Config
	mb          *metadata.MetricsBuilder
	host        component.Host
}

func (s *topicScraperFranz) start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *topicScraperFranz) shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *topicScraperFranz) ensureClients(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *topicScraperFranz) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// 1) list topics (with metadata details)

// filter topic names first

// 2) offsets for matched topics (newest & oldest)

// 3) per-topic configs & replication factor

// replication factor: derive from first partition's replica count

// first partition is enough; RF should be consistent

// seconds = ms / 1000

// 4) per-topic partitions & per-partition metrics

// partitions count

// iterate partitions without copying large structs

// replicas

// in-sync replicas

// offsets: newest/current and oldest (use .Offset)

// Factory helper for franz-go path (selected via feature gate elsewhere).
func createTopicsScraperFranz(_ context.Context, cfg Config, settings receiver.Settings) (scraper.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(scraper.Metrics), nil
}
