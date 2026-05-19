// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package tcpcheckreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tcpcheckreceiver"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tcpcheckreceiver/internal/metadata"
)

type scraper struct {
	cfg                *Config
	settings           component.TelemetrySettings
	mb                 *metadata.MetricsBuilder
	getConnectionState func(tcpConfig *confignet.TCPAddrConfig) (tcpConnectionState, error)
	errorCounts        map[string]int64
}

type tcpConnectionState struct {
	LocalAddr  string // Local address of the connection
	RemoteAddr string // Remote address of the connection
	Network    string // Network type (e.g., "tcp")
}

func getConnectionState(tcpConfig *confignet.TCPAddrConfig) (tcpConnectionState, error) {
	_ = "STUB: not implemented"
	return *new(tcpConnectionState), nil
}

func (*scraper) errorListener(ctx context.Context, eQueue <-chan error, eOut chan<- *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *scraper) scrapeEndpoint(ctx context.Context, tcpConfig *confignet.TCPAddrConfig, wg *sync.WaitGroup, mux *sync.Mutex, errChan chan<- error) {
	_ = "STUB: not implemented"
	return
}

// Convert error to appropriate error code

// Increment error count for this endpoint

// Record error metrics with the current error count

// Send error to channel

// Record success metrics

func (s *scraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Initialize error counts if not already done

// Start error listener

func newScraper(cfg *Config, settings receiver.Settings) *scraper {
	_ = "STUB: not implemented"
	return nil
}
