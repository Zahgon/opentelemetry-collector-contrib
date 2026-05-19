// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package pagingscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/pagingscraper"

import (
	"io"
)

const swapsFilePath = "/proc/swaps"

// swaps file column indexes
const (
	nameCol = 0
	// typeCol     = 1
	totalCol = 2
	usedCol  = 3
	// priorityCol = 4

	minimumColCount = usedCol + 1
)

func getPageFileStats() ([]*pageFileStats, error) { _ = "STUB: not implemented"; return nil, nil }

func parseSwapsFile(r io.Reader) ([]*pageFileStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check header headerFields are as expected
