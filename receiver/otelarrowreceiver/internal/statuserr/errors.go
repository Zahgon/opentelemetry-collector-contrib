// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package statuserr // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/otelarrowreceiver/internal/statuserr"

// GetStatusFromError converts a pipeline error into an appropriate gRPC status
// error. If the error already has a gRPC status attached, it is returned as-is.
// Otherwise, non-permanent errors are mapped to codes.Unavailable (retryable)
// and permanent errors are mapped to codes.Internal.
//
// This mirrors the logic in the core OTLP receiver:
// go.opentelemetry.io/collector/receiver/otlpreceiver/internal/errors
//
// Without this conversion, gRPC defaults plain Go errors to codes.Unknown,
// which upstream exporters treat as a permanent (non-retryable) failure.
// See: https://github.com/grpc/grpc-go/blob/v1.59.0/server.go#L1345
func GetStatusFromError(err error) error { _ = "STUB: not implemented"; return nil }

// Default to a retryable error.

// If an error is permanent but doesn't have an attached gRPC
// status, assume it is server-side.
