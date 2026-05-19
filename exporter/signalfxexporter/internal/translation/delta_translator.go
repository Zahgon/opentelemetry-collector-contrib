// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translation // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/translation"

import (
	sfxpb "github.com/signalfx/com_signalfx_metrics_protobuf/model"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/ttlmap"
)

type deltaTranslator struct {
	prevPts *ttlmap.TTLMap
}

func newDeltaTranslator(ttl int64, done chan struct{}) *deltaTranslator {
	_ = "STUB: not implemented"
	return nil
}

func (t *deltaTranslator) start() { _ = "STUB: not implemented"; return }

func (t *deltaTranslator) translate(pts []*sfxpb.DataPoint, tr *Rule) []*sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}

// only metrics defined in Rule.Mapping get translated

func (t *deltaTranslator) deltaPt(deltaMetricName string, currPt *sfxpb.DataPoint) *sfxpb.DataPoint {
	_ = "STUB: not implemented"
	// check if we have a previous point for this metric + dimensions
	return nil
}

// without proto.Clone here, points' DoubleValue are converted into IntValues, presumably by other translators

// no previous point, so we can't calculate a delta

func (t *deltaTranslator) shutdown() { _ = "STUB: not implemented"; return }

func doubleDeltaPt(currPt, prevPt *sfxpb.DataPoint, deltaMetricName string) *sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}

// assume a reset, emit the current value

func intDeltaPt(currPt, prevPt *sfxpb.DataPoint, deltaMetricName string) *sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}

// assume a reset, emit the current value

var metricTypeGauge = sfxpb.MetricType_GAUGE

func basePt(currPt *sfxpb.DataPoint, deltaMetricName string) *sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}
