// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs

import (
	"github.com/spf13/pflag"

	"github.com/open-telemetry/opentelemetry-collector-contrib/cmd/telemetrygen/internal/config"
)

// Config describes the test scenario.
type Config struct {
	config.Config
	NumLogs        int
	SeverityText   string
	SeverityNumber int32
	Body           string
	TraceID        string
	SpanID         string
}

func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// Flags registers config flags.
func (c *Config) Flags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// SetDefaults sets the default values for the configuration
// This is called before parsing the command line flags and when
// calling NewConfig()
func (c *Config) SetDefaults() { _ = "STUB: not implemented"; return }

// Validate validates the test scenario parameters.
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
