// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsxrayreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver/internal/errors"

// ErrRecoverable represents an error that can be ignored
// so that the receiver can continue to function.
type ErrRecoverable struct {
	Err error
}

func (e *ErrRecoverable) Error() string { _ = "STUB: not implemented"; return "" }

// Unwrap implements the new error feature introduced in Go 1.13
func (e *ErrRecoverable) Unwrap() error {
	_ = "STUB: not implemented"

	// ErrIrrecoverable represents an error that should
	// stop the receiver.
	return nil
}

type ErrIrrecoverable struct {
	Err error
}

func (e *ErrIrrecoverable) Error() string { _ = "STUB: not implemented"; return "" }

// Unwrap implements the new error feature introduced in Go 1.13
func (e *ErrIrrecoverable) Unwrap() error { _ = "STUB: not implemented"; return nil }
