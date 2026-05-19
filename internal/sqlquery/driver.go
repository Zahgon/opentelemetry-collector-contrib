// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlquery // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/sqlquery"

const (
	DriverClickHouse = "clickhouse"
	DriverHDB        = "hdb"
	DriverMySQL      = "mysql"
	DriverOracle     = "oracle"
	DriverPostgres   = "postgres"
	DriverSnowflake  = "snowflake"
	DriverSQLServer  = "sqlserver"
	DriverTDS        = "tds"
)

// IsValidDriver checks if the given driver name is supported
func IsValidDriver(driver string) bool { _ = "STUB: not implemented"; return false }
