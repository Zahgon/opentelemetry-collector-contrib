// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package provider // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata/provider"

import (
	"context"

	"go.uber.org/zap"
)

type ClusterNameProvider interface {
	ClusterName(context.Context) (string, error)
}

var _ ClusterNameProvider = (*chainClusterProvider)(nil)

type chainClusterProvider struct {
	logger       *zap.Logger
	providers    map[string]ClusterNameProvider
	priorityList []string
}

func (p *chainClusterProvider) ClusterName(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Chain providers into a single provider that returns the first available hostname.
func ChainCluster(logger *zap.Logger, providers map[string]ClusterNameProvider, priorityList []string) (ClusterNameProvider, error) {
	_ = "STUB: not implemented"
	return *new(ClusterNameProvider), nil
}
