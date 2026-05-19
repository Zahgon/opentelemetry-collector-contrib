// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package entry // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"

// copyValue will deep copy a value based on its type.
func copyValue(v any) any { _ = "STUB: not implemented"; return *new(any) }

// copyStringMap will deep copy a map of strings.
func copyStringMap(m map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

// copyInterfaceMap will deep copy a map of interfaces.
func copyInterfaceMap(m map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

// copyStringArray will deep copy an array of strings.
func copyStringArray(a []string) []string { _ = "STUB: not implemented"; return nil }

// copyByteArray will deep copy an array of bytes.
func copyByteArray(a []byte) []byte { _ = "STUB: not implemented"; return nil }

// copyIntArray will deep copy an array of ints.
func copyIntArray(a []int) []int { _ = "STUB: not implemented"; return nil }

// copyInterfaceArray will deep copy an array of interfaces.
func copyInterfaceArray(a []any) []any { _ = "STUB: not implemented"; return nil }

// copyUnknown will copy an unknown value using json encoding.
// If this process fails, the result will be an empty interface.
func copyUnknown(value any) any { _ = "STUB: not implemented"; return *new(any) }
