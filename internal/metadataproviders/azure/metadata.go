// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// This file contains code based on the Azure IMDS samples, https://github.com/microsoft/azureimds
// under the Apache License 2.0

package azure // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/azure"

import (
	"context"
	"net/http"
)

const (
	// Azure IMDS compute endpoint, see https://aka.ms/azureimds
	metadataEndpoint = "http://169.254.169.254/metadata/instance/compute"
)

// Provider gets metadata from the Azure IMDS.
type Provider interface {
	Metadata(context.Context) (*ComputeMetadata, error)
}

type azureProviderImpl struct {
	endpoint string
	client   *http.Client
}

// NewProvider creates a new metadata provider
func NewProvider() Provider { _ = "STUB: not implemented"; return *new(Provider) }

type ComputeTagsListMetadata struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type OSProfile struct {
	ComputerName string `json:"computerName"`
}

// ComputeMetadata is the Azure IMDS compute metadata response format
type ComputeMetadata struct {
	Location          string                    `json:"location"`
	Name              string                    `json:"name"`
	VMID              string                    `json:"vmID"`
	VMSize            string                    `json:"vmSize"`
	SubscriptionID    string                    `json:"subscriptionID"`
	ResourceGroupName string                    `json:"resourceGroupName"`
	VMScaleSetName    string                    `json:"vmScaleSetName"`
	AvailabilityZone  string                    `json:"zone"`
	OSProfile         OSProfile                 `json:"osProfile"`
	TagsList          []ComputeTagsListMetadata `json:"tagsList"`
}

// Metadata queries a given endpoint and parses the output to the Azure IMDS format
func (p *azureProviderImpl) Metadata(ctx context.Context) (*ComputeMetadata, error) {
	_ = "STUB: not implemented"

	// API version used
	return nil, nil
}

// format used

// As per the Azure IMDS documentation, the Metadata header must be set to "true" (note lowercase).

//lint:ignore ST1005 Azure is a capitalized proper noun here
