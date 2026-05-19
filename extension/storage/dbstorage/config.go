// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dbstorage // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/dbstorage"

const (
	driverPostgreSQL   = "pgx"
	driverSQLite       = "sqlite"
	driverSQLiteLegacy = "sqlite3"
)

// Config defines configuration for dbstorage extension.
type Config struct {
	DriverName string `mapstructure:"driver,omitempty"`
	DataSource string `mapstructure:"datasource,omitempty"`
}

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }
