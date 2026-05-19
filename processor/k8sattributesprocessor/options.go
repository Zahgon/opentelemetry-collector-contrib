// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sattributesprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor"

import (
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor/internal/kube"
)

const (
	filterOPEquals       = "equals"
	filterOPNotEquals    = "not-equals"
	filterOPExists       = "exists"
	filterOPDoesNotExist = "does-not-exist"
	containerImageTag    = "container.image.tag"
)

// option represents a configuration option that can be passes.
// to the k8s-tagger
type option func(*kubernetesprocessor) error

// withAPIConfig provides k8s API related configuration to the processor.
// It defaults the authentication method to in-cluster auth using service accounts.
func withAPIConfig(cfg k8sconfig.APIConfig) option { _ = "STUB: not implemented"; return *new(option) }

// withPassthrough enables passthrough mode. In passthrough mode, the processor
// only detects and tags the pod IP and does not invoke any k8s APIs.
func withPassthrough() option { _ = "STUB: not implemented"; return *new(option) }

// enabledAttributes returns the list of resource attributes enabled by default.
func enabledAttributes() (attributes []string) { _ = "STUB: not implemented"; return nil }

// withExtractMetadata allows specifying options to control extraction of pod metadata.
// If no fields explicitly provided, the defaults are pulled from metadata.yaml.
func withExtractMetadata(fields ...string) option { _ = "STUB: not implemented"; return *new(option) }

func withOtelAnnotations(enabled bool) option { _ = "STUB: not implemented"; return *new(option) }

func withDeploymentNameFromReplicaSet(enabled bool) option {
	_ = "STUB: not implemented"
	return *new(option)
}

// withExtractLabels allows specifying options to control extraction of pod labels.
func withExtractLabels(labels ...FieldExtractConfig) option {
	_ = "STUB: not implemented"
	return *new(option)
}

// withExtractAnnotations allows specifying options to control extraction of pod annotations tags.
func withExtractAnnotations(annotations ...FieldExtractConfig) option {
	_ = "STUB: not implemented"
	return *new(option)
}

func extractFieldRules(fields ...FieldExtractConfig) ([]kube.FieldExtractionRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// withFilterNode allows specifying options to control filtering pods by a node/host.
func withFilterNode(node, nodeFromEnvVar string) option {
	_ = "STUB: not implemented"
	return *new(option)
}

// withFilterNamespace allows specifying options to control filtering pods by a namespace.
func withFilterNamespace(ns string) option { _ = "STUB: not implemented"; return *new(option) }

// withFilterLabels allows specifying options to control filtering pods by pod labels.
func withFilterLabels(filters ...FieldFilterConfig) option {
	_ = "STUB: not implemented"
	return *new(option)
}

// withFilterFields allows specifying options to control filtering pods by pod fields.
func withFilterFields(filters ...FieldFilterConfig) option {
	_ = "STUB: not implemented"
	return *new(option)
}

// withExtractPodAssociations allows specifying options to associate pod metadata with incoming resource
func withExtractPodAssociations(podAssociations ...PodAssociationConfig) option {
	_ = "STUB: not implemented"
	return *new(option)
}

// withExcludes allows specifying pods to exclude
func withExcludes(podExclude ExcludeConfig) option { _ = "STUB: not implemented"; return *new(option) }

// withWaitForMetadata allows specifying whether to wait for pod metadata to be synced.
func withWaitForMetadata(wait bool) option { _ = "STUB: not implemented"; return *new(option) }

// withWaitForMetadataTimeout allows specifying the timeout for waiting for pod metadata to be synced.
func withWaitForMetadataTimeout(timeout time.Duration) option {
	_ = "STUB: not implemented"
	return *new(option)
}

// withWatchSyncPeriod allows specifying the resync period for informer.
func withWatchSyncPeriod(duration time.Duration) option {
	_ = "STUB: not implemented"
	return *new(option)
}
