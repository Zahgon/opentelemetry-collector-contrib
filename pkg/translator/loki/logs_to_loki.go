// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loki // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/loki"

import (
	"github.com/grafana/loki/pkg/push"
	"github.com/prometheus/common/model"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

type PushRequest struct {
	*push.PushRequest
	Report *PushReport
}

// PushReport contains the summary for the outcome of a LogsToLoki operation
type PushReport struct {
	Errors       []error
	NumSubmitted int
	NumDropped   int
}

const (
	levelAttributeName = "level"
)

// LogsToLokiRequests converts a Logs pipeline data into Loki PushRequests grouped
// by tenant. The tenant value is inferred from the `loki.tenant` resource or log
// attribute hint. If the `loki.tenant` attribute is present in both resource or
// log attributes, then the resource attribute takes precedence.
// Labels for each record are inferred based on the hints "loki.attribute.labels"
// and "loki.resource.labels". Each hint might contain a comma-separated list of
// attributes (resource or record) that should be promoted to a Loki label. Those
// attributes are removed from the body as a result, otherwise they would be shown
// in duplicity in Loki.
// PushStreams are created based on the labels: all records containing the same
// set of labels are part of the same stream. All streams are then packed within
// the resulting PushRequest.
// When this function isn't able to marshal a log record, the log record is dropped
// and processing continues, so that the caller can decide to either skip the entire
// batch or send only the data that could be parsed. The caller can use the PushReport
// to make this decision, as it includes all of the errors that were encountered,
// as well as the number of items dropped and submitted.
func LogsToLokiRequests(ld plog.Logs, defaultLabelsEnabled map[string]bool) map[string]PushRequest {
	_ = "STUB: not implemented"
	return nil
}

// Couldn't convert so dropping log.

// create the stream name based on the labels

// PushEntry is Loki log entry enriched with labels
type PushEntry struct {
	Entry  *push.Entry
	Labels model.LabelSet
}

// LogToLokiEntry converts LogRecord into Loki log entry enriched with normalized labels
func LogToLokiEntry(lr plog.LogRecord, rl pcommon.Resource, scope pcommon.InstrumentationScope, defaultLabelsEnabled map[string]bool) (*PushEntry, error) {
	_ = "STUB: not implemented"
	// we may remove attributes, so change only our version
	return nil, nil
}

// similarly, we may remove attributes, so we make a copy and change our version

// adds level attribute from log.severityNumber

// remove the attributes that were promoted to labels

// Loki doesn't support dots in label names
// labelName is normalized label name to follow Prometheus label names standard

func getFormatFromFormatHint(logAttr, resourceAttr pcommon.Map) string {
	_ = "STUB: not implemented"
	return ""
}

// GetTenantFromTenantHint extract an attribute based on the tenant hint.
// it looks up for the attribute first in resource attributes and fallbacks to
// record attributes if it is not found.
func GetTenantFromTenantHint(logAttr, resourceAttr pcommon.Map) string {
	_ = "STUB: not implemented"
	return ""
}

type pushRequestGroup struct {
	streams map[string]*push.Stream
	report  *PushReport
}

func addLogLevelAttributeAndHint(log plog.LogRecord) { _ = "STUB: not implemented"; return }

func addHint(log plog.LogRecord) { _ = "STUB: not implemented"; return }

var severityNumberToLevel = map[string]string{
	plog.SeverityNumberUnspecified.String(): "UNSPECIFIED",
	plog.SeverityNumberTrace.String():       "TRACE",
	plog.SeverityNumberTrace2.String():      "TRACE2",
	plog.SeverityNumberTrace3.String():      "TRACE3",
	plog.SeverityNumberTrace4.String():      "TRACE4",
	plog.SeverityNumberDebug.String():       "DEBUG",
	plog.SeverityNumberDebug2.String():      "DEBUG2",
	plog.SeverityNumberDebug3.String():      "DEBUG3",
	plog.SeverityNumberDebug4.String():      "DEBUG4",
	plog.SeverityNumberInfo.String():        "INFO",
	plog.SeverityNumberInfo2.String():       "INFO2",
	plog.SeverityNumberInfo3.String():       "INFO3",
	plog.SeverityNumberInfo4.String():       "INFO4",
	plog.SeverityNumberWarn.String():        "WARN",
	plog.SeverityNumberWarn2.String():       "WARN2",
	plog.SeverityNumberWarn3.String():       "WARN3",
	plog.SeverityNumberWarn4.String():       "WARN4",
	plog.SeverityNumberError.String():       "ERROR",
	plog.SeverityNumberError2.String():      "ERROR2",
	plog.SeverityNumberError3.String():      "ERROR3",
	plog.SeverityNumberError4.String():      "ERROR4",
	plog.SeverityNumberFatal.String():       "FATAL",
	plog.SeverityNumberFatal2.String():      "FATAL2",
	plog.SeverityNumberFatal3.String():      "FATAL3",
	plog.SeverityNumberFatal4.String():      "FATAL4",
}
