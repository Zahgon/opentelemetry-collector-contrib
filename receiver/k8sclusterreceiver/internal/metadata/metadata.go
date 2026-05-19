// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	metadataPkg "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
)

// KubernetesMetadata associates a resource to a set of properties.
type KubernetesMetadata struct {
	// The type of the entity, e.g. k8s.pod
	EntityType string
	// resourceIDKey is the label key of UID label for the resource.
	ResourceIDKey string
	// resourceID is the Kubernetes UID of the resource. In case of
	// containers, this value is the container id.
	ResourceID metadataPkg.ResourceID
	// metadata is a set of key-value pairs that describe a resource.
	Metadata map[string]string
}

func TransformObjectMeta(om v1.ObjectMeta) v1.ObjectMeta {
	_ = "STUB: not implemented"
	return *new(v1.ObjectMeta)
}

// GetGenericMetadata is responsible for collecting metadata from K8s resources that
// live on v1.ObjectMeta.
func GetGenericMetadata(om *v1.ObjectMeta, resourceType string) *KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}

func GetOTelUIDFromKind(kind string) string { _ = "STUB: not implemented"; return "" }

func GetOTelNameFromKind(kind string) string { _ = "STUB: not implemented"; return "" }

func getOTelEntityTypeFromKind(kind string) string { _ = "STUB: not implemented"; return "" }

// mergeKubernetesMetadataMaps merges maps of string (resource id) to
// KubernetesMetadata into a single map.
func MergeKubernetesMetadataMaps(maps ...map[metadataPkg.ResourceID]*KubernetesMetadata) map[metadataPkg.ResourceID]*KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}

// GetMetadataUpdate processes metadata updates and returns
// a map of a delta of metadata mapped to each resource.
func GetMetadataUpdate(oldMetadata, newMetadata map[metadataPkg.ResourceID]*KubernetesMetadata) []*metadataPkg.MetadataUpdate {
	_ = "STUB: not implemented"
	return nil
}

// if an object with the same id has a previous revision, take a delta
// of the metadata.

// In case there are resources in the current revision, that was not in the
// previous revision, collect metadata to be added

// if an id is seen for the first time, all metadata need to be added.

// getMetadataDelta returns MetadataDelta between two sets for properties.
// If the delta between old (oldProps) and new (newProps) revisions of a
// resource end up being empty, nil is returned.
func getMetadataDelta(oldProps, newProps map[string]string) *metadataPkg.MetadataDelta {
	_ = "STUB: not implemented"
	return nil
}

// If metadata exist in the previous revision as well, collect if
// the new values are different. Otherwise, the property is a new value
// and has to be added.

// Properties that don't exist in the latest revision should be removed
