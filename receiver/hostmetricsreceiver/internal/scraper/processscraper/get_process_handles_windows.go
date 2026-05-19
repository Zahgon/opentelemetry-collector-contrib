// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package processscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper"

import (
	"context"
)

func getGopsutilProcessHandles(ctx context.Context) (processHandles, error) {
	_ = "STUB: not implemented"
	return *new(processHandles), nil
}

// Ignoring any errors here to keep same behavior as the legacy implementation
// based on the `process.ProcessesWithContext` from the `gopsutil` package.
