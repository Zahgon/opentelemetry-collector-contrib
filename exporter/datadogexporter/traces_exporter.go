// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !aix

package datadogexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter"

import (
	"context"
	"sync"

	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/inframetadata"
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes"
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes/source"
	"github.com/DataDog/datadog-agent/pkg/trace/agent"
	traceconfig "github.com/DataDog/datadog-agent/pkg/trace/config"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/DataDog/datadog-go/v5/statsd"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/clientutil"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/scrub"
	datadogconfig "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/config"
)

type traceExporter struct {
	params exporter.Settings
	cfg    *datadogconfig.Config
	ctx    context.Context // ctx triggers shutdown upon cancellation

	metricsAPI       *datadogV2.MetricsApi    // client sends running metrics to backend
	scrubber         scrub.Scrubber           // scrubber scrubs sensitive information from error messages
	onceMetadata     *sync.Once               // onceMetadata ensures that metadata is sent only once across all exporters
	agent            *agent.Agent             // agent processes incoming traces
	sourceProvider   source.Provider          // is able to source the origin of a trace (hostname, container, etc)
	metadataReporter *inframetadata.Reporter  // reports host metadata from resource attributes and metrics
	retrier          *clientutil.Retrier      // retrier handles retries on requests
	gatewayUsage     *attributes.GatewayUsage // gatewayUsage stores the gateway usage metrics
}

func newTracesExporter(
	ctx context.Context,
	params exporter.Settings,
	cfg *datadogconfig.Config,
	onceMetadata *sync.Once,
	sourceProvider source.Provider,
	agent *agent.Agent,
	metadataReporter *inframetadata.Reporter,
	gatewayUsage *attributes.GatewayUsage,
) (*traceExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// client to send running metric to the backend & perform API key validation

var _ consumer.ConsumeTracesFunc = (*traceExporter)(nil).consumeTraces

// headerComputedStats specifies the HTTP header which indicates whether APM stats
// have already been computed for a payload.
const headerComputedStats = "Datadog-Client-Computed-Stats"

func (exp *traceExporter) consumeTraces(
	ctx context.Context,
	td ptrace.Traces,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// start host metadata with resource attributes from
// the first payload.

// Consume resources for host metadata

func (exp *traceExporter) exportUsageMetrics(ctx context.Context, hosts, tags map[string]struct{}) {
	_ = "STUB: not implemented"
	return
}

func newTraceAgent(ctx context.Context, params exporter.Settings, cfg *datadogconfig.Config, sourceProvider source.Provider, metricsClient statsd.ClientInterface, attrsTranslator *attributes.Translator) (*agent.Agent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTraceAgentConfig(ctx context.Context, params exporter.Settings, cfg *datadogconfig.Config, sourceProvider source.Provider, attrsTranslator *attributes.Translator) (*traceconfig.AgentConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// disable HTTP receiver

// TODO: This shouldn't be a singleton
