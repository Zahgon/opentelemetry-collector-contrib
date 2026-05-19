// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package podmanreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/podmanreceiver"

import (
	"context"
	"net/url"
	"sync"

	"go.uber.org/zap"
)

type clientFactory func(logger *zap.Logger, cfg *Config) (PodmanClient, error)

type PodmanClient interface {
	ping(context.Context) error
	stats(context.Context, url.Values) ([]containerStats, error)
	list(context.Context, url.Values) ([]container, error)
	events(context.Context, url.Values) (<-chan event, <-chan error)
}

type containerScraper struct {
	client         PodmanClient
	containers     map[string]container
	containersLock sync.Mutex
	logger         *zap.Logger
	config         *Config
}

func newContainerScraper(engineClient PodmanClient, logger *zap.Logger, config *Config) *containerScraper {
	_ = "STUB: not implemented"
	return nil
}

// containers provides a slice of container to use for individual fetchContainerStats calls.
func (pc *containerScraper) getContainers() []container { _ = "STUB: not implemented"; return nil }

// loadContainerList will load the initial running container maps for
// inspection and establishing which containers warrant stat gathering calls
// by the receiver.
func (pc *containerScraper) loadContainerList(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (pc *containerScraper) events(ctx context.Context, options url.Values) (<-chan event, <-chan error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (pc *containerScraper) containerEventLoop(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// We are only interested when the context hasn't been canceled since requests made
// with a closed context are guaranteed to fail.

// Either decoding or connection error has occurred, so we should resume the event loop after
// waiting a moment.  In cases of extended daemon unavailability this will retry until
// collector teardown or background context is closed.

// inspectAndPersistContainer queries inspect api and returns *container and true when container should be queried for stats,
// nil and false otherwise. Persists the container in the cache if container is
// running and not excluded.
func (pc *containerScraper) inspectAndPersistContainer(ctx context.Context, cid string) (*container, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// fetchContainerStats will query the desired container stats
func (pc *containerScraper) fetchContainerStats(ctx context.Context, c container) (containerStats, error) {
	_ = "STUB: not implemented"
	return *new(containerStats), nil
}

func (pc *containerScraper) persistContainer(c container) { _ = "STUB: not implemented"; return }

func (pc *containerScraper) removeContainer(cid string) { _ = "STUB: not implemented"; return }
