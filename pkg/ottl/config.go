// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"

// ErrorMode is the way OTTL should handle errors.
type ErrorMode string

const (
	// IgnoreError means OTTL will only log errors.
	IgnoreError ErrorMode = "ignore"
	// PropagateError means OTTL will log and return errors.
	PropagateError ErrorMode = "propagate"
	// SilentError means OTTL will not log or return errors.
	SilentError ErrorMode = "silent"
)

func (e *ErrorMode) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// LogicOperation represents the logical operations OTTL understands.
type LogicOperation string

const (
	// And is the logical operator "and".
	And LogicOperation = "and"
	// Or is the logical operator "or".
	Or LogicOperation = "or"
)

// UnmarshalText unmarshals a string into a LogicOperation. It errors if the string is not "and" or "or".
func (l *LogicOperation) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }
