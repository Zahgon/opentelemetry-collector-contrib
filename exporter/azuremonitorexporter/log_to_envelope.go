// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"

import (
	"github.com/microsoft/ApplicationInsights-Go/appinsights/contracts"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

type logPacker struct {
	logger *zap.Logger
	config *Config
}

func (*logPacker) initEnvelope(logRecord plog.LogRecord) (*contracts.Envelope, *contracts.Data) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (packer *logPacker) handleEventData(envelope *contracts.Envelope, data *contracts.Data, logRecord plog.LogRecord) {
	_ = "STUB: not implemented"
	return
}

func (packer *logPacker) handleMessageData(envelope *contracts.Envelope, data *contracts.Data, logRecord plog.LogRecord, resource pcommon.Resource, instrumentationScope pcommon.InstrumentationScope) {
	_ = "STUB: not implemented"
	return
}

func (packer *logPacker) sanitizeAll(envelope *contracts.Envelope, data any) {
	_ = "STUB: not implemented"
	return
}

func (packer *logPacker) LogRecordToEnvelope(logRecord plog.LogRecord, resource pcommon.Resource, instrumentationScope pcommon.InstrumentationScope) *contracts.Envelope {
	_ = "STUB: not implemented"
	return nil
}

func (packer *logPacker) handleExceptionData(envelope *contracts.Envelope, data *contracts.Data, logRecord plog.LogRecord, resource pcommon.Resource, instrumentationScope pcommon.InstrumentationScope) {
	_ = "STUB: not implemented"
	return
}

func (packer *logPacker) sanitize(sanitizeFunc func() []string) { _ = "STUB: not implemented"; return }

func (*logPacker) toAiSeverityLevel(sn plog.SeverityNumber) contracts.SeverityLevel {
	_ = "STUB: not implemented"
	return *new(contracts.SeverityLevel)
}

func newLogPacker(logger *zap.Logger, config *Config) *logPacker {
	_ = "STUB: not implemented"
	return nil
}

func timestampFromLogRecord(lr plog.LogRecord) pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

func hasOneOfKeys(attrMap pcommon.Map, keys ...string) bool {
	_ = "STUB: not implemented"
	return false
}

func isEventData(attrMap pcommon.Map) bool { _ = "STUB: not implemented"; return false }

func isExceptionData(attributes pcommon.Map) bool { _ = "STUB: not implemented"; return false }

func mapIncomingAttributeMapExceptionDetail(attributemap pcommon.Map) *contracts.ExceptionDetails {
	_ = "STUB: not implemented"
	return nil
}
