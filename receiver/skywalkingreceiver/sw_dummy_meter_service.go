// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package skywalkingreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/skywalkingreceiver"

import (
	meter "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

type meterService struct {
	meter.UnimplementedMeterReportServiceServer
}

func (*meterService) Collect(meter.MeterReportService_CollectServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (*meterService) CollectBatch(meter.MeterReportService_CollectBatchServer) error {
	_ = "STUB: not implemented"
	return nil
}
