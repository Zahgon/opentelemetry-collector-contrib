// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics

import (
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.uber.org/zap"
)

// Start starts the metric telemetry generator
func Start(cfg *Config) error { _ = "STUB: not implemented"; return nil }

// run executes the test scenario.
func run(c *Config, expF exporterFunc, logger *zap.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

type exporterFunc func() (sdkmetric.Exporter, error)

func exporterFactory(cfg *Config, logger *zap.Logger) exporterFunc {
	_ = "STUB: not implemented"
	return *new(exporterFunc)
}

func createExporter(cfg *Config, logger *zap.Logger) (sdkmetric.Exporter, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Exporter), nil
}

func exemplarsFromConfig(c *Config) []metricdata.Exemplar[int64] {
	_ = "STUB: not implemented"
	return nil
}

// we validated this already during the Validate() function for config

// we validated this already during the Validate() function for config
