// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sampling // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/sampling"

import (
	"context"
	"regexp"

	"github.com/golang/groupcache/lru"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/pkg/samplingpolicy"
)

const defaultCacheSize = 128

type stringAttributeFilter struct {
	key    string
	logger *zap.Logger
	// matcher defines the func to match the attribute values in strict string
	// or in regular expression
	matcher     func(string) bool
	invertMatch bool
}

type regexStrSetting struct {
	matchedAttrs *lru.Cache
	filterList   []*regexp.Regexp
}

var _ samplingpolicy.Evaluator = (*stringAttributeFilter)(nil)

// NewStringAttributeFilter creates a policy evaluator that samples all traces with
// the given attribute in the given numeric range.
func NewStringAttributeFilter(settings component.TelemetrySettings, key string, values []string, regexMatchEnabled bool, evictSize int, invertMatch bool) (samplingpolicy.Evaluator, error) {
	_ = "STUB: not implemented"
	// initialize regex filter rules and LRU cache for matched results
	return *new(samplingpolicy.Evaluator), nil
}

// matcher returns true if the given string matches the regex rules defined in string attribute filters

// initialize the exact value map

// matcher returns true if the given string matches any of the string attribute filters

// Evaluate looks at the trace data and returns a corresponding SamplingDecision.
// The SamplingDecision is made by comparing the attribute values with the matching values,
// which might be static strings or regular expressions.
func (saf *stringAttributeFilter) Evaluate(_ context.Context, _ pcommon.TraceID, trace *samplingpolicy.TraceData) (samplingpolicy.Decision, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision), nil
}

// Invert Match returns true by default, except when key and value are matched

func (*stringAttributeFilter) IsStateful() bool {
	_ = "STUB: not implemented"

	// addFilters compiles all the given filters and stores them as regexes.
	// All regexes are automatically anchored to enforce full string matches.
	return false
}

func addFilters(exprs []string) ([]*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
