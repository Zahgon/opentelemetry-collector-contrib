// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
// Originally copied from https://github.com/signalfx/signalfx-agent/blob/fbc24b0fdd3884bd0bbfbd69fe3c83f49d4c0b77/pkg/apm/tracetracker/tracker.go

package tracetracker // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/apm/tracetracker"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/apm/correlations"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/apm/log"
)

// fallbackEnvironment is the environment value to use if no environment is found in the span.
// This is the same value that is being set on the backend on spans that don't have an environment.
const fallbackEnvironment = "unknown"

// DefaultDimsToSyncSource are the default dimensions to sync correlated environment and services onto.
var DefaultDimsToSyncSource = map[string]string{
	"container_id":       "container_id",
	"kubernetes_pod_uid": "kubernetes_pod_uid",
}

// ActiveServiceTracker keeps track of which services are seen in the trace
// spans passed through ProcessSpans.  It supports expiry of service names if
// they are not seen for a certain amount of time.
type ActiveServiceTracker struct {
	log log.Logger

	// hostIDDims is the map of key/values discovered by the agent that identify the host
	hostIDDims map[string]string

	// hostServiceCache is a cache of services associated with the host
	hostServiceCache *TimeoutCache

	// hostEnvironmentCache is a cache of environments associated with the host
	hostEnvironmentCache *TimeoutCache

	// tenantServiceCache is the cache for services related to containers/pods
	tenantServiceCache *TimeoutCache

	// tenantEnvironmentCache is the cache for environments related to containers/pods
	tenantEnvironmentCache *TimeoutCache

	// tenantEnvironmentEmptyCache is the cache to make sure we don't send multiple requests to mitigate a very
	// specific corner case.  This is a separate cache so we don't impact metrics.  See the note in processEnvironment()
	// for more information
	tenantEmptyEnvironmentCache *TimeoutCache

	timeNow func() time.Time

	// correlationClient is the client used for updating infrastructure correlation properties
	correlationClient correlations.CorrelationClient

	// Map of dimensions to sync to with the key being the span attribute to lookup and the value being
	// the dimension to sync to.
	dimsToSyncSource map[string]string
}

// LoadHostIDDimCorrelations asynchronously retrieves all known correlations from the backend
// for all known hostIDDims.  This allows the agent to timeout and manage correlation
// deletions on restart.
func (a *ActiveServiceTracker) LoadHostIDDimCorrelations() {
	_ = "STUB: not implemented"
	// asynchronously fetch all services and environments for each hostIDDim at startup
	return
}

// Note that only the value is set for the host service cache because we only track services for the host
// therefore there we don't need to include the dim key and value on the cache key

// Note that only the value is set for the host environment cache because we only track environments for the host
// therefore there we don't need to include the dim key and value on the cache key

// New creates a new initialized service tracker
func New(
	log log.Logger,
	timeout time.Duration,
	correlationClient correlations.CorrelationClient,
	hostIDDims map[string]string,
	dimsToSyncSource map[string]string,
) *ActiveServiceTracker {
	_ = "STUB: not implemented"
	return nil
}

// ProcessTraces accepts a list of trace spans and uses them to update the
// current list of active services.  This is thread-safe.
func (a *ActiveServiceTracker) ProcessTraces(_ context.Context, traces ptrace.Traces) {
	_ = "STUB: not implemented"
	// Take current time once since this is a system call.
	return
}

func (a *ActiveServiceTracker) processEnvironment(res pcommon.Resource, now time.Time) {
	_ = "STUB: not implemented"
	return
}

// Determine the environment value from the incoming spans.
// First check "deployment.environment.name" attribute (new OTel standard).
// Then check "deployment.environment" attribute (deprecated).
// Then, try "environment" attribute (SignalFx schema).
// Otherwise, use the same fallback value as set on the backend.

// update the environment for the hostIDDims
// Note that only the value is set for the host environment cache because we only track environments for the host
// therefore there we don't need to include the dim key and value on the cache key

//nolint:errorlint

// container / pod level stuff
// this cache is necessary to identify environments associated with a kubernetes pod or container id

// Note that the value is not set on the cache key.  We only send the first environment received for a
// given pod/container, and we never delete the values set on the container/pod dimension.
// So we only need to cache the dim name and dim value that have been associated with an environment.

func (a *ActiveServiceTracker) processService(res pcommon.Resource, now time.Time) {
	_ = "STUB: not implemented"
	// Can't do anything if the spans don't have a local service name
	return
}

// Handle host level service and environment correlation
// Note that only the value is set for the host service cache because we only track services for the host
// therefore there we don't need to include the dim key and value on the cache key

// all of the host id dims need to be correlated with the service

//nolint:errorlint

// container / pod level stuff (this should not directly affect the active service count)
// this cache is necessary to identify services associated with a kubernetes pod or container id

// Note that the value is not set on the cache key.  We only send the first service received for a
// given pod/container, and we never delete the values set on the container/pod dimension.
// So we only need to cache the dim name and dim value that have been associated with a service.

// Purges caches on the ActiveServiceTracker
func (a *ActiveServiceTracker) Purge() { _ = "STUB: not implemented"; return }

// delete the correlation from all host id dims

// delete the correlation from all host id dims

// Purge the caches for containers and pods, but don't do the deletions.
// These values aren't expected to change, and can be overwritten.
// The onPurge() function doesn't need to do anything
