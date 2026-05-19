// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !aix

package datadogexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter"

import (
	"context"
	"sync"
	"time"

	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/inframetadata"
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes"
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes/source"
	"github.com/DataDog/datadog-agent/pkg/trace/agent"
	"github.com/DataDog/datadog-agent/pkg/trace/writer"
	"github.com/DataDog/datadog-agent/pkg/util/log"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/featuregate"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata"
	datadogconfig "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/config"
)

func init() {
	log.SetupLogger(log.Disabled(), "off")
}

func isMetricExportSerializerEnabled() bool { _ = "STUB: not implemented"; return false }

func consumeResource(metadataReporter *inframetadata.Reporter, res pcommon.Resource, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

type factory struct {
	onceMetadata sync.Once

	onceProvider   sync.Once
	sourceProvider source.Provider
	providerErr    error

	onceReporter     sync.Once
	onceStopReporter sync.Once
	reporter         *inframetadata.Reporter
	reporterErr      error

	onceAttributesTranslator sync.Once
	attributesTranslator     *attributes.Translator
	attributesErr            error

	registry *featuregate.Registry

	gatewayUsage *attributes.GatewayUsage
}

func (f *factory) SourceProvider(set component.TelemetrySettings, configHostname string, timeout time.Duration) (source.Provider, error) {
	_ = "STUB: not implemented"
	return *new(source.Provider), nil
}

func (f *factory) AttributesTranslator(set component.TelemetrySettings) (*attributes.Translator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reporter builds and returns an *inframetadata.Reporter.
func (f *factory) Reporter(params exporter.Settings, pcfg hostmetadata.PusherConfig) (*inframetadata.Reporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StopReporter stops the host metadata reporter.
func (f *factory) StopReporter() {
	_ = "STUB: not implemented"
	// Use onceReporter or wait until it is done
	return
}

// Stop the reporter

func (*factory) TraceAgent(ctx context.Context, wg *sync.WaitGroup, params exporter.Settings, cfg *datadogconfig.Config, sourceProvider source.Provider, attrsTranslator *attributes.Translator) (*agent.Agent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newFactoryWithRegistry(registry *featuregate.Registry) exporter.Factory {
	_ = "STUB: not implemented"
	return *new(exporter.Factory)
}

// NewFactory creates a Datadog exporter factory
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func defaultClientConfig() confighttp.ClientConfig {
	_ = "STUB: not implemented"
	return *new(confighttp.ClientConfig)
}

// createDefaultConfig creates the default exporter configuration
func (*factory) createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func (*factory) consumeStatsPayload(ctx context.Context, wg *sync.WaitGroup, statsIn <-chan []byte, statsWriter *writer.DatadogStatsWriter, tracerVersion, agentVersion string, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// The DD Connector doesn't set the agent version, so we'll set it here

// createMetricsExporter creates a metrics exporter based on this config.
func (f *factory) createMetricsExporter(
	ctx context.Context,
	set exporter.Settings,
	c component.Config,
) (exporter.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Metrics), nil
}

// cancel() runs on shutdown

// waits for consumeStatsPayload to exit

// Don't start a `Reporter` if host metadata is disabled.

// only sending metadata use only metrics

// Consume resources for host metadata

// first cancel context
// then wait for shutdown

// Start the hostmetadata pusher once.
// It sends the hostmetadata for the host where the collector is running.

// first cancel context
// then wait for shutdown

// first cancel context
// then wait for shutdown

// explicitly disable since we rely on http.Client timeout logic.

// We use our own custom mechanism for retries, since we hit several endpoints.

// The metrics remapping code mutates data

// first cancel context
// then wait for shutdown

// createTracesExporter creates a trace exporter based on this config.
func (f *factory) createTracesExporter(
	ctx context.Context,
	set exporter.Settings,
	c component.Config,
) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

// waits for agent to exit

// cancel() runs on shutdown

// Don't start a `Reporter` if host metadata is disabled.

// only host metadata needs to be sent, once.

// Consume resources for host metadata

// then wait for shutdown

// first cancel context

// explicitly disable since we rely on http.Client timeout logic.

// We don't do retries on traces because of deduping concerns on APM Events.

// createLogsExporter creates a logs exporter based on the config.
func (f *factory) createLogsExporter(
	ctx context.Context,
	set exporter.Settings,
	c component.Config,
) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}

// cancel() runs on shutdown

// Don't start a `Reporter` if host metadata is disabled.

// only host metadata needs to be sent, once.

// explicitly disable since we rely on http.Client timeout logic.
