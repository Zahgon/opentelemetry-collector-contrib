// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stanzaerrors // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/stanzaerrors"

import "go.uber.org/zap/zapcore"

// ErrorDetails is a map of details for an agent error.
type ErrorDetails map[string]string

// MarshalLogObject will define the representation of details when logging.
func (d ErrorDetails) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

// createDetails will create details for an error from key/value pairs.
func createDetails(keyValues []string) ErrorDetails {
	_ = "STUB: not implemented"
	return *new(ErrorDetails)
}
