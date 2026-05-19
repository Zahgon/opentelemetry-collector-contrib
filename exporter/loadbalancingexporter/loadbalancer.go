// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loadbalancingexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"

import (
	"context"
	"errors"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter/internal/metadata"
)

const (
	defaultPort = "4317"
)

var (
	errNoResolver                = errors.New("no resolvers specified for the exporter")
	errMultipleResolversProvided = errors.New("only one resolver should be specified")
)

type componentFactory func(ctx context.Context, endpoint string) (component.Component, error)

type loadBalancer struct {
	logger *zap.Logger
	host   component.Host

	res  resolver
	ring *hashRing

	componentFactory componentFactory
	exporters        map[string]*wrappedExporter

	stopped    bool
	updateLock sync.RWMutex
}

// Create new load balancer
func newLoadBalancer(logger *zap.Logger, cfg component.Config, factory componentFactory, telemetry *metadata.TelemetryBuilder) (*loadBalancer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lb *loadBalancer) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (lb *loadBalancer) onBackendChanges(resolved []string) { _ = "STUB: not implemented"; return }

// TODO: set a timeout?

// add the missing exporters first

func (lb *loadBalancer) addMissingExporters(ctx context.Context, endpoints []string) {
	_ = "STUB: not implemented"
	return
}

func endpointWithPort(endpoint string) string { _ = "STUB: not implemented"; return "" }

func (lb *loadBalancer) removeExtraExporters(ctx context.Context, endpoints []string) {
	_ = "STUB: not implemented"
	return
}

// Shutdown the exporter asynchronously to avoid blocking the resolver

func (lb *loadBalancer) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// exporterAndEndpoint returns the exporter and the endpoint for the given identifier.
func (lb *loadBalancer) exporterAndEndpoint(identifier []byte) (*wrappedExporter, string, error) {
	_ = "STUB: not implemented"
	// NOTE: make rolling updates of next tier of collectors work. currently, this may cause
	// data loss because the latest batches sent to outdated backend will never find their way out.
	// for details: https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/1690
	return nil, "", nil
}

// something is really wrong... how come we couldn't find the exporter??
