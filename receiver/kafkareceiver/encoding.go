// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver"

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/receiver"
)

var (
	errUnknownEncodingExtension = errors.New("unknown encoding extension")
	errInvalidComponentType     = errors.New("invalid component type")
)

func newTracesUnmarshaler(encoding string, _ receiver.Settings, host component.Host) (ptrace.Unmarshaler, error) {
	_ = "STUB: not implemented"
	// Extensions take precedence.
	return *new(ptrace.Unmarshaler), nil
}

func newLogsUnmarshaler(encoding string, set receiver.Settings, host component.Host) (plog.Unmarshaler, error) {
	_ = "STUB: not implemented"
	// Extensions take precedence.
	return *new(plog.Unmarshaler), nil
}

// There is a special case for text-based encodings, where you can specify
// the text encoding (e.g. utf8, utf16) as a suffix in the encoding name.

func newMetricsUnmarshaler(encoding string, _ receiver.Settings, host component.Host) (pmetric.Unmarshaler, error) {
	_ = "STUB: not implemented"
	// Extensions take precedence.
	return *new(pmetric.Unmarshaler), nil
}

// loadEncodingExtension tries to load an available extension for the given encoding.
func loadEncodingExtension[T any](host component.Host, encoding, signalType string) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func newProfilesUnmarshaler(encoding string, _ receiver.Settings, host component.Host) (pprofile.Unmarshaler, error) {
	_ = "STUB: not implemented"
	// Extensions take precedence.
	return *new(pprofile.Unmarshaler), nil
}

// encodingToComponentID attempts to parse the encoding string as a component ID.
func encodingToComponentID(encoding string) (*component.ID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
