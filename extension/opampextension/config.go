// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package opampextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/opampextension"

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/open-telemetry/opamp-go/client"
	"github.com/open-telemetry/opamp-go/protobufs"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configtls"
	"go.uber.org/zap"
)

// Default value for HTTP client's polling interval, set to 30 seconds in
// accordance with the OpAMP spec.
const httpPollingIntervalDefault = 30 * time.Second

// Config contains the configuration for the opamp extension. Trying to mirror
// the OpAMP supervisor config for some consistency.
type Config struct {
	Server *OpAMPServer `mapstructure:"server"`

	// InstanceUID is a UUID formatted as a 36 character string in canonical
	// representation. Auto-generated on start if missing.
	InstanceUID string `mapstructure:"instance_uid"`

	// Capabilities contains options to enable a particular OpAMP capability
	Capabilities Capabilities `mapstructure:"capabilities"`

	// Agent descriptions contains options to modify the AgentDescription message
	AgentDescription AgentDescription `mapstructure:"agent_description"`

	// PPID is the process ID of the parent for the collector. If the PPID is specified,
	// the extension will continuously poll for the status of the parent process, and emit a fatal error
	// when the parent process is no longer running.
	// If unspecified, the orphan detection logic does not run.
	PPID int32 `mapstructure:"ppid"`

	// PPIDPollInterval is the time between polling for whether PPID is running.
	PPIDPollInterval time.Duration `mapstructure:"ppid_poll_interval"`
}

type AgentDescription struct {
	// NonIdentifyingAttributes are a map of key-value pairs that may be specified to provide
	// extra information about the agent to the OpAMP server.
	NonIdentifyingAttributes map[string]string `mapstructure:"non_identifying_attributes"`
	// IncludeResourceAttributes determines whether the agent should copy its resource attributes
	// to the non identifying attributes. (default: false)
	IncludeResourceAttributes bool `mapstructure:"include_resource_attributes"`
}

type Capabilities struct {
	// ReportsEffectiveConfig enables the OpAMP ReportsEffectiveConfig Capability. (default: true)
	ReportsEffectiveConfig bool `mapstructure:"reports_effective_config"`
	// ReportsHealth enables the OpAMP ReportsHealth Capability. (default: true)
	ReportsHealth bool `mapstructure:"reports_health"`
	// ReportsAvailableComponents enables the OpAMP ReportsAvailableComponents Capability (default: true)
	ReportsAvailableComponents bool `mapstructure:"reports_available_components"`
	// AcceptsRestartCommand enables the OpAMP AcceptsRestartCommand Capability (default: false)
	AcceptsRestartCommand bool `mapstructure:"accepts_restart_command"`
}

func (caps Capabilities) toAgentCapabilities() protobufs.AgentCapabilities {
	_ = "STUB: not implemented"
	// All Agents MUST report status.
	return *new(protobufs.AgentCapabilities)
}

type commonFields struct {
	Endpoint string                         `mapstructure:"endpoint"`
	TLS      configtls.ClientConfig         `mapstructure:"tls,omitempty"`
	Headers  map[string]configopaque.String `mapstructure:"headers,omitempty"`
	Auth     component.ID                   `mapstructure:"auth,omitempty"`
}

func (c *commonFields) Scheme() string { _ = "STUB: not implemented"; return "" }

func (c *commonFields) Validate() error { _ = "STUB: not implemented"; return nil }

type httpFields struct {
	commonFields `mapstructure:",squash"`

	PollingInterval time.Duration `mapstructure:"polling_interval"`
}

func (h *httpFields) Validate() error { _ = "STUB: not implemented"; return nil }

// OpAMPServer contains the OpAMP transport configuration.
type OpAMPServer struct {
	WS   *commonFields `mapstructure:"ws,omitempty"`
	HTTP *httpFields   `mapstructure:"http,omitempty"`
}

func (s OpAMPServer) GetClient(logger *zap.Logger) client.OpAMPClient {
	_ = "STUB: not implemented"
	return *new(client.OpAMPClient)
}

func (s OpAMPServer) GetHeaders() map[string]configopaque.String {
	_ = "STUB: not implemented"
	return nil
}

// GetTLSConfig returns a TLS config if the endpoint is secure (wss or https)
func (s OpAMPServer) GetTLSConfig(ctx context.Context) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s OpAMPServer) getTLS() configtls.ClientConfig {
	_ = "STUB: not implemented"
	return *new(configtls.ClientConfig)
}

func (s OpAMPServer) GetEndpoint() string { _ = "STUB: not implemented"; return "" }

func (s OpAMPServer) GetAuthExtensionID() component.ID {
	_ = "STUB: not implemented"
	return *new(component.ID)
}

func (s OpAMPServer) GetPollingInterval() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Validate checks if the extension configuration is valid
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }
