// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package yaml provides a YAML file-based lookup source.
package yaml // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/lookupprocessor/internal/source/yaml"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/lookupprocessor/lookupsource"
)

const sourceType = "yaml"

// Config is the configuration for the YAML lookup source.
type Config struct {
	// Path is the path to the YAML file containing key-value mappings.
	// The file should contain a flat map of string keys to values.
	// Required.
	Path string `mapstructure:"path"`
}

// Validate implements lookupsource.SourceConfig.
func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// NewFactory creates a factory for the YAML source.
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
	cfg lookupsource.SourceConfig,
) (lookupsource.Source, error) {
	_ = "STUB: not implemented"
	return *new(lookupsource.Source), nil
}

// no shutdown needed

// yamlSource holds the loaded YAML data.
//
// The file is read once during Start and never reloaded. The RWMutex is
// present to allow a future file-watch/reload mechanism to swap the data
// map safely while concurrent lookups are in progress.
// TODO: support periodic or file-watch-based reload.
type yamlSource struct {
	path string
	mu   sync.RWMutex
	data map[string]any
}

// start loads the YAML file.
func (s *yamlSource) start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// lookup retrieves a value from the loaded YAML data.
func (s *yamlSource) lookup(_ context.Context, key string) (any, bool, error) {
	_ = "STUB: not implemented"
	return *new(any), false, nil
}
