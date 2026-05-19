// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package opensearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opensearchexporter"

import (
	"context"
	"time"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/opensearch-project/opensearch-go/v4/opensearchutil"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type traceBulkIndexer struct {
	bulkAction  string
	pipeline    string
	model       mappingModel
	errs        []error
	bulkIndexer opensearchutil.BulkIndexer
}

func newTraceBulkIndexer(bulkAction string, model mappingModel, pipeline string) *traceBulkIndexer {
	_ = "STUB: not implemented"
	return nil
}

func (tbi *traceBulkIndexer) joinedError() error { _ = "STUB: not implemented"; return nil }

func (tbi *traceBulkIndexer) start(client *opensearchapi.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func (tbi *traceBulkIndexer) close(ctx context.Context) { _ = "STUB: not implemented"; return }

func (tbi *traceBulkIndexer) onIndexerError(_ context.Context, indexerErr error) {
	_ = "STUB: not implemented"
	return
}

func (tbi *traceBulkIndexer) appendPermanentError(e error) { _ = "STUB: not implemented"; return }

func (tbi *traceBulkIndexer) appendRetryTraceError(err error, trace ptrace.Traces) {
	_ = "STUB: not implemented"
	return
}

func (tbi *traceBulkIndexer) submit(ctx context.Context, td ptrace.Traces, ir *indexResolver, cfg *Config, timestamp time.Time) {
	_ = "STUB: not implemented"
	return
}

func (tbi *traceBulkIndexer) processItem(ctx context.Context, indexName string, resource pcommon.Resource, resourceSchemaURL string, scope pcommon.InstrumentationScope, scopeSchemaURL string, span ptrace.Span) {
	_ = "STUB: not implemented"
	return
}

// Setup error handler. The handler handles the per item response status based on the
// selective ACKing in the bulk response.

func makeTrace(resource pcommon.Resource, resourceSchemaURL string, scope pcommon.InstrumentationScope, scopeSchemaURL string, span ptrace.Span) ptrace.Traces {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces)
}

func (tbi *traceBulkIndexer) processItemFailure(resp opensearchapi.BulkRespItem, itemErr error, traces ptrace.Traces) {
	_ = "STUB: not implemented"
	return
}

// Recoverable OpenSearch error

// Non-recoverable OpenSearch error while indexing document

// Encoding error. We didn't even attempt to send the event

// responseAsError converts an opensearchapi.BulkRespItem.Error into an error
func responseAsError(item opensearchapi.BulkRespItem) error { _ = "STUB: not implemented"; return nil }

func attributesToMapString(attributes pcommon.Map) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func shouldRetryEvent(status int) bool { _ = "STUB: not implemented"; return false }

func (tbi *traceBulkIndexer) newBulkIndexerItem(document []byte, indexName string) opensearchutil.BulkIndexerItem {
	_ = "STUB: not implemented"
	return *new(opensearchutil.BulkIndexerItem)
}

func newOpenSearchBulkIndexer(client *opensearchapi.Client, onIndexerError func(context.Context, error), pipeline string) (opensearchutil.BulkIndexer, error) {
	_ = "STUB: not implemented"
	return *new(opensearchutil.BulkIndexer), nil
}
