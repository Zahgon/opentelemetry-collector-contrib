// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package classic // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/ibmcloud/classic"

import (
	"context"
	"net/http"
)

const (
	defaultEndpoint = "https://api.service.softlayer.com/rest/v3.1/SoftLayer_Resource_Metadata"
	maxResponseSize = 1 << 20 // 1 MB limit for metadata responses
)

// Provider gathers IBM Cloud Classic (SoftLayer) instance metadata.
type Provider interface {
	InstanceMetadata(ctx context.Context) (*InstanceMetadata, error)
}

// InstanceMetadata represents the metadata collected from the SoftLayer Resource Metadata API.
type InstanceMetadata struct {
	ID               string
	Hostname         string
	Datacenter       string
	AccountID        string
	GlobalIdentifier string
}

type metadataClient struct {
	endpoint string
	client   *http.Client
}

var _ Provider = (*metadataClient)(nil)

// NewProvider returns a new IBM Cloud Classic metadata provider
// that queries the SoftLayer Resource Metadata API.
func NewProvider() Provider { _ = "STUB: not implemented"; return *new(Provider) }

// newProvider is the internal constructor that accepts an arbitrary endpoint.
// It is used by unit tests to point at httptest servers.
func newProvider(endpoint string) Provider { _ = "STUB: not implemented"; return *new(Provider) }

// InstanceMetadata retrieves instance metadata from the SoftLayer Resource Metadata API.
// It makes parallel GET requests for each metadata field using the plain-text (.txt)
// format, which returns unquoted values directly.
func (c *metadataClient) InstanceMetadata(ctx context.Context) (*InstanceMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fetchText makes a GET request to the given method path and returns the
// response body as a trimmed string. The SoftLayer metadata API .txt format
// returns plain-text values without JSON encoding.
func fetchText(ctx context.Context, c *metadataClient, method string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
