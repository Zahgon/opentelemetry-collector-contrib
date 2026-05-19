// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"time"
)

// DurationWithInf is a custom type that can handle both regular durations and "inf" value
type DurationWithInf time.Duration

// NewDurationWithInf creates a new DurationWithInf from a string value
// It accepts duration strings (e.g., "5s", "1m") or "inf" for infinite duration
func NewDurationWithInf(s string) (DurationWithInf, error) {
	_ = "STUB: not implemented"
	return *new(DurationWithInf), nil
}

// MustDurationWithInf creates a new DurationWithInf from a string value
// It panics if the string cannot be parsed as a valid duration or "inf"
func MustDurationWithInf(s string) DurationWithInf {
	_ = "STUB: not implemented"
	return *new(DurationWithInf)
}

func (d *DurationWithInf) String() string { _ = "STUB: not implemented"; return "" }

func (d *DurationWithInf) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (*DurationWithInf) Type() string { _ = "STUB: not implemented"; return "" }

func (d *DurationWithInf) IsInf() bool { _ = "STUB: not implemented"; return false }

func (d *DurationWithInf) Duration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
