// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package opensearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opensearchexporter"

import (
	"errors"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opensearchexporter/internal/pool"
)

var errInvalidTypeForBodyMapMode = errors.New("invalid log record body type for 'bodymap' mapping mode")

type mappingModel interface {
	encodeLog(resource pcommon.Resource,
		scope pcommon.InstrumentationScope,
		schemaURL string,
		record plog.LogRecord) ([]byte, error)
	encodeTrace(resource pcommon.Resource,
		scope pcommon.InstrumentationScope,
		schemaURL string,
		record ptrace.Span) ([]byte, error)
}

type bodyMapMappingModel struct {
	bufferPool *pool.BufferPool
}

func (*bodyMapMappingModel) encodeTrace(
	_ pcommon.Resource,
	_ pcommon.InstrumentationScope,
	_ string,
	_ ptrace.Span,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *bodyMapMappingModel) encodeLog(
	_ pcommon.Resource,
	_ pcommon.InstrumentationScope,
	_ string,
	record plog.LogRecord,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Copy bytes to avoid holding reference to pooled buffer

// encodeModel supports multiple encoding OpenTelemetry signals to multiple schemas.
type encodeModel struct {
	dedup             bool
	dedot             bool
	sso               bool
	flattenAttributes bool
	timestampField    string
	unixTime          bool

	dataset   string
	namespace string
}

func (m *encodeModel) encodeLog(resource pcommon.Resource,
	scope pcommon.InstrumentationScope,
	schemaURL string,
	record plog.LogRecord,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// encodeLogSSO encodes a plog.LogRecord following the Simple Schema for Observability.
// See: https://github.com/opensearch-project/opensearch-catalog/tree/main/docs/schema/observability
func (m *encodeModel) encodeLogSSO(
	resource pcommon.Resource,
	scope pcommon.InstrumentationScope,
	schemaURL string,
	record plog.LogRecord,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// encodeLogDataModel encodes a plog.LogRecord following the Log Data Model.
// See: https://github.com/open-telemetry/oteps/blob/main/text/logs/0097-log-data-model.md
func (m *encodeModel) encodeLogDataModel(resource pcommon.Resource, record plog.LogRecord) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// encodeTrace encodes a ptrace.Span following the Simple Schema For Observability
// See: https://github.com/opensearch-project/opensearch-catalog/tree/main/docs/schema/observability
func (m *encodeModel) encodeTrace(
	resource pcommon.Resource,
	scope pcommon.InstrumentationScope,
	schemaURL string,
	span ptrace.Span,
) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func epochMilliTimestamp(record plog.LogRecord) int64 { _ = "STUB: not implemented"; return 0 }
