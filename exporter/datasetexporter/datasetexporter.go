// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasetexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datasetexporter"

import (
	"context"

	"github.com/scalyr/dataset-go/pkg/api/add_events"
	"github.com/scalyr/dataset-go/pkg/client"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"
)

type datasetExporter struct {
	client      *client.DataSetClient
	logger      *zap.Logger
	session     string
	exporterCfg *exporterConfig
	serverHost  string
}

func newDatasetExporter(entity string, config *Config, set exporter.Settings) (*datasetExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *datasetExporter) shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func sendBatch(events []*add_events.EventBundle, client *client.DataSetClient) error {
	_ = "STUB: not implemented"
	return nil
}

func buildKey(prefix, separator, key string, depth int) string {
	_ = "STUB: not implemented"
	return ""
}

func updateWithPrefixedValuesMap(target map[string]any, prefix, separator, suffix string, source map[string]any, depth int) {
	_ = "STUB: not implemented"
	return
}

func updateWithPrefixedValuesArray(target map[string]any, prefix, separator, suffix string, source []any, depth int) {
	_ = "STUB: not implemented"
	return
}

func updateWithPrefixedValues(target map[string]any, prefix, separator, suffix string, source any, depth int) {
	_ = "STUB: not implemented"
	return
}

// now the last value wins
// Should the first value win?

func inferServerHost(
	resource pcommon.Resource,
	attrs map[string]any,
	serverHost string,
) string {
	_ = "STUB: not implemented"
	// first use value from the attribute serverHost
	return ""
}

// then use value from resource attributes - serverHost and host.name
