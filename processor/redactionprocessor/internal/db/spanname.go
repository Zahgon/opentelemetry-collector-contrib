// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package db // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor/internal/db"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// SanitizeSpanName obfuscates the span name if it represents a database statement.
// Returns the obfuscated name, whether a change was made, and any error encountered.
func SanitizeSpanName(span ptrace.Span, obfuscator *Obfuscator) (string, bool, error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

func GetDBSystem(attributes pcommon.Map) string { _ = "STUB: not implemented"; return "" }

func getStringAttrLower(attributes pcommon.Map, key string) string {
	_ = "STUB: not implemented"
	return ""
}
