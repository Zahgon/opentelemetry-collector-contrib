// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver"

import (
	v1 "k8s.io/api/core/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
)

// convertNodeToEndpoint converts a node instance into a k8s.node observer.Endpoint. It will determine the
// Target by the first address match of InternalIP, InternalDNS, HostName, ExternalIP, and ExternalDNS in that
// order.
func convertNodeToEndpoint(idNamespace string, node *v1.Node) observer.Endpoint {
	_ = "STUB: not implemented"
	return *new(observer.Endpoint)
}
