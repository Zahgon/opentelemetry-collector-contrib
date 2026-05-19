// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxmetric // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxmetric"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func PathGetSetter[K Context](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accessName[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessDescription[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessUnit[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessType[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

// TODO Implement methods so correctly convert data types.
// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/10130

func accessAggTemporality[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessIsMonotonic[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessDataPoints[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessMetadata[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessMetadataKey[K Context](keys []ottl.Key[K]) ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}
