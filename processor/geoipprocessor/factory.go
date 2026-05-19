// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package geoipprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/otel/attribute"
	conventions "go.opentelemetry.io/otel/semconv/v1.40.0"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor/internal/provider"
	maxmind "github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor/internal/provider/maxmindprovider"
)

var (
	processorCapabilities = consumer.Capabilities{MutatesData: true}
	// defaultAttributes holds a list of default resource attribute keys.
	// These keys are used to identify an IP address attribute associated with the resource.
	defaultAttributes = []attribute.Key{
		// The client attributes are in use by the HTTP semantic conventions
		conventions.ClientAddressKey,
		// The source attributes are used when there is no client/server relationship between the two sides, or when that relationship is unknown
		conventions.SourceAddressKey,
	}
)

// providerFactories is a map that stores GeoIPProviderFactory instances, keyed by the provider type.
var providerFactories = map[string]provider.GeoIPProviderFactory{
	maxmind.TypeStr: &maxmind.Factory{},
}

// NewFactory creates a new processor factory with default configuration,
// and registers the processors for metrics, traces, and logs.
func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

// getProviderFactory retrieves the GeoIPProviderFactory for the given key.
// It returns the factory and a boolean indicating whether the factory was found.
func getProviderFactory(key string) (provider.GeoIPProviderFactory, bool) {
	_ = "STUB: not implemented"
	return *new(provider.GeoIPProviderFactory), false
}

// createDefaultConfig returns a default configuration for the processor.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// createGeoIPProviders creates a list of GeoIPProvider instances based on the provided configuration and providers factories.
func createGeoIPProviders(
	ctx context.Context,
	set processor.Settings,
	config *Config,
	factories map[string]provider.GeoIPProviderFactory,
) ([]provider.GeoIPProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createMetricsProcessor(ctx context.Context, set processor.Settings, cfg component.Config, nextConsumer consumer.Metrics) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}

func createTracesProcessor(ctx context.Context, set processor.Settings, cfg component.Config, nextConsumer consumer.Traces) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}

func createLogsProcessor(ctx context.Context, set processor.Settings, cfg component.Config, nextConsumer consumer.Logs) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}
