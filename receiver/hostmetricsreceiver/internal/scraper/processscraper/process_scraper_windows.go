// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package processscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper"

import (
	"context"
	"regexp"

	"github.com/shirou/gopsutil/v4/cpu"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper/ucal"
)

func (s *processScraper) recordCPUTimeMetric(now pcommon.Timestamp, cpuTime *cpu.TimesStat) {
	_ = "STUB: not implemented"
	return
}

func (s *processScraper) recordCPUUtilization(now pcommon.Timestamp, cpuUtilization ucal.CPUUtilization) {
	_ = "STUB: not implemented"
	return
}

func getProcessName(_ context.Context, _ processHandle, exePath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getProcessCgroup(_ context.Context, _ processHandle) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getProcessExecutable(ctx context.Context, proc processHandle) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// matches the first argument before an unquoted space or slash
var cmdRegex = regexp.MustCompile(`^((?:[^"]*?"[^"]*?")*?[^"]*?)(?:[ /]|$)`)

func getProcessCommand(ctx context.Context, proc processHandle) (*commandMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
