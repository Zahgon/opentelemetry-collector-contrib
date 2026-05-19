// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxresource // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxresource"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func PathGetSetter[K Context](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accessResourceAttributes[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessResourceAttributesKey[K Context](keys []ottl.Key[K]) ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessResourceDroppedAttributesCount[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessResourceSchemaURLItem[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}
