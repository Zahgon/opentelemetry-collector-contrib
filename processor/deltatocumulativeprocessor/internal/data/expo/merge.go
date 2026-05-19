// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package expo // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/data/expo"

// Merge combines the counts of buckets a and b into a.
// Both buckets MUST be of same scale
func Merge(arel, brel Buckets) { _ = "STUB: not implemented"; return }

// Skip leading and trailing zeros to reduce number of buckets.
// As we cap number of buckets this allows us to have higher scale.
