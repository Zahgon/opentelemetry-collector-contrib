// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package podmanreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/podmanreceiver"

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

var errNoStatsFound = errors.New("No stats found")

type libpodClient struct {
	conn     *http.Client
	endpoint string
}

func newLibpodClient(logger *zap.Logger, cfg *Config) (PodmanClient, error) {
	_ = "STUB: not implemented"
	return *new(PodmanClient), nil
}

func (c *libpodClient) request(ctx context.Context, path string, params url.Values) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *libpodClient) stats(ctx context.Context, options url.Values) ([]containerStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *libpodClient) list(ctx context.Context, options url.Values) ([]container, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *libpodClient) ping(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// events returns a stream of events. It's up to the caller to close the stream by canceling the context.
func (c *libpodClient) events(ctx context.Context, options url.Values) (<-chan event, <-chan error) {
	_ = "STUB: not implemented"
	return nil, nil
}
