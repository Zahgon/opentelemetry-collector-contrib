// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxcache // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxcache"

const cacheErrorMessage = "access to cache must be performed using the same context, please replace %q with %q"

func NewError(lowerContext, pathContext, fullPath string) error {
	_ = "STUB: not implemented"
	return nil
}
