// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"net/http"
	"time"

	"github.com/open-telemetry/opamp-go/protobufs"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configtelemetry"
	"go.opentelemetry.io/collector/config/configtls"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/service/telemetry/otelconftelemetry"
	config "go.opentelemetry.io/contrib/otelconf/v0.3.0"
	"go.uber.org/zap/zapcore"

	"github.com/open-telemetry/opentelemetry-collector-contrib/cmd/opampsupervisor/supervisor/extensions"
)

var _ confmap.Validator = (*Supervisor)(nil)

// Supervisor is the Supervisor config file format.
type Supervisor struct {
	Server       OpAMPServer       `mapstructure:"server"`
	Agent        Agent             `mapstructure:"agent"`
	Capabilities Capabilities      `mapstructure:"capabilities"`
	Storage      Storage           `mapstructure:"storage"`
	Telemetry    Telemetry         `mapstructure:"telemetry"`
	HealthCheck  HealthCheck       `mapstructure:"healthcheck"`
	Extensions   extensions.Config `mapstructure:"extensions,omitempty"`
}

// Load loads the Supervisor config from a file.
func Load(configFile string) (Supervisor, error) {
	_ = "STUB: not implemented"
	return *new(Supervisor), nil
}

func (Supervisor) Validate() error { _ = "STUB: not implemented"; return nil }

type Storage struct {
	// Directory is the directory where the Supervisor will store its data.
	Directory string `mapstructure:"directory"`
	// prevent unkeyed literal initialization
	_ struct{}
}

// Capabilities is the set of capabilities that the Supervisor supports.
type Capabilities struct {
	AcceptsRemoteConfig            bool `mapstructure:"accepts_remote_config"`
	AcceptsRestartCommand          bool `mapstructure:"accepts_restart_command"`
	AcceptsOpAMPConnectionSettings bool `mapstructure:"accepts_opamp_connection_settings"`
	ReportsEffectiveConfig         bool `mapstructure:"reports_effective_config"`
	ReportsOwnMetrics              bool `mapstructure:"reports_own_metrics"`
	ReportsOwnLogs                 bool `mapstructure:"reports_own_logs"`
	ReportsOwnTraces               bool `mapstructure:"reports_own_traces"`
	ReportsHealth                  bool `mapstructure:"reports_health"`
	ReportsRemoteConfig            bool `mapstructure:"reports_remote_config"`
	ReportsAvailableComponents     bool `mapstructure:"reports_available_components"`
	ReportsHeartbeat               bool `mapstructure:"reports_heartbeat"`
	AcceptsPackages                bool `mapstructure:"accepts_packages"`
	ReportsPackageStatuses         bool `mapstructure:"reports_package_statuses"`
}

func (c Capabilities) SupportedCapabilities() protobufs.AgentCapabilities {
	_ = "STUB: not implemented"
	return *new(protobufs.AgentCapabilities)
}

// AcceptsPackages and ReportsPackageStatuses are not yet fully implemented.
// They are included here for completeness.
// See https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/47272

type OpAMPServer struct {
	Endpoint string                 `mapstructure:"endpoint"`
	Headers  http.Header            `mapstructure:"headers"`
	TLS      configtls.ClientConfig `mapstructure:"tls,omitempty"`
	// prevent unkeyed literal initialization
	_ struct{}
}

func (o OpAMPServer) OpaqueHeaders() map[string][]configopaque.String {
	_ = "STUB: not implemented"
	return nil
}

func (o OpAMPServer) Validate() error { _ = "STUB: not implemented"; return nil }

type Agent struct {
	Executable              string            `mapstructure:"executable"`
	InstanceID              string            `mapstructure:"instance_id"`
	OrphanDetectionInterval time.Duration     `mapstructure:"orphan_detection_interval"`
	Description             AgentDescription  `mapstructure:"description"`
	ConfigApplyTimeout      time.Duration     `mapstructure:"config_apply_timeout"`
	BootstrapTimeout        time.Duration     `mapstructure:"bootstrap_timeout"`
	OpAMPServerPort         int               `mapstructure:"opamp_server_port"`
	PassthroughLogs         bool              `mapstructure:"passthrough_logs"`
	UseHUPConfigReload      bool              `mapstructure:"use_hup_config_reload"`
	ValidateConfig          bool              `mapstructure:"validate_config"`
	ConfigFiles             []string          `mapstructure:"config_files"`
	Arguments               []string          `mapstructure:"args"`
	Env                     map[string]string `mapstructure:"env"`
}

func (a Agent) Validate() error { _ = "STUB: not implemented"; return nil }

type SpecialConfigFile string

const (
	SpecialConfigFileOwnTelemetry   SpecialConfigFile = "$OWN_TELEMETRY_CONFIG"
	SpecialConfigFileOpAMPExtension SpecialConfigFile = "$OPAMP_EXTENSION_CONFIG"
	SpecialConfigFileRemoteConfig   SpecialConfigFile = "$REMOTE_CONFIG"
)

var SpecialConfigFiles = []SpecialConfigFile{
	SpecialConfigFileOwnTelemetry,
	SpecialConfigFileOpAMPExtension,
	SpecialConfigFileRemoteConfig,
}

type AgentDescription struct {
	IdentifyingAttributes    map[string]string `mapstructure:"identifying_attributes"`
	NonIdentifyingAttributes map[string]string `mapstructure:"non_identifying_attributes"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type Telemetry struct {
	// TODO: Add more telemetry options
	// Issue here: https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/35582
	Logs    Logs                           `mapstructure:"logs"`
	Metrics Metrics                        `mapstructure:"metrics"`
	Traces  otelconftelemetry.TracesConfig `mapstructure:"traces"`

	Resource otelconftelemetry.ResourceConfig `mapstructure:"resource"`
	// prevent unkeyed literal initialization
	_ struct{}
}

type ResourceConfig = otelconftelemetry.ResourceConfig

type HealthCheck struct {
	confighttp.ServerConfig `mapstructure:",squash"`
	// prevent unkeyed literal initialization
	_ struct{}
}

func (h HealthCheck) Port() int64 { _ = "STUB: not implemented"; return 0 }

func (h HealthCheck) Validate() error { _ = "STUB: not implemented"; return nil }

type Logs struct {
	Level            zapcore.Level `mapstructure:"level"`
	ErrorOutputPaths []string      `mapstructure:"error_output_paths"`
	OutputPaths      []string      `mapstructure:"output_paths"`
	// Processors allow configuration of log record processors to emit logs to
	// any number of supported backends.
	Processors []config.LogRecordProcessor `mapstructure:"processors,omitempty"`
}

type Metrics struct {
	Level   configtelemetry.Level `mapstructure:"level"`
	Readers []config.MetricReader `mapstructure:"readers"`
	// prevent unkeyed literal initialization
	_ struct{}
}

// DefaultSupervisor returns the default supervisor config
func DefaultSupervisor() Supervisor { _ = "STUB: not implemented"; return *new(Supervisor) }

// Windows default is "%ProgramData%\Otelcol\Supervisor"
// If the ProgramData environment variable is not set,
// it falls back to C:\ProgramData
