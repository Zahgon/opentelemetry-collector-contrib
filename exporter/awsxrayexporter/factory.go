// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsxrayexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/featuregate"
)

var skipTimestampValidationFeatureGate = featuregate.GlobalRegistry().MustRegister(
	"exporter.awsxray.skiptimestampvalidation",
	featuregate.StageBeta,
	featuregate.WithRegisterDescription("Remove XRay's timestamp validation on first 32 bits of trace ID"),
	featuregate.WithRegisterFromVersion("v0.84.0"))

// NewFactory creates a factory for AWS-Xray exporter.
func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesExporter(ctx context.Context, params exporter.Settings, cfg component.Config) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}
