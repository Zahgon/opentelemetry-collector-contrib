// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxcache // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxcache"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

const Name = "cache"

type Getter[K any] func(K) pcommon.Map

func PathExpressionParser[K any](cacheGetter Getter[K]) ottl.PathExpressionParser[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessCache[K any](cacheGetter Getter[K]) ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessCacheKey[K any](cacheGetter Getter[K], key []ottl.Key[K]) ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}
