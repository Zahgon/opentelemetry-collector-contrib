// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package traces // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/traces"

import (
	"github.com/Masterminds/semver/v3"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/otel/attribute"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
)

var (
	minKnownSemConvVersion = semver.MustParse("1.37.0")
	maxKnownSemConvVersion = semver.MustParse("1.40.0")
)

type setSemconvSpanNameArguments struct {
	SemconvVersion            string
	OriginalSpanNameAttribute ottl.Optional[string]
}

func NewSetSemconvSpanNameFactory() ottl.Factory[*ottlspan.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func createSetSemconvSpanNameFunction(_ ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[*ottlspan.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseSemconvSpanNameArguments(oArgs ottl.Arguments) (*setSemconvSpanNameArguments, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setSemconvSpanName(originalSpanNameAttribute ottl.Optional[string], span ptrace.Span) {
	_ = "STUB: not implemented"
	return
}

func deriveSemconvSpanName(span ptrace.Span) string { _ = "STUB: not implemented"; return "" }

// If no semantic convention defines the span name, default to the original span name

// https://opentelemetry.io/docs/specs/semconv/http/http-spans/
func httpSpanName(span ptrace.Span, subject attribute.Key) string {
	_ = "STUB: not implemented"
	return ""
}

// https://opentelemetry.io/docs/specs/semconv/rpc/rpc-spans/
func rpcSpanName(span ptrace.Span) string { _ = "STUB: not implemented"; return "" }

// https://opentelemetry.io/docs/specs/semconv/database/database-spans/
func dbSpanName(span ptrace.Span) string { _ = "STUB: not implemented"; return "" }

func databaseTarget(span ptrace.Span) string { _ = "STUB: not implemented"; return "" }

// https://opentelemetry.io/docs/specs/semconv/messaging/messaging-spans/#span-name
func messagingSpanName(span ptrace.Span) string { _ = "STUB: not implemented"; return "" }

// https://opentelemetry.io/docs/specs/semconv/messaging/messaging-spans/#span-name
func messagingDestination(span ptrace.Span) string { _ = "STUB: not implemented"; return "" }

// anonymous destinations should also be marked as temporary by the messaging instrumentation
// double check in case the instrumentation forgot to mark the anonymous destination as temporary

func attributeValue(span ptrace.Span, name attribute.Key, alias string) (pcommon.Value, bool) {
	_ = "STUB: not implemented"
	return *new(pcommon.Value), false
}
