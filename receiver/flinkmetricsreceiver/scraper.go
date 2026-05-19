// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package flinkmetricsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/flinkmetricsreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/flinkmetricsreceiver/internal/metadata"
)

var (
	errClientNotInit       = errors.New("client not initialized")
	jobmanagerFailedFetch  = "Failed to fetch jobmanager metrics"
	taskmanagerFailedFetch = "Failed to fetch taskmanager metrics"
	jobsFailedFetch        = "Failed to fetch jobs metrics"
	subtasksFailedFetch    = "Failed to fetch subtasks metrics"
)

type flinkmetricsScraper struct {
	client   client
	cfg      *Config
	settings component.TelemetrySettings
	mb       *metadata.MetricsBuilder
}

func newflinkScraper(config *Config, settings receiver.Settings) *flinkmetricsScraper {
	_ = "STUB: not implemented"
	return nil
}

func (s *flinkmetricsScraper) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *flinkmetricsScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	// Validate we don't attempt to scrape without initializing the client
	return *new(pmetric.Metrics), nil
}
