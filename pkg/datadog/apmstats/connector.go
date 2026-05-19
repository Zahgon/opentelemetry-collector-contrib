// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package apmstats // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/apmstats"

import (
	"context"
	"sync"

	"github.com/DataDog/datadog-agent/comp/core/tagger/types"
	"github.com/DataDog/datadog-agent/pkg/obfuscate"
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes"
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/metrics"
	pb "github.com/DataDog/datadog-agent/pkg/proto/pbgo/trace"
	"github.com/DataDog/datadog-agent/pkg/trace/config"
	"github.com/DataDog/datadog-agent/pkg/trace/stats"
	"github.com/DataDog/datadog-agent/pkg/util/option"
	"github.com/DataDog/datadog-go/v5/statsd"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	datadogconfig "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/config"
)

// traceToMetricConnector is the schema for connector
type traceToMetricConnector struct {
	metricsConsumer consumer.Metrics // the next component in the pipeline to ingest metrics after connector
	logger          *zap.Logger

	wg sync.WaitGroup

	// concentrator ingests spans and produces APM stats
	concentrator *stats.Concentrator

	// tcfg is the trace agent config
	tcfg *config.AgentConfig

	// ctagKeys are container tag keys
	ctagKeys []string

	// peerTagKeys are peer tag keys to group APM stats
	peerTagKeys []string

	// translator specifies the translator used to transform APM Stats Payloads
	// from the agent to OTLP Metrics.
	// We use the deprecated Translator type because it's the only one that provides
	// the StatsToMetrics method needed for APM stats conversion.
	//nolint:staticcheck // SA1019: Using deprecated Translator type for StatsToMetrics functionality
	translator *metrics.Translator

	// statsout specifies the channel through which the agent will output Stats Payloads
	// resulting from ingested traces.
	statsout chan *pb.StatsPayload

	// obfuscator is used to obfuscate sensitive data from various span
	// tags based on their type.
	obfuscator *obfuscate.Obfuscator

	// exit specifies the exit channel, which will be closed upon shutdown.
	exit chan struct{}

	// isStarted tracks whether Start() has been called.
	isStarted bool
}

// newTraceToMetricConnector creates a new connector with native OTel span ingestion
func newTraceToMetricConnector(set component.TelemetrySettings, cfg component.Config, metricsConsumer consumer.Metrics, metricsClient statsd.ClientInterface, concentrator *stats.Concentrator, tagger types.TaggerClient, hostnameOpt option.Option[string]) (*traceToMetricConnector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// disable metrics for the connector

// We use the deprecated NewTranslator because the new NewDefaultTranslator
// doesn't provide the StatsToMetrics method which is required for APM stats conversion.
//nolint:staticcheck // SA1019: Using deprecated NewTranslator for StatsToMetrics functionality

func getTraceAgentCfg(logger *zap.Logger, cfg datadogconfig.TracesConnectorConfig, attributesTranslator *attributes.Translator, tagger types.TaggerClient, hostnameOpt option.Option[string]) *config.AgentConfig {
	_ = "STUB: not implemented"
	return nil
}

// ==== START for DDOT

// ==== END for DDOT

var _ component.Component = (*traceToMetricConnector)(nil) // testing that the connectorImp properly implements the type Component interface

// Start implements the component.Component interface.
func (c *traceToMetricConnector) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown implements the component.Component interface.
func (c *traceToMetricConnector) Shutdown(context.Context) error {
	_ = "STUB: not implemented"

	// Note: it is not necessary to manually close c.exit, c.in and c.concentrator.exit channels as these are unused.
	return nil
}

// stop the obfuscator and concentrator and wait for the run loop to exit

// wait for close

// Capabilities implements the consumer interface.
// tells use whether the component(connector) will mutate the data passed into it. if set to true the connector does modify the data
func (*traceToMetricConnector) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (c *traceToMetricConnector) ConsumeTraces(_ context.Context, traces ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// run awaits incoming stats resulting from the agent's ingestion, converts them
// to metrics and flushes them using the configured metrics exporter.
func (c *traceToMetricConnector) run() { _ = "STUB: not implemented"; return }

// APM stats as metrics

// send metrics to the consumer or next component in pipeline

var _ stats.Writer = (*otelStatsWriter)(nil)

type otelStatsWriter struct {
	out chan *pb.StatsPayload
}

// Write this payload to the `out` channel
func (a *otelStatsWriter) Write(payload *pb.StatsPayload) { _ = "STUB: not implemented"; return }
