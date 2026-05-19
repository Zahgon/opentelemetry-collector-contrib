// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package redactionprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor"

//nolint:gosec
import (
	"context"
	"hash"
	"regexp"

	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor/internal/db"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor/internal/url"
)

const attrValuesSeparator = ","

type redaction struct {
	// Attribute keys allowed in a span
	allowList map[string]string
	// Attribute keys ignored in a span
	ignoreList map[string]string
	// Attribute key patterns ignored in a span
	ignoreKeyRegexList map[string]*regexp.Regexp
	// Attribute values blocked in a span
	blockRegexList map[string]*regexp.Regexp
	// Attribute values allowed in a span
	allowRegexList map[string]*regexp.Regexp
	// Attribute keys blocked in a span
	blockKeyRegexList map[string]*regexp.Regexp
	// Hash function to hash blocked values
	hashFunction HashFunction
	// Redaction processor configuration
	config *Config
	// Logger
	logger *zap.Logger
	// URL sanitizer
	urlSanitizer *url.URLSanitizer
	// Database obfuscator
	dbObfuscator *db.Obfuscator
}

// newRedaction creates a new instance of the redaction processor
func newRedaction(ctx context.Context, config *Config, logger *zap.Logger) (*redaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Placeholder for an error metric in the next PR

// TODO: Placeholder for an error metric in the next PR

// TODO: Placeholder for an error metric in the next PR

// TODO: Placeholder for an error metric in the next PR

// processTraces implements ProcessMetricsFunc. It processes the incoming data
// and returns the data to be sent to the next component
func (s *redaction) processTraces(ctx context.Context, batch ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func (s *redaction) processLogs(ctx context.Context, logs plog.Logs) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func (s *redaction) processMetrics(ctx context.Context, metrics pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// processResourceSpan processes the RS and all of its spans and then returns the last
// view metric context. The context can be used for tests
func (s *redaction) processResourceSpan(ctx context.Context, rs ptrace.ResourceSpans) {
	_ = "STUB: not implemented"
	return
}

// Attributes can be part of a resource span

// Attributes can also be part of span

// Attributes can also be part of span events

func (s *redaction) processSpanEvents(ctx context.Context, events ptrace.SpanEventSlice) {
	_ = "STUB: not implemented"
	return
}

// processResourceLog processes the log resource and all of its logs and then returns the last
// view metric context. The context can be used for tests
func (s *redaction) processResourceLog(ctx context.Context, rl plog.ResourceLogs) {
	_ = "STUB: not implemented"
	return
}

func (s *redaction) processLogBody(ctx context.Context, body pcommon.Value, attributes pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func (s *redaction) redactLogBodyRecursive(ctx context.Context, key string, value pcommon.Value, redactedKeys, maskedKeys, allowedKeys, ignoredKeys *[]string) {
	_ = "STUB: not implemented"
	return
}

func (s *redaction) processResourceMetric(ctx context.Context, rm pmetric.ResourceMetrics) {
	_ = "STUB: not implemented"
	return
}

// processAttrs redacts the attributes of a resource span or a span
func (s *redaction) processAttrs(_ context.Context, attributes pcommon.Map) {
	_ = "STUB: not implemented"
	// TODO: Use the context for recording metrics
	return
}

// Identify attributes to redact and mask in the following sequence
// 1. Make a list of attribute keys to redact
// 2. Mask any blocked values for the other attributes
// 3. Delete the attributes from 1
//
// This sequence satisfies these performance constraints:
// - Only range through all attributes once
// - Don't mask any values if the whole attribute is slated for deletion

// Delete the attributes on the redaction list

// Add diagnostic information to the span

//nolint:gosec
func (s *redaction) maskValue(val string, regex *regexp.Regexp) string {
	_ = "STUB: not implemented"
	return ""
}

func hashString(input string, hasher hash.Hash) string { _ = "STUB: not implemented"; return "" }

func hashStringHMAC(input string, key configopaque.String, newHash func() hash.Hash) string {
	_ = "STUB: not implemented"
	return ""
}

// addMetaAttrs adds diagnostic information about redacted or masked attribute keys
func (s *redaction) addMetaAttrs(redactedAttrs []string, attributes pcommon.Map, valuesAttr, countAttr string) {
	_ = "STUB: not implemented"
	return
}

// Record summary as span attributes, empty string for ignored items

func (s *redaction) processStringValueForAttribute(strVal, attributeKey string) string {
	_ = "STUB: not implemented"
	return ""
}

func (s *redaction) processStringValueForLogBody(strVal string) string {
	_ = "STUB: not implemented"
	// Mask any blocked values for the other attributes
	return ""
}

func (s *redaction) shouldMaskKey(k string) bool {
	_ = "STUB: not implemented"
	// Mask any blocked keys for the other attributes
	return false
}

func (s *redaction) shouldAllowValue(strVal string) bool {
	_ = "STUB: not implemented"
	// Allow any values matching the allowed list regex
	return false
}

func (s *redaction) shouldIgnoreKey(k string) bool { _ = "STUB: not implemented"; return false }

// Check if key matches any of the ignored key patterns

func (s *redaction) shouldRedactKey(k string) bool { _ = "STUB: not implemented"; return false }

func (s *redaction) sanitizeSpanName(span ptrace.Span) { _ = "STUB: not implemented"; return }

func applySpanName(span ptrace.Span, original, candidate string) { _ = "STUB: not implemented"; return }

func (s *redaction) shouldSanitizeSpanNameForURL() bool { _ = "STUB: not implemented"; return false }

func (s *redaction) shouldSanitizeSpanNameForDB() bool { _ = "STUB: not implemented"; return false }

const (
	debug                      = "debug"
	info                       = "info"
	redactionRedactedKeys      = "redaction.redacted.keys"
	redactionRedactedCount     = "redaction.redacted.count"
	redactionMaskedKeys        = "redaction.masked.keys"
	redactionMaskedCount       = "redaction.masked.count"
	redactionAllowedKeys       = "redaction.allowed.keys"
	redactionAllowedCount      = "redaction.allowed.count"
	redactionIgnoredCount      = "redaction.ignored.count"
	redactionBodyRedactedKeys  = "redaction.body.redacted.keys"
	redactionBodyRedactedCount = "redaction.body.redacted.count"
	redactionBodyMaskedKeys    = "redaction.body.masked.keys"
	redactionBodyMaskedCount   = "redaction.body.masked.count"
	redactionBodyAllowedKeys   = "redaction.body.allowed.keys"
	redactionBodyAllowedCount  = "redaction.body.allowed.count"
	redactionBodyIgnoredCount  = "redaction.body.ignored.count"
)

// makeAllowList sets up a lookup table of allowed span attribute keys
func makeAllowList(c *Config) map[string]string {
	_ = "STUB: not implemented"
	// redactionKeys are additional span attributes created by the processor to
	// summarize the changes it made to a span. If the processor removes
	// 2 attributes from a span (e.g. `birth_date`, `mothers_maiden_name`),
	// then it will list them in the `redaction.redacted.keys` span attribute
	// and set the `redaction.redacted.count` attribute to 2
	//
	// If the processor finds and masks values matching a blocked regex in 2
	// span attributes (e.g. `notes`, `description`), then it will those
	// attribute keys in `redaction.masked.keys` and set the
	// `redaction.masked.count` to 2
	return nil
}

// allowList consists of the keys explicitly allowed by the configuration
// as well as of the new span attributes that the processor creates to
// summarize its changes

func makeIgnoreList(c *Config) map[string]string { _ = "STUB: not implemented"; return nil }

// makeRegexList precompiles all the regex patterns in the defined list
func makeRegexList(_ context.Context, valuesList []string) (map[string]*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Placeholder for an error metric in the next PR
