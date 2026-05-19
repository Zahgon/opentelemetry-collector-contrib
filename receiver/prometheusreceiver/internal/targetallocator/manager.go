// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package targetallocator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver/internal/targetallocator"

import (
	"context"
	"net/http"
	"sync"
	"sync/atomic"

	promconfig "github.com/prometheus/prometheus/config"
	"github.com/prometheus/prometheus/discovery"
	"github.com/prometheus/prometheus/scrape"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"
)

type Manager struct {
	settings             receiver.Settings
	shutdown             chan struct{}
	cfg                  *Config
	promCfg              *promconfig.Config
	initialScrapeConfigs []*promconfig.ScrapeConfig
	scrapeManager        *scrape.Manager
	discoveryManager     *discovery.Manager
	wg                   sync.WaitGroup

	// configUpdateCount tracks how many times the config has changed, for
	// testing.
	configUpdateCount *atomic.Int64
	// configUpdated is signaled after each config update, for testing.
	configUpdated chan struct{}
}

func NewManager(set receiver.Settings, cfg *Config, promCfg *promconfig.Config) *Manager {
	_ = "STUB: not implemented"
	return nil
}

func (m *Manager) Start(ctx context.Context, host component.Host, sm *scrape.Manager, dm *discovery.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

// the target allocator is disabled

// immediately sync jobs, not waiting for the first tick

func (m *Manager) Shutdown() { _ = "STUB: not implemented"; return }

// sync request jobs from targetAllocator and update underlying receiver, if the response does not match the provided compareHash.
// baseDiscoveryCfg can be used to provide additional ScrapeConfigs which will be added to the retrieved jobs.
func (m *Manager) sync(compareHash uint64, httpClient *http.Client) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// no update needed

// Copy initial scrape configurations

// Validate the scrape config and also fill in the defaults from the global config as needed.

func (m *Manager) applyCfg() error { _ = "STUB: not implemented"; return nil }

func getScrapeConfigsResponse(httpClient *http.Client, baseURL string) (map[string]*promconfig.ScrapeConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if valid

// instantiateShard inserts the SHARD environment variable in the returned configuration
func instantiateShard(body []byte, lookup func(string) (string, bool)) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Calculate a hash for a scrape config map.
// This is done by marshaling to YAML because it's the most straightforward and doesn't run into problems with unexported fields.
func getScrapeConfigHash(jobToScrapeConfig map[string]*promconfig.ScrapeConfig) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
