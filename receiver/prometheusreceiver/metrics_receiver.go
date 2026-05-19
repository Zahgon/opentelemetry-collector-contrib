// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver"

import (
	"context"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/prometheus/discovery"
	"github.com/prometheus/prometheus/scrape"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver/internal/targetallocator"
)

const (
	defaultGCInterval = 2 * time.Minute
	gcIntervalDelta   = 1 * time.Minute

	// Use same settings as Prometheus web server
	maxConnections     = 512
	readTimeoutMinutes = 10
)

// pReceiver is the type that provides Prometheus scraper/receiver functionality.
type pReceiver struct {
	cfg            *Config
	consumer       consumer.Metrics
	cancelFunc     context.CancelFunc
	configLoaded   chan struct{}
	loadConfigOnce sync.Once

	settings               receiver.Settings
	scrapeManager          *scrape.Manager
	discoveryManager       *discovery.Manager
	targetAllocatorManager *targetallocator.Manager
	apiServer              *http.Server
	registry               *prometheus.Registry
	registerer             prometheus.Registerer
	unregisterMetrics      func()
}

// New creates a new prometheus.Receiver reference.
func newPrometheusReceiver(set receiver.Settings, cfg *Config, next consumer.Metrics) (*pReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start is the method that starts Prometheus scraping. It
// is controlled by having previously defined a Configuration using perhaps New.
func (r *pReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *pReceiver) start(ctx context.Context, host component.Host, opts prometheusComponentTestOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *pReceiver) initPrometheusComponents(
	ctx context.Context, logger *slog.Logger, host component.Host,
	opts prometheusComponentTestOptions,
) error {
	_ = "STUB: not implemented"
	// Register the metrics needed by service discovery mechanisms.
	return nil
}

// NewManager can sometimes return nil if it encountered an error, but
// the error message is logged separately.

// for testing only

// The scrape manager needs to wait for the configuration to be loaded before beginning

func (r *pReceiver) initScrapeOptions(o prometheusScrapeTestOptions) *scrape.Options {
	_ = "STUB: not implemented"
	return nil
}

func (r *pReceiver) initAPIServer(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// If allowed CORS origins are provided in the receiver config, combine them into a single regex since the Prometheus API server requires this format.

// If read timeout is not set in the receiver config, use the default Prometheus value.

// Creates the API object in the same way as the Prometheus web package: https://github.com/prometheus/prometheus/blob/6150e1ca0ede508e56414363cc9062ef522db518/web/web.go#L314-L354
// Anything not defined by the options above will be nil, such as o.QueryEngine, o.Storage, etc. IsAgent=true, so these being nil is expected by Prometheus.

// This ensures that any changes to the config made, even by the target allocator, are reflected in the API.

// nil

// TSDBAdminStats
// ""
// false

// 0
// 0
// 0

// StatsRenderer

// LookbackDelta - Using the default value of 5 minutes

// appendMetadata from remote write
// OverrideErrorCode
// FeatureRegistry

// Create listener and monitor with conntrack in the same way as the Prometheus web package: https://github.com/prometheus/prometheus/blob/6150e1ca0ede508e56414363cc9062ef522db518/web/web.go#L564-L579

// Run the API server in the same way as the Prometheus web package: https://github.com/prometheus/prometheus/blob/6150e1ca0ede508e56414363cc9062ef522db518/web/web.go#L582-L630

// This is the path the web package uses, but the router above with no prefix can also be Registered by apiV1 instead.

// Helper function from the Prometheus web package: https://github.com/prometheus/prometheus/blob/6150e1ca0ede508e56414363cc9062ef522db518/web/web.go#L582-L630
func setPathWithPrefix(prefix string) func(handlerName string, handler http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return nil
}

// gcInterval returns the longest scrape interval used by a scrape config,
// plus a delta to prevent race conditions.
// This ensures jobs are not garbage collected between scrapes.
func gcInterval(cfg *PromConfig) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Shutdown stops and cancels the underlying Prometheus scrapers.
func (r *pReceiver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type prometheusComponentTestOptions struct {
	discovery prometheusDiscoveryTestOptions
	scrape    prometheusScrapeTestOptions
}

type prometheusDiscoveryTestOptions struct {
	// updatert is the interval for updating targets.
	//
	// If zero, the default (5s) from Prometheus is used.
	// This option is primarily for testing.
	updatert time.Duration
}

type prometheusScrapeTestOptions struct {
	// discoveryReloadInterval is the interval for reloading
	// scrape configurations.
	//
	// If zero, the default (5s) from Prometheus is used.
	// This option is primarily for testing.
	discoveryReloadInterval time.Duration
}
