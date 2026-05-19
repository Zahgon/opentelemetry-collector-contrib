// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector/internal/metrics"

import (
	"encoding"
)

const (
	Milliseconds Unit = iota
	Seconds

	MillisecondsStr = "ms"
	SecondsStr      = "s"
)

type Unit int8

var (
	_ encoding.TextMarshaler   = (*Unit)(nil)
	_ encoding.TextUnmarshaler = (*Unit)(nil)
)

func (u Unit) String() string { _ = "STUB: not implemented"; return "" }

// MarshalText marshals Unit to text.
func (u Unit) MarshalText() (text []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText unmarshalls text to a Unit.
func (u *Unit) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }
