// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package grafanacloudconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/grafanacloudconnector"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/connector"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/connector/grafanacloudconnector/internal/metadata"
)

const (
	hostInfoMetric     = "traces_host_info"
	hostIdentifierAttr = "grafana.host.id"
)

var _ connector.Traces = (*connectorImp)(nil)

type connectorImp struct {
	config Config
	logger *zap.Logger

	started      bool
	done         chan struct{}
	shutdownOnce sync.Once

	metricsConsumer consumer.Metrics
	hostMetrics     *hostMetrics

	telemetryBuilder *metadata.TelemetryBuilder
}

func newConnector(logger *zap.Logger, set component.TelemetrySettings, config component.Config) (*connectorImp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Capabilities implements connector.Traces.
func (*connectorImp) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

// ConsumeTraces implements connector.Traces.
func (c *connectorImp) ConsumeTraces(_ context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// Start implements connector.Traces.
func (c *connectorImp) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown implements connector.Traces.
func (c *connectorImp) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// flush metrics on shutdown

func (c *connectorImp) flush(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
