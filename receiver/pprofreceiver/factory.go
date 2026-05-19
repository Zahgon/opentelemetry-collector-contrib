// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pprofreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pprofreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/xreceiver"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
)

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func defaultControllerConfig() scraperhelper.ControllerConfig {
	_ = "STUB: not implemented"
	return *new(scraperhelper.ControllerConfig)
}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createProfilesReceiver(
	_ context.Context,
	settings receiver.Settings,
	cfg component.Config,
	consumer xconsumer.Profiles,
) (xreceiver.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xreceiver.Profiles), nil
}
