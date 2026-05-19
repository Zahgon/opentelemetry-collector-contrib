// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sattributesprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor/internal/kube"
)

// extractPodIds returns pod identifier for first association matching all sources
func extractPodID(ctx context.Context, attrs pcommon.Map, associations []kube.Association) kube.PodIdentifier {
	_ = "STUB: not implemented"
	// If pod association is not set
	return *new(kube.PodIdentifier)
}

// If association configured to take IP address from connection

// Extract values based on configured resource_attribute.

// If association configured by resource_attribute
// In k8s environment, host.name label set to a pod IP address.
// If the value doesn't represent an IP address, we skip it.

// If all association sources has been resolved, return result

// extractPodIds returns pod identifier for first association matching all sources
func extractPodIDNoAssociations(ctx context.Context, attrs pcommon.Map) kube.PodIdentifier {
	_ = "STUB: not implemented"
	return *new(kube.PodIdentifier)
}

// buildPodIdentifierString returns a low-cardinality string representing which sources
// were used to build the PodIdentifier, formatted as "from" or "from/name" for each
// non-empty slot, joined by "+". Actual identifier values are intentionally excluded
// to avoid unbounded metric cardinality.
// Examples: "connection", "resource_attribute/k8s.pod.ip",
// "resource_attribute/k8s.pod.uid+resource_attribute/container.id"
func buildPodIdentifierString(id kube.PodIdentifier) string { _ = "STUB: not implemented"; return "" }

func stringAttributeFromMap(attrs pcommon.Map, key string) string {
	_ = "STUB: not implemented"
	return ""
}
