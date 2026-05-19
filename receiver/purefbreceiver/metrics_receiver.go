// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package purefbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefbreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var _ receiver.Metrics = (*purefbMetricsReceiver)(nil)

type purefbMetricsReceiver struct {
	cfg  *Config
	set  receiver.Settings
	next consumer.Metrics

	wrapped receiver.Metrics
}

func newReceiver(cfg *Config, set receiver.Settings, next consumer.Metrics) *purefbMetricsReceiver {
	_ = "STUB: not implemented"
	return nil
}

func (r *purefbMetricsReceiver) Start(ctx context.Context, compHost component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *purefbMetricsReceiver) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
