// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awscloudwatchmetricstreamsencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awscloudwatchmetricstreamsencodingextension"

import (
	"errors"
	"io"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
)

const (
	attributeAWSCloudWatchMetricStreamName = "aws.cloudwatch.metric_stream_name"
	dimensionInstanceID                    = "InstanceId"
	namespaceDelimiter                     = "/"
)

var (
	_ pmetric.Unmarshaler = (*formatJSONUnmarshaler)(nil)
	_ streamUnmarshal     = (*formatJSONUnmarshaler)(nil)
)

var (
	errNoMetricName      = errors.New("cloudwatch metric is missing metric name field")
	errNoMetricNamespace = errors.New("cloudwatch metric is missing namespace field")
	errNoMetricUnit      = errors.New("cloudwatch metric is missing unit field")
	errNoMetricValue     = errors.New("cloudwatch metric is missing value")
)

type formatJSONUnmarshaler struct {
	buildInfo component.BuildInfo
}

func (r *formatJSONUnmarshaler) UnmarshalMetrics(record []byte) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	// Decode as a stream but flush all at once using flush options
	return *new(pmetric.Metrics), nil
}

// we must check for EOF with direct comparison and avoid wrapped EOF that can come from stream itself
//nolint:errorlint

// EOF indicates no metrics were found, return any metrics that's available

func (r *formatJSONUnmarshaler) NewMetricsDecoder(reader io.Reader, options ...encoding.DecoderOption) (encoding.MetricsDecoder, error) {
	_ = "STUB: not implemented"
	return *new(encoding.MetricsDecoder), nil
}

// addMetricToResource adds a new cloudwatchMetric to the resource it belongs to according to resourceKey.
// It then sets the data point for the cloudwatchMetric.
func (*formatJSONUnmarshaler) addMetricToResource(
	byResource map[resourceKey]map[metricKey]pmetric.Metric,
	cwMetric cloudwatchMetric,
) {
	_ = "STUB: not implemented"
	return
}

// Convert percentile to quantile

// createMetrics creates pmetric.Metrics based on the extracted metrics of each resource.
func (r *formatJSONUnmarshaler) createMetrics(
	byResource map[resourceKey]map[metricKey]pmetric.Metric,
) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

// The cloudwatchMetric is the format for the CloudWatch metric stream records.
//
// More details can be found at:
// https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-metric-streams-formats-json.html
type cloudwatchMetric struct {
	// MetricStreamName is the name of the CloudWatch metric stream.
	MetricStreamName string `json:"metric_stream_name"`
	// AccountID is the AWS account ID associated with the metric.
	AccountID string `json:"account_id"`
	// Region is the AWS region for the metric.
	Region string `json:"region"`
	// Namespace is the CloudWatch namespace the metric is in.
	Namespace string `json:"namespace"`
	// MetricName is the name of the metric.
	MetricName string `json:"metric_name"`
	// Dimensions is a map of name/value pairs that help to
	// differentiate a metric.
	Dimensions map[string]string `json:"dimensions"`
	// Timestamp is the milliseconds since epoch for
	// the metric.
	Timestamp int64 `json:"timestamp"`
	// Value is the cloudwatchMetricValue, which has the min, max,
	// sum, and count.
	Value cloudwatchMetricValue `json:"value"`
	// Unit is the unit for the metric.
	//
	// More details can be found at:
	// https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html
	Unit string `json:"unit"`
}

// The cloudwatchMetricValue is the actual values of the CloudWatch metric.
type cloudwatchMetricValue struct {
	isSet bool

	// Max is the highest value observed.
	Max float64
	// Min is the lowest value observed.
	Min float64
	// Sum is the sum of data points collected.
	Sum float64
	// Count is the number of data points.
	Count float64
	// Percentiles contains percentile fields (e.g., p50, p99, p99.9).
	Percentiles map[string]float64
}

func (v *cloudwatchMetricValue) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	// All CloudWatch metric values are float64, so use a typed map
	return nil
}

// Other statistics (TM, WM, TC, TS, PR, IQM) are silently ignored.

// resourceKey stores the metric attributes
// that make a cloudwatchMetric unique to
// a resource
type resourceKey struct {
	metricStreamName string
	namespace        string
	accountID        string
	region           string
}

// metricKey stores the metric attributes
// that make a metric unique within
// a resource
type metricKey struct {
	name string
	unit string
}

// validateMetric validates that the cloudwatch metric has been unmarshalled correctly
func validateMetric(metric cloudwatchMetric) error { _ = "STUB: not implemented"; return nil }

// setResourceAttributes sets attributes on a pcommon.Resource from a cloudwatchMetric.
func setResourceAttributes(rKey resourceKey, resource pcommon.Resource) {
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

// setDataPointAttributes sets attributes on a metric data point from a cloudwatchMetric.
func setDataPointAttributes(metric cloudwatchMetric, dp pmetric.SummaryDataPoint) {
	_ = "STUB: not implemented"
	return
}
