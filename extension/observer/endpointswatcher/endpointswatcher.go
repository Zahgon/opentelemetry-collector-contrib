// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package endpointswatcher // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/endpointswatcher"

import (
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
)

var _ observer.Observable = (*EndpointsWatcher)(nil)

// EndpointsWatcher provides a generic mechanism to run EndpointsLister.ListEndpoints every
// RefreshInterval and report any new or removed endpoints to Notify instances registered
// via ListAndWatch. Any observer that lists endpoints can make use of EndpointsWatcher
// to poll for endpoints by embedding this struct and using NewEndpointsWatcher().
type EndpointsWatcher struct {
	EndpointsLister EndpointsLister
	RefreshInterval time.Duration

	// subscribed Notify instances ~sync.Map(map[NotifyID]Notify)
	toNotify sync.Map
	// map of NotifyID to known endpoints for that Notify (subscriptions can occur at different times in service startup).
	// ~sync.Map(map[NotifyID]map[EndpointID]Endpoint)
	existingEndpoints sync.Map
	stop              chan struct{}
	once              *sync.Once
	logger            *zap.Logger
}

func New(endpointsLister EndpointsLister, refreshInterval time.Duration, logger *zap.Logger) *EndpointsWatcher {
	_ = "STUB: not implemented"
	return nil
}

// ListAndWatch runs EndpointsLister.ListEndpoints() on a regular interval and keeps track of the results
// for alerting all subscribed Notify's of the based on the differences from the previous call.
func (ew *EndpointsWatcher) ListAndWatch(notify observer.Notify) { _ = "STUB: not implemented"; return }

func (ew *EndpointsWatcher) Unsubscribe(notify observer.Notify) { _ = "STUB: not implemented"; return }

// notifyOfLatestEndpoints alerts subscribed Notify instances by their NotifyID of latest Endpoint events,
// updating their internal store with results of ListEndpoints() call.
func (ew *EndpointsWatcher) notifyOfLatestEndpoints(notifyIDs ...observer.NotifyID) {
	_ = "STUB: not implemented"
	return
}

// an Unsubscribe() must have occurred during this call

func (ew *EndpointsWatcher) updateAndNotifyOfEndpoints(notify observer.Notify, endpoints []observer.Endpoint, done *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

func (ew *EndpointsWatcher) updateEndpoints(notify observer.Notify, endpoints []observer.Endpoint) (removed, added, changed []observer.Endpoint) {
	_ = "STUB: not implemented"
	return nil,

		// Create map from ID to endpoint for lookup.
		nil, nil
}

// copy to not modify sync.Map value directly (will be reloaded)

// Iterate over the latest endpoints obtained. An endpoint needs
// to be added or updated in case it is not already available in existingEndpoints or doesn't match
// the latest value.

// Collect updated endpoints.

// If endpoint present in existingEndpoints does not exist in the latest
// list, it needs to be removed.

// StopListAndWatch polling the ListEndpoints.
func (ew *EndpointsWatcher) StopListAndWatch() { _ = "STUB: not implemented"; return }

// EndpointsLister that provides a list of endpoints.
type EndpointsLister interface {
	// ListEndpoints provides a list of endpoints and is expected to be
	// implemented by an observer looking for endpoints.
	ListEndpoints() []observer.Endpoint
}

func (ew *EndpointsWatcher) logEndpointEvent(msg string, notify observer.Notify, endpoints []observer.Endpoint) {
	_ = "STUB: not implemented"
	return
}

func endpointsEqual(lhs, rhs observer.Endpoint) bool { _ = "STUB: not implemented"; return false }
