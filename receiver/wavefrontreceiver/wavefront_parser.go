// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package wavefrontreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/wavefrontreceiver"

import (
	"strings"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver/protocol"
)

// wavefrontParser converts metrics in the Wavefront format, see
// https://docs.wavefront.com/wavefront_data_format.html#metrics-data-format-syntax,
// into the internal format of the Collector
type wavefrontParser struct {
	ExtractCollectdTags bool `mapstructure:"extract_collectd_tags"`

	// prevent unkeyed literal initialization
	_ struct{}
}

var (
	_ protocol.Parser       = (*wavefrontParser)(nil)
	_ protocol.ParserConfig = (*wavefrontParser)(nil)
)

// Only two chars can be escaped per Wavefront SDK, see
// https://github.com/wavefrontHQ/wavefront-sdk-go/blob/2c5891318fcd83c35c93bba2b411640495473333/senders/formatter.go#L20
var escapedCharReplacer = strings.NewReplacer(
	`\"`, `"`, // Replaces escaped double-quotes
	`\n`, "\n", // Replaces escaped new-line.
)

// BuildParser creates a new Parser instance that receives Wavefront metric data.
func (wp *wavefrontParser) BuildParser() (protocol.Parser, error) {
	_ = "STUB: not implemented"

	// Parse receives the string with Wavefront metric data, and transforms it to
	// the collector metric format. See
	// https://docs.wavefront.com/wavefront_data_format.html#metrics-data-format-syntax.
	//
	// Each line received represents a Wavefront metric in the following format:
	//
	//	"<metricName> <metricValue> [<timestamp>] source=<source> [pointTags]"
	//
	// Detailed description of each element is available on the link above.
	return *new(protocol.Parser), nil
}

func (wp *wavefrontParser) Parse(line string) (pmetric.Metric, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metric), nil
}

// Timestamp can be omitted so it is only correct if the string was a tag.

// Assume timestamp was omitted, get current time and adjust index.

// no need for special treatment for source, treat it as a normal tag since
// tags are separated by space and are optionally double-quoted.

func (*wavefrontParser) injectCollectDLabels(
	metricName string,
	attributes pcommon.Map,
) string {
	_ = "STUB: not implemented"
	// This comes from SignalFx Gateway code that has the capability to
	// remove CollectD tags from the name of the metric.
	return ""
}

func buildLabels(attributes pcommon.Map, tags string) error { _ = "STUB: not implemented"; return nil }

// First we need to find the key, find first '='

// Quoted value, skip until non-escaped double quote.

// Non-escaped double-quote, it is the end of the value.

// If we didn't find non-escaped double-quote then this is an error.

// Per implementation of Wavefront SDK only double-quotes and
// newline characters are escaped. See the link below:
// https://github.com/wavefrontHQ/wavefront-sdk-go/blob/2c5891318fcd83c35c93bba2b411640495473333/senders/formatter.go#L20

// The value is up to the end.

func unDoubleQuote(s string) string { _ = "STUB: not implemented"; return "" }
