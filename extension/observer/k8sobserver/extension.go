// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"k8s.io/client-go/tools/cache"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/endpointswatcher"
)

var (
	_ extension.Extension = (*k8sObserver)(nil)
	_ observer.Observable = (*k8sObserver)(nil)
)

type k8sObserver struct {
	*endpointswatcher.EndpointsWatcher
	telemetry             component.TelemetrySettings
	podListerWatchers     []cache.ListerWatcher
	serviceListerWatchers []cache.ListerWatcher
	ingressListerWatchers []cache.ListerWatcher
	nodeListerWatcher     cache.ListerWatcher
	handler               *handler
	once                  *sync.Once
	stop                  chan struct{}
	config                *Config
}

// Start will populate the cache.SharedInformers for pods and nodes as configured and run them as goroutines.
func (k *k8sObserver) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown tells any cache.SharedInformers to stop running.
func (k *k8sObserver) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

// newObserver creates a new k8s observer extension.
func newObserver(config *Config, set extension.Settings) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}
