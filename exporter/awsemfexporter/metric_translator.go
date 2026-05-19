// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsemfexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/cwlogs"
)

const (
	// OTel instrumentation lib name as dimension
	oTellibDimensionKey = "OTelLib"
	defaultNamespace    = "default"

	// DimensionRollupOptions
	zeroAndSingleDimensionRollup = "ZeroAndSingleDimensionRollup"
	singleDimensionRollupOnly    = "SingleDimensionRollupOnly"

	prometheusReceiver        = "prometheus"
	attributeReceiver         = "receiver"
	fieldPrometheusMetricType = "prom_metric_type"

	// metric attributes for AWS EMF, not to be treated as metric labels
	emfStorageResolutionAttribute = "aws.emf.storage_resolution"
)

var fieldPrometheusTypes = map[pmetric.MetricType]string{
	pmetric.MetricTypeEmpty:     "",
	pmetric.MetricTypeGauge:     "gauge",
	pmetric.MetricTypeSum:       "counter",
	pmetric.MetricTypeHistogram: "histogram",
	pmetric.MetricTypeSummary:   "summary",
}

type cWMetrics struct {
	measurements []cWMeasurement
	timestampMs  int64
	fields       map[string]any
}

type cWMetricInfo struct {
	Name              string
	Unit              string
	StorageResolution int
}

type cWMeasurement struct {
	Namespace  string
	Dimensions [][]string
	Metrics    []cWMetricInfo
}

type cWMetricStats struct {
	Max   float64
	Min   float64
	Count uint64
	Sum   float64
}

// The SampleCount of CloudWatch metrics will be calculated by the sum of the 'Counts' array.
// The 'Count' field should be same as the sum of the 'Counts' array and will be ignored in CloudWatch.
type cWMetricHistogram struct {
	Values []float64
	Counts []float64
	Max    float64
	Min    float64
	Count  uint64
	Sum    float64
}

type groupedMetricMetadata struct {
	namespace                  string
	timestampMs                int64
	logGroup                   string
	logStream                  string
	metricDataType             pmetric.MetricType
	batchIndex                 int
	retainInitialValueForDelta bool
}

// cWMetricMetadata represents the metadata associated with a given CloudWatch metric
type cWMetricMetadata struct {
	groupedMetricMetadata
	instrumentationScopeName string
	receiver                 string
}

type metricTranslator struct {
	metricDescriptor map[string]MetricDescriptor
	calculators      *emfCalculators
}

func newMetricTranslator(config Config) metricTranslator {
	_ = "STUB: not implemented"
	return *new(metricTranslator)
}

func (mt metricTranslator) Shutdown() error { _ = "STUB: not implemented"; return nil }

// translateOTelToGroupedMetric converts OT metrics to Grouped Metric format.
func (mt metricTranslator) translateOTelToGroupedMetric(rm pmetric.ResourceMetrics, groupedMetrics map[any]*groupedMetric, config *Config) error {
	_ = "STUB: not implemented"
	return nil
}

// translateGroupedMetricToCWMetric converts Grouped Metric format to CloudWatch Metric format.
func translateGroupedMetricToCWMetric(groupedMetric *groupedMetric, config *Config) *cWMetrics {
	_ = "STUB: not implemented"
	return nil
}

// Add labels to fields

// Add metrics to fields

// If there are no metric declarations defined, translate grouped metric
// into the corresponding CW Measurement

// If metric declarations are defined, filter grouped metric's metrics using
// metric declarations and translate into the corresponding list of CW Measurements

// groupedMetricToCWMeasurement creates a single CW Measurement from a grouped metric.
func groupedMetricToCWMeasurement(groupedMetric *groupedMetric, config *Config) cWMeasurement {
	_ = "STUB: not implemented"
	return *new(cWMeasurement)
}

// Create a dimension set containing list of label names

// Apply single/zero dimension rollup to labels

// Perform duplication check for edge case with a single label and single dimension roll-up

// Remove duplicated dimension set before adding on rolled-up dimensions

// Add on rolled-up dimensions

// groupedMetricToCWMeasurementsWithFilters filters the grouped metric using the given list of metric
// declarations and returns the corresponding list of CW Measurements.
func groupedMetricToCWMeasurementsWithFilters(groupedMetric *groupedMetric, config *Config) (cWMeasurements []cWMeasurement) {
	_ = "STUB: not implemented"
	return nil
}

// Filter metric declarations by labels

// If the whole batch of metrics don't match any metric declarations, drop them

// Group metrics by matched metric declarations

// Filter metric declarations by metric name

// Apply single/zero dimension rollup to labels

// Translate each group into a CW Measurement

// Extract dimensions from matched metric declarations

// De-duplicate dimensions

// Export metrics only with non-empty dimensions list

// translateCWMetricToEMF converts CloudWatch Metric format to EMF.
func translateCWMetricToEMF(cWMetric *cWMetrics, config *Config) (*cwlogs.Event, error) {
	_ = "STUB: not implemented"
	// convert CWMetric into map format for compatible with PLE input
	return nil, nil
}

// restore the json objects that are stored as string in attributes

// Create EMF metrics if there are measurements
// https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Embedded_Metric_Format_Specification.html#CloudWatch_Embedded_Metric_Format_Specification_structure

/* 	EMF V1
	"Version": "1",
	"_aws": {
		"CloudWatchMetrics": [
		{
			"Namespace": "ECS",
			"Dimensions": [ ["ClusterName"] ],
			"Metrics": [{"Name": "memcached_commands_total"}]
		}
		],
		"Timestamp": 1668387032641
  	}
*/

/* 	EMF V0
	{
		"Version": "0",
		"CloudWatchMetrics": [
		{
			"Namespace": "ECS",
			"Dimensions": [ ["ClusterName"] ],
			"Metrics": [{"Name": "memcached_commands_total"}]
		}
		],
		"Timestamp": "1668387032641"
  	}
*/

// Utility function that converts from groupedMetric to a cloudwatch event
func translateGroupedMetricToEmf(groupedMetric *groupedMetric, config *Config, defaultLogStream string) (*cwlogs.Event, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func filterAWSEMFAttributes(labels map[string]string) map[string]string {
	_ = "STUB: not implemented"
	// remove any labels that are attributes specific to AWS EMF Exporter
	return nil
}
