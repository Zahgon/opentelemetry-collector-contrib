// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"

import (
	"context"

	"go.uber.org/zap"
)

// serviceDiscovery runs the discovery loop.
// It writes discovered targets as prometheus file sd format (for now).
type serviceDiscovery struct {
	logger   *zap.Logger
	cfg      Config
	fetcher  *taskFetcher
	filter   *taskFilter
	exporter *taskExporter
}

type serviceDiscoveryOptions struct {
	Logger  *zap.Logger
	Fetcher *taskFetcher // mock server in test, otherwise call new newTaskFetcherFromConfig to use AWS API
}

func newDiscovery(cfg Config, opts serviceDiscoveryOptions) (*serviceDiscovery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// runAndWriteFile writes the output to Config.ResultFile.
func (s *serviceDiscovery) runAndWriteFile(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop on critical error

// Print all the minor errors for debugging, e.g. user config etc.

// We may get 0 targets form some recoverable errors
// e.g. throttled, in that case we keep existing exported file.

// We already printed th error

// As long as we have some targets, export them regardless of errors.
// A better approach might be keep previous targets in memory and do a diff and merge on error.
// For now we just replace entire exported file.

// Encoding and file write error should never happen,
// so we stop extension by returning error.

// NOTE: We assume the folder already exists and does NOT try to create one.

// discover fetch tasks, filter by matching result and export them.
func (s *serviceDiscovery) discover(ctx context.Context) ([]prometheusECSTarget, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
