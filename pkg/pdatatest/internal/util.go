// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/internal"

import (
	"regexp"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/xpdata/entity"
)

func MaskResourceAttributeValue(res pcommon.Resource, attr string) {
	_ = "STUB: not implemented"
	return
}

func ChangeResourceAttributeValue(res pcommon.Resource, attr string, changeFn func(string) string) {
	_ = "STUB: not implemented"
	return
}

func MatchResourceAttributeValue(res pcommon.Resource, attr string, re *regexp.Regexp) {
	_ = "STUB: not implemented"
	return
}

// AddErrPrefix adds a prefix to every multierr error.
func AddErrPrefix(prefix string, in error) error { _ = "STUB: not implemented"; return nil }

func CompareResource(expected, actual pcommon.Resource) error {
	_ = "STUB: not implemented"
	return nil
}

func MaskResourceEntityRefs(res pcommon.Resource) { _ = "STUB: not implemented"; return }

func CompareEntityRefs(expected, actual entity.EntityRefSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func CompareInstrumentationScope(expected, actual pcommon.InstrumentationScope) error {
	_ = "STUB: not implemented"
	return nil
}

func CompareSchemaURL(expected, actual string) error { _ = "STUB: not implemented"; return nil }

func CompareAttributes(expected, actual pcommon.Map) error { _ = "STUB: not implemented"; return nil }

func CompareDroppedAttributesCount(expected, actual uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func OrderMapByKey(input map[string]any) map[string]any {
	_ = "STUB: not implemented"
	// Create a slice to hold the keys
	return nil
}

// Sort the keys

// Create a new map to hold the sorted key-value pairs
