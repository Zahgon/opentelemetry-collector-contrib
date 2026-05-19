// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package failoverconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/failoverconnector"
import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

type metricsRouter struct {
	*baseFailoverRouter[consumer.Metrics]
}

func newMetricsRouter(provider consumerProvider[consumer.Metrics], cfg *Config) (*metricsRouter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Consume is the metrics-specific consumption method
func (f *metricsRouter) Consume(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// consumeByHealthyPipeline will consume the metrics by the current healthy level
func (f *metricsRouter) consumeByHealthyPipeline(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// sampleRetryConsumers iterates through all unhealthy consumers to re-establish a healthy connection
func (f *metricsRouter) sampleRetryConsumers(ctx context.Context, md pmetric.Metrics) bool {
	_ = "STUB: not implemented"
	return false
}

type metricsFailover struct {
	component.StartFunc
	component.ShutdownFunc

	config   *Config
	failover *metricsRouter
	logger   *zap.Logger
}

func (*metricsFailover) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

// ConsumeMetrics will try to export to the current set priority level and handle failover in the case of an error
func (f *metricsFailover) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *metricsFailover) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func newMetricsToMetrics(set connector.Settings, cfg component.Config, metrics consumer.Metrics) (connector.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(connector.Metrics), nil
}
