// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package expo // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/data/expo"

import (
	"cmp"
)

// WidenZero widens the zero-bucket to span at least [-width,width], possibly wider
// if min falls in the middle of a bucket.
//
// Both buckets counts MUST be of same scale.
func WidenZero(dp DataPoint, width float64) { _ = "STUB: not implemented"; return }

// the largest bucket index inside the zero width

// right next to the new zero bucket, constrained to slice range

// Slice drops data outside the range from <= i < to from the bucket counts. It behaves the same as Go's [a:b]
//
// Limitations:
//   - due to a limitation of the pcommon package, slicing cannot happen in-place and allocates
//   - in consequence, data outside the range is garbage collected
func (a Absolute) Slice(from, to int) { _ = "STUB: not implemented"; return }

// clamp constraints v to the range up..=lo
func clamp[N cmp.Ordered](v, lo, up N) N { _ = "STUB: not implemented"; return *new(N) }
