// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxdatapoint // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxdatapoint"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func PathGetSetter[K Context](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accessAttributes[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessAttributesKey[K Context](key []ottl.Key[K]) ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessStartTimeUnixNano[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessStartTime[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessTimeUnixNano[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessTime[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessDoubleValue[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessIntValue[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessExemplars[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessFlags[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessCount[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessSum[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessExplicitBounds[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessBucketCounts[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessScale[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessZeroCount[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessPositive[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessPositiveOffset[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessPositiveBucketCounts[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessNegative[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessNegativeOffset[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessNegativeBucketCounts[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessQuantileValues[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}
