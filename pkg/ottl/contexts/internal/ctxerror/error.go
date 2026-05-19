// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxerror // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxerror"

const defaultErrorMessage = "segment %q from path %q is not a valid path nor a valid OTTL keyword for the %v context - review %v to see all valid paths"

func New(pathSegment, fullPath, context, ref string) error { _ = "STUB: not implemented"; return nil }
