// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxprofile // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxprofile"

import (
	"go.opentelemetry.io/collector/pdata/pprofile"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type valueTypeSource[K Context] = func(ctx K) pprofile.ValueType

func valueTypeGetterSetter[K Context](
	path ottl.Path[K],
	source valueTypeSource[K],
) (ottl.GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accessValueType[K Context](path ottl.Path[K], getValueType valueTypeSource[K]) ottl.GetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessValueTypeType[K Context](path ottl.Path[K], getValueType valueTypeSource[K]) ottl.GetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func accessValueTypeUnit[K Context](path ottl.Path[K], getValueType valueTypeSource[K]) ottl.GetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func getValueTypeString[K Context](
	path ottl.Path[K],
	dict pprofile.ProfilesDictionary,
	currIndex int32,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func setValueTypeString[K Context](
	path ottl.Path[K],
	dict pprofile.ProfilesDictionary,
	currIndex int32,
	val any,
) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
