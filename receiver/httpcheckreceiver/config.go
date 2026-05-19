// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package httpcheckreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/httpcheckreceiver"

import (
	"errors"

	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/httpcheckreceiver/internal/metadata"
)

// Predefined error responses for configuration validation failures
var (
	errInvalidEndpoint = errors.New(`"endpoint" must be in the form of <scheme>://<hostname>[:<port>]`)
	errMissingEndpoint = errors.New("at least one of 'endpoint' or 'endpoints' must be specified")
)

// Config defines the configuration for the various elements of the receiver agent.
type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	Targets                        []*targetConfig `mapstructure:"targets"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// validationConfig defines configuration for response validation
type validationConfig struct {
	// String matching
	Contains    string `mapstructure:"contains"`
	NotContains string `mapstructure:"not_contains"`

	// JSON path validation
	JSONPath string `mapstructure:"json_path"`
	Equals   string `mapstructure:"equals"`

	// Size validation
	MaxSize *int64 `mapstructure:"max_size"`
	MinSize *int64 `mapstructure:"min_size"`

	// Regex validation
	Regex string `mapstructure:"regex"`
}

// targetConfig defines configuration for individual HTTP checks.
type targetConfig struct {
	confighttp.ClientConfig `mapstructure:",squash"`
	Method                  string             `mapstructure:"method"`
	Endpoints               []string           `mapstructure:"endpoints"`         // Field for a list of endpoints
	Body                    string             `mapstructure:"body"`              // Request body content
	AutoContentType         bool               `mapstructure:"auto_content_type"` // Whether to automatically set Content-Type based on body
	Validations             []validationConfig `mapstructure:"validations"`       // Response validation rules
}

// Validate validates an individual targetConfig.
func (cfg *targetConfig) Validate() error {
	_ = "STUB: not implemented"

	// Ensure at least one of 'endpoint' or 'endpoints' is specified.
	return nil
}

// Validate the single endpoint in ClientConfig.

// Validate each endpoint in the Endpoints list.

// Validate validates the top-level Config by checking each targetConfig.
func (cfg *Config) Validate() error {
	_ = "STUB: not implemented"

	// Ensure at least one target is configured.
	return nil
}

// Validate each targetConfig.
