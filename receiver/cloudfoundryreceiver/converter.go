// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cloudfoundryreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudfoundryreceiver"

import (
	"time"

	"code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

const (
	attributeNamePrefix = "org.cloudfoundry."
)

var ResourceAttributesKeys = []string{
	"index",
	"ip",
	"deployment",
	"id",
	"job",
	"product",
	"instance_group",
	"instance_id",
	"origin",
	"system_domain",
	"source_id",
	"source_type",
	"process_type",
	"process_id",
	"process_instance_id",
}

func convertEnvelopeToMetrics(envelope *loggregator_v2.Envelope, metricSlice pmetric.MetricSlice, startTime time.Time) {
	_ = "STUB: not implemented"
	return
}

func convertEnvelopeToLogs(envelope *loggregator_v2.Envelope, logSlice plog.LogRecordSlice, startTime time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

//exhaustive:enforce

func copyEnvelopeAttributes(attributes pcommon.Map, envelope *loggregator_v2.Envelope) {
	_ = "STUB: not implemented"
	return
}

func getEnvelopeDataAttributes(envelope *loggregator_v2.Envelope) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func getEnvelopeResourceAttributes(envelope *loggregator_v2.Envelope) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}
