// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// github.com/DataDog/datadog-agent/comp/core/config is not suppported on AIX
//go:build !aix

package agentcomponents // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/agentcomponents"

import (
	coreconfig "github.com/DataDog/datadog-agent/comp/core/config"
	corelog "github.com/DataDog/datadog-agent/comp/core/log/def"
	"github.com/DataDog/datadog-agent/comp/forwarder/defaultforwarder"
	pkgconfigmodel "github.com/DataDog/datadog-agent/pkg/config/model"
	"github.com/DataDog/datadog-agent/pkg/metrics"
	"github.com/DataDog/datadog-agent/pkg/serializer"
	"go.opentelemetry.io/collector/component"

	datadogconfig "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/config"
)

// ConfigOption is a function that configures the Datadog agent config component.
// This allows for flexible configuration by different modules.
type ConfigOption func(pkgconfigmodel.Config)

// SerializerWithForwarder is an interface that extends the MetricSerializer interface
// with ability to interact directly with the underlying forwarder's lifecycle methods.
type SerializerWithForwarder interface {
	serializer.MetricSerializer
	Start() error
	State() uint32
	Stop()
	// SendSeriesWithMetadata sends a metrics series to Datadog
	SendSeriesWithMetadata(series metrics.Series) error
}

// forwarderWithLifecycle extends the defaultforwarder.Forwarder interface
// with lifecycle management methods
type forwarderWithLifecycle interface {
	defaultforwarder.Forwarder
	Start() error
	State() uint32
	Stop()
}

// Compile-time check to ensure DefaultForwarder implements ForwarderWithLifecycle
var _ forwarderWithLifecycle = (*defaultforwarder.DefaultForwarder)(nil)

// Compile-time check to ensure datadogSerializer implements SerializerWithForwarder
var _ SerializerWithForwarder = (*datadogSerializer)(nil)

// datadogSerializer is a concrete implementation of SerializerWithForwarder that wraps
// a MetricSerializer and provides access to the underlying forwarder's lifecycle methods
type datadogSerializer struct {
	serializer.MetricSerializer
	forwarder forwarderWithLifecycle
}

// Start delegates to the underlying forwarder's Start method
func (ds *datadogSerializer) Start() error { _ = "STUB: not implemented"; return nil }

// State delegates to the underlying forwarder's State method
func (ds *datadogSerializer) State() uint32 { _ = "STUB: not implemented"; return 0 }

// Stop delegates to the underlying forwarder's Stop method
func (ds *datadogSerializer) Stop() { _ = "STUB: not implemented"; return }

// seriesSource wraps a metrics.Series to implement the metrics.SerieSource interface
type seriesSource struct {
	series metrics.Series
	index  int
}

// MoveNext moves to the next serie in the collection
func (s *seriesSource) MoveNext() bool { _ = "STUB: not implemented"; return false }

// Current returns the current serie
func (s *seriesSource) Current() *metrics.Serie { _ = "STUB: not implemented"; return nil }

// Count returns the total number of series
func (s *seriesSource) Count() uint64 { _ = "STUB: not implemented"; return 0 }

// SendSeriesWithMetadata sends a metrics series to Datadog
// This method wraps the series in a SerieSource and uses the serializer's SendIterableSeries
func (ds *datadogSerializer) SendSeriesWithMetadata(series metrics.Series) error {
	_ = "STUB: not implemented"
	return nil
}

// Wrap the series in a SerieSource

// Start before the first element

// NewLogComponent creates a new log component for collector that uses the provided telemetry settings.
func NewLogComponent(set component.TelemetrySettings) corelog.Component {
	_ = "STUB: not implemented"
	return *new(corelog.Component)
}

// NewSerializerComponent creates a new serializer that serializes and compresses payloads prior to being forwarded
func NewSerializerComponent(cfg coreconfig.Component, logger corelog.Component, hostname string) SerializerWithForwarder {
	_ = "STUB: not implemented"
	return *new(SerializerWithForwarder)
}

// NewConfigComponent creates a new Datadog agent config component with the given options.
// This function uses the options pattern to allow different modules to configure
// the component with their specific needs.
func NewConfigComponent(options ...ConfigOption) coreconfig.Component {
	_ = "STUB: not implemented"
	return *new(coreconfig.Component)
}

// Apply all configuration options

// WithAPIConfig configures API-related settings
func WithAPIConfig(cfg *datadogconfig.Config) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

// WithForwarderConfig configures forwarder-related settings
func WithForwarderConfig() ConfigOption { _ = "STUB: not implemented"; return *new(ConfigOption) }

// WithLogsEnabled enables logs for agent config
func WithLogsEnabled() ConfigOption { _ = "STUB: not implemented"; return *new(ConfigOption) }

// WithLogsConfig configures logs-related settings (requires WithLogsEnabled)
func WithLogsConfig(cfg *datadogconfig.Config) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

// WithLogsDefaults configures logs default settings (requires WithLogsEnabled)
func WithLogsDefaults() ConfigOption { _ = "STUB: not implemented"; return *new(ConfigOption) }

// add logs config pipelines config value, see https://github.com/DataDog/datadog-agent/pull/31190

// WithLogLevel configures log level settings (requires WithLogsEnabled)
func WithLogLevel(set component.TelemetrySettings) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

// WithPayloadsConfig configures payload settings
func WithPayloadsConfig() ConfigOption { _ = "STUB: not implemented"; return *new(ConfigOption) }

// WithProxy configures proxy settings from config or environment variables
func WithProxy(cfg *datadogconfig.Config) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

// WithCustomConfig allows setting arbitrary configuration values
func WithCustomConfig(key string, value any, source pkgconfigmodel.Source) ConfigOption {
	_ = "STUB: not implemented"
	return *new(ConfigOption)
}

func setProxy(cfg *datadogconfig.Config, pkgconfig pkgconfigmodel.Config) {
	_ = "STUB: not implemented"
	return
}

// proxy_url takes precedence over proxy environment variables if set

// If this is set to an empty []string, viper will have a type conflict when merging
// this config during secrets resolution. It unmarshals empty yaml lists to type
// []any, which will then conflict with type []string and fail to merge.

// newForwarderComponent creates a new forwarder that sends payloads to Datadog backend
func newForwarderComponent(cfg coreconfig.Component, log corelog.Component) forwarderWithLifecycle {
	_ = "STUB: not implemented"
	return *new(forwarderWithLifecycle)
}
