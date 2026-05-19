// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver"

import (
	"sync"

	"go.uber.org/zap"
	"k8s.io/client-go/tools/cache"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/endpointswatcher"
)

var (
	_ cache.ResourceEventHandler       = (*handler)(nil)
	_ endpointswatcher.EndpointsLister = (*handler)(nil)
)

// handler handles k8s cache informer callbacks.
type handler struct {
	// idNamespace should be some unique token to distinguish multiple handler instances.
	idNamespace string
	// endpoints is a map[observer.EndpointID]observer.Endpoint all existing endpoints at any given moment
	endpoints *sync.Map

	logger *zap.Logger
}

func (h *handler) ListEndpoints() []observer.Endpoint { _ = "STUB: not implemented"; return nil }

// OnAdd is called in response to a new pod or node being detected.
func (h *handler) OnAdd(objectInterface any, _ bool) { _ = "STUB: not implemented"; return }

// unsupported

// OnUpdate is called in response to an existing pod or node changing.
func (h *handler) OnUpdate(oldObjectInterface, newObjectInterface any) {
	_ = "STUB: not implemented"
	return
}

// unsupported

// Find endpoints that are present in oldPod and newPod and see if they've
// changed. Otherwise if it wasn't in oldPod it's a new endpoint.

// If an endpoint is present in the oldPod but not in the newPod then
// send as removed.

// OnDelete is called in response to a pod or node being deleted.
func (h *handler) OnDelete(objectInterface any) { _ = "STUB: not implemented"; return }

// Assuming we never saw the pod state where new endpoints would have been created
// to begin with it seems that we can't leak endpoints here.

// unsupported
