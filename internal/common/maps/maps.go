// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package maps // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/maps"

// MergeRawMaps merges n maps with a later map's keys overriding earlier maps.
func MergeRawMaps(maps ...map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }

// MergeStringMaps merges n maps with a later map's keys overriding earlier maps.
func MergeStringMaps(maps ...map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// CloneStringMap makes a shallow copy of a map[string]string.
func CloneStringMap(m map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }
