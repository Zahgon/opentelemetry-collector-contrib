// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter/internal"

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func GetServiceName(resAttr pcommon.Map) string { _ = "STUB: not implemented"; return "" }

func AttributesToMap(attributes pcommon.Map) column.IterableOrderedMap {
	_ = "STUB: not implemented"
	return *new(column.IterableOrderedMap)
}

// UniqueFlattenedAttributes converts a pcommon.Map into a slice of attributes. Paths are flattened and sorted.
func UniqueFlattenedAttributes(m pcommon.Map) []string { _ = "STUB: not implemented"; return nil }

func uniqueFlattenedAttributesNested(pathPrefix string, pathsSet *map[string]struct{}, paths *[]string, m pcommon.Map) {
	_ = "STUB: not implemented"
	return
}
