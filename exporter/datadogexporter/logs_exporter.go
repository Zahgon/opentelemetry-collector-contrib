// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !aix

package datadogexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter"

import (
	"context"

	"github.com/DataDog/datadog-agent/comp/otelcol/logsagentpipeline"
	"github.com/DataDog/datadog-agent/comp/otelcol/otlp/components/exporter/logsagentexporter"
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes"
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes/source"
	"go.opentelemetry.io/collector/exporter"

	datadogconfig "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/datadog/config"
)

const (
	// logSourceName specifies the Datadog source tag value to be added to logs sent from the Datadog exporter.
	logSourceName = "otlp_log_ingestion"
	// otelSource specifies a source to be added to all logs sent from the Datadog exporter. The tag has key `otel_source` and the value specified on this constant.
	otelSource = "datadog_exporter"
)

// newLogsAgentExporter creates new instances of the logs agent and the logs agent exporter
func newLogsAgentExporter(
	ctx context.Context,
	params exporter.Settings,
	cfg *datadogconfig.Config,
	sourceProvider source.Provider,
	_ *attributes.GatewayUsage,
) (logsagentpipeline.LogsAgent, *logsagentexporter.Exporter, error) {
	_ = "STUB: not implemented"
	return *new(logsagentpipeline.LogsAgent), nil, nil
}
