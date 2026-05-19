// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package prometheusremotewritereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusremotewritereceiver"

import (
	"context"
	"net/http"
	"sync"

	lru "github.com/hashicorp/golang-lru/v2"
	remoteapi "github.com/prometheus/client_golang/exp/api/remote"
	"github.com/prometheus/prometheus/model/labels"
	writev2 "github.com/prometheus/prometheus/prompb/io/prometheus/write/v2"
	promremote "github.com/prometheus/prometheus/storage/remote"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/exp/metrics/identity"
)

func newRemoteWriteReceiver(settings receiver.Settings, cfg *Config, nextConsumer consumer.Metrics) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

// Pre-allocate 4KiB

type prometheusRemoteWriteReceiver struct {
	settings     receiver.Settings
	nextConsumer consumer.Metrics

	config *Config
	server *http.Server
	wg     sync.WaitGroup

	rmCache *lru.Cache[uint64, pmetric.ResourceMetrics]
	obsrecv *receiverhelper.ObsReport

	bodyBufferPool *sync.Pool
}

// scopeInfo holds instrumentation scope fields extracted from otel_scope_* labels.
type scopeInfo struct {
	Name       string
	Version    string
	SchemaURL  string
	scopeAttrs []attribute // scope attributes with the "otel_scope_" prefix stripped
}

// attribute is a simple key-value pair for scope attributes.
type attribute struct {
	Key   string
	Value string
}

// metricIdentity contains all the components that uniquely identify a metric
// according to the OpenTelemetry Protocol data model.
// The definition of the metric uniqueness is based on the following document. Ref: https://opentelemetry.io/docs/specs/otel/metrics/data-model/#opentelemetry-protocol-data-model
type metricIdentity struct {
	resource       identity.Resource
	ScopeName      string
	ScopeVersion   string
	ScopeSchemaURL string
	ScopeAttrs     []attribute
	MetricName     string
	Unit           string
	Type           writev2.Metadata_MetricType
}

// createMetricIdentity creates a metricIdentity struct from the required components
func createMetricIdentity(resource identity.Resource, metricName, unit string, si scopeInfo, metricType writev2.Metadata_MetricType) metricIdentity {
	_ = "STUB: not implemented"
	return *new(metricIdentity)
}

// Hash generates a unique hash for the metric identity using the identity library's hasher
// as a foundation, extended with scope and metric fields.
func (mi metricIdentity) Hash() uint64 { _ = "STUB: not implemented"; return 0 }

func (prw *prometheusRemoteWriteReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (prw *prometheusRemoteWriteReceiver) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Only wait if Shutdown returns successfully,
// otherwise we may block indefinitely.

func (prw *prometheusRemoteWriteReceiver) handlePRW(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// After parsing the content-type header, the next step would be to handle content-encoding.
// Luckly confighttp's Server has middleware that already decompress the request body for us.

// Following instructions at https://prometheus.io/docs/specs/remote_write_spec_2_0/#invalid-samples

// Return early if metric count is 0.

// parseProto parses the content-type header and returns the version of the remote-write protocol.
// We can't expect that senders of remote-write v1 will add the "proto=" parameter since it was not
// a requirement in v1. So, if the parameter is not found, we assume v1.
func (*prometheusRemoteWriteReceiver) parseProto(contentType string) (remoteapi.WriteMessageType, error) {
	_ = "STUB: not implemented"
	return *new(remoteapi.WriteMessageType), nil
}

// No "proto=" parameter found, assume v1.

// getOrCreateRM returns or creates the ResourceMetrics for a job/instance pair within an HTTP request.
//
// Two-level cache:
//
//  1. reqRM (per-request): groups samples with the same job/instance into a single ResourceMetrics
//     during current request processing, avoiding duplication in the output.
//
//  2. prw.rmCache (global LRU): stores snapshots of previously seen resource attributes (from target_info),
//     allowing future requests to reuse enriched attributes.
//
// This function always creates new ResourceMetrics per request, only copying attributes
// from the LRU cache when available. Never returns cached objects to avoid shared
// mutation across concurrent requests.
func (prw *prometheusRemoteWriteReceiver) getOrCreateRM(ls labels.Labels, otelMetrics pmetric.Metrics, reqRM map[uint64]pmetric.ResourceMetrics) (pmetric.ResourceMetrics, uint64) {
	_ = "STUB: not implemented"
	// Hash job+instance directly to avoid allocating a temporary pcommon.Resource
	// on every call (which happens once per time series).
	return *new(pmetric.ResourceMetrics), 0
}

// When the ResourceMetrics already exists in the global cache, we can reuse the previous snapshots and perpass the already seen attributes to the current request.

// When the ResourceMetrics does not exist in the global cache, we need to create a new one and add it to the request map.
// Saving the new ResourceMetrics in the global cache to avoid creating duplicates in the next requests.

// translateV2 translates a v2 remote-write request into OTLP metrics.
// translate is not feature complete.
func (prw *prometheusRemoteWriteReceiver) translateV2(_ context.Context, req *writev2.Request) (pmetric.Metrics, promremote.WriteResponseStats, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), *new(promremote.WriteResponseStats), nil
}

// otelMetrics represents the final metrics, after all the processing, that will be returned by the receiver.

// More about stats: https://github.com/prometheus/docs/blob/main/docs/specs/prw/remote_write_spec_2_0.md#required-written-response-headers
// TODO: add exemplars to the stats.

// The key is composed by: resource_hash:scope_name:scope_version:metric_name:unit:type

// modifiedResourceMetric keeps track, for each request, of which resources (identified by the job/instance hash) had their resource attributes modified — for example, through target_info.
// Once the request is fully processed, only the resource attributes contained in the request’s ResourceMetrics are snapshotted back into the LRU cache.
// This ensures that future requests start with the enriched resource attributes already applied.

// exemplarMap keeps track of exemplars and key is composed by scope_name:scope_version:metric_name:type

// If the metric name is equal to target_info, we use its labels as attributes of the resource
// Ref: https://opentelemetry.io/docs/specs/otel/compatibility/prometheus_and_openmetrics/#resource-attributes-1

// Add the remaining labels as resource attributes

// target_info is not stored as a metric but PRW requires the response
// to report all received samples, including target_info, to avoid a stats mismatch

// Handle histograms separately due to their complex mixed-schema processing

// Handle regular metrics (gauge, counter, summary)

// Resource identity
// Metric name
// Unit
// Scope info
// Metric type

// Find or create scope

// Get or create metric

// Ref: https://opentelemetry.io/docs/specs/otel/compatibility/prometheus_and_openmetrics/#otlp-to-prometheus
// Info and StateSet metrics MUST be converted to an OTLP Non-Monotonic Sum.

// Drop summary series as we will not handle them.

// When the new description is longer than the existing one, we should update the metric description.
// Reference to this behavior: https://opentelemetry.io/docs/specs/otel/metrics/data-model/#opentelemetry-protocol-data-model-producer-recommendations

// add exemplars to counter datapoints.

// Drop summary series as we will not handle them.

// processHistogramTimeSeries handles all histogram processing, including validation and mixed schemas.
func (prw *prometheusRemoteWriteReceiver) processHistogramTimeSeries(
	otelMetrics pmetric.Metrics,
	ls labels.Labels,
	ts *writev2.TimeSeries,
	si scopeInfo,
	metricName, unit, description string,
	metricCache map[uint64]pmetric.Metric,
	stats *promremote.WriteResponseStats,
	modifiedRM map[uint64]pmetric.ResourceMetrics,
	exemplarMap map[uint64]pmetric.ExemplarSlice,
) {
	_ = "STUB: not implemented"
	// Drop classic histogram series (those with samples)
	return
}

// Determine histogram type based on schema
// See https://prometheus.io/docs/specs/native_histograms/#schema

// Skip invalid schema - log at debug level for details

// Create resource if needed (only for the first valid histogram)

// Find or create scope (search each time since different histograms might need different scopes)

// This default case should not be reached as this function is only called when:
// 1. ts.Metadata.Type == METRIC_TYPE_HISTOGRAM, or
// 2. ts.Metadata.Type == METRIC_TYPE_UNSPECIFIED && len(ts.Histograms) > 0

// When the new description is longer than the existing one, we should update the metric description.
// Reference to this behavior: https://opentelemetry.io/docs/specs/otel/metrics/data-model/#opentelemetry-protocol-data-model-producer-recommendations

// all the exemplars for a given histogram are attached to first data point.

// Process the individual histogram

// setMetric append a new empty metric and assign the name, unit and description to it.
func setMetric(scope pmetric.ScopeMetrics, metricName, unit, description string) pmetric.Metric {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric)
}

// parseJobAndInstance turns the job and instance labels service resource attributes.
// Following the specification at https://opentelemetry.io/docs/specs/otel/compatibility/prometheus_and_openmetrics/
func parseJobAndInstance(dest pcommon.Map, job, instance string) { _ = "STUB: not implemented"; return }

// addNumberDatapoints adds the labels to the datapoints attributes.
func addNumberDatapoints(datapoints pmetric.NumberDataPointSlice, ls labels.Labels, ts *writev2.TimeSeries, stats *promremote.WriteResponseStats) {
	_ = "STUB: not implemented"
	// Add samples from the timeseries
	return
}

// Set timestamp in nanoseconds (Prometheus uses milliseconds)

func (prw *prometheusRemoteWriteReceiver) addExponentialHistogramDatapoint(datapoints pmetric.ExponentialHistogramDataPointSlice, histogram *writev2.Histogram, attrs pcommon.Map, ls labels.Labels, stats *promremote.WriteResponseStats) {
	_ = "STUB: not implemented"
	// Drop Native Histogram with negative counts
	return
}

// Set count and sum using common helper

// The maximum bucket index is derived from the formula (2^(2^-n))^i <= MaxFloat64.
// MaxFloat64 is approx 2^1024. So (2^-n) * i <= 1024 => i <= 1024 * 2^n.
// The bucket containing MaxFloat64 has index i_max = 1024 * 2^n.
// The next bucket (i_max + 1) is the +Inf overflow bucket, which is also allowed.
// Buckets with an index strictly greater than i_max + 1 must be dropped.
// See https://prometheus.io/docs/specs/native_histograms/#schema for more information.

// The difference between float and integer histograms is that float histograms are stored as absolute counts
// while integer histograms are stored as deltas.

// Float histograms

// -1 because OTEL offset are for the lower bound, not the upper bound

// -1 because OTEL offset are for the lower bound, not the upper bound

// Integer histograms

// -1 because OTEL offset are for the lower bound, not the upper bound

// -1 because OTEL offset are for the lower bound, not the upper bound

// hasNegativeCounts checks if a histogram has any negative counts
func hasNegativeCounts(histogram *writev2.Histogram) bool { _ = "STUB: not implemented"; return false }

// Check overall count

// Check zero count

// Check positive bucket counts

// Check negative bucket counts

// Integer histograms

// convertDeltaBuckets converts Prometheus native histogram spans and deltas to OpenTelemetry bucket counts
// For integer buckets, the values are deltas between the buckets. i.e a bucket list of 1,2,-2 would correspond to a bucket count of 1,3,1
func convertDeltaBuckets(spans []writev2.BucketSpan, deltas []int64, buckets pcommon.UInt64Slice, overflowLimit int32) uint64 {
	_ = "STUB: not implemented"
	// The total capacity is the sum of the deltas and the offsets of the spans.
	return 0
}

// convertAbsoluteBuckets converts Prometheus native histogram spans and absolute counts to OpenTelemetry bucket counts
// For float buckets, the values are absolute counts, and must be 0 or positive.
func convertAbsoluteBuckets(spans []writev2.BucketSpan, counts []float64, buckets pcommon.UInt64Slice, overflowLimit int32) uint64 {
	_ = "STUB: not implemented"
	// The total capacity is the sum of the counts and the offsets of the spans.
	return 0
}

// extractAttributes returns metric data point attributes, excluding job, instance, metric name, and all otel_scope_* labels.
func extractAttributes(ls labels.Labels) pcommon.Map {
	_ = "STUB: not implemented"
	return *

	// job, instance and metric name will always become labels
	new(pcommon.Map)
}

// Become resource attributes
// Becomes metric name
// Become instrumentation scope fields

// extractScopeInfo extracts all otel_scope_* labels into a scopeInfo per the Prometheus/OTLP compatibility spec.
// Falls back to receiver build info when otel_scope_name is absent.
func (prw *prometheusRemoteWriteReceiver) extractScopeInfo(ls labels.Labels) scopeInfo {
	_ = "STUB: not implemented"
	return *new(scopeInfo)
}

func scopeMatchesInfo(sm pmetric.ScopeMetrics, si scopeInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func applyScopeInfo(sm pmetric.ScopeMetrics, si scopeInfo) { _ = "STUB: not implemented"; return }

// addNHCBDatapoint converts a single Native Histogram Custom Buckets (NHCB) to OpenTelemetry histogram datapoints
func (*prometheusRemoteWriteReceiver) addNHCBDatapoint(datapoints pmetric.HistogramDataPointSlice, histogram *writev2.Histogram, attrs pcommon.Map, stats *promremote.WriteResponseStats) {
	_ = "STUB: not implemented"
	return
}

// convertNHCBBuckets converts NHCB bucket data to OpenTelemetry bucket counts
func convertNHCBBuckets(histogram *writev2.Histogram) []uint64 {
	_ = "STUB: not implemented"
	// For NHCB, we need numExplicitBounds + 1 buckets (including the final +inf bucket)
	return nil
}

// NHCB uses the positive bucket list and spans for all buckets

// Float histograms: values are absolute counts

// Skip empty buckets based on offset

// Fill buckets for this span

// Integer histograms: values are deltas between buckets

// Skip empty buckets based on offset

// Fill buckets for this span

// setCountAndSum sets count and sum for histogram datapoints (common interface)
type countSumSetter interface {
	SetSum(float64)
	SetCount(uint64)
}

func setCountAndSum(histogram *writev2.Histogram, dp countSumSetter) {
	_ = "STUB: not implemented"
	return
}
