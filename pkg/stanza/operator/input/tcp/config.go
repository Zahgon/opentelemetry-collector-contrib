// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tcp // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/tcp"

import (
	"bufio"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configtls"
	"golang.org/x/text/encoding"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/split"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/trim"
)

const (
	operatorType = "tcp_input"

	// minMaxLogSize is the minimal size which can be used for buffering
	// TCP input
	minMaxLogSize = 64 * 1024

	// DefaultMaxLogSize is the max buffer sized used
	// if MaxLogSize is not set
	DefaultMaxLogSize = 1024 * 1024
)

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new TCP input config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new TCP input config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a tcp input operator.
type Config struct {
	helper.InputConfig `mapstructure:",squash"`
	BaseConfig         `mapstructure:",squash"`
}

// BaseConfig is the detailed configuration of a tcp input operator.
type BaseConfig struct {
	MaxLogSize       helper.ByteSize         `mapstructure:"max_log_size,omitempty"`
	ListenAddress    string                  `mapstructure:"listen_address,omitempty"`
	TLS              *configtls.ServerConfig `mapstructure:"tls,omitempty"`
	AddAttributes    bool                    `mapstructure:"add_attributes,omitempty"`
	OneLogPerPacket  bool                    `mapstructure:"one_log_per_packet,omitempty"`
	Encoding         string                  `mapstructure:"encoding,omitempty"`
	SplitConfig      split.Config            `mapstructure:"multiline,omitempty"`
	TrimConfig       trim.Config             `mapstructure:",squash"`
	SplitFuncBuilder SplitFuncBuilder        `mapstructure:"-"`
}

type SplitFuncBuilder func(enc encoding.Encoding) (bufio.SplitFunc, error)

func (c Config) defaultSplitFuncBuilder(enc encoding.Encoding) (bufio.SplitFunc, error) {
	_ = "STUB: not implemented"
	return *new(bufio.SplitFunc), nil
}

// Build will build a tcp input operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}

// If MaxLogSize not set, set sane default

// validate the input address

// Build split func
