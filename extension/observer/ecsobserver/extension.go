// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

var _ extension.Extension = (*ecsObserver)(nil)

// ecsObserver implements component.ServiceExtension interface.
type ecsObserver struct {
	telemetrySettings component.TelemetrySettings
	sd                *serviceDiscovery

	// for Shutdown
	cancel func()
}

// Start runs the service discovery in background
func (e *ecsObserver) Start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Ignore the ctx parameter as it is not for long running operation

func (e *ecsObserver) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }
