// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsecscontainermetrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver/internal/awsecscontainermetrics"

import (
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil"
)

// StatsProvider wraps a RestClient, returning an unmarshaled metadata and docker stats
type StatsProvider struct {
	rc               ecsutil.RestClient
	metadataProvider ecsutil.MetadataProvider
}

// NewStatsProvider returns a new stats provider
func NewStatsProvider(rc ecsutil.RestClient, logger *zap.Logger) *StatsProvider {
	_ = "STUB: not implemented"
	return nil
}

// GetStats calls the ecs task metadata endpoint and unmarshals the data
func (p *StatsProvider) GetStats() (map[string]*ContainerStats, ecsutil.TaskMetadata, error) {
	_ = "STUB: not implemented"
	return nil, *new(ecsutil.TaskMetadata), nil
}
