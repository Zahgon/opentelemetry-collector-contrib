// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package vultr // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/vultr"

import (
	"context"
	"net/http"
)

const (
	// Vultr IMDS compute endpoint, see https://www.vultr.com/metadata/
	metadataEndpoint = "http://169.254.169.254/v1.json"
)

// Provider gets metadata from the Vultr metadata service.
type Provider interface {
	Metadata(context.Context) (*Metadata, error)
}

type vultrProviderImpl struct {
	endpoint string
	client   *http.Client
}

// NewProvider creates a new Vultr metadata provider.
func NewProvider() Provider { _ = "STUB: not implemented"; return *new(Provider) }

type Region struct {
	RegionCode string `json:"regioncode"`
}

type Metadata struct {
	Hostname     string `json:"hostname"`
	InstanceID   string `json:"instanceid"`
	InstanceV2ID string `json:"instance-v2-id"`
	Region       Region `json:"region"`
}

// Metadata fetches and decodes Vultr instance metadata.
func (p *vultrProviderImpl) Metadata(ctx context.Context) (*Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
