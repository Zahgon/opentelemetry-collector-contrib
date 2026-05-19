// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azurefunctionsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azurefunctionsreceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
)

// loadEncodingExtension loads an extension by ID from the host
// Returns an error if the extension is missing or does not implement the expected type.
func loadEncodingExtension[T any](host component.Host, id component.ID, signalType string) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// loadLogsUnmarshalers builds a map of binding name to plog.Unmarshaler by loading
// each encoding extension from the host
func loadLogsUnmarshalers(host component.Host, bindings []LogsEncodingConfig) (map[string]plog.Unmarshaler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
