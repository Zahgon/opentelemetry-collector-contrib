// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pprofreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pprofreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/xreceiver"
	"go.opentelemetry.io/collector/scraper"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
	"go.opentelemetry.io/collector/scraper/xscraper"
)

type pprofReceiver struct {
	subComponents []component.Component
}

var _ xreceiver.Profiles = (*pprofReceiver)(nil)

func (r *pprofReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Roll back any already-started sub-components so we don't leak
// goroutines or hold a listening port on partial start failure.

func (r *pprofReceiver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func newReceiver(cfg *Config, settings receiver.Settings, consumer xconsumer.Profiles) (xreceiver.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xreceiver.Profiles), nil
}

func newScraperController(
	settings receiver.Settings,
	consumer xconsumer.Profiles,
	controllerCfg *scraperhelper.ControllerConfig,
	scraperFn func(scraper.Settings) (xscraper.Profiles, error),
) (xreceiver.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xreceiver.Profiles), nil
}
