// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oraclecloud // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/oraclecloud"

import (
	"context"
	"net/http"
)

// OracleCloud IMDS compute endpoint
var metadataEndpoint = "http://169.254.169.254/opc/v2/instance/"

// Provider gets metadata from the OracleCloud IMDS.
type Provider interface {
	Metadata(context.Context) (*ComputeMetadata, error)
}

// IsRunningOnOracleCloud performs a fast probe to the OCI metadata endpoint with a short timeout.
// Returns true if the endpoint responds with HTTP 200 OK, false otherwise.
func IsRunningOnOracleCloud(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

// IsRunningOnOracleCloudFunc can be overridden in tests to simulate probe presence/absence.
var IsRunningOnOracleCloudFunc = IsRunningOnOracleCloud

type oraclecloudProviderImpl struct {
	endpoint string
	client   *http.Client
}

// NewProvider creates a new metadata provider
func NewProvider() Provider { _ = "STUB: not implemented"; return *new(Provider) }

type ComputeTagsListMetadata struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// ComputeMetadata is the OracleCloud IMDS compute metadata response format
type ComputeMetadata struct {
	HostID             string `json:"id"`
	HostDisplayName    string `json:"displayName"`
	HostType           string `json:"shape"`
	RegionID           string `json:"canonicalRegionName"`
	AvailabilityDomain string `json:"availabilityDomain"`

	Metadata InstanceMetadata `json:"metadata"`
}

type InstanceMetadata struct {
	OKEClusterDisplayName string `json:"oke-cluster-display-name"`
	Realm                 string `json:"realm"`
}

// Metadata queries a given endpoint and parses the output to the OracleCloud IMDS format
func (p *oraclecloudProviderImpl) Metadata(ctx context.Context) (*ComputeMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
