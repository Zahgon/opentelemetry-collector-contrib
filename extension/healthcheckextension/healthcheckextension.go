// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package healthcheckextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension/extensioncapabilities"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension/internal/healthcheck"
)

type healthCheckExtension struct {
	config   Config
	logger   *zap.Logger
	state    *healthcheck.HealthCheck
	server   *http.Server
	stopCh   chan struct{}
	settings component.TelemetrySettings
}

var _ extensioncapabilities.PipelineWatcher = (*healthCheckExtension)(nil)

func (hc *healthCheckExtension) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Mount HC handler

// The listener ownership goes to the server.

// base handler function
func (hc *healthCheckExtension) baseHandler() http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (hc *healthCheckExtension) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (hc *healthCheckExtension) Ready() error { _ = "STUB: not implemented"; return nil }

func (hc *healthCheckExtension) NotReady() error { _ = "STUB: not implemented"; return nil }

func newServer(config Config, settings component.TelemetrySettings) *healthCheckExtension {
	_ = "STUB: not implemented"
	return nil
}
