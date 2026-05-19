// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package identity // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/identity"

import (
	"hash"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

type Stream struct {
	metric Metric
	attrs  [16]byte
}

func (s Stream) Hash() hash.Hash64 { _ = "STUB: not implemented"; return *new(hash.Hash64) }

func (s Stream) Metric() Metric { _ = "STUB: not implemented"; return *new(Metric) }

func (s Stream) String() string { _ = "STUB: not implemented"; return "" }

func OfStream[DataPoint attrPoint](m Metric, dp DataPoint) Stream {
	_ = "STUB: not implemented"
	return *new(Stream)
}

type attrPoint interface {
	Attributes() pcommon.Map
}
