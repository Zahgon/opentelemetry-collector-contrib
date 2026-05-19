// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package identity // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/identity"

import (
	"hash"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

type scope = Scope

type Scope struct {
	resource resource

	name    string
	version string
	attrs   [16]byte
}

func (s Scope) Hash() hash.Hash64 { _ = "STUB: not implemented"; return *new(hash.Hash64) }

func (s Scope) Resource() Resource { _ = "STUB: not implemented"; return *new(Resource) }

func (s Scope) String() string { _ = "STUB: not implemented"; return "" }

func OfScope(res Resource, scope pcommon.InstrumentationScope) Scope {
	_ = "STUB: not implemented"
	return *new(Scope)
}
