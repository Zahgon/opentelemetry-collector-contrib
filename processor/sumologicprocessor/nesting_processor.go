// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumologicprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/sumologicprocessor"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type NestingProcessorConfig struct {
	Separator          string   `mapstructure:"separator"`
	Enabled            bool     `mapstructure:"enabled"`
	Include            []string `mapstructure:"include"`
	Exclude            []string `mapstructure:"exclude"`
	SquashSingleValues bool     `mapstructure:"squash_single_values"`
}

type nestingProcessor struct {
	separator          string
	enabled            bool
	allowlist          []string
	denylist           []string
	squashSingleValues bool
}

func newNestingProcessor(config *NestingProcessorConfig) *nestingProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (proc *nestingProcessor) processLogs(logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (proc *nestingProcessor) processMetrics(metrics pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (proc *nestingProcessor) processTraces(traces ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (proc *nestingProcessor) processAttributes(attributes pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

// If key is not on allow list or is on deny list, skip translating it.

// Split returns empty slice only if both string and separator are empty
// set map[""] = v and return

// If previous value was not a map, change it into a map.
// The former value will be set under the key "".

// If we're checking the last key, insert empty value, to which v will be copied.

// If we're not checking the last key, put a map.

// Now check the value we want to copy. If it is a map, we should merge both maps.
// Else, just place the value under the key "".

// Checks if given key fulfills the following conditions:
// - has a prefix that exists in the allowlist (if it's not empty)
// - does not have a prefix that exists in the denylist
func (proc *nestingProcessor) shouldTranslateKey(k string) bool {
	_ = "STUB: not implemented"
	return false
}

// Squashes maps that have single values, eg. map {"a": {"b": {"c": "C", "d": "D"}}}}
// gets squashes into {"a.b": {"c": "C", "d": "D"}}}
func (proc *nestingProcessor) squash(attributes pcommon.Map) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

// A function that squashes keys in a value.
// If this value contained a map with one element, it gets squished and its key gets returned.
//
// If this value contained a map with many elements, this function is called on these elements,
// and the key gets replaced if needed, "" is returned.
//
// Else, nothing happens and "" is returned.
func (proc *nestingProcessor) squashAttribute(value pcommon.Value) string {
	_ = "STUB: not implemented"
	return ""
}

// If the map contains only one key-value pair, squash it.

// This will iterate only over one value (the only one)

// This map doesn't get squashed, but its content might have keys replaced.

// If "" was returned, the value was not a one-element map and did not get squashed.

func (proc *nestingProcessor) squashKey(key, keySuffix string) string {
	_ = "STUB: not implemented"
	return ""
}

func (proc *nestingProcessor) isEnabled() bool { _ = "STUB: not implemented"; return false }

func (*nestingProcessor) ConfigPropertyName() string { _ = "STUB: not implemented"; return "" }
