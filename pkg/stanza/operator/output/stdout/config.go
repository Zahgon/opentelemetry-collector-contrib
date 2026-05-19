// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stdout // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/output/stdout"

import (
	"io"
	"os"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "stdout"

// Stdout is a global handle to standard output
var Stdout io.Writer = os.Stdout

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig("") })
}

// NewConfig creates a new stdout config with default values
func NewConfig(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of the Stdout operator
type Config struct {
	helper.OutputConfig `mapstructure:",squash"`
}

// Build will build a stdout operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
