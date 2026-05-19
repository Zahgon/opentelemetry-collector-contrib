// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package noop provides a no-operation lookup source for testing.
package noop // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/lookupprocessor/internal/source/noop"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/lookupprocessor/lookupsource"
)

const sourceType = "noop"

type Config struct{}

func (*Config) Validate() error { _ = "STUB: not implemented"; return nil }

func NewFactory() lookupsource.SourceFactory {
	_ = "STUB: not implemented"
	return *new(lookupsource.SourceFactory)
}

func createDefaultConfig() lookupsource.SourceConfig {
	_ = "STUB: not implemented"
	return *new(lookupsource.SourceConfig)
}

func createSource(
	_ context.Context,
	_ lookupsource.CreateSettings,
	_ lookupsource.SourceConfig,
) (lookupsource.Source, error) {
	_ = "STUB: not implemented"
	return *new(lookupsource.Source), nil
}

// no start needed
// no shutdown needed

// noopLookup always returns not found.
func noopLookup(_ context.Context, _ string) (any, bool, error) {
	_ = "STUB: not implemented"
	return *new(any), false, nil
}
