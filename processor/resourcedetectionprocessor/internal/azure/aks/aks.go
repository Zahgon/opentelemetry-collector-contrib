// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package aks // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/azure/aks"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/azure"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/azure/aks/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "aks"

	// Environment variable that is set when running on Kubernetes
	kubernetesServiceHostEnvVar = "KUBERNETES_SERVICE_HOST"
)

type Detector struct {
	provider           azure.Provider
	resourceAttributes metadata.ResourceAttributesConfig
}

// NewDetector creates a new AKS detector
func NewDetector(_ processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

func (d *Detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

// If we can't get a response from the metadata endpoint, we're not running in Azure

func onK8s() bool { _ = "STUB: not implemented"; return false }

// parseClusterName parses the cluster name from the infrastructure
// resource group name. AKS IMDS returns the resource group name in
// the following formats:
//
// 1. Generated group: MC_<resource group>_<cluster name>_<location>
//   - Example:
//   - Resource group: my-resource-group
//   - Cluster name:   my-cluster
//   - Location:       eastus
//   - Generated name: MC_my-resource-group_my-cluster_eastus
//
// 2. Custom group: custom-infra-resource-group-name
//
// When using the generated infrastructure resource group, the resource
// group will include the cluster name. If the cluster's resource group
// or cluster name contains underscores, parsing will fall back on the
// unparsed infrastructure resource group name.
//
// When using a custom infrastructure resource group, the resource group name
// does not contain the cluster name. The custom infrastructure resource group
// name is returned instead.
//
// It is safe to use the infrastructure resource group name as a unique identifier
// because Azure will not allow the user to create multiple AKS clusters with the same
// infrastructure resource group name.
func parseClusterName(resourceGroup string) string {
	_ = "STUB: not implemented"
	// Code inspired by https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/exporter/datadogexporter/internal/hostmetadata/internal/azure/provider.go#L36
	return ""
}
