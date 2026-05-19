// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package udp // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/udp"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/split"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/trim"
)

const (
	operatorType = "udp_input"

	// Maximum UDP packet size
	MaxUDPSize = 64 * 1024

	defaultReaders        = 1
	defaultProcessors     = 1
	defaultMaxQueueLength = 100
)

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new UDP input config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new UDP input config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Use never matching regex to not split data by default

// Config is the configuration of a udp input operator.
type Config struct {
	helper.InputConfig `mapstructure:",squash"`
	BaseConfig         `mapstructure:",squash"`
}

type AsyncConfig struct {
	Readers        int `mapstructure:"readers,omitempty"`
	Processors     int `mapstructure:"processors,omitempty"`
	MaxQueueLength int `mapstructure:"max_queue_length,omitempty"`
}

// BaseConfig is the details configuration of a udp input operator.
type BaseConfig struct {
	ListenAddress   string       `mapstructure:"listen_address,omitempty"`
	OneLogPerPacket bool         `mapstructure:"one_log_per_packet,omitempty"`
	AddAttributes   bool         `mapstructure:"add_attributes,omitempty"`
	Encoding        string       `mapstructure:"encoding,omitempty"`
	SplitConfig     split.Config `mapstructure:"multiline,omitempty"`
	TrimConfig      trim.Config  `mapstructure:",squash"`
	AsyncConfig     *AsyncConfig `mapstructure:"async,omitempty"`
}

// Build will build a udp input operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}

// Build split func
