// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver"

import (
	v1 "k8s.io/api/core/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
)

// convertPodToEndpoints converts a pod instance into a slice of endpoints. The endpoints
// include the pod itself as well as an endpoint for each container port that is mapped
// to a container that is in a running state.
func convertPodToEndpoints(idNamespace string, pod *v1.Pod) []observer.Endpoint {
	_ = "STUB: not implemented"
	return nil
}

// Return no endpoints if the Pod is not running

// Map of running containers by name.

// Create endpoint for each named container port.

// Create endpoint for each named container port.

func getTransport(protocol v1.Protocol) observer.Transport {
	_ = "STUB: not implemented"
	return *new(observer.Transport)
}

// containerIDWithRuntime parses the container ID to get the actual ID string
func containerIDWithRuntime(c *v1.ContainerStatus) runningContainer {
	_ = "STUB: not implemented"
	return *new(runningContainer)
}

type runningContainer struct {
	ID      string
	Runtime string
}
