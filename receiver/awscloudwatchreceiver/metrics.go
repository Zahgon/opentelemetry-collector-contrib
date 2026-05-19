// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awscloudwatchreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscloudwatchreceiver"

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

const (
	// maxGetMetricDataQueries is the AWS limit per GetMetricData request.
	maxGetMetricDataQueries = 500

	// statsPerMetric is the number of CloudWatch statistics we request per metric (Sum, SampleCount, Minimum, Maximum).
	statsPerMetric = 4

	// stat index constants within a metric's sub-queries
	statIdxSum   = 0
	statIdxCount = 1
	statIdxMin   = 2
	statIdxMax   = 3
)

type cloudWatchMetricsScraper struct {
	settings           receiver.Settings
	cfg                *Config
	period             time.Duration
	delay              time.Duration
	collectionInterval time.Duration
	metrics            []MetricQuery
	discovery          *MetricsDiscoveryConfig
	client             metricsClient
}

type metricsClient interface {
	ListMetrics(ctx context.Context, params *cloudwatch.ListMetricsInput, optFns ...func(*cloudwatch.Options)) (*cloudwatch.ListMetricsOutput, error)
	GetMetricData(ctx context.Context, params *cloudwatch.GetMetricDataInput, optFns ...func(*cloudwatch.Options)) (*cloudwatch.GetMetricDataOutput, error)
}

func newCloudWatchMetricsScraper(cfg *Config, settings receiver.Settings) *cloudWatchMetricsScraper {
	_ = "STUB: not implemented"
	return nil
}

func (s *cloudWatchMetricsScraper) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *cloudWatchMetricsScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Shift endTime back by delay to account for CloudWatch metric publication latency.
// Query exactly one collectionInterval worth of data so consecutive scrapes do not overlap.

// Build the largest batch whose total sub-query count stays within the API limit.

// alignTimeToPeriod rounds t down to the nearest period boundary (in seconds from Unix epoch).
// Per GetMetricData docs, aligning StartTime and EndTime to the metric's Period improves
// performance: https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html
func alignTimeToPeriod(t time.Time, periodSec int64) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// listMetrics discovers metrics via ListMetrics API, respecting discovery config (namespace, metric name, limit).
func (s *cloudWatchMetricsScraper) listMetrics(ctx context.Context) ([]MetricQuery, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dimensionsToMap(dims []types.Dimension) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// metricStats are the four CloudWatch statistics fetched for every metric to build a Summary.
var metricStats = [statsPerMetric]string{"Sum", "SampleCount", "Minimum", "Maximum"}

// numSubQueries returns the number of GetMetricData sub-queries needed for one metric.
// Metrics without an explicit Stats list use all four summary statistics.
func numSubQueries(q MetricQuery) int { _ = "STUB: not implemented"; return 0 }

// pollBatch runs GetMetricData for a batch of metrics and returns the converted pdata.Metrics.
// Each metric generates four sub-queries (Sum, SampleCount, Minimum, Maximum) so that the results
// can be combined into an OpenTelemetry Summary metric aligned with the CloudWatch Metric Streams
// OpenTelemetry 1.0.0 format.
// It follows pagination via NextToken to collect all data points for the requested time window.
func (s *cloudWatchMetricsScraper) pollBatch(ctx context.Context, batch []MetricQuery, startTime, endTime time.Time) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func dimensionsFromMap(d map[string]string) []types.Dimension {
	_ = "STUB: not implemented"
	return nil
}

// setResourceAttributes sets the resource-level attributes that identify the AWS source.
// Aligned with the CloudWatch Metric Streams OpenTelemetry 1.0.0 format: only cloud.provider
// and cloud.region are set on the resource; namespace and dimensions go on data points.
func (s *cloudWatchMetricsScraper) setResourceAttributes(resource pcommon.Resource) {
	_ = "STUB: not implemented"
	return
}

// parseQueryID parses a sub-query ID of the form "q{metricIdx}_{statIdx}".
func parseQueryID(id string) (metricIdx, statIdx int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// convertGetMetricDataToPdata converts GetMetricData results to pdata.Metrics.
//
// When a MetricQuery has no Stats (empty), all four statistics are fetched and combined into an
// OpenTelemetry Summary metric, aligned with the CloudWatch Metric Streams OpenTelemetry 1.0.0 format:
//   - Metric name: "amazonaws.com/{Namespace}/{MetricName}".
//   - Metric type: Summary with count (SampleCount), sum (Sum), min quantile (Minimum), max quantile (Maximum).
//   - Data point attributes: Namespace (string), MetricName (string), Dimensions (kvlist).
//
// When a MetricQuery has explicit Stats, each selected statistic is emitted as a Gauge data point on the
// same metric, with an additional "stat" attribute identifying the statistic.
func (s *cloudWatchMetricsScraper) convertGetMetricDataToPdata(results []types.MetricDataResult, metricsList []MetricQuery, endTime time.Time) pmetric.Metrics {
	_ = "STUB: not implemented"
	// statMaps[metricIdx][statIdx] holds a map from timestamp → value for that metric/stat combination.
	// The inner slice length equals numSubQueries(metricsList[i]).
	return *new(pmetric.Metrics)
}

// Skip metrics with no data across any stat.

// Summary mode: combine Sum, SampleCount, Minimum, Maximum into one Summary data point per timestamp.

// Gauge mode: one Gauge data point per (stat, timestamp), tagged with a "stat" attribute.

// Drop the ResourceMetrics if no metric had any data.

// applyQueryAttrs sets the common data point attributes: Namespace, MetricName, and Dimensions (kvlist).
func applyQueryAttrs(attrs pcommon.Map, q MetricQuery) { _ = "STUB: not implemented"; return }
