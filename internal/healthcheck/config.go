// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package healthcheck // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/healthcheck"

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/healthcheck/internal/common"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/healthcheck/internal/grpc"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/healthcheck/internal/http"
)

// Type aliases to expose internal types publicly
type (
	HTTPLegacyConfig             = http.LegacyConfig
	HTTPConfig                   = http.Config
	PathConfig                   = http.PathConfig
	GRPCConfig                   = grpc.Config
	ComponentHealthConfig        = common.ComponentHealthConfig
	CheckCollectorPipelineConfig = http.CheckCollectorPipelineConfig
	ResponseBodyConfig           = http.ResponseBodyConfig
)

const (
	httpConfigKey   = "http"
	grpcConfigKey   = "grpc"
	DefaultGRPCPort = 13132
	DefaultHTTPPort = 13133
)

var (
	ErrMissingProtocol      = errors.New("must specify at least one protocol")
	ErrGRPCEndpointRequired = errors.New("grpc endpoint required")
	ErrHTTPEndpointRequired = errors.New("http endpoint required")
	ErrInvalidPath          = errors.New("path must start with /")
)

// endpointForPort returns a localhost endpoint for the given port.
func endpointForPort(port int) string { _ = "STUB: not implemented"; return "" }

// Config has the configuration for the extension enabling the health check
// extension, used to report the health status of the service.
type Config struct {
	// LegacyConfig contains the config for the existing healthcheck extension.
	http.LegacyConfig `mapstructure:",squash"`

	// GRPCConfig is v2 config for the grpc healthcheck service.
	GRPCConfig *grpc.Config `mapstructure:"grpc"`

	// HTTPConfig is v2 config for the http healthcheck service.
	HTTPConfig *http.Config `mapstructure:"http"`

	// ComponentHealthConfig is v2 config shared between http and grpc services
	ComponentHealthConfig *common.ComponentHealthConfig `mapstructure:"component_health"`
}

var _ component.Config = (*Config)(nil)

// Validate checks if the extension configuration is valid
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// Unmarshal a confmap.Conf into the config struct.
func (c *Config) Unmarshal(conf *confmap.Conf) error {
	_ = "STUB: not implemented"
	// Initialize with default values to enable unmarshaling into nested structs.
	// For healthcheckextension: the feature gate determines behavior, not these fields.
	// For healthcheckv2extension: these fields control which protocols are enabled.
	// We conditionally initialize and then clear to preserve "user specified" vs "not specified".
	return nil
}

// Clear configs that weren't actually set in the confmap.
// This preserves the distinction between "user didn't specify" vs "user specified with defaults".

func NewDefaultConfig() component.Config { _ = "STUB: not implemented"; return *new(component.Config) }
