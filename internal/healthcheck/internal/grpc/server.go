// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package grpc // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/healthcheck/internal/grpc"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/healthcheck/internal/common"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/status"
)

type Server struct {
	healthpb.UnimplementedHealthServer
	grpcServer            *grpc.Server
	aggregator            *status.Aggregator
	config                *Config
	componentHealthConfig *common.ComponentHealthConfig
	telemetry             component.TelemetrySettings
	doneCh                chan struct{}
	doneOnce              sync.Once
}

var _ component.Component = (*Server)(nil)

func NewServer(
	config *Config,
	componentHealthConfig *common.ComponentHealthConfig,
	telemetry component.TelemetrySettings,
	aggregator *status.Aggregator,
) *Server {
	_ = "STUB: not implemented"
	return nil
}

// Start implements the component.Component interface.
func (s *Server) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Server never started, ensure doneCh is closed so shutdown doesn't block

// Shutdown implements the component.Component interface.
func (s *Server) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Stop the server - this will eventually release the port even if context times out

// Context timed out, but server is stopping. Force stop to ensure port is released.
