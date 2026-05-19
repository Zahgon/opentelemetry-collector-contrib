// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package maxmind // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor/internal/provider/maxmindprovider"

import (
	"context"

	"go.opentelemetry.io/collector/processor"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor/internal/provider"
)

const (
	// TypeStr the value of "type" key in configuration.
	TypeStr = "maxmind"
)

// Factory is the Factory for the MaxMind GeoIP provider.
type Factory struct{}

var _ provider.GeoIPProviderFactory = (*Factory)(nil)

// CreateDefaultConfig creates the default configuration for the Provider.
func (*Factory) CreateDefaultConfig() provider.Config {
	_ = "STUB: not implemented"

	// CreateGeoIPProvider creates a provider based on this config.
	return *new(provider.Config)
}

func (*Factory) CreateGeoIPProvider(_ context.Context, _ processor.Settings, cfg provider.Config) (provider.GeoIPProvider, error) {
	_ = "STUB: not implemented"
	return *new(provider.GeoIPProvider), nil
}
