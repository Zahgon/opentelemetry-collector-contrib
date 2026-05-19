// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package healthcheck // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/healthcheck"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componentstatus"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensioncapabilities"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/status"
)

type eventSourcePair struct {
	source *componentstatus.InstanceID
	event  *componentstatus.Event
}

type HealthCheckExtension struct {
	config        Config
	telemetry     component.TelemetrySettings
	aggregator    *status.Aggregator
	subcomponents []component.Component
	eventCh       chan *eventSourcePair
	readyCh       chan struct{}
	host          component.Host
	shutdownOnce  sync.Once
}

var (
	_ component.Component                   = (*HealthCheckExtension)(nil)
	_ extensioncapabilities.ConfigWatcher   = (*HealthCheckExtension)(nil)
	_ extensioncapabilities.PipelineWatcher = (*HealthCheckExtension)(nil)
)

func NewHealthCheckExtension(
	config Config,
	set extension.Settings,
) *HealthCheckExtension {
	_ = "STUB: not implemented"
	return nil
}

// Start processing events in the background so that our status watcher doesn't
// block others before the extension starts.

// Start implements the component.Component interface.
func (hc *HealthCheckExtension) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown implements the component.Component interface.
func (hc *HealthCheckExtension) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Preemptively send the stopped event, so it can be exported before shutdown

// ComponentStatusChanged implements the extension.StatusWatcher interface.
func (hc *HealthCheckExtension) ComponentStatusChanged(
	source *componentstatus.InstanceID,
	event *componentstatus.Event,
) {
	_ = "STUB: not implemented"
	// There can be late arriving events after shutdown. We need to close
	// the event channel so that this function doesn't block and we release all
	// goroutines, but attempting to write to a closed channel will panic; log
	// and recover.
	return
}

// NotifyConfig implements the extensioncapabilities.ConfigWatcher interface.
func (hc *HealthCheckExtension) NotifyConfig(ctx context.Context, conf *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil
}

// Ready implements the extension.PipelineWatcher interface.
func (hc *HealthCheckExtension) Ready() error { _ = "STUB: not implemented"; return nil }

// NotReady implements the extension.PipelineWatcher interface.
func (*HealthCheckExtension) NotReady() error { _ = "STUB: not implemented"; return nil }

func (hc *HealthCheckExtension) eventLoop() {
	_ = "STUB: not implemented"
	// Record events with component.StatusStarting, but queue other events until
	// PipelineWatcher.Ready is called. This prevents aggregate statuses from
	// flapping between StatusStarting and StatusOK as components are started
	// individually by the service.
	return
}

// After PipelineWatcher.Ready, record statuses as they are received.
