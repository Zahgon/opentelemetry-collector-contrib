// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"go.opentelemetry.io/collector/pdata/pprofile"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

const profileIDFuncName = "ProfileID"

type ProfileIDArguments[K any] struct {
	Target ottl.ByteSliceLikeGetter[K]
}

func NewProfileIDFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createProfileIDFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func profileID[K any](target ottl.ByteSliceLikeGetter[K]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodeHexToProfileID(b []byte) (pprofile.ProfileID, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.ProfileID), nil
}
