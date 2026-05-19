// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"errors"
	"fmt"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pprofile"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

var (
	errDecodeID        = errors.New("could not decode ID")
	errIDInvalidLength = fmt.Errorf("%w: %w", errDecodeID, errors.New("invalid length"))
	errIDHexDecode     = fmt.Errorf("%w: %w", errDecodeID, errors.New("invalid hex"))
)

type idByteArray interface {
	pcommon.SpanID | pcommon.TraceID | pprofile.ProfileID
}

// newIDExprFunc builds an expression function that accepts either a byte slice
// of the target length or a hex string twice that size.
// If the target is a literal getter, the ID is pre-computed once for optimal performance.
// We pass the hex decoder function as a parameter to allow implementations to decode directly into the ID type.
// This reduces allocations.
func newIDExprFunc[K any, R idByteArray](funcName string, target ottl.ByteSliceLikeGetter[K], hexDecoder func([]byte) (R, error)) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check if target is a literal getter, just grab the raw bytes if so

// Dynamic path: evaluate on every call

// bytesToID converts a byte slice to an ID of the specified type.
// It accepts either raw bytes of length idLen or hex-encoded bytes of length idHexLen.
func bytesToID[R idByteArray](funcName string, b []byte, idLen, idHexLen int, hexDecoder func([]byte) (R, error)) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// copyToFixedLenID copies the bytes from the source slice to the destination fixed length array.
func copyToFixedLenID[R idByteArray](dst *R, src []byte) { _ = "STUB: not implemented"; return }
