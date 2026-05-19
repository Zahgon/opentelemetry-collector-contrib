// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecs // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/alibaba/ecs"

import (
	"context"
	"net/http"
	"sync"
	"time"
)

const (
	metadataBaseURL    = "http://100.100.100.200/latest/meta-data/"
	tokenURL           = "http://100.100.100.200/latest/api/token" //nolint:gosec // not a credential, metadata service URL
	tokenTTLHeader     = "X-aliyun-ecs-metadata-token-ttl-seconds" //nolint:gosec // not a credential, header name
	tokenHeader        = "X-aliyun-ecs-metadata-token"             //nolint:gosec // not a credential, header name
	defaultTokenTTL    = 21600
	tokenRefreshBuffer = 5 * time.Minute
	maxResponseSize    = 1 << 20 // 1 MB limit for metadata responses
)

// Metadata represents Alibaba Cloud ECS instance metadata.
type Metadata struct {
	Hostname       string
	ImageID        string
	InstanceID     string
	InstanceType   string
	OwnerAccountID string
	RegionID       string
	ZoneID         string
}

// Provider is the interface for retrieving Alibaba Cloud ECS metadata.
type Provider interface {
	// Metadata retrieves all ECS instance metadata.
	Metadata(ctx context.Context) (*Metadata, error)
}

type metadataClient struct {
	client *http.Client

	tokenMu     sync.Mutex
	token       string
	tokenExpiry time.Time
}

var _ Provider = (*metadataClient)(nil)

// NewProvider returns a new Alibaba Cloud ECS metadata provider.
func NewProvider() Provider { _ = "STUB: not implemented"; return *new(Provider) }

// getToken retrieves a metadata token, refreshing if necessary.
func (c *metadataClient) getToken(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Return cached token if still valid

// getMetadata retrieves a single metadata value from the given path.
func (c *metadataClient) getMetadata(ctx context.Context, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Metadata retrieves all ECS instance metadata.
func (c *metadataClient) Metadata(ctx context.Context) (*Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
