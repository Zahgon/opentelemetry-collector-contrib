// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pslice // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/putil/pslice"

type Slice[E any] interface {
	At(int) E
	Len() int
}

func Equal[E comparable, S Slice[E]](a, b S) bool { _ = "STUB: not implemented"; return false }

func All[E any, S Slice[E]](slice S) func(func(E) bool) { _ = "STUB: not implemented"; return nil }
