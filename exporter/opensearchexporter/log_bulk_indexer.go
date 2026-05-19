// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package opensearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opensearchexporter"

import (
	"context"
	"time"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/opensearch-project/opensearch-go/v4/opensearchutil"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

type logBulkIndexer struct {
	bulkAction  string
	pipeline    string
	model       mappingModel
	errs        []error
	bulkIndexer opensearchutil.BulkIndexer
}

func newLogBulkIndexer(bulkAction string, model mappingModel, pipeline string) *logBulkIndexer {
	_ = "STUB: not implemented"
	return nil
}

func (lbi *logBulkIndexer) start(client *opensearchapi.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func (lbi *logBulkIndexer) joinedError() error { _ = "STUB: not implemented"; return nil }

func (lbi *logBulkIndexer) close(ctx context.Context) { _ = "STUB: not implemented"; return }

func (lbi *logBulkIndexer) onIndexerError(_ context.Context, indexerErr error) {
	_ = "STUB: not implemented"
	return
}

func (lbi *logBulkIndexer) appendPermanentError(e error) { _ = "STUB: not implemented"; return }

func (lbi *logBulkIndexer) appendRetryLogError(err error, log plog.Logs) {
	_ = "STUB: not implemented"
	return
}

func (lbi *logBulkIndexer) submit(ctx context.Context, ld plog.Logs, ir *indexResolver, cfg *Config, timestamp time.Time) {
	_ = "STUB: not implemented"
	return
}

func (lbi *logBulkIndexer) processItem(ctx context.Context, indexName string, resource pcommon.Resource, resourceSchemaURL string, scope pcommon.InstrumentationScope, scopeSchemaURL string, logRecord plog.LogRecord) {
	_ = "STUB: not implemented"
	return
}

// Setup error handler. The handler handles the per item response status based on the
// selective ACKing in the bulk response.

func makeLog(resource pcommon.Resource, resourceSchemaURL string, scope pcommon.InstrumentationScope, scopeSchemaURL string, log plog.LogRecord) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

func (lbi *logBulkIndexer) processItemFailure(resp opensearchapi.BulkRespItem, itemErr error, logs plog.Logs) {
	_ = "STUB: not implemented"
	return
}

// Recoverable OpenSearch error

// Non-recoverable OpenSearch error while indexing document

// Encoding error. We didn't even attempt to send the event

func (lbi *logBulkIndexer) newBulkIndexerItem(document []byte, indexName string) opensearchutil.BulkIndexerItem {
	_ = "STUB: not implemented"
	return *new(opensearchutil.BulkIndexerItem)
}

func newLogOpenSearchBulkIndexer(client *opensearchapi.Client, onIndexerError func(context.Context, error), pipeline string) (opensearchutil.BulkIndexer, error) {
	_ = "STUB: not implemented"
	return *new(opensearchutil.BulkIndexer), nil
}
