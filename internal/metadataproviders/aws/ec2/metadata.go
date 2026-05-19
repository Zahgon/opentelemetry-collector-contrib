// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ec2 // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/aws/ec2"

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
)

type Provider interface {
	Get(ctx context.Context) (imds.InstanceIdentityDocument, error)
	Hostname(ctx context.Context) (string, error)
	InstanceID(ctx context.Context) (string, error)
	Tags(ctx context.Context) ([]string, error)
	Tag(ctx context.Context, key string) (string, error)
}

type metadataClient struct {
	client *imds.Client
}

var _ Provider = (*metadataClient)(nil)

func NewProvider(cfg aws.Config) Provider { _ = "STUB: not implemented"; return *new(Provider) }

func (c *metadataClient) getMetadata(ctx context.Context, path string) (string, error) {
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

func (c *metadataClient) Get(ctx context.Context) (imds.InstanceIdentityDocument, error) {
	_ = "STUB: not implemented"
	return *new(imds.InstanceIdentityDocument), nil
}

func (c *metadataClient) Tags(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *metadataClient) Tag(ctx context.Context, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
