// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelserializer // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/serializer/otelserializer"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

const tsLayout = "2006-01-02T15:04:05.000000000Z"

func isGeoAttribute(k string, val pcommon.Value) bool { _ = "STUB: not implemented"; return false }
