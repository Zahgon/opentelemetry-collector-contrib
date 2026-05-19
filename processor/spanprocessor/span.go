// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package spanprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanprocessor"

import (
	"context"
	"regexp"

	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
)

type spanProcessor struct {
	config           Config
	toAttributeRules []toAttributeRule
	skipExpr         expr.BoolExpr[*ottlspan.TransformContext]
}

// toAttributeRule is the compiled equivalent of config.ToAttributes field.
type toAttributeRule struct {
	// Compiled regexp.
	re *regexp.Regexp

	// Attribute names extracted from the regexp's subexpressions.
	attrNames []string
}

// newSpanProcessor returns the span processor.
func newSpanProcessor(config Config) (*spanProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Compile ToAttributes regexp and extract attributes names.

// Subexpression names will become attribute names during extraction.

func (sp *spanProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func (sp *spanProcessor) processFromAttributes(span ptrace.Span) { _ = "STUB: not implemented"; return }

// There is FromAttributes rule.

// There are no attributes to create span name from.

// Note: There was a separate proposal for creating the string.
// With benchmarking, strings.Builder is faster than the proposal.
// For full context, refer to this PR comment:
// https://go.opentelemetry.io/collector/pull/301#discussion_r318357678

// If one of the keys isn't found, the span name is not updated.

// Note: WriteString() always return a nil error so there is no error checking
// for this method call.
// https://golang.org/src/strings/builder.go?s=3425:3477#L110

// Include the separator before appending an attribute value if:
// this isn't the first value(ie i == 0) loop through the FromAttributes
// and
// the separator isn't an empty string.

func (sp *spanProcessor) processToAttributes(span ptrace.Span) { _ = "STUB: not implemented"; return }

// There is no span name to work on.

// No rules to apply.

// Process rules one by one. Store results of processing in the span
// so that each subsequent rule works on the span name that is the output
// after processing the previous rule.

// Match the regular expression and extract matched subexpressions.

// There is a match. We will also need positions of subexpression matches.

// A place to accumulate new span name.

// Index in the oldName until which we traversed.

// TODO: Pre-allocate len(submatches) space in the attributes.

// Start from index 1, which is the first submatch (index 0 is the entire match).
// We will go over submatches and will simultaneously build a new span name,
// replacing matched subexpressions by attribute names.

// Add part of span name from end of previous match to start of this match
// and then add attribute name wrapped in curly brackets.
// start of i'th submatch.

// Advance the index to the end of current match.
// end of i'th submatch.

// Append the remainder, from the end of last match until end of span name.

// Set new span name.

// Stop processing, break after first match is requested.

func (sp *spanProcessor) processUpdateStatus(span ptrace.Span) { _ = "STUB: not implemented"; return }
