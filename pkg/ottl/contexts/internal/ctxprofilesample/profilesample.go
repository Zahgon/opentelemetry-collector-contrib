// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxprofilesample // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxprofilesample"

import (
	"errors"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

var errMaxValueExceed = errors.New("exceeded max value")

func PathGetSetter[K Context](path ottl.Path[K]) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accessValues[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessAttributeIndices[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessLinkIndex[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }

func accessTimestampsUnixNano[K Context]() ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessTimestamps[K Context]() ottl.StandardGetSetter[K] { _ = "STUB: not implemented"; return nil }
