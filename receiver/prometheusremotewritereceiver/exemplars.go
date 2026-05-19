// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheusremotewritereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusremotewritereceiver"

import (
	"github.com/prometheus/prometheus/model/labels"
	writev2 "github.com/prometheus/prometheus/prompb/io/prometheus/write/v2"
	promremote "github.com/prometheus/prometheus/storage/remote"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

// collectExemplars extracts Prometheus exemplars from a writev2 request and
// groups them into ExemplarSlices keyed by metric identity.
//
// Exemplars are grouped by a hash composed of:
//   - instrumentation scope name
//   - instrumentation scope version
//   - metric name
//   - metric type
//
// TODO:
//
//	Right now, remote-write 2.0 sends disconnected exemplars without histogram, which requires
//	caching exemplars and associating them later with histogram data points.
//	Once https://github.com/prometheus/prometheus/issues/17857 is resolved, we can optimize this
func collectExemplars(
	req *writev2.Request,
	settings receiver.Settings,
	stats *promremote.WriteResponseStats,
) map[uint64]pmetric.ExemplarSlice {
	_ = "STUB: not implemented"
	return nil
}

func extractScopeFromLabels(settings receiver.Settings, ls labels.Labels) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// setTraceAndSpan extracts trace ID and span ID from exemplar labels
// and sets them on the provided Exemplar.
//
// The function expects hexadecimal-encoded IDs using Prometheus
// exemplar label keys and silently ignores invalid values.
func setTraceAndSpan(exemplar pmetric.Exemplar, labels labels.Labels) {
	_ = "STUB: not implemented"
	return
}

// copyExemplarAttributes copies all labels into the destination attribute map
// except for trace ID and span ID labels, which are handled separately.
//
// The destination map is typically the exemplar's filtered attributes.
func copyExemplarAttributes(dest pcommon.Map, labels labels.Labels) {
	_ = "STUB: not implemented"
	return
}

type exemplarKey struct {
	ScopeName    string
	ScopeVersion string
	MetricName   string
	MetricType   writev2.Metadata_MetricType
}

// sep is a byte that is not valid UTF-8, used as a field separator to prevent
// hash collisions between different field boundary combinations (e.g. "ab"+"c" vs "a"+"bc").
var sep = []byte{0xff}

func (k exemplarKey) hash() uint64 { _ = "STUB: not implemented"; return 0 }
