// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package consul // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/consul"

import (
	"context"

	"github.com/hashicorp/consul/api"
)

type Provider interface {
	Metadata(context.Context) (*Metadata, error)
}

type consulMetadataImpl struct {
	consulClient  *api.Client
	allowedLabels map[string]any
}

type Metadata struct {
	NodeID       string
	Hostname     string
	Datacenter   string
	HostMetadata map[string]string
}

func NewProvider(client *api.Client, allowedLabels map[string]any) Provider {
	_ = "STUB: not implemented"
	return *new(Provider)
}

func (d *consulMetadataImpl) Metadata(_ context.Context) (*Metadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
