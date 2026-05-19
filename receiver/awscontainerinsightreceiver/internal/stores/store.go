// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stores // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/stores"

import (
	"context"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor/extractors"
)

var _ cadvisor.Decorator = &K8sDecorator{}

// CIMetric represents the raw metric interface for container insights
type CIMetric interface {
	HasField(key string) bool
	AddField(key string, val any)
	GetField(key string) any
	HasTag(key string) bool
	AddTag(key, val string)
	GetTag(key string) string
	RemoveTag(key string)
}

type K8sStore interface {
	Decorate(ctx context.Context, metric CIMetric, kubernetesBlob map[string]any) bool
	RefreshTick(ctx context.Context)
}

type K8sDecorator struct {
	stores []K8sStore
	// We save ctx in the struct because it is used in Decorate(...) function when calling K8sStore.Decorate(...)
	// It would be easier to keep the ctx here than passing it as a parameter for Decorate(...) function.
	// The K8sStore (e.g. podstore) does network request in Decorate function, thus needs to take a context
	// object for canceling the request
	ctx context.Context
	// the pod store needs to be saved here because the map it is stateful and needs to be shut down.
	podStore *PodStore
}

func NewK8sDecorator(ctx context.Context, tagService, prefFullPodName, addFullPodNameMetricLabel bool, logger *zap.Logger) (*K8sDecorator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *K8sDecorator) Decorate(metric *extractors.CAdvisorMetric) *extractors.CAdvisorMetric {
	_ = "STUB: not implemented"
	return nil
}

func (k *K8sDecorator) Shutdown() error { _ = "STUB: not implemented"; return nil }
