// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package nginxreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver"

import (
	"context"

	"github.com/nginx/nginx-prometheus-exporter/client"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver/internal/metadata"
)

type nginxScraper struct {
	client   *client.NginxClient
	settings component.TelemetrySettings
	cfg      *Config
	mb       *metadata.MetricsBuilder
}

func newNginxScraper(
	settings receiver.Settings,
	cfg *Config,
) *nginxScraper {
	_ = "STUB: not implemented"
	return nil
}

func (r *nginxScraper) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *nginxScraper) scrape(context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}
