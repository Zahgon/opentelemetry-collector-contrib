// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !aix

package datadogextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/datadogextension"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/datadogextension/internal/httpserver"
	datadogconfig "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/config"
)

var _ component.Config = (*Config)(nil)

// Config contains the information necessary for enabling the Datadog Extension.
type Config struct {
	confighttp.ClientConfig `mapstructure:",squash"`
	// Define the site and API key (and whether to fail on invalid API key) in API.
	API datadogconfig.APIConfig `mapstructure:"api"`
	// If Hostname is empty extension will use available system APIs and cloud provider endpoints.
	Hostname string `mapstructure:"hostname"`
	// HTTPConfig is v2 config for the http metadata service.
	HTTPConfig *httpserver.Config `mapstructure:"http"`
	// DeploymentType indicates the type of deployment (gateway, daemonset, or unknown).
	// Defaults to "unknown" if not set.
	DeploymentType string `mapstructure:"deployment_type"`
	// InstallationMethod indicates how the collector was installed.
	// Valid values: "", "kubernetes", "bare-metal", "docker", "ecs-fargate", "eks-fargate".
	// Defaults to "" (unset) if not configured.
	InstallationMethod string `mapstructure:"installation_method"`
	// GatewayService is the k8s Service fronting the gateway collector pods.
	// Set by gateway collectors. Format: "service" or "namespace/service".
	// Together with cluster_name, it forms the join key for fleet topology queries.
	GatewayService string `mapstructure:"gateway_service"`
	// GatewayDestination is the k8s Service that this collector forwards telemetry to.
	// Set by agent/daemonset collectors. Format: "service" or "namespace/service".
	// Must match gateway_service on the receiving gateway collector.
	GatewayDestination string `mapstructure:"gateway_destination"`
}

// Validate ensures that the configuration is valid.
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// Validate deployment_type if set

// Set default if not provided

// Validate installation_method if set
