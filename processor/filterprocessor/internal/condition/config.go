// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package condition // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor/internal/condition"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

var _ ottl.ConditionsGetter = (*ContextConditions)(nil)

type ContextID string

const (
	Resource  ContextID = "resource"
	Scope     ContextID = "scope"
	Span      ContextID = "span"
	SpanEvent ContextID = "spanevent"
	Metric    ContextID = "metric"
	DataPoint ContextID = "datapoint"
	Log       ContextID = "log"
	Profile   ContextID = "profile"
)

func (c *ContextID) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// ContextConditions is a wrapper struct for OTTL conditions.
type ContextConditions struct {
	Context    ContextID `mapstructure:"context"`
	Conditions []string  `mapstructure:"conditions"`
	// ErrorMode determines how the processor reacts to errors that occur while processing
	// this group of conditions. When provided, it overrides the default Config ErrorMode.
	ErrorMode ottl.ErrorMode `mapstructure:"error_mode"`
}

func (c ContextConditions) GetConditions() []string { _ = "STUB: not implemented"; return nil }

func toContextConditions(conditions any) (*ContextConditions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getErrorMode[T any](pc *ottl.ParserCollection[T], contextConditions *ContextConditions) ottl.ErrorMode {
	_ = "STUB: not implemented"
	return *new(ottl.ErrorMode)
}
