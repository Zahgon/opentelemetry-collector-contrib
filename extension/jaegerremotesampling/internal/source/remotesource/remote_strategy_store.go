// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package remotesource // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling/internal/source/remotesource"

import (
	"context"
	"io"
	"time"

	"github.com/jaegertracing/jaeger-idl/proto-gen/api_v2"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configopaque"
	"google.golang.org/grpc"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling/internal/source"
)

type grpcRemoteStrategyStore struct {
	headerAdditions configopaque.MapList
	delegate        *ConfigManagerProxy
	cache           serviceStrategyCache
}

// NewRemoteSource returns a StrategyStore that delegates to the configured Jaeger gRPC endpoint, making
// extension-configured enhancements (header additions only for now) to the gRPC context of every outbound gRPC call.
// Note: it would be nice to expand the configuration surface to include an optional TTL-based caching behavior
// for service-specific outbound GetSamplingStrategy calls.
func NewRemoteSource(
	conn *grpc.ClientConn,
	grpcClientSettings *configgrpc.ClientConfig,
	reloadInterval time.Duration,
) (source.Source, io.Closer) {
	_ = "STUB: not implemented"
	return *new(source.Source), *new(io.Closer)
}

func (g *grpcRemoteStrategyStore) GetSamplingStrategy(
	ctx context.Context,
	serviceName string,
) (*api_v2.SamplingStrategyResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This function is used to add the extension configuration defined HTTP headers to a given outbound gRPC call's context.
func (g *grpcRemoteStrategyStore) enhanceContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (g *grpcRemoteStrategyStore) Close() error { _ = "STUB: not implemented"; return nil }
