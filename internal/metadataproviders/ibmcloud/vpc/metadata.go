// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package vpc // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/ibmcloud/vpc"

import (
	"context"
	"net/http"
	"sync"
	"time"
)

const (
	metadataHost        = "api.metadata.cloud.ibm.com"
	tokenPath           = "/identity/v1/token"
	instancePath        = "/metadata/v1/instance"
	apiVersion          = "2026-01-30"
	metadataFlavorKey   = "Metadata-Flavor"
	metadataFlavorValue = "ibm"
	defaultTokenTTL     = 300 // seconds
	tokenRefreshBuffer  = 30 * time.Second
	maxResponseSize     = 1 << 20 // 1 MB limit for metadata responses
)

// Provider gathers IBM Cloud VPC instance metadata.
type Provider interface {
	InstanceMetadata(ctx context.Context) (*InstanceMetadata, error)
}

// InstanceMetadata represents the response from the VPC IMDS instance endpoint.
type InstanceMetadata struct {
	ID      string `json:"id"`
	CRN     string `json:"crn"`
	Name    string `json:"name"`
	Profile struct {
		Name string `json:"name"`
	} `json:"profile"`
	Zone struct {
		Name string `json:"name"`
	} `json:"zone"`
	VPC struct {
		ID   string `json:"id"`
		CRN  string `json:"crn"`
		Name string `json:"name"`
	} `json:"vpc"`
	ResourceGroup struct {
		ID string `json:"id"`
	} `json:"resource_group"`
	Image struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"image"`
}

// tokenResponse represents the IMDS token response.
type tokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type metadataClient struct {
	endpoint string
	client   *http.Client

	tokenMu     sync.Mutex
	token       string
	tokenExpiry time.Time
}

var _ Provider = (*metadataClient)(nil)

// NewProvider returns a new IBM Cloud VPC metadata provider.
// The protocol parameter selects the scheme ("http" or "https").
// Both schemes target the same host (api.metadata.cloud.ibm.com).
func NewProvider(protocol string) Provider { _ = "STUB: not implemented"; return *new(Provider) }

// newProvider is the internal constructor that accepts an arbitrary endpoint.
// It is used by unit tests to point at httptest servers.
func newProvider(endpoint string) Provider { _ = "STUB: not implemented"; return *new(Provider) }

// InstanceMetadata retrieves instance metadata from the IBM Cloud VPC IMDS.
func (c *metadataClient) InstanceMetadata(ctx context.Context) (*InstanceMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getToken retrieves an instance identity token, refreshing if necessary.
func (c *metadataClient) getToken(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Return cached token if still valid

// getInstanceMetadata retrieves instance metadata using the given token.
func (c *metadataClient) getInstanceMetadata(ctx context.Context, token string) (*InstanceMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
