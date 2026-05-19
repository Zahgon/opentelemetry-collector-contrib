// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package systemdreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/systemdreceiver"

import (
	"errors"

	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/systemdreceiver/internal/metadata"
)

var (
	errInvalidScope = errors.New(`"scope" must be one of "systemd" or "user"`)
	errNoUnits      = errors.New("no units configured")
)

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`

	Scope string   `mapstructure:"scope"`
	Units []string `mapstructure:"units"`
}

func (c Config) Validate() error {
	_ = "STUB: not implemented"

	// Ensure we have a valid scope.
	return nil
}
