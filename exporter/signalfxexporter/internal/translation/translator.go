// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translation // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/translation"

import (
	sfxpb "github.com/signalfx/com_signalfx_metrics_protobuf/model"
	"go.uber.org/zap"
)

// Action is the enum to capture actions to perform on metrics.
type Action string

const (
	// ActionRenameMetrics renames metrics using Rule.Mapping.
	ActionRenameMetrics Action = "rename_metrics"

	// ActionMultiplyInt scales integer metrics by multiplying their values using
	// Rule.ScaleFactorsInt key/values as metric_name/multiplying_factor
	ActionMultiplyInt Action = "multiply_int"

	// ActionDivideInt scales integer metric by dividing their values using
	// Rule.ScaleFactorsInt key/values as metric_name/divisor
	ActionDivideInt Action = "divide_int"

	// ActionMultiplyFloat scales integer metric by multiplying their values using
	// Rule.ScaleFactorsFloat key/values as metric_name/multiplying_factor
	// This rule can only be applied to metrics that are a float value
	ActionMultiplyFloat Action = "multiply_float"

	// ActionConvertValues converts metric values from int to float or float to int
	// Rule.TypesMapping key/values as metric_name/new_type.
	ActionConvertValues Action = "convert_values"

	// ActionCopyMetrics copies metrics using Rule.Mapping.
	// Rule.DimensionKey and Rule.DimensionValues can be used to filter datapoints that must be copied,
	// if these fields are set, only metrics having a dimension with key == Rule.DimensionKey and
	// value in Rule.DimensionValues will be copied.
	ActionCopyMetrics Action = "copy_metrics"

	// ActionSplitMetric splits a metric with Rule.MetricName into multiple metrics
	// based on a dimension specified in Rule.DimensionKey.
	// Rule.Mapping represents "dimension value" -> "new metric name" for the translation.
	// For example, having the following translation rule:
	//   - action: split_metric
	// 	   metric_name: k8s.pod.network.io
	//     dimension_key: direction
	//     mapping:
	//       receive: pod_network_receive_bytes_total
	//       transmit: pod_network_transmit_bytes_total
	// The following translations will be performed:
	// k8s.pod.network.io{direction="receive"} -> pod_network_receive_bytes_total{}
	// k8s.pod.network.io{direction="transmit"} -> pod_network_transmit_bytes_total{}
	ActionSplitMetric Action = "split_metric"

	// ActionAggregateMetric aggregates metrics excluding dimensions set in tr.WithoutDimensions.
	// This method is equivalent of "without" clause in Prometheus aggregation:
	// https://prometheus.io/docs/prometheus/latest/querying/operators/#aggregation-operators
	// It takes datapoints with name tr.MetricName and aggregates them to a smaller set keeping the same name.
	// It drops the dimensions provided in tr.WithoutDimensions and keeps others as is.
	// tr.AggregationMethod is used to specify a method to aggregate the values.
	// For example, having the following translation rule:
	// - action: aggregate_metric
	//   metric_name: machine_cpu_cores
	//   aggregation_method: count
	//   without_dimensions:
	//   - cpu
	// The following translations will be performed:
	// Original datapoints:
	//   machine_cpu_cores{cpu="cpu1",host="host1"} 0.22
	//   machine_cpu_cores{cpu="cpu2",host="host1"} 0.11
	//   machine_cpu_cores{cpu="cpu1",host="host2"} 0.33
	// Transformed datapoints:
	//   machine_cpu_cores{host="host1"} 2
	//   machine_cpu_cores{host="host2"} 1
	ActionAggregateMetric Action = "aggregate_metric"

	// ActionCalculateNewMetric calculates a new metric based on two existing metrics.
	// It takes two operand metrics, an operator, and a metric name and produces a new metric with the given
	// metric name, but with the attributes of the first operand metric.
	// For example, for the following translation rule:
	// - action: calculate_new_metric
	//  metric_name: memory.utilization
	//  operand1_metric: memory.used
	//  operand2_metric: memory.total
	//  operator: /
	// the integer value of the 'memory.used' metric will be divided by the integer value of 'memory.total'. The
	// result will be a new float metric with the name 'memory.utilization' and the value of the quotient. The
	// new metric will also get any attributes of the 'memory.used' metric except for its value and metric name.
	// Currently only integer inputs are handled and only division is supported.
	ActionCalculateNewMetric Action = "calculate_new_metric"

	// ActionDropMetrics drops datapoints with metric name defined in "metric_names".
	ActionDropMetrics Action = "drop_metrics"

	// ActionDeltaMetric creates a new delta (cumulative) metric from an existing non-cumulative int or double
	// metric. It takes mappings of names of the existing metrics to the names of the new, delta metrics to be
	// created. All dimensions will be preserved.
	ActionDeltaMetric Action = "delta_metric"
)

type MetricOperator string

const (
	// MetricOperatorDivision is the MetricOperator division.
	MetricOperatorDivision MetricOperator = "/"
)

// MetricValueType is the enum to capture valid metric value types that can be converted
type MetricValueType string

const (
	// MetricValueTypeInt represents integer metric value type
	MetricValueTypeInt MetricValueType = "int"
	// MetricValueTypeDouble represents double metric value type
	MetricValueTypeDouble MetricValueType = "double"
)

// AggregationMethod is the enum used to capture aggregation method
type AggregationMethod string

// Values for enum AggregationMethodCount.
const (
	AggregationMethodCount AggregationMethod = "count"
	AggregationMethodAvg   AggregationMethod = "avg"
	AggregationMethodSum   AggregationMethod = "sum"
)

type Rule struct {
	// Action specifies the translation action to be applied on metrics.
	// This is a required field.
	Action Action `mapstructure:"action"`

	// Mapping specifies key/value mapping that is used by rename_dimension_keys,
	// rename_metrics, copy_metrics, and split_metric actions.
	Mapping map[string]string `mapstructure:"mapping"`

	// ScaleFactorsInt is used by multiply_int and divide_int action to scale
	// integer metric values, key/value format: metric_name/scale_factor
	ScaleFactorsInt map[string]int64 `mapstructure:"scale_factors_int"`

	// ScaleFactorsInt is used by multiply_float action to scale
	// float metric values, key/value format: metric_name/scale_factor
	ScaleFactorsFloat map[string]float64 `mapstructure:"scale_factors_float"`

	// MetricName is used by "split_metric" translation rule to specify a name
	// of a metric that will be split.
	MetricName string `mapstructure:"metric_name"`
	// DimensionKey is used by "split_metric" translation rule action to specify dimension key
	// that will be used to translate the metric datapoints. Datapoints that don't have
	// the specified dimension key will not be translated.
	// DimensionKey is also used by "copy_metrics" for filtering.
	DimensionKey string `mapstructure:"dimension_key"`

	// DimensionValues is used by "copy_metrics" to filter out datapoints with dimensions values
	// not matching values set in this field
	DimensionValues map[string]bool `mapstructure:"dimension_values"`

	// TypesMapping is represents metric_name/metric_type key/value pairs,
	// used by ActionConvertValues.
	TypesMapping map[string]MetricValueType `mapstructure:"types_mapping"`

	// AggregationMethod specifies method used by "aggregate_metric" translation rule
	AggregationMethod AggregationMethod `mapstructure:"aggregation_method"`

	// WithoutDimensions used by "aggregate_metric" translation rule to specify dimensions to be
	// excluded by aggregation.
	WithoutDimensions []string `mapstructure:"without_dimensions"`

	// AddDimensions used by "rename_metrics" translation rule to add dimensions that are necessary for
	// existing SFx content for desired metric name
	AddDimensions map[string]string `mapstructure:"add_dimensions"`

	// CopyDimensions used by "rename_metrics" translation rule to copy dimensions that are necessary for
	// existing SFx content for desired metric name.  This will duplicate the dimension value and isn't a rename.
	CopyDimensions map[string]string `mapstructure:"copy_dimensions"`

	// MetricNames is used by "drop_metrics" translation rule.
	MetricNames map[string]bool `mapstructure:"metric_names"`

	Operand1Metric string         `mapstructure:"operand1_metric"`
	Operand2Metric string         `mapstructure:"operand2_metric"`
	Operator       MetricOperator `mapstructure:"operator"`
}

type MetricTranslator struct {
	rules []Rule

	deltaTranslator *deltaTranslator
}

func NewMetricTranslator(rules []Rule, ttl int64, done chan struct{}) (*MetricTranslator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateTranslationRules(rules []Rule) error { _ = "STUB: not implemented"; return nil }

func (mp *MetricTranslator) Start() { _ = "STUB: not implemented"; return }

// TranslateDataPoints transforms datapoints to a format compatible with signalfx backend
// sfxDataPoints represents one metric converted to signalfx protobuf datapoints
func (mp *MetricTranslator) TranslateDataPoints(logger *zap.Logger, sfxDataPoints []*sfxpb.DataPoint) []*sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: Based on the usage of TranslateDataPoints we can assume that the datapoints batch []*sfxpb.DataPoint
// represents only one metric and all the datapoints can be aggregated together.

// This slice can contain additional datapoints from a different metric
// for example copied in a translation step before

func (mp *MetricTranslator) Shutdown() { _ = "STUB: not implemented"; return }

func calcNewMetricInputPairs(processedDataPoints []*sfxpb.DataPoint, tr *Rule) [][2]*sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}

func dimensionsEqual(d1, d2 []*sfxpb.Dimension) bool { _ = "STUB: not implemented"; return false }

// avoid allocating a map

func calculateNewMetric(
	logger *zap.Logger,
	operand1 *sfxpb.DataPoint,
	operand2 *sfxpb.DataPoint,
	tr *Rule,
) *sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}

// We can get here if, for example, in the denominator we get multiple
// datapoints that have the same counter value, which will yield a delta of
// zero.

// only supporting divide operator for now

func ptToFloatVal(pt *sfxpb.DataPoint) *float64 { _ = "STUB: not implemented"; return nil }

// aggregateDatapoints aggregates datapoints assuming that they have
// the same Timestamp, MetricType, Metric and Source fields.
func aggregateDatapoints(
	dps []*sfxpb.DataPoint,
	withoutDimensions []string,
	aggregation AggregationMethod,
) []*sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}

// group datapoints by dimension values

// set slice capacity to the possible maximum = len(dps)-i to avoid reallocations

// Get aggregated results

// stringifyDimensions turns the passed-in `dimensions` into a string while
// ignoring the passed-in `exclusions`. The result has the following form:
// dim1:val1//dim2:val2. Order is deterministic so this function can be used to
// generate map keys.
func stringifyDimensions(dimensions []*sfxpb.Dimension, exclusions []string) string {
	_ = "STUB: not implemented"
	return ""
}

// filterDimensions returns list of dimension excluding withoutDimensions
func filterDimensions(dimensions []*sfxpb.Dimension, withoutDimensions []string) []*sfxpb.Dimension {
	_ = "STUB: not implemented"
	return nil
}

// splitMetric renames a metric with "dimension key" == dimensionKey to mapping["dimension value"],
// datapoint not changed if not dimension found equal to dimensionKey:mapping->key.
func splitMetric(dp *sfxpb.DataPoint, dimensionKey string, mapping map[string]string) {
	_ = "STUB: not implemented"
	return
}

// The dimension value matches the mapping, proceeding

// The dimension value doesn't match the mapping, keep the datapoint as is

// No dimension key found for the specified dimensionKey, keep the datapoint as is

func convertMetricValue(logger *zap.Logger, dp *sfxpb.DataPoint, newType MetricValueType) {
	_ = "STUB: not implemented"
	return
}

func copyMetric(tr *Rule, dp *sfxpb.DataPoint, newMetricName string) *sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}
