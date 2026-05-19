// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kubelet // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/kubelet"

import (
	"regexp"

	conventions "go.opentelemetry.io/otel/semconv/v1.40.0"
	v1 "k8s.io/api/core/v1"
	stats "k8s.io/kubelet/pkg/apis/stats/v1alpha1"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/metadata"
)

type MetadataLabel string

// Values for MetadataLabel enum.
const MetadataLabelVolumeType MetadataLabel = labelVolumeType

var (
	MetadataLabelContainerID MetadataLabel = MetadataLabel(conventions.ContainerIDKey)
	supportedLabels                        = map[MetadataLabel]bool{
		MetadataLabelContainerID: true,
		MetadataLabelVolumeType:  true,
	}
)

// ValidateMetadataLabelsConfig validates that provided list of metadata labels is supported
func ValidateMetadataLabelsConfig(labels []MetadataLabel) error {
	_ = "STUB: not implemented"
	return nil
}

type Metadata struct {
	Labels                    map[MetadataLabel]bool
	PodsMetadata              *v1.PodList
	DetailedPVCResourceSetter func(rb *metadata.ResourceBuilder, volCacheID, volumeClaim, namespace string) error
	podResources              map[string]resources
	containerResources        map[string]resources
	nodeInfo                  NodeInfo
}

type resources struct {
	cpuRequest    float64
	cpuLimit      float64
	memoryRequest int64
	memoryLimit   int64
}

type NodeInfo struct {
	Name string
	// node's CPU capacity in cores
	CPUCapacity float64
	// node's Memory capacity in bytes
	MemoryCapacity float64
}

func getContainerResources(r *v1.ResourceRequirements) resources {
	_ = "STUB: not implemented"
	return *new(resources)
}

func NewMetadata(labels []MetadataLabel, podsMetadata *v1.PodList, nodeInfo NodeInfo,
	detailedPVCResourceSetter func(rb *metadata.ResourceBuilder, volCacheID, volumeClaim, namespace string) error,
) Metadata {
	_ = "STUB: not implemented"
	return *new(Metadata)
}

func getLabelsMap(metadataLabels []MetadataLabel) map[MetadataLabel]bool {
	_ = "STUB: not implemented"
	return nil
}

// getExtraResources gets extra resources based on provided metadata label.
func (m *Metadata) setExtraResources(rb *metadata.ResourceBuilder, podRef stats.PodReference,
	extraMetadataLabel MetadataLabel, extraMetadataFrom string,
) error {
	_ = "STUB: not implemented"
	// Ensure MetadataLabel exists before proceeding.
	return nil
}

// Cannot proceed, if metadata is unavailable.

// Get more labels from PersistentVolumeClaim volume type.

// getContainerID retrieves container id from metadata for given pod UID and container name,
// returns an error if no container found in the metadata that matches the requirements
// or if the apiServer returned a newly created container with empty containerID.
func (m *Metadata) getContainerID(podUID, containerName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

//nolint:gocritic // appendAssign: append result not assigned to the same slice

var containerSchemeRegexp = regexp.MustCompile(`^[\w_-]+://`)

// stripContainerID returns a pure container id without the runtime scheme://
func stripContainerID(id string) string { _ = "STUB: not implemented"; return "" }

func (m *Metadata) getPodVolume(podUID, volumeName string) (*v1.Volume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
