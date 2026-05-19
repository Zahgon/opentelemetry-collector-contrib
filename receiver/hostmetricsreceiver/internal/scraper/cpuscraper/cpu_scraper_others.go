// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !linux

package cpuscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/cpuscraper"

import (
	"github.com/shirou/gopsutil/v4/cpu"
	"go.opentelemetry.io/collector/pdata/pcommon"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/cpuscraper/ucal"
)

func (s *cpuScraper) recordCPUTimeStateDataPoints(now pcommon.Timestamp, cpuTime cpu.TimesStat) {
	_ = "STUB: not implemented"
	return
}

func (s *cpuScraper) recordCPUUtilization(now pcommon.Timestamp, cpuUtilization ucal.CPUUtilization) {
	_ = "STUB: not implemented"
	return
}

func (*cpuScraper) getCPUInfo() ([]cpuInfo, error) { _ = "STUB: not implemented"; return nil, nil }
