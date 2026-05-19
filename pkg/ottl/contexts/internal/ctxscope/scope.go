// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxscope // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxscope"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func PathGetSetter[K Context](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accessInstrumentationScopeAttributes[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessInstrumentationScopeAttributesKey[K Context](keys []ottl.Key[K]) ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessInstrumentationScopeName[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessInstrumentationScopeVersion[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessInstrumentationScopeDroppedAttributesCount[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessInstrumentationScopeSchemaURLItem[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}
