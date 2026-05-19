// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package gopsutilenv // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/gopsutilenv"

import (
	"context"

	"github.com/shirou/gopsutil/v4/common"
)

var gopsutilEnvVars = map[common.EnvKeyType]string{
	common.HostProcEnvKey: "/proc",
	common.HostSysEnvKey:  "/sys",
	common.HostEtcEnvKey:  "/etc",
	common.HostVarEnvKey:  "/var",
	common.HostRunEnvKey:  "/run",
	common.HostDevEnvKey:  "/dev",
}

// This exists to validate that different components that use this module do not
// have inconsistent root_path configurations. The root_path is passed down to gopsutil
// through context, so it must be consistent across the process.
var globalRootPath string

func ValidateRootPath(rootPath string) error { _ = "STUB: not implemented"; return nil }

func SetGoPsutilEnvVars(rootPath string) common.EnvMap {
	_ = "STUB: not implemented"
	return *new(common.EnvMap)
}

// don't override if existing env var is set

// SetGlobalRootPath mainly used for unit tests
func SetGlobalRootPath(rootPath string) { _ = "STUB: not implemented"; return }

// copied from gopsutil:
// GetEnvWithContext retrieves the environment variable key. If it does not exist it returns the default.
// The context may optionally contain a map superseding os.EnvKey.
func GetEnvWithContext(ctx context.Context, key, dfault string, combineWith ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func combine(value string, combineWith []string) string { _ = "STUB: not implemented"; return "" }
