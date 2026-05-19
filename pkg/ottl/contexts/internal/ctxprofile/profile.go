// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxprofile // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxprofile"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func PathGetSetter[K Context](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accessSample[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessTimeUnixNano[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessTime[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessDurationUnixNano[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessDuration[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessPeriodType[K Context](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accessPeriod[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessSampleType[K Context](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accessProfileID[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessStringProfileID[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessAttributeIndices[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessDroppedAttributesCount[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessOriginalPayloadFormat[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessOriginalPayload[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}
