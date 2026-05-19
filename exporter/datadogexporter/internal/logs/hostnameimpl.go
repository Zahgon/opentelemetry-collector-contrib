// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/logs"

import (
	"context"

	"github.com/DataDog/datadog-agent/comp/core/hostname/hostnameinterface"
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes/source"
)

type service struct {
	provider source.Provider
}

var _ hostnameinterface.Component = (*service)(nil)

// Get returns the hostname.
func (hs *service) Get(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetSafe returns the hostname, or 'unknown host' if anything goes wrong.
func (hs *service) GetSafe(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// GetWithProvider returns the hostname for the Agent and the provider that was use to retrieve it.
func (hs *service) GetWithProvider(ctx context.Context) (hostnameinterface.Data, error) {
	_ = "STUB: not implemented"
	return *new(hostnameinterface.Data), nil
}

// NewHostnameService creates a new instance of the component hostname
func NewHostnameService(provider source.Provider) hostnameinterface.Component {
	_ = "STUB: not implemented"
	return *new(hostnameinterface.Component)
}
