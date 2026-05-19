// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cwmetricstream // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver/internal/unmarshaler/cwmetricstream"

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"
)

const (
	TypeStr = "cwmetrics"

	attributeAWSCloudWatchMetricStreamName = "aws.cloudwatch.metric_stream_name"
	dimensionInstanceID                    = "InstanceId"
	namespaceDelimiter                     = "/"
)

var errInvalidRecords = errors.New("record format invalid")

// Unmarshaler for the CloudWatch Metric Stream JSON record format.
//
// More details can be found at:
// https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats-json.html
type Unmarshaler struct {
	logger *zap.Logger

	buildInfo component.BuildInfo
}

var _ pmetric.Unmarshaler = (*Unmarshaler)(nil)

// NewUnmarshaler creates a new instance of the Unmarshaler.
func NewUnmarshaler(logger *zap.Logger, buildInfo component.BuildInfo) *Unmarshaler {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalMetrics deserializes the record in CloudWatch Metric Stream JSON
// format into a pmetric.Metrics, grouping metrics by resource and metric
// name and unit.
func (u Unmarshaler) UnmarshalMetrics(record []byte) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Multiple metrics in each record separated by newline character

// Only process percentile fields (those starting with 'p')

// Extract the percentile value from the field name (e.g., "p95" -> 0.95)

// Skip if we can't parse the percentile value

// Calculate the quantile value (divide by 100 to get a value between 0 and 1)

// Treat this as a non-fatal error, and handle the data below.

// isValid validates that the cWMetric has been unmarshalled correctly.
func (Unmarshaler) isValid(metric cWMetric) bool { _ = "STUB: not implemented"; return false }

// Type of the serialized messages.
func (Unmarshaler) Type() string { _ = "STUB: not implemented"; return "" }

type resourceKey struct {
	metricStreamName string
	namespace        string
	accountID        string
	region           string
}

// setResourceAttributes sets attributes on a pcommon.Resource from a cwMetric.
func setResourceAttributes(key resourceKey, resource pcommon.Resource) {
	_ = "STUB: not implemented"
	return
}

// toServiceAttributes splits the CloudWatch namespace into service namespace/name
// if prepended by AWS/. Otherwise, it returns the CloudWatch namespace as the
// service name with an empty service namespace
func toServiceAttributes(namespace string) (serviceNamespace, serviceName string) {
	_ = "STUB: not implemented"
	return "", ""
}

// setResourceAttributes sets attributes on a metric data point from a cwMetric.
func setDataPointAttributes(m cWMetric, dp pmetric.SummaryDataPoint) {
	_ = "STUB: not implemented"
	return
}
