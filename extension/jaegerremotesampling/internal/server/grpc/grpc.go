// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package grpc // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling/internal/server/grpc"

import (
	"context"
	"errors"
	"net"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling/internal/source"
)

var _ component.Component = (*SamplingGRPCServer)(nil)

var (
	errMissingStrategyStore = errors.New("the strategy store has not been provided")
	errGRPCServerNotRunning = errors.New("gRPC server is not running")
)

type grpcServer interface {
	Serve(lis net.Listener) error
	GracefulStop()
	Stop()
}

// NewGRPC returns a new sampling gRPC Server.
func NewGRPC(
	telemetry component.TelemetrySettings,
	settings configgrpc.ServerConfig,
	strategyStore source.Source,
) (*SamplingGRPCServer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SamplingGRPCServer implements component.Component to make the life cycle easy to manage.
type SamplingGRPCServer struct {
	telemetry     component.TelemetrySettings
	settings      configgrpc.ServerConfig
	strategyStore source.Source

	grpcServer grpcServer
}

func (s *SamplingGRPCServer) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown tries to terminate connections gracefully as long as the passed context is valid.
func (s *SamplingGRPCServer) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}
