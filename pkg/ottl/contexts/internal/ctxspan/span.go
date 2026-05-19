// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxspan // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxspan"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func PathGetSetter[K Context](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

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

func accessTraceState[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessTraceStateKey[K Context](keys []ottl.Key[K]) (ottl.StandardGetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accessParentSpanID[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessStringParentSpanID[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessSpanName[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessKind[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessStringKind[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessDeprecatedStringKind[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessStartTimeUnixNano[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessEndTimeUnixNano[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessStartTime[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessEndTime[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessAttributes[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessAttributesKey[K Context](keys []ottl.Key[K]) ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessSpanDroppedAttributesCount[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessEvents[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessDroppedEventsCount[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessLinks[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessDroppedLinksCount[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessStatus[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessStatusCode[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessStatusMessage[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessFlags[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }
