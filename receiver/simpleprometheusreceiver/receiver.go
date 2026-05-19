// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package simpleprometheusreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/simpleprometheusreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver"
)

type prometheusReceiverWrapper struct {
	params             receiver.Settings
	config             *Config
	consumer           consumer.Metrics
	prometheusReceiver receiver.Metrics
}

// newPrometheusReceiverWrapper returns a prometheusReceiverWrapper
func newPrometheusReceiverWrapper(params receiver.Settings, cfg *Config, consumer consumer.Metrics) *prometheusReceiverWrapper {
	_ = "STUB: not implemented"
	return nil
}

// Start creates and starts the prometheus receiver.
func (prw *prometheusReceiverWrapper) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Deprecated: [v0.55.0] Use getPrometheusConfig instead.
func getPrometheusConfigWrapper(cfg *Config, params receiver.Settings) (*prometheusreceiver.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getPrometheusConfig(cfg *Config) (*prometheusreceiver.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Shutdown stops the underlying Prometheus receiver.
func (prw *prometheusReceiverWrapper) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
