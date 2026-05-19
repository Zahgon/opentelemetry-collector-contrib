// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuredataexplorerexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuredataexplorerexporter"

import (
	"context"

	"github.com/Azure/azure-kusto-go/azkustodata"
	"github.com/Azure/azure-kusto-go/azkustoingest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

// adxDataProducer uses the ADX client to perform ingestion
type adxDataProducer struct {
	ingestor      azkustoingest.Ingestor     // ingestion for logs, traces and metrics
	ingestOptions []azkustoingest.FileOption // options for the ingestion
	logger        *zap.Logger                // logger for tracing the flow
}

const nextline = "\n"

const (
	// Scope name
	scopename = "scope.name"
	// Scope version
	scopeversion = "scope.version"
)

// given the full metrics, extract each metric, resource attributes and scope attributes. Individual metric mapping is sent on to metricdata mapping
func (e *adxDataProducer) metricsDataPusher(ctx context.Context, metrics pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// since the transform succeeded, using the option for ingestion ingest the data into ADX

func (e *adxDataProducer) ingestData(b []string) error { _ = "STUB: not implemented"; return nil }

func (e *adxDataProducer) logsDataPusher(_ context.Context, logData plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *adxDataProducer) tracesDataPusher(_ context.Context, traceData ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *adxDataProducer) Close(context.Context) error {
	_ = "STUB: not implemented"
	// Close the ingestor and client connections
	return nil
}

// Create an exporter. The exporter instantiates a client , creates the ingestor and then sends data through it

func newExporter(config *Config, logger *zap.Logger, telemetryDataType int, version string) (*adxDataProducer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Expect that this mapping is already existent

// The exporter could be configured to run in either modes. Using managedstreaming or batched queueing

// Fetches the corresponding ingestionRef if the mapping is provided
func getMappingRef(config *Config, telemetryDataType int) azkustoingest.FileOption {
	_ = "STUB: not implemented"
	return *new(azkustoingest.FileOption)
}

func createKcsb(config *Config, version string) *azkustodata.ConnectionStringBuilder {
	_ = "STUB: not implemented"
	return nil
}

// If the user has managed identity done, use it. For System managed identity use the MI as system

// Depending on the table, create separate ingestors
func createManagedStreamingIngestor(config *Config, version, tablename string) (*azkustoingest.Managed, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A queued ingestor in case that is provided as the config option
func createQueuedIngestor(config *Config, version, tablename string) (*azkustoingest.Ingestion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getScopeMap(sc pcommon.InstrumentationScope) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func getTableName(config *Config, telemetrydatatype int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
