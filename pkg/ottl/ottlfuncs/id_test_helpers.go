// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

type idExprBuilder func(ottl.ByteSliceLikeGetter[any]) (ottl.ExprFunc[any], error)

type idSuccessTestCase struct {
	name  string
	value []byte
	want  any
}

type idErrorTestCase struct {
	name  string
	value []byte
	err   error
}

// makeIDGetter creates a ByteSliceLikeGetter for testing purposes.
// This is a shared helper used by TraceID, SpanID, and ProfileID tests.
func makeIDGetter(bytes []byte) ottl.ByteSliceLikeGetter[any] {
	_ = "STUB: not implemented"
	return nil
}

func runIDSuccessTests(t *testing.T, builder idExprBuilder, cases []idSuccessTestCase) {
	_ = "STUB: not implemented"
	return
}

func runIDErrorTests(t *testing.T, builder idExprBuilder, funcName string, cases []idErrorTestCase) {
	_ = "STUB: not implemented"
	return
}

// Dynamic getters succeed at init, fail at execution

// assertErrorIsFuncDecode asserts that the error message contains the function name.
func assertErrorIsFuncDecode(t *testing.T, err error, funcName string) {
	_ = "STUB: not implemented"
	return
}
