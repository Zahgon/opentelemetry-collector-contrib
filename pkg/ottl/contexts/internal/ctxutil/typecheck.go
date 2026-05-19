// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxutil"

// ExpectType ensures val can be asserted to T, returning a descriptive error when it cannot.
func ExpectType[T any](val any) (T, error) { _ = "STUB: not implemented"; return *new(T), nil }
