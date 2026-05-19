// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package statsdreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver"

import (
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/internal/metadata"
)

// reporter struct implements the transport.Reporter interface to give consistent
// observability per Collector metric observability package.
type reporter struct {
	logger           *zap.Logger
	sugaredLogger    *zap.SugaredLogger // Used for generic debug logging
	receiverAttr     attribute.KeyValue
	telemetryBuilder *metadata.TelemetryBuilder
}

var (
	parseSuccessAttr = attribute.String("parse_success", "true")
	parseFailureAttr = attribute.String("parse_success", "false")
)

func newReporter(set receiver.Settings) (*reporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *reporter) OnDebugf(template string, args ...any) { _ = "STUB: not implemented"; return }

func (r *reporter) RecordParseFailure() { _ = "STUB: not implemented"; return }

func (r *reporter) RecordParseSuccess(count int64) { _ = "STUB: not implemented"; return }
