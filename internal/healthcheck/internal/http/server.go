// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package http // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/healthcheck/internal/http"

import (
	"context"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/extension/extensioncapabilities"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/healthcheck/internal/common"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/status"
)

type Server struct {
	telemetry      component.TelemetrySettings
	httpConfig     confighttp.ServerConfig
	httpServer     *http.Server
	mux            *http.ServeMux
	responder      responder
	colconf        atomic.Value
	aggregator     *status.Aggregator
	startTimestamp time.Time
	doneWg         sync.WaitGroup
	doneCh         chan struct{}
	doneOnce       sync.Once
}

var (
	_ component.Component                 = (*Server)(nil)
	_ extensioncapabilities.ConfigWatcher = (*Server)(nil)
)

func NewServer(
	config *Config,
	legacyConfig LegacyConfig,
	componentHealthConfig *common.ComponentHealthConfig,
	telemetry component.TelemetrySettings,
	aggregator *status.Aggregator,
) *Server {
	_ = "STUB: not implemented"
	return nil
}

// default for backward compatibility

// Start implements the component.Component interface.
func (s *Server) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Server never started, ensure doneCh is closed so shutdown doesn't block

// Shutdown implements the component.Component interface.
func (s *Server) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// NotifyConfig implements the extension.ConfigWatcher interface.
func (s *Server) NotifyConfig(_ context.Context, conf *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil
}
