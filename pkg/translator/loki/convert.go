// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loki // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/loki"

import (
	"time"

	"github.com/grafana/loki/pkg/push"
	"github.com/prometheus/common/model"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

const (
	hintAttributes = "loki.attribute.labels"
	hintResources  = "loki.resource.labels"
	hintTenant     = "loki.tenant"
	hintFormat     = "loki.format"
)

const (
	formatJSON   string = "json"
	formatLogfmt string = "logfmt"
	formatRaw    string = "raw"
)

const (
	exporterLabel string = "exporter"
	levelLabel    string = "level"
)

const attrSeparator = "."

func convertAttributesAndMerge(logAttrs, resAttrs pcommon.Map, defaultLabelsEnabled map[string]bool) model.LabelSet {
	_ = "STUB: not implemented"
	return *new(model.LabelSet)
}

// get the hint from the log attributes, not from the resource
// the value can be a single resource name to use as label
// or a slice of string values

func getDefaultLabels(resAttrs pcommon.Map, defaultLabelsEnabled map[string]bool) model.LabelSet {
	_ = "STUB: not implemented"
	return *new(model.LabelSet)
}

// Map service.namespace + service.name to job

// Map service.instance.id to instance

func convertAttributesToLabels(attributes pcommon.Map, attrsToSelect pcommon.Value) model.LabelSet {
	_ = "STUB: not implemented"
	return *new(model.LabelSet)
}

func getAttribute(attr string, attributes pcommon.Map) (pcommon.Value, bool) {
	_ = "STUB: not implemented"
	return *new(pcommon.Value), false
}

// couldn't find the attribute under the given name directly
// perhaps it's a nested attribute?

func parseAttributeNames(attrsToSelect pcommon.Value) []string {
	_ = "STUB: not implemented"
	return nil
}

// trying to make the most of bad data

func removeAttributes(attrs pcommon.Map, labels model.LabelSet) { _ = "STUB: not implemented"; return }

func convertLogToJSONEntry(lr plog.LogRecord, res pcommon.Resource, scope pcommon.InstrumentationScope) (*push.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertLogToLogfmtEntry(lr plog.LogRecord, res pcommon.Resource, scope pcommon.InstrumentationScope) (*push.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertLogToLogRawEntry(lr plog.LogRecord) (*push.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertLogToLokiEntry(lr plog.LogRecord, res pcommon.Resource, format string, scope pcommon.InstrumentationScope) (*push.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func timestampFromLogRecord(lr plog.LogRecord) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
