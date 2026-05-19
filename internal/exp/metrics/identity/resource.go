// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package identity // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/identity"

import (
	"hash"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

type resource = Resource

type Resource struct {
	attrs [16]byte
}

func (r Resource) Hash() hash.Hash64 { _ = "STUB: not implemented"; return *new(hash.Hash64) }

func (r Resource) String() string { _ = "STUB: not implemented"; return "" }

func OfResource(r pcommon.Resource) Resource { _ = "STUB: not implemented"; return *new(Resource) }
