// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ackextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/ackextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

var defaultStorageType = (*component.ID)(nil)

const (
	defaultMaxNumPartition               uint64 = 1_000_000
	defaultMaxNumPendingAcksPerPartition uint64 = 1_000_000
)

// NewFactory creates a factory for ack extension.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createExtension(_ context.Context, _ extension.Settings, cfg component.Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}
