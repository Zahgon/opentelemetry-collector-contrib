// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxcommon // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxcommon"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pprofile"
)

func ParseSpanID(spanIDStr string) (pcommon.SpanID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID), nil
}

func ParseTraceID(traceIDStr string) (pcommon.TraceID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceID), nil
}

func ParseProfileID(profileIDStr string) (pprofile.ProfileID, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.ProfileID), nil
}
