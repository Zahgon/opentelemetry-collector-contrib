// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package syslog // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/syslog"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const (
	operatorType = "syslog_parser"

	RFC3164 = "rfc3164"
	RFC5424 = "rfc5424"

	NULTrailer = "NUL"
	LFTrailer  = "LF"
)

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new syslog parser config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new syslog parser config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a syslog parser operator.
type Config struct {
	helper.ParserConfig `mapstructure:",squash"`
	BaseConfig          `mapstructure:",squash"`
}

// BaseConfig is the detailed configuration of a syslog parser.
type BaseConfig struct {
	Protocol                     string  `mapstructure:"protocol,omitempty"`
	Location                     string  `mapstructure:"location,omitempty"`
	EnableOctetCounting          bool    `mapstructure:"enable_octet_counting,omitempty"`
	AllowSkipPriHeader           bool    `mapstructure:"allow_skip_pri_header,omitempty"`
	NonTransparentFramingTrailer *string `mapstructure:"non_transparent_framing_trailer,omitempty"`
	MaxOctets                    int     `mapstructure:"max_octets,omitempty"`
}

// Build will build a JSON parser operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
