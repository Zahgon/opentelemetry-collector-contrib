// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tinybirdexporter/internal"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func getServiceName(resAttr pcommon.Map) string { _ = "STUB: not implemented"; return "" }

func convertAttributes(attributes pcommon.Map) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
