// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver"

import (
	v1 "k8s.io/api/core/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
)

// convertServiceToEndpoints converts a service instance into a slice of endpoints. The endpoints
// include the service itself only.
func convertServiceToEndpoints(idNamespace string, service *v1.Service) []observer.Endpoint {
	_ = "STUB: not implemented"
	return nil
}

func generateServiceTarget(service *observer.K8sService) string {
	_ = "STUB: not implemented"
	return ""
}
