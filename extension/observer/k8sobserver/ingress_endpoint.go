// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver"

import (
	v1 "k8s.io/api/networking/v1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
)

// convertIngressToEndpoints converts a ingress instance into a slice of endpoints. The endpoints
// include an endpoint for each path that is mapped to an ingress.
func convertIngressToEndpoints(idNamespace string, ingress *v1.Ingress) []observer.Endpoint {
	_ = "STUB: not implemented"
	return nil
}

// Loop through every ingress rule to get every defined path.

// Create endpoint for each ingress rule.

// getTLSHosts return a list of tls hosts for an ingress resource.
func getTLSHosts(i *v1.Ingress) []string { _ = "STUB: not implemented"; return nil }

// matchesHostPattern returns true if the host matches the host pattern or wildcard pattern.
func matchesHostPattern(pattern, host string) bool {
	_ = "STUB: not implemented"
	// if host match the pattern (host pattern).
	return false
}

// if string does not contains any dot, don't do the next part as it's for wildcard pattern.

// If the first part of the pattern is not a wildcard pattern.

// If host and pattern without wildcard part does not match.

// getScheme return the scheme of an ingress host based on tls configuration.
func getScheme(host string, tlsHosts []string) string { _ = "STUB: not implemented"; return "" }
