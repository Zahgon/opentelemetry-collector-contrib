// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pod // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/pod"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/cache"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

const (
	// Keys for pod metadata and entity attributes. These are NOT used by resource attributes.
	podCreationTime = "pod.creation_timestamp"
	podPhase        = "k8s.pod.phase"
	podStatusReason = "k8s.pod.status_reason"
)

// Transform transforms the pod to remove the fields that we don't use to reduce RAM utilization.
// IMPORTANT: Make sure to update this function before using new pod fields.
func Transform(pod *corev1.Pod) *corev1.Pod { _ = "STUB: not implemented"; return nil }

func RecordMetrics(logger *zap.Logger, mb *metadata.MetricsBuilder, pod *corev1.Pod, ts pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func reasonToInt(reason string) int32 { _ = "STUB: not implemented"; return 0 }

func phaseToInt(phase corev1.PodPhase) int32 { _ = "STUB: not implemented"; return 0 }

// GetMetadata returns all metadata associated with the pod.
func GetMetadata(pod *corev1.Pod, mc *metadata.Store, logger *zap.Logger) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}

// defer syncing replicaset and job workload metadata.

// collectPodJobProperties checks if pod owner of type Job is cached. Check owners reference
// on Job to see if it was created by a CronJob. Sync metadata accordingly.
func collectPodJobProperties(pod *corev1.Pod, jobStores map[string]cache.Store, logger *zap.Logger) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// in practice the conversion should not fail, but checking just to be safe

// collectPodReplicaSetProperties checks if pod owner of type ReplicaSet is cached. Check owners reference
// on ReplicaSet to see if it was created by a Deployment. Sync metadata accordingly.
func collectPodReplicaSetProperties(pod *corev1.Pod, replicaSetStores map[string]cache.Store, logger *zap.Logger) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// in practice the conversion should not fail, but checking just to be safe

func logDebug(ref *v1.OwnerReference, podUID types.UID, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

func logError(err error, ref *v1.OwnerReference, podUID types.UID, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// getWorkloadProperties returns workload metadata for provided owner reference.
func getWorkloadProperties(ref *v1.OwnerReference, labelKey string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func getPodContainerProperties(pod *corev1.Pod, logger *zap.Logger) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}

// getIDForCache returns keys to lookup resources from the cache exposed
// by shared informers.
func getIDForCache(namespace, resourceName string) string { _ = "STUB: not implemented"; return "" }

// getObjectFromStore retrieves the requested object from the given stores.
// first, the object is attempted to be retrieved from the store for all namespaces,
// and if it is not found there, the namespace-specific store is used
func getObjectFromStore(namespace, objName string, stores map[string]cache.Store) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// findOwnerWithKind returns the OwnerReference of the matching kind from
// the provided list of owner references.
func findOwnerWithKind(ors []v1.OwnerReference, kind string) *v1.OwnerReference {
	_ = "STUB: not implemented"
	return nil
}
