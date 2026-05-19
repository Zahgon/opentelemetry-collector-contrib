// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package nova // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/openstack/nova"

import (
	"context"
	"net/http"
)

const (
	openstackMetaURL = "http://169.254.169.254/openstack/latest/meta_data.json"
	ec2MetaBaseURL   = "http://169.254.169.254/latest/meta-data/"
)

type Provider interface {
	Get(ctx context.Context) (Document, error)
	Hostname(ctx context.Context) (string, error)
	InstanceID(ctx context.Context) (string, error)
	InstanceType(ctx context.Context) (string, error)
}

type metadataClient struct {
	client *http.Client
}

var _ Provider = (*metadataClient)(nil)

// Document is a minimal representation of OpenStack's meta_data.json
type Document struct {
	AvailabilityZone string            `json:"availability_zone"`
	Hostname         string            `json:"hostname"`
	Name             string            `json:"name"`
	Meta             map[string]string `json:"meta"`
	ProjectID        string            `json:"project_id"`
	UUID             string            `json:"uuid"`
}

// NewProvider returns a new Nova metadata provider with a short timeout.
func NewProvider() Provider { _ = "STUB: not implemented"; return *new(Provider) }

func (c *metadataClient) getMetadata(ctx context.Context) (Document, error) {
	_ = "STUB: not implemented"
	return *new(Document), nil
}

func (c *metadataClient) getEc2Metadata(ctx context.Context, fullURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *metadataClient) InstanceID(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *metadataClient) Hostname(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *metadataClient) InstanceType(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *metadataClient) Get(ctx context.Context) (Document, error) {
	_ = "STUB: not implemented"
	return *new(Document), nil
}
