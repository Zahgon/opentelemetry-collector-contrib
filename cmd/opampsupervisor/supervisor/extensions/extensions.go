// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package extensions // import "github.com/open-telemetry/opentelemetry-collector-contrib/cmd/opampsupervisor/supervisor/extensions"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.uber.org/zap"
)

// TODO: Replace this implementation with 'go.opentelemetry.io/collector/service/extensions' once
// this issue is completed: https://github.com/open-telemetry/opentelemetry-collector/issues/15216

// Extensions manages the lifecycle of extensions configured in the supervisor.
type Extensions struct {
	extensions map[component.ID]extension.Extension
	// order is the deterministic start/shutdown order of extensions.
	order  []component.ID
	host   *host
	logger *zap.Logger
}

// host is the minimal component.Host implementation exposed to extensions.
type host struct {
	extensions map[component.ID]component.Component
}

func (h *host) GetExtensions() map[component.ID]component.Component {
	_ = "STUB: not implemented"
	return nil

	// New creates and returns a ready-to-start Extensions from already-validated
	// configs. Extension instances are created eagerly so that factory errors are
	// surfaced before any lifecycle begins. Start must be called to transition the
	// extensions to the running state; Shutdown must be called to stop them.
}

func New(
	ctx context.Context,
	cfg Config,
	factories map[component.Type]extension.Factory,
	telemetry component.TelemetrySettings,
) (*Extensions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start starts all extensions in deterministic order. If any extension fails
// to start, already-started extensions are shut down in reverse order.
func (e *Extensions) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Roll back already-started extensions in reverse order.

// Shutdown stops all extensions in reverse order. Errors from individual
// extensions are joined and returned.
func (e *Extensions) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// GetExtensions returns the extensions keyed by component ID.
func (e *Extensions) GetExtensions() map[component.ID]component.Component {
	_ = "STUB: not implemented"
	return nil
}
