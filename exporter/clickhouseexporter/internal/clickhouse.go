// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter/internal"

import (
	"context"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

const DefaultDatabase = "default"

// DatabaseFromDSN returns the database specified in the DSN. Empty if unset.
func DatabaseFromDSN(dsn string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// NewClickhouseClientFromOptions creates a new ClickHouse client from a clickhouse.Options struct.
func NewClickhouseClientFromOptions(opt *clickhouse.Options) (driver.Conn, error) {
	_ = "STUB: not implemented"
	// Always connect to default database since configured database may not exist yet.
	// TODO: only do this if createSchema is true
	return *new(driver.Conn), nil
}

// GenerateTTLExpr generates a TTL expression for a ClickHouse table.
func GenerateTTLExpr(ttl time.Duration, timeField string) string {
	_ = "STUB: not implemented"
	return ""
}

// CreateDatabase runs the DDL for creating a database, with optional cluster string
func CreateDatabase(ctx context.Context, db driver.Conn, database, clusterStr string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetTableColumns returns the column names on a table for schema detection
func GetTableColumns(ctx context.Context, db driver.Conn, database, table string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
