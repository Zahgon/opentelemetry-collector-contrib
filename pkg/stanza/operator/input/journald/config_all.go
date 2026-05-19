// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package journald // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/journald"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

const operatorType = "journald_input"

// NewConfig creates a new input config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// NewConfigWithID creates a new input config with default values
func NewConfigWithID(operatorID string) *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a journald input operator
type Config struct {
	helper.InputConfig `mapstructure:",squash"`

	RootPath            string        `mapstructure:"root_path,omitempty"`
	JournalctlPath      string        `mapstructure:"journalctl_path,omitempty"`
	Directory           *string       `mapstructure:"directory,omitempty"`
	Files               []string      `mapstructure:"files,omitempty"`
	StartAt             string        `mapstructure:"start_at,omitempty"`
	Units               []string      `mapstructure:"units,omitempty"`
	Priority            string        `mapstructure:"priority,omitempty"`
	Matches             []MatchConfig `mapstructure:"matches,omitempty"`
	Identifiers         []string      `mapstructure:"identifiers,omitempty"`
	Grep                string        `mapstructure:"grep,omitempty"`
	Dmesg               bool          `mapstructure:"dmesg,omitempty"`
	All                 bool          `mapstructure:"all,omitempty"`
	Namespace           string        `mapstructure:"namespace,omitempty"`
	ConvertMessageBytes bool          `mapstructure:"convert_message_bytes,omitempty"`
	Merge               bool          `mapstructure:"merge,omitempty"`
}

type MatchConfig map[string]string
