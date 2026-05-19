// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stanzaerrors // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/stanzaerrors"

import (
	"go.uber.org/zap/zapcore"
)

// AgentError is an error that occurs in the agent.
type AgentError struct {
	Description string
	Suggestion  string
	Details     ErrorDetails
}

// Error will return the error message.
func (e AgentError) Error() string { _ = "STUB: not implemented"; return "" }

// MarshalLogObject will define the representation of this error when logging.
func (e AgentError) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// WithDetails will return the error with additional details
func (e AgentError) WithDetails(keyValues ...string) AgentError {
	_ = "STUB: not implemented"
	return *new(AgentError)
}

// WithDetails will add details to an agent error
func WithDetails(err error, keyValues ...string) AgentError {
	_ = "STUB: not implemented"
	return *new(AgentError)
}

// Deprecated: [v0.143.0] use fmt.Errorf with %w.
func Wrap(err error, context string) AgentError { _ = "STUB: not implemented"; return *new(AgentError) }

// NewError will create a new agent error.
func NewError(description, suggestion string, keyValues ...string) AgentError {
	_ = "STUB: not implemented"
	return *new(AgentError)
}
