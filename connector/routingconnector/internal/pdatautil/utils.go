// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pdatautil // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector/internal/pdatautil"

// OnceValue to set a given value only once.
type OnceValue[K any] struct {
	val    K
	isInit bool
}

func (ov *OnceValue[K]) IsInit() bool { _ = "STUB: not implemented"; return false }

func (ov *OnceValue[K]) Init(val K) { _ = "STUB: not implemented"; return }

func (ov *OnceValue[K]) Value() K { _ = "STUB: not implemented"; return *new(K) }
