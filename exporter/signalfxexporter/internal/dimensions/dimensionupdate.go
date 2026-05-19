// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dimensions // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/dimensions"

type DimensionUpdate struct {
	Name       string
	Value      string
	Properties map[string]*string
	Tags       map[string]bool
}

func (d *DimensionUpdate) String() string { _ = "STUB: not implemented"; return "" }

func (d *DimensionUpdate) Key() DimensionKey { _ = "STUB: not implemented"; return *new(DimensionKey) }

// DimensionKey is what uniquely identifies a dimension, its name and value
// together.
type DimensionKey struct {
	Name  string
	Value string
}

func (dk DimensionKey) String() string { _ = "STUB: not implemented"; return "" }
