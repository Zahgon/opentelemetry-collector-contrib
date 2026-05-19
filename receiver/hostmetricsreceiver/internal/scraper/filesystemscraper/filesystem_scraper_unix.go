// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux || darwin || freebsd || openbsd || solaris

package filesystemscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/filesystemscraper"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

const fileSystemStatesLen = 3

func (s *filesystemsScraper) recordFileSystemUsageMetric(now pcommon.Timestamp, deviceUsages []*deviceUsage) {
	_ = "STUB: not implemented"
	return
}

const systemSpecificMetricsLen = 1

func (s *filesystemsScraper) recordSystemSpecificMetrics(now pcommon.Timestamp, deviceUsages []*deviceUsage) {
	_ = "STUB: not implemented"
	return
}
