// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/huaweicloudcesreceiver/internal"

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v1/model"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type MetricData struct {
	MetricName string
	Dimensions []model.MetricsDimension
	Namespace  string
	Unit       string
	Datapoints []model.Datapoint
}

func GetMetricKey(m model.MetricInfoList) string { _ = "STUB: not implemented"; return "" }

func GetDimension(dimensions []model.MetricsDimension, index int) *string {
	_ = "STUB: not implemented"
	return nil
}

func ConvertCESMetricsToOTLP(projectID, regionID, filter string, cesMetrics map[string][]*MetricData) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}
