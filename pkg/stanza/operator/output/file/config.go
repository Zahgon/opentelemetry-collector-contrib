// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package file // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/output/file"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "file_output"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig("") })
}

// NewConfig creates a new file output config with default values
func NewConfig(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a file output operatorn.
type Config struct {
	helper.OutputConfig `mapstructure:",squash"`

	Path   string `mapstructure:"path"`
	Format string `mapstructure:"format"`
}

// Build will build a file output operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
