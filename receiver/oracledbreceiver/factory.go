// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oracledbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver"

import (
	"database/sql"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"
)

// NewFactory creates a new Oracle receiver factory.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

type sqlOpenerFunc func(dataSourceName string) (*sql.DB, error)

func createReceiverFunc(sqlOpenerFunc sqlOpenerFunc, clientProviderFunc clientProviderFunc) receiver.CreateMetricsFunc {
	_ = "STUB: not implemented"
	return *new(receiver.CreateMetricsFunc)
}

func createLogsReceiverFunc(sqlOpenerFunc sqlOpenerFunc, clientProviderFunc clientProviderFunc) receiver.CreateLogsFunc {
	_ = "STUB: not implemented"
	return *new(receiver.CreateLogsFunc)
}

// cacheSize is kept at 2 times MaxQuerySampleCount to keep queries of adjacent collections available for delta calculation.

func getDataSource(cfg Config) string { _ = "STUB: not implemented"; return "" }

// Don't need to worry about errors here as config validation already checked.

func getInstanceName(datasource string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func getHostName(datasource string) (string, error) { _ = "STUB: not implemented"; return "", nil }
