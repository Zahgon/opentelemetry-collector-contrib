// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunkhecexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter"

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

// perScopeBatcher is a consumer.Logs that rebatches logs by a type found in the scope name: profiling or regular logs.
type perScopeBatcher struct {
	logsEnabled      bool
	profilingEnabled bool
	logger           *zap.Logger
	next             consumer.Logs
}

// Capabilities returns capabilities of the next consumer because perScopeBatcher doesn't mutate data itself.
func (rb *perScopeBatcher) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (rb *perScopeBatcher) ConsumeLogs(ctx context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// if we don't have both types of logs, just call next if enabled

func copyResourceLogs(src, dest plog.ResourceLogs, isProfiling bool) {
	_ = "STUB: not implemented"
	return
}
