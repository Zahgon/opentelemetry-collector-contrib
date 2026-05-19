// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package azure contains the Azure hostname provider
package azure // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata/internal/azure"

import (
	"context"

	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes/source"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata/provider"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/azure"
)

var (
	_ source.Provider              = (*Provider)(nil)
	_ provider.ClusterNameProvider = (*Provider)(nil)
)

type Provider struct {
	detector azure.Provider
}

// Hostname returns the Azure cloud integration hostname.
func (p *Provider) Source(ctx context.Context) (source.Source, error) {
	_ = "STUB: not implemented"
	return *new(source.Source), nil
}

// ClusterName gets the AKS cluster name from the resource group name.
func (p *Provider) ClusterName(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Code comes from https://github.com/DataDog/datadog-agent/blob/1b4afdd6a/pkg/util/cloudproviders/azure/azure.go#L72
// It expects the resource group name to have the format (MC|mc)_resource-group_cluster-name_zone.

// NewProvider creates a new Azure hostname provider.
func NewProvider() *Provider { _ = "STUB: not implemented"; return nil }
