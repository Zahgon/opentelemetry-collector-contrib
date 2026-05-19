// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loki // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/loki"

import (
	"github.com/grafana/loki/pkg/push"
	"github.com/prometheus/common/model"
	"go.opentelemetry.io/collector/pdata/plog"
)

// PushRequestToLogs converts loki push request to logs pipeline data
func PushRequestToLogs(pushRequest *push.PushRequest, keepTimestamp bool) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *

	// Return early if request does not contain any streams
	new(plog.Logs), nil
}

// Return early if stream does not contain any entries

// Get stream labels
// Stream contains labels in string format: `{label1="value1", label2="value2"}`
// Here we parse such a string into labels.Labels

// Convert to model.LabelSet

// Labels started from __ are considered internal and should be ignored

// ConvertEntryToLogRecord converts loki log entry to otlp log record
func ConvertEntryToLogRecord(entry *push.Entry, lr *plog.LogRecord, labelSet model.LabelSet, keepTimestamp bool) {
	_ = "STUB: not implemented"
	return
}
