// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sdktest // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/testing/sdktest"

import (
	"go.opentelemetry.io/otel/attribute"
	sdk "go.opentelemetry.io/otel/sdk/metric/metricdata"
)

// Spec is the partial metric specification. To be used with [Compare]
type Spec = map[string]Metric

type Type string

const (
	TypeSum   Type = "sum"
	TypeGauge Type = "gauge"
)

type Metric struct {
	Type
	Name string

	Numbers     []Number
	Monotonic   bool
	Temporality sdk.Temporality
}

type Number struct {
	Int   *int64
	Float *float64
	Attr  attributes
}

// Unmarshal specification in [Format] into the given [Spec].
func Unmarshal(data Format, into *Spec) error { _ = "STUB: not implemented"; return nil }

type attributes map[string]string

func (attr attributes) Into() attribute.Set { _ = "STUB: not implemented"; return *new(attribute.Set) }

// Format defines the yaml-based format to be used with [Unmarshal] for specifying [Spec].
//
// It looks as follows:
//
//	<instrument> <name> [ delta|cumulative ]:
//	- int: <int64> | float: <float64>
//	  attr:
//	    [string]: <string>
//
// The supported instruments are:
//   - counter: [TypeSum], monotonic
//   - updown: [TypeSum], non-monotonic
//   - gauge: [TypeGauge]
//
// Temporality is optional and defaults to [sdk.CumulativeTemporality]
type Format = []byte
