// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sshcheckreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sshcheckreceiver"

import (
	"errors"

	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sshcheckreceiver/internal/configssh"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sshcheckreceiver/internal/metadata"
)

// Predefined error responses for configuration validation failures
var (
	errMissingEndpoint           = errors.New(`"endpoint" not specified in config`)
	errInvalidEndpoint           = errors.New(`"endpoint" is invalid`)
	errMissingUsername           = errors.New(`"username" not specified in config`)
	errMissingPasswordAndKeyFile = errors.New(`either "password" or "key_file" is required`)

	errConfigNotSSHCheck = errors.New("config was not a SSH check receiver config")
)

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	configssh.SSHClientSettings    `mapstructure:",squash"`

	CheckSFTP            bool                          `mapstructure:"check_sftp"`
	MetricsBuilderConfig metadata.MetricsBuilderConfig `mapstructure:",squash"`
}

// SFTPEnabled tells whether SFTP metrics are Enabled in MetricsSettings.
func (c Config) SFTPEnabled() bool { _ = "STUB: not implemented"; return false }

func (c Config) Validate() (err error) { _ = "STUB: not implemented"; return nil }
