// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal"

import (
	"go.mongodb.org/atlas/mongodbatlas"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal/metadata"
)

func processMeasurements(
	mb *metadata.MetricsBuilder,
	measurements []*mongodbatlas.Measurements,
) error {
	_ = "STUB: not implemented"
	return nil
}

func calculateTotalMetrics(
	mb *metadata.MetricsBuilder,
	measurements []*mongodbatlas.Measurements,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Combine data point values with matching timestamps

func cloneMeasurement(meas *mongodbatlas.Measurements) *mongodbatlas.Measurements {
	_ = "STUB: not implemented"
	return nil
}
