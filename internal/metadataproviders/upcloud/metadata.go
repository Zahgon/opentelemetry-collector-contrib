// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package upcloud // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/upcloud"

import (
	"context"
	"net/http"
)

const (
	// Upcloud IMDS compute endpoint, see https://upcloud.com/docs/guides/upcloud-metadata-service/
	metadataEndpoint = "http://169.254.169.254/metadata/v1.json"
)

// Provider gets metadata from the Upcloud metadata service.
type Provider interface {
	Metadata(context.Context) (*Metadata, error)
}

type upcloudProviderImpl struct {
	endpoint string
	client   *http.Client
}

// NewProvider creates a new Upcloud metadata provider.
func NewProvider() Provider { _ = "STUB: not implemented"; return *new(Provider) }

type Metadata struct {
	CloudName  string `json:"cloud_name"`
	Hostname   string `json:"hostname"`
	InstanceID string `json:"instance_id"`
	Region     string `json:"region"`
}

// Metadata fetches and decodes Upcloud instance metadata.
func (p *upcloudProviderImpl) Metadata(ctx context.Context) (*Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
