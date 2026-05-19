// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package expotest // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/data/expo/expotest"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/data/expo"
)

type Histogram struct {
	Ts pcommon.Timestamp

	Pos, Neg expo.Buckets
	PosNeg   expo.Buckets

	Scale int
	Count uint64
	Sum   *float64

	Min, Max *float64

	Zt float64
	Zc uint64
}

func (hist Histogram) Into() expo.DataPoint { _ = "STUB: not implemented"; return *new(expo.DataPoint) }

func zero[T comparable](v T) bool { _ = "STUB: not implemented"; return false }
