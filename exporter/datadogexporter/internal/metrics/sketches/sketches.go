// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Parts of this package are based on the code from the datadog-agent,
// https://github.com/DataDog/datadog-agent/blob/main/pkg/metrics/sketch_series.go

// Package sketches is a copy of part from github.com/DataDog/datadog-agent/pkg/metrics.
// TODO(mx-psi): import pkg/metrics from datadog-agent directly
package sketches // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metrics/sketches"

import (
	"github.com/DataDog/datadog-agent/pkg/util/quantile"
)

const (
	SketchSeriesEndpoint string = "/api/beta/sketches"
)

// A SketchSeries is a timeseries of quantile sketches.
type SketchSeries struct {
	Name     string        `json:"metric"`
	Tags     []string      `json:"tags"`
	Host     string        `json:"host"`
	Interval int64         `json:"interval"`
	Points   []SketchPoint `json:"points"`
}

// A SketchPoint represents a quantile sketch at a specific time
type SketchPoint struct {
	Sketch *quantile.Sketch `json:"sketch"`
	Ts     int64            `json:"ts"`
}

// A SketchSeriesList implements marshaler.Marshaler
type SketchSeriesList []SketchSeries

// Marshal encodes this series list.
func (sl SketchSeriesList) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
