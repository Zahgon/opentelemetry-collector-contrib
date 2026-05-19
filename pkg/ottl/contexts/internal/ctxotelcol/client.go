// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxotelcol // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxotelcol"

import (
	"go.opentelemetry.io/collector/client"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func accessClient[K any](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accessClientMetadata[K any](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accessClientAddr[K any](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getAuthAttributeValue(authData client.AuthData, key string) (pcommon.Value, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Value), nil
}

func convertAuthDataToMap(authData client.AuthData) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func accessClientAuth[K any](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accessClientAuthAttributesKeys[K any]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessClientAuthAttributesKey[K any](keys []ottl.Key[K]) ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func convertClientMetadataToMap(md client.Metadata) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func accessClientMetadataKeys[K any]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessClientMetadataKey[K any](keys []ottl.Key[K]) ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}
