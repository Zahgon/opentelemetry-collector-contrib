// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package purefareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefareceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var _ receiver.Metrics = (*purefaReceiver)(nil)

type purefaReceiver struct {
	cfg  *Config
	set  receiver.Settings
	next consumer.Metrics

	wrapped receiver.Metrics
}

func newReceiver(cfg *Config, set receiver.Settings, next consumer.Metrics) *purefaReceiver {
	_ = "STUB: not implemented"
	return nil
}

func (r *purefaReceiver) Start(ctx context.Context, compHost component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Extracting environment & fa_array_name from commonLabel

func (r *purefaReceiver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
