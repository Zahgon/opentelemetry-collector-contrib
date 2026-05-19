// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureauthextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/azureauthextension"

import (
	"errors"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configoptional"
)

var (
	validOptions = []string{
		"use_default",
		"managed_identity",
		"workload_identity",
		"service_principal",
	}

	errEmptyTenantID           = errors.New(`empty "tenant_id" field`)
	errEmptyClientID           = errors.New(`empty "client_id" field`)
	errEmptyClientCredential   = errors.New(`both "client_secret" and "client_certificate_path" fields are empty`)
	errEmptyFederatedTokenFile = errors.New(`empty "federated_token_file" field`)
	errEmptyAuthentication     = fmt.Errorf("authentication configuration is empty, please choose one of %s", validOptions)
	errMutuallyExclusiveAuth   = errors.New(`"client_secret" and "client_certificate_path" are mutually exclusive`)
	errEmptyServerIssuerURL    = errors.New(`empty "issuer_url" field`)
	errEmptyServerAudience     = errors.New(`empty "audience" field`)
)

type Config struct {
	Managed          configoptional.Optional[ManagedIdentity]  `mapstructure:"managed_identity"`
	Workload         configoptional.Optional[WorkloadIdentity] `mapstructure:"workload_identity"`
	ServicePrincipal configoptional.Optional[ServicePrincipal] `mapstructure:"service_principal"`
	Server           configoptional.Optional[Server]           `mapstructure:"server"`
	UseDefault       bool                                      `mapstructure:"use_default"`
	Scopes           []string                                  `mapstructure:"scopes"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type Server struct {
	IssuerURL string `mapstructure:"issuer_url"`
	Audience  string `mapstructure:"audience"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type ManagedIdentity struct {
	// if left empty, then it is system managed
	ClientID string `mapstructure:"client_id"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type WorkloadIdentity struct {
	ClientID           string `mapstructure:"client_id"`
	TenantID           string `mapstructure:"tenant_id"`
	FederatedTokenFile string `mapstructure:"federated_token_file"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type ServicePrincipal struct {
	TenantID              string `mapstructure:"tenant_id"`
	ClientID              string `mapstructure:"client_id"`
	ClientSecret          string `mapstructure:"client_secret"`
	ClientCertificatePath string `mapstructure:"client_certificate_path"`
	// prevent unkeyed literal initialization
	_ struct{}
}

var _ component.Config = (*Config)(nil)

func (*ManagedIdentity) Validate() error { _ = "STUB: not implemented"; return nil }

func (cfg *WorkloadIdentity) Validate() error { _ = "STUB: not implemented"; return nil }

func (cfg *ServicePrincipal) Validate() error { _ = "STUB: not implemented"; return nil }

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }
