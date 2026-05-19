// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package servicegraphconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/servicegraphconnector"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func findServiceName(attributes pcommon.Map) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func getFirstMatchingValue(keys []string, attributes ...pcommon.Map) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}
