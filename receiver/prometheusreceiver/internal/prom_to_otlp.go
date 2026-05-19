// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver/internal"

import (
	"github.com/prometheus/prometheus/model/labels"
	"go.opentelemetry.io/collector/pdata/pcommon"
	conventions "go.opentelemetry.io/otel/semconv/v1.40.0"
)

// isDiscernibleHost checks if a host can be used as a value for the 'host.name' key.
// localhost-like hosts and unspecified (0.0.0.0) hosts are not discernible.
func isDiscernibleHost(host string) bool { _ = "STUB: not implemented"; return false }

// An IP is discernible if
//  - it's not local (e.g. belongs to 127.0.0.0/8 or ::1/128) and
//  - it's not unspecified (e.g. the 0.0.0.0 address).

// not an IP, not 'localhost', assume it is discernible.

// CreateResource creates the resource data added to OTLP payloads.
func CreateResource(job, instance string, serviceDiscoveryLabels labels.Labels) pcommon.Resource {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource)
}

// kubernetesDiscoveryToResourceAttributes maps from metadata labels discovered
// through the kubernetes implementation of service discovery to opentelemetry
// resource attribute keys.
var kubernetesDiscoveryToResourceAttributes = map[string]string{
	"__meta_kubernetes_pod_name":           string(conventions.K8SPodNameKey),
	"__meta_kubernetes_pod_uid":            string(conventions.K8SPodUIDKey),
	"__meta_kubernetes_pod_container_name": string(conventions.K8SContainerNameKey),
	"__meta_kubernetes_namespace":          string(conventions.K8SNamespaceNameKey),
	// Only one of the node name service discovery labels will be present
	"__meta_kubernetes_pod_node_name":      string(conventions.K8SNodeNameKey),
	"__meta_kubernetes_node_name":          string(conventions.K8SNodeNameKey),
	"__meta_kubernetes_endpoint_node_name": string(conventions.K8SNodeNameKey),
}

// addKubernetesResource adds resource information detected by prometheus'
// kubernetes service discovery.
func addKubernetesResource(attrs pcommon.Map, serviceDiscoveryLabels labels.Labels) {
	_ = "STUB: not implemented"
	return
}
