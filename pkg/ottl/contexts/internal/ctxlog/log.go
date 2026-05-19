// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxlog // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxlog"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func PathGetSetter[K Context](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accessTimeUnixNano[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessObservedTimeUnixNano[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessTime[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessObservedTime[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessSeverityNumber[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessSeverityText[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessBody[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessBodyKey[K Context](key []ottl.Key[K]) ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessStringBody[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessAttributes[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessAttributesKey[K Context](key []ottl.Key[K]) ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessDroppedAttributesCount[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessFlags[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessTraceID[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessStringTraceID[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessSpanID[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessStringSpanID[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessEventName[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }
