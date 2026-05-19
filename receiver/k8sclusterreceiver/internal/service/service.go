// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package service // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/service"
import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

func shouldSkipAnnotation(key string) bool { _ = "STUB: not implemented"; return false }

// Transform transforms the service to remove the fields that we don't use to reduce RAM utilization.
// IMPORTANT: Make sure to update this function before using new service fields.
func Transform(service *corev1.Service) *corev1.Service { _ = "STUB: not implemented"; return nil }

// Only keep annotations that aren't excluded to save RAM.

func RecordMetrics(_ *zap.Logger, mb *metadata.MetricsBuilder, svc *corev1.Service, endpointCounts EndpointCountsByKey, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

//nolint:staticcheck

//nolint:staticcheck

//nolint:staticcheck

//nolint:staticcheck

//nolint:staticcheck

// EndpointCounts tracks the counts of endpoints in different conditions
type EndpointCounts struct {
	Ready       int
	Serving     int
	Terminating int
}

// EndpointKey represents a unique combination of address type and zone for aggregation
type EndpointKey struct {
	AddressType string
	Zone        string
}

// EndpointCountsByKey maps (addressType, zone) to endpoint counts
type EndpointCountsByKey map[EndpointKey]EndpointCounts

// GetMetadata returns entity metadata for a Service.
func GetMetadata(svc *corev1.Service) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}

// GetPodServiceTags returns a set of services associated with the pod.
func GetPodServiceTags(pod *corev1.Pod, services map[string]cache.Store) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
