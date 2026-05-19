// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxotelcol // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxotelcol"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"google.golang.org/grpc/metadata"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

func accessGRPC[K any](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertGRPCMetadataToMap(md metadata.MD) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func accessGRPCMetadataKeys[K any]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessGRPCMetadataKey[K any](keys []ottl.Key[K]) ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}
