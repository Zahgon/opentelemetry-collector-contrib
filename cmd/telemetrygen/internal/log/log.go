// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package log

import (
	"go.uber.org/zap"
)

// CreateLogger creates a logger for use by telemetrygen
func CreateLogger(skipSettingGRPCLogger bool) (*zap.Logger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
