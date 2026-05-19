// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package alibabacloudlogserviceexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter"

import (
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

const (
	slsLogTimeUnixNano   = "timeUnixNano"
	slsLogSeverityNumber = "severityNumber"
	slsLogSeverityText   = "severityText"
	slsLogContent        = "content"
	slsLogAttribute      = "attribute"
	slsLogFlags          = "flags"
	slsLogResource       = "resource"
	slsLogHost           = "host"
	slsLogService        = "service"
	// shortcut for "otlp.instrumentation.library.name" "otlp.instrumentation.library.version"
	slsLogInstrumentationName    = "otlp.name"
	slsLogInstrumentationVersion = "otlp.version"
)

func logDataToLogService(ld plog.Logs) []*sls.Log { _ = "STUB: not implemented"; return nil }

func resourceToLogContents(resource pcommon.Resource) []*sls.LogContent {
	_ = "STUB: not implemented"
	return nil
}

func instrumentationScopeToLogContents(instrumentationScope pcommon.InstrumentationScope) []*sls.LogContent {
	_ = "STUB: not implemented"
	return nil
}

func mapLogRecordToLogService(lr plog.LogRecord,
	resourceContents,
	instrumentationLibraryContents []*sls.LogContent,
) *sls.Log {
	_ = "STUB: not implemented"
	return nil
}

// pre alloc, refine if logContent's len > 16

// convert time nano to time seconds
