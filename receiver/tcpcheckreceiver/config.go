// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tcpcheckreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tcpcheckreceiver"

import (
	"errors"

	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tcpcheckreceiver/internal/metadata"
)

// Predefined error responses for configuration validation failures
var (
	errInvalidEndpoint = errors.New(`"Endpoint" must be in the form of <hostname>:<port>`)
	errMissingTargets  = errors.New(`No targets specified`)
	errConfigTCPCheck  = errors.New(`Invalid Config`)
)

// Config defines the configuration for the various elements of the receiver agent.
type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	Targets                        []*confignet.TCPAddrConfig `mapstructure:"targets"`

	// prevent unkeyed literal initialization
	_ struct{}
}

func validatePort(port string) error { _ = "STUB: not implemented"; return nil }

func validateTarget(cfg *confignet.TCPAddrConfig) error { _ = "STUB: not implemented"; return nil }

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }
