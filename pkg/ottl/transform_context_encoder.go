// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	transformContextKey = "TransformContext"
	deletedReplacement  = "<deleted>"
)

// newTransformContextField creates a zapcore.Field that can be used with any OTTL parser
// transform context object. It marshals the PData safely by filtering out invalid values,
// avoiding panics. It is intended for debug logging only and should not be used in
// production code.
func newTransformContextField(tCtx any) zap.Field {
	_ = "STUB: not implemented"
	return *new(zap.Field)
}

type transformContextMarshaller struct {
	tCtx any
}

func (m *transformContextMarshaller) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type transformContextEncoder struct {
	zapcore.ObjectEncoder
	zapcore.ArrayEncoder
}

func (m *transformContextEncoder) withArrayEncoder(enc zapcore.ArrayEncoder) *transformContextEncoder {
	_ = "STUB: not implemented"
	return nil
}

func (m *transformContextEncoder) withObjectEncoder(enc zapcore.ObjectEncoder) *transformContextEncoder {
	_ = "STUB: not implemented"
	return nil
}

func (m *transformContextEncoder) AddArray(k string, v zapcore.ArrayMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *transformContextEncoder) AddObject(k string, v zapcore.ObjectMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *transformContextEncoder) AddReflected(k string, v any) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *transformContextEncoder) OpenNamespace(k string) { _ = "STUB: not implemented"; return }

func (m *transformContextEncoder) AppendArray(v zapcore.ArrayMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *transformContextEncoder) AppendObject(v zapcore.ObjectMarshaler) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *transformContextEncoder) AppendReflected(v any) error {
	_ = "STUB: not implemented"
	return nil
}

// isInvalidPData reports whether `data` refers to an invalid PData value.
// A PData becomes invalid after it is deleted from its parent; in that state the underlying
// `orig` field is set to nil, and attempting to marshal it may panic.
// There is no stable public API to detect this today, so this function uses reflection to
// look for an `orig` field and treat a nil `orig` as invalid.
func isInvalidPData(data any) bool { _ = "STUB: not implemented"; return false }
