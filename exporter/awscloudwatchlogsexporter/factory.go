// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

// Package awscloudwatchlogsexporter provides an AWS CloudWatch Logs exporter for the OpenTelemetry collector.
// This package is subject to change and may break configuration settings and behavior.
package awscloudwatchlogsexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awscloudwatchlogsexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

func NewFactory() exporter.Factory { _ = "STUB: not implemented"; return *new(exporter.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// For backwards compatibilitiy, we default to 1 consumer

func createLogsExporter(ctx context.Context, params exporter.Settings, config component.Config) (exporter.Logs, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Logs), nil
}
