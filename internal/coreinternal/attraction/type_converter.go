// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package attraction // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/attraction"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"
)

const (
	stringConversionTarget = "string"
	intConversionTarget    = "int"
	doubleConversionTarget = "double"
)

func convertValue(logger *zap.Logger, key, to string, v pcommon.Value) {
	_ = "STUB: not implemented"
	return
}

// No-op
