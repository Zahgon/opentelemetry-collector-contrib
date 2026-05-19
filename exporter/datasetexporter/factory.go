// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package datasetexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datasetexporter"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

// NewFactory created new factory with DataSet exporters.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// castConfig casts it to the Dataset Config struct.
func castConfig(c component.Config) *Config { _ = "STUB: not implemented"; return nil }
