// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package k8s contains the Kubernetes hostname provider
package k8s // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata/internal/k8s"

import (
	"context"

	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes/source"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata/provider"
)

var _ source.Provider = (*Provider)(nil)

type Provider struct {
	logger              *zap.Logger
	nodeNameProvider    nodeNameProvider
	clusterNameProvider provider.ClusterNameProvider
}

// Hostname returns the Kubernetes node name followed by the cluster name if available.
func (p *Provider) Source(ctx context.Context) (source.Source, error) {
	_ = "STUB: not implemented"
	return *new(source.Source), nil
}

// NewProvider creates a new Kubernetes hostname provider.
func NewProvider(logger *zap.Logger, clusterProvider provider.ClusterNameProvider) (*Provider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
