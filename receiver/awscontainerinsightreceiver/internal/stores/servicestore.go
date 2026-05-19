// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stores // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/stores"

import (
	"context"
	"time"

	"go.uber.org/zap"
)

const (
	refreshIntervalService = 10 * time.Second
)

type endpointInfo interface {
	PodKeyToServiceNames() map[string][]string
}

type ServiceStore struct {
	podKeyToServiceNamesMap map[string][]string
	endpointInfo            endpointInfo
	lastRefreshed           time.Time
	logger                  *zap.Logger
}

func NewServiceStore(logger *zap.Logger) (*ServiceStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *ServiceStore) RefreshTick(ctx context.Context) { _ = "STUB: not implemented"; return }

// Decorate decorates metrics and update kubernetesBlob
// service info is not mandatory
func (s *ServiceStore) Decorate(_ context.Context, metric CIMetric, _ map[string]any) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *ServiceStore) refresh(ctx context.Context) { _ = "STUB: not implemented"; return }

func addServiceNameTag(metric CIMetric, serviceNames []string) {
	_ = "STUB: not implemented"
	// TODO handle serviceNames len is larger than 1. We need to duplicate the metric object
	return
}
