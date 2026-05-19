// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package expotest // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/data/expo/expotest"

import (
	"math"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/data/expo"
)

const (
	Empty = math.MaxUint64
	ø     = Empty
)

// index:  0  1  2 3 4 5 6 7
// bucket: -3 -2 -1 0 1 2 3 4
// bounds: (0.125,0.25], (0.25,0.5], (0.5,1], (1,2], (2,4], (4,8], (8,16], (16,32]
type Bins [8]uint64

func (bins Bins) Into() expo.Buckets { _ = "STUB: not implemented"; return *new(expo.Buckets) }

func ObserveInto(bs expo.Buckets, scale expo.Scale, pts ...float64) {
	_ = "STUB: not implemented"
	return
}

func Observe(scale expo.Scale, pts ...float64) expo.Buckets {
	_ = "STUB: not implemented"
	return *new(expo.Buckets)
}

func Observe0(pts ...float64) expo.Buckets { _ = "STUB: not implemented"; return *new(expo.Buckets) }
