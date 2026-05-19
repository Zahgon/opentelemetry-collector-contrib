// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tencentcloudlogserviceexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tencentcloudlogserviceexporter"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"

	cls "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tencentcloudlogserviceexporter/internal/proto"
)

const (
	traceIDField = "traceID"
	spanIDField  = "spanID"

	clsLogTimeUnixNano   = "timeUnixNano"
	clsLogSeverityNumber = "severityNumber"
	clsLogSeverityText   = "severityText"
	clsLogContent        = "content"
	clsLogAttribute      = "attribute"
	clsLogFlags          = "flags"
	clsLogResource       = "resource"
	clsLogHost           = "host"
	clsLogService        = "service"
	// shortcut for "otlp.instrumentation.library.name" "otlp.instrumentation.library.version"
	clsLogInstrumentationName    = "otlp.name"
	clsLogInstrumentationVersion = "otlp.version"
)

func convertLogs(ld plog.Logs) []*cls.Log { _ = "STUB: not implemented"; return nil }

func resourceToLogContents(resource pcommon.Resource) []*cls.Log_Content {
	_ = "STUB: not implemented"
	return nil
}

func instrumentationLibraryToLogContents(scope pcommon.InstrumentationScope) []*cls.Log_Content {
	_ = "STUB: not implemented"
	return nil
}

func mapLogRecordToLogService(lr plog.LogRecord,
	resourceContents,
	instrumentationLibraryContents []*cls.Log_Content,
) *cls.Log {
	_ = "STUB: not implemented"
	return nil
}

// pre alloc, refine if logContent's len > 16

// convert time nano to time seconds
