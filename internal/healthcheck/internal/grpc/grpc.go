// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package grpc // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/healthcheck/internal/grpc"

import (
	"context"

	"go.opentelemetry.io/collector/component/componentstatus"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	grpcstatus "google.golang.org/grpc/status"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/status"
)

var (
	errNotFound     = grpcstatus.Error(codes.NotFound, "Service not found.")
	errShuttingDown = grpcstatus.Error(codes.Canceled, "Server shutting down.")
	errStreamSend   = grpcstatus.Error(codes.Canceled, "Error sending; stream terminated.")
	errStreamEnded  = grpcstatus.Error(codes.Canceled, "Stream has ended.")

	statusToServingStatusMap = map[componentstatus.Status]healthpb.HealthCheckResponse_ServingStatus{
		componentstatus.StatusNone:             healthpb.HealthCheckResponse_NOT_SERVING,
		componentstatus.StatusStarting:         healthpb.HealthCheckResponse_NOT_SERVING,
		componentstatus.StatusOK:               healthpb.HealthCheckResponse_SERVING,
		componentstatus.StatusRecoverableError: healthpb.HealthCheckResponse_SERVING,
		componentstatus.StatusPermanentError:   healthpb.HealthCheckResponse_SERVING,
		componentstatus.StatusFatalError:       healthpb.HealthCheckResponse_NOT_SERVING,
		componentstatus.StatusStopping:         healthpb.HealthCheckResponse_NOT_SERVING,
		componentstatus.StatusStopped:          healthpb.HealthCheckResponse_NOT_SERVING,
	}
)

func (s *Server) Check(
	_ context.Context,
	req *healthpb.HealthCheckRequest,
) (*healthpb.HealthCheckResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Server) Watch(req *healthpb.HealthCheckRequest, stream healthpb.Health_WatchServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) toServingStatus(
	ev status.Event,
) healthpb.HealthCheckResponse_ServingStatus {
	_ = "STUB: not implemented"
	return *new(healthpb.HealthCheckResponse_ServingStatus)
}
