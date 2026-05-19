// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package icmpcheckreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/icmpcheckreceiver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/icmpcheckreceiver/internal/metadata"
)

const (
	DefaultPingCount    = 3
	DefaultPingTimeout  = time.Second * 5
	DefaultPingInterval = time.Second * 1
)

type pingResult struct {
	stats      *pingStats
	targetHost string
	targetIP   string
	err        error
}

type icmpCheckScraper struct {
	cfg           *Config
	settings      component.TelemetrySettings
	mb            *metadata.MetricsBuilder
	pingerFactory func(target PingTarget) (pinger, error)
}

// apply defaults on targets if not set
func (scr *icmpCheckScraper) start(_ context.Context, _ component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (scr *icmpCheckScraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func ping(p pinger, results chan<- pingResult) { _ = "STUB: not implemented"; return }

func addMetrics(result pingResult, mb *metadata.MetricsBuilder, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// Record all metrics - they will be filtered by MetricsBuilderConfig

// Record resource attributes

func newScraper(cfg *Config, settings receiver.Settings) *icmpCheckScraper {
	_ = "STUB: not implemented"
	return nil
}
