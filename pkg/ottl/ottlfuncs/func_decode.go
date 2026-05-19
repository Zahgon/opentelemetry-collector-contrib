// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"context"
	"encoding/base64"

	"golang.org/x/text/encoding"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type decoder interface {
	Decode(src []byte) (any, error)
	DecodeString(src string) (any, error)
}

type textDecoder struct {
	enc encoding.Encoding
}

func (td textDecoder) Decode(src []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (td textDecoder) DecodeString(src string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type base64Decoder struct {
	enc *base64.Encoding
}

func (bd base64Decoder) Decode(src []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (bd base64Decoder) DecodeString(src string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type DecodeArguments[K any] struct {
	Target   ottl.Getter[K]
	Encoding ottl.StringGetter[K]
}

func NewDecodeFactory[K any]() ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func createDecodeFunction[K any](_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decode[K any](target ottl.Getter[K], encoding ottl.StringGetter[K]) (ottl.ExprFunc[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type decoderGetter[K any] struct {
	decoder decoder
	getter  ottl.StringGetter[K]
}

func newDecoderGetter[K any](getter ottl.StringGetter[K]) (*decoderGetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *decoderGetter[K]) Get(ctx context.Context, tCtx K) (decoder, error) {
	_ = "STUB: not implemented"
	return *new(decoder), nil
}

func asDecoder(encodingVal string) (decoder, error) {
	_ = "STUB: not implemented"
	return *

	// base64 is not in IANA index, so we have to deal with this encoding separately
	new(decoder), nil
}
