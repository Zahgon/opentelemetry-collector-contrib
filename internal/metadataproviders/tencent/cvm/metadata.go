// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cvm // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/tencent/cvm"

import (
	"context"
	"net/http"
)

const (
	metadataBaseURL = "http://metadata.tencentyun.com/latest/meta-data/"
	maxResponseSize = 1 << 20 // 1 MB limit for metadata responses
)

// Metadata represents Tencent Cloud CVM instance metadata.
type Metadata struct {
	InstanceName string
	ImageID      string
	InstanceID   string
	InstanceType string
	AppID        string
	RegionID     string
	ZoneID       string
}

// Provider is the interface for retrieving Tencent Cloud CVM metadata.
type Provider interface {
	// Metadata retrieves all CVM instance metadata.
	Metadata(ctx context.Context) (*Metadata, error)
}

type metadataClient struct {
	client *http.Client
}

var _ Provider = (*metadataClient)(nil)

// NewProvider returns a new Tencent Cloud CVM metadata provider.
func NewProvider() Provider { _ = "STUB: not implemented"; return *new(Provider) }

// getMetadata retrieves a single metadata value from the given path.
func (c *metadataClient) getMetadata(ctx context.Context, path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Metadata retrieves all CVM instance metadata.
func (c *metadataClient) Metadata(ctx context.Context) (*Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
