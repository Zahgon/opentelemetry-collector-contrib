// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxspanevent // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxspanevent"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func PathGetSetter[K Context](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accessSpanEventTimeUnixNano[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessSpanEventTime[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessSpanEventName[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessSpanEventAttributes[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessSpanEventAttributesKey[K Context](key []ottl.Key[K]) ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessSpanEventDroppedAttributeCount[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}
