// Copyright The OpenTelemetry Authors
// Copyright (c) 2018 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package remotesource // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling/internal/source/remotesource"

import (
	"context"

	"github.com/jaegertracing/jaeger-idl/proto-gen/api_v2"
	"google.golang.org/grpc"
)

// ConfigManagerProxy returns sampling decisions from collector over gRPC.
type ConfigManagerProxy struct {
	client api_v2.SamplingManagerClient
}

// NewConfigManager creates gRPC sampling manager.
func NewConfigManager(conn *grpc.ClientConn) *ConfigManagerProxy {
	_ = "STUB: not implemented"
	return nil
}

// GetSamplingStrategy returns sampling strategies from collector.
func (s *ConfigManagerProxy) GetSamplingStrategy(ctx context.Context, serviceName string) (*api_v2.SamplingStrategyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
