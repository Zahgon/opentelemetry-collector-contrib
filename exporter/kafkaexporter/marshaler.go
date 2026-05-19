// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkaexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter"

import (
	"errors"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter/internal/marshaler"
)

var errUnknownEncodingExtension = errors.New("unknown encoding extension")

func getTracesMarshaler(encoding string, host component.Host) (marshaler.TracesMarshaler, error) {
	_ = "STUB: not implemented"
	return *new(marshaler.TracesMarshaler), nil
}

func getMetricsMarshaler(encoding string, host component.Host) (marshaler.MetricsMarshaler, error) {
	_ = "STUB: not implemented"
	return *new(marshaler.MetricsMarshaler), nil
}

func getLogsMarshaler(encoding string, host component.Host) (marshaler.LogsMarshaler, error) {
	_ = "STUB: not implemented"
	return *new(marshaler.LogsMarshaler), nil
}

func getProfilesMarshaler(encoding string, host component.Host) (marshaler.ProfilesMarshaler, error) {
	_ = "STUB: not implemented"
	return *new(marshaler.ProfilesMarshaler), nil
}

// loadEncodingExtension tries to load an available extension for the given encoding.
func loadEncodingExtension[T any](host component.Host, encoding, signalType string) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// encodingToComponentID attempts to parse the encoding string as a component ID.
func encodingToComponentID(encoding string) (*component.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
