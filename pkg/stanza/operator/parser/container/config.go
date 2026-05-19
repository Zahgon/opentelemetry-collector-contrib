// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package container // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/container"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/attrs"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/transformer/recombine"
)

const (
	operatorType              = "container"
	recombineSourceIdentifier = attrs.LogFilePath
	recombineIsLastEntry      = "attributes.logtag == 'F'"
	defaultMaxLogSize         = 1024 * 1024
)

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new JSON parser config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new JSON parser config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a Container parser operator.
type Config struct {
	helper.ParserConfig `mapstructure:",squash"`

	Format                  string          `mapstructure:"format"`
	AddMetadataFromFilePath bool            `mapstructure:"add_metadata_from_filepath"`
	MaxLogSize              helper.ByteSize `mapstructure:"max_log_size,omitempty"`
}

// Build will build a Container parser operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}

// createRecombine creates an internal recombine operator which outputs to an async helper.LogEmitter
// the equivalent recombine config:
//
//	combine_field: body
//	combine_with: ""
//	is_last_entry: attributes.logtag == 'F'
//	max_log_size: 1048576 (1MiB)
//	source_identifier: attributes["log.file.path"]
//	type: recombine
func createRecombine(set component.TelemetrySettings, c Config, cLogEmitter *helper.BatchingLogEmitter) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}

// set the LogEmmiter as the output of the recombine parser

func createRecombineConfig(c Config) *recombine.Config { _ = "STUB: not implemented"; return nil }

// Set batch sizes to 0 (unlimited) - rely on max_log_size for protection
