// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func resourceSignature(attributes pcommon.Map) string { _ = "STUB: not implemented"; return "" }

func extractInstance(attributes pcommon.Map) (string, bool) {
	_ = "STUB: not implemented"
	// Map service.instance.id to instance
	return "", false
}

func extractJob(attributes pcommon.Map) (string, bool) {
	_ = "STUB: not implemented"
	// Map service.name + service.namespace to job
	return "", false
}
