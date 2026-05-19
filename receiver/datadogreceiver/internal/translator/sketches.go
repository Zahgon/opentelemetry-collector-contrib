// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver/internal/translator"

import (
	"math"
	"net/http"

	"github.com/DataDog/agent-payload/v5/gogen"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

const (
	// The relativeAccuracy (also called epsilon or eps) comes from DDSketch's logarithmic mapping, which is used for sketches
	// in the Datadog agent. The Datadog agent uses the default value from opentelemetry-go-mapping configuration
	// See:
	// https://github.com/DataDog/datadog-agent/blob/fcb59435e45053bcb53a1eec482104290f1dd166/pkg/util/quantile/config.go#L15
	relativeAccuracy = 1.0 / 128

	// The gamma value comes from the default values of the epsilon/relative accuracy from opentelemetry-go-mapping. This value is used for
	// finding the lower boundary of the bucket at a specific index
	// See:
	// https://github.com/DataDog/datadog-agent/blob/fcb59435e45053bcb53a1eec482104290f1dd166/pkg/util/quantile/config.go#L138
	gamma = 1 + 2*relativeAccuracy

	// Since the default bucket factor for Sketches (gamma value) is 1.015625, this corresponds to a scale between 5 (2^2^-5=1.0219)
	// and 6 (2^2^-6=1.01088928605). However, the lower resolution of 5 will produce larger buckets which allows for easier mapping
	scale = 5

	// The agentSketchOffset value comes from the following calculation:
	// min = 1e-9
	// emin = math.Floor((math.Log(min)/math.Log1p(2*relativeAccuracy))
	// offset = -emin + 1
	// The resulting value is 1338.
	// See: https://github.com/DataDog/datadog-agent/blob/fcb59435e45053bcb53a1eec482104290f1dd166/pkg/util/quantile/config.go#L154
	// (Note: in Datadog's code, it is referred to as 'bias')
	agentSketchOffset int32 = 1338

	// The max limit for the index of a sketch bucket
	// See https://github.com/DataDog/datadog-agent/blob/fcb59435e45053bcb53a1eec482104290f1dd166/pkg/util/quantile/ddsketch.go#L21
	// and https://github.com/DataDog/datadog-agent/blob/fcb59435e45053bcb53a1eec482104290f1dd166/pkg/util/quantile/ddsketch.go#L138
	maxIndex = math.MaxInt16
)

// Unmarshal the sketch payload, which contains the underlying Dogsketch structure used for the translation
func (*MetricsTranslator) HandleSketchesPayload(req *http.Request) (sp []gogen.SketchPayload_Sketch, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mt *MetricsTranslator) TranslateSketches(sketches []gogen.SketchPayload_Sketch) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

// The dogsketches field of the payload contains the sketch data

// If a sketch is invalid, remove this datapoint

func sketchToDatapoint(sketch gogen.SketchPayload_Sketch_Dogsketch, dp pmetric.ExponentialHistogramDataPoint, attributes pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

// OTel uses nanoseconds, while Datadog uses seconds

// See https://github.com/DataDog/sketches-go/blob/7546f8f95179bb41d334d35faa281bfe97812a86/ddsketch/mapping/logarithmic_mapping.go#L48

// mapSketchBucketsToHistogramBuckets attempts to map the counts in each Sketch bucket to the closest equivalent Exponential Histogram
// bucket(s). It works by first calculating an Exponential Histogram key that corresponds most closely with the Sketch key (using the lower
// bound of the sketch bucket the key corresponds to), calculates differences in the range of the Sketch bucket and exponential histogram bucket,
// and distributes the count to the corresponding bucket, and the bucket(s) after it, based on the proportion of overlap between the
// exponential histogram buckets and the Sketch bucket. Note that the Sketch buckets are not separated into positive and negative buckets, but exponential
// histograms store positive and negative buckets separately. Negative buckets in exponential histograms are mapped in the same way as positive buckets.
// Note that negative indices in exponential histograms do not necessarily correspond to negative values; they correspond with values between 0 and 1,
// on either the negative or positive side
func mapSketchBucketsToHistogramBuckets(sketchKeys []int32, sketchCounts []uint32) (map[int]uint64, map[int]uint64, uint64, error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, nil
}

// The data format for the sketch received from the sketch payload does not have separate positive and negative buckets,
// and instead just uses a single list of sketch keys that are in order by increasing bucket index, starting with negative indices,
// which correspond to negative buckets

// A sketch key of 0 corresponds to the zero bucket

// This should not happen, as sketches that contain bucket(s) with an index greater than the max
// limit should have already been discarded. However, if there happens to be an index > maxIndex,
// it can cause an infinite loop within the below inner for loop on some operating systems. Therefore,
// throw an error for sketches that have an index above the max limit

// The approach here is to use the Datadog sketch index's lower bucket boundary to find the
// OTel exponential histogram bucket that with the closest range to the sketch bucket. Then,
// the buckets before and after that bucket are also checked for overlap with the sketch bucket.
// A count proportional to the intersection of the sketch bucket with the OTel bucket(s) is then
// added to the OTel bucket(s). After looping through all possible buckets that are within the Sketch
// bucket range, the bucket with the highest proportion of overlap is given the remaining count

// TODO: look into better algorithms for applying fractional counts

// In this case, the bucket does not overlap with the sketch bucket, so continue to the next bucket

// OTel exponential histograms only support integer bucket counts, so rounding needs to be done here

// Add the difference between the original sketch bucket's count and the total count that has been
// added to the matching OTel bucket(s) thus far to the bucket that had the highest proportion of
// overlap between the original sketch bucket and the corresponding exponential histogram buckets

// convertBucketLayout populates the count for positive or negative buckets in the resulting OTel
// exponential histogram structure. The bucket layout is dense and consists of an offset, which is the
// index of the first populated bucket, and a list of counts, which correspond to the counts at the offset
// bucket's index, and the counts of each bucket after. Unpopulated/empty buckets must be represented with
// a count of 0. After assigning bucket counts, it sets the offset for the bucket layout
func convertBucketLayout(inputBuckets map[int]uint64, outputBuckets pmetric.ExponentialHistogramDataPointBuckets) {
	_ = "STUB: not implemented"
	return
}

// find total number of buckets needed

// getSketchBounds calculates the lower and upper bounds of a sketch bucket based on the index of the bucket.
// This is based on sketch buckets placing values in bucket so that γ^k <= v < γ^(k+1)
// See https://github.com/DataDog/datadog-agent/blob/0ada7a97fed6727838a6f4d9c87123d2aafde735/pkg/quantile/config.go#L83
// and https://github.com/DataDog/sketches-go/blob/8a1961cf57f80fbbe26e7283464fcc01ebf17d5c/ddsketch/ddsketch.go#L468
func getSketchBounds(index int32) (float64, float64) { _ = "STUB: not implemented"; return 0, 0 }

// sketchLowerBound calculates the lower bound of a sketch bucket based on the index of the bucket.
// It uses the index offset and multiplier (represented by (1 / math.Log(gamma))). The logic behind this
// is based on the DD agent using logarithmic mapping for definition DD agent sketches
// See:
// https://github.com/DataDog/datadog-agent/blob/fcb59435e45053bcb53a1eec482104290f1dd166/pkg/util/quantile/config.go#L54
// https://github.com/DataDog/sketches-go/blob/8a1961cf57f80fbbe26e7283464fcc01ebf17d5c/ddsketch/mapping/logarithmic_mapping.go#L39
func sketchLowerBound(index int32) float64 { _ = "STUB: not implemented"; return 0 }

// getHistogramBounds returns the lower and upper boundaries of the histogram bucket that
// corresponds to the specified bucket index
func getHistogramBounds(histIndex int) (float64, float64) { _ = "STUB: not implemented"; return 0, 0 }

// This equation for finding the lower bound of the exponential histogram bucket
// Based on: https://github.com/open-telemetry/opentelemetry-go/blob/3a72c5ea94bf843beeaa044b0dda2ce4d627bb7b/sdk/metric/internal/aggregate/exponential_histogram.go#L122
// See also: https://github.com/open-telemetry/opentelemetry-go/blob/3a72c5ea94bf843beeaa044b0dda2ce4d627bb7b/sdk/metric/internal/aggregate/exponential_histogram.go#L139
func histogramLowerBound(histIndex int) float64 { _ = "STUB: not implemented"; return 0 }

// sketchLowerBoundToHistogramIndex takes the lower boundary of a sketch bucket and computes the
// closest equivalent exponential histogram index that corresponds to an exponential histogram
// bucket that has a range covering that lower bound
// See: https://opentelemetry.io/docs/specs/otel/metrics/data-model/#all-scales-use-the-logarithm-function
func sketchLowerBoundToHistogramIndex(value float64) int { _ = "STUB: not implemented"; return 0 }
