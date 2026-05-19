// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !linux

package gopsutilenv // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/gopsutilenv"

import (
	"context"

	"github.com/shirou/gopsutil/v4/common"
)

func ValidateRootPath(rootPath string) error { _ = "STUB: not implemented"; return nil }

func SetGoPsutilEnvVars(_ string) common.EnvMap {
	_ = "STUB: not implemented"
	return *new(common.EnvMap)
}

func GetEnvWithContext(_ context.Context, _, dfault string, _ ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func SetGlobalRootPath(_ string) { _ = "STUB: not implemented"; return }
