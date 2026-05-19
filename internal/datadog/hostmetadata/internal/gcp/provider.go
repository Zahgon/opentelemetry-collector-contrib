// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package gcp contains the GCP hostname provider
package gcp // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata/internal/gcp"

import (
	"context"

	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes/source"
	"github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata/provider"
)

var (
	_ source.Provider              = (*Provider)(nil)
	_ provider.ClusterNameProvider = (*Provider)(nil)
)

var _ gcpDetector = gcp.NewDetector()

type gcpDetector interface {
	ProjectID() (string, error)
	CloudPlatform() gcp.Platform
	GCEHostName() (string, error)
	GKEClusterName() (string, error)
}

type Provider struct {
	detector gcpDetector
}

func platformDescription(platform gcp.Platform) string { _ = "STUB: not implemented"; return "" }

// Hostname returns the GCP cloud integration hostname.
func (p *Provider) Source(context.Context) (source.Source, error) {
	_ = "STUB: not implemented"
	return *new(source.Source), nil
}

// Use the same logic as in the metadata from attributes logic.

func (p *Provider) ClusterName(_ context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NewProvider creates a new GCP hostname provider.
func NewProvider() *Provider { _ = "STUB: not implemented"; return nil }
