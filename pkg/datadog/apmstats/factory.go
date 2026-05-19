// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package apmstats // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/apmstats"

import (
	"context"

	"github.com/DataDog/datadog-agent/comp/core/tagger/types"
	"github.com/DataDog/datadog-agent/pkg/trace/stats"
	"github.com/DataDog/datadog-agent/pkg/util/option"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
)

type factory struct {
	tagger       types.TaggerClient
	concentrator *stats.Concentrator
	hostname     option.Option[string]
}

// SourceProviderFunc is a function that returns the source of the host.
type SourceProviderFunc func(context.Context) (string, error)

// NewConnectorFactory creates a factory for datadog connector for use in OTel agent
func NewConnectorFactory(componentType component.Type, metricsStability, traceStability component.StabilityLevel, tagger types.TaggerClient, hostGetter SourceProviderFunc, concentrator *stats.Concentrator) connector.Factory {
	_ = "STUB: not implemented"
	return *new(connector.Factory)
}

//  OTel connector factory to make a factory for connectors

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// defines the consumer type of the connector
// we want to consume traces and export metrics therefore define nextConsumer as metrics, consumer is the next component in the pipeline
func (f *factory) createTracesToMetricsConnector(_ context.Context, params connector.Settings, cfg component.Config, nextConsumer consumer.Metrics) (c connector.Traces, err error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}

func createTracesToTracesConnector(_ context.Context, params connector.Settings, _ component.Config, nextConsumer consumer.Traces) (connector.Traces, error) {
	_ = "STUB: not implemented"
	return *new(connector.Traces), nil
}
