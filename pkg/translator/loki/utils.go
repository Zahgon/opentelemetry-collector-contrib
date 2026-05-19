// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loki // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/loki"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func extractInstance(attributes pcommon.Map) (string, bool) {
	_ = "STUB: not implemented"
	// Map service.instance.id to instance
	return "", false
}

func extractJob(attributes pcommon.Map) (string, bool) {
	_ = "STUB: not implemented"
	// Map service.namespace + service.name to job
	return "", false
}
