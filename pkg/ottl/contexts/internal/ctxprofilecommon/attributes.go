// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxprofilecommon // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxprofilecommon"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pprofile"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type ProfileAttributable interface {
	AttributeIndices() pcommon.Int32Slice
}

type attributeSource[K any] = func(ctx K) (pprofile.ProfilesDictionary, ProfileAttributable)

func AccessAttributes[K any](source attributeSource[K]) ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func AccessAttributesKey[K any](key []ottl.Key[K], source attributeSource[K]) ottl.StandardGetSetter[K] {
	_ = "STUB: not implemented"
	return nil
}

func getAttributeValue(dict pprofile.ProfilesDictionary, indices pcommon.Int32Slice, key string) pcommon.Value {
	_ = "STUB: not implemented"
	return *new(pcommon.Value)
}

// Copy the value because OTTL expects to do inplace updates for the values.
