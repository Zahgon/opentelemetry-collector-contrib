// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package hostmetricsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const entityType = "host"

type hostEntitiesReceiver struct {
	cfg *Config

	nextLogs consumer.Logs
	cancel   context.CancelFunc

	settings *receiver.Settings
}

func (hmr *hostEntitiesReceiver) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (hmr *hostEntitiesReceiver) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (hmr *hostEntitiesReceiver) sendEntityEvent(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// Note: receiver contract says that we need to retry sending if the
// returned error is not Permanent. However, we are not doing it here.
// Instead, we rely on the fact the metadata is collected periodically
// and the entity events will be delivered on the next cycle. This is
// fine because we deliver cumulative entity state.
// This allows us to avoid stressing the Collector or its destination
// unnecessarily (typically non-Permanent errors happen in stressed conditions).
