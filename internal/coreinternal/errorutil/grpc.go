// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package errorutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/errorutil"

func GrpcError(err error) error { _ = "STUB: not implemented"; return nil }

// Default to a retryable error
// https://github.com/open-telemetry/opentelemetry-proto/blob/main/docs/specification.md#failures

// non-retryable error
