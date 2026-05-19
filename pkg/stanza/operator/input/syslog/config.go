// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package syslog // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/syslog"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/tcp"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/udp"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/syslog"
)

const operatorType = "syslog_input"

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new input config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new input config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

type Config struct {
	helper.InputConfig `mapstructure:",squash"`
	syslog.BaseConfig  `mapstructure:",squash"`
	TCP                *tcp.BaseConfig `mapstructure:"tcp"`
	UDP                *udp.BaseConfig `mapstructure:"udp"`
	OnError            string          `mapstructure:"on_error"`
}

func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}

// Octet counting and Non-Transparent-Framing are invalid for UDP connections
