// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlquery // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/sqlquery"

import (
	"context"
	"database/sql"
)

// These are wrappers and interfaces around SQL.DB so that it can be swapped out for testing.

type Db interface {
	QueryContext(ctx context.Context, query string, args ...any) (rows, error)
}

type rows interface {
	ColumnTypes() ([]colType, error)
	Next() bool
	Scan(dest ...any) error
}

type colType interface {
	Name() string
}

type DbWrapper struct {
	Db *sql.DB
}

func (d DbWrapper) QueryContext(ctx context.Context, query string, args ...any) (rows, error) {
	_ = "STUB: not implemented"
	return *new(rows), nil
}

type rowsWrapper struct {
	rows *sql.Rows
}

func (r rowsWrapper) ColumnTypes() ([]colType, error) { _ = "STUB: not implemented"; return nil, nil }

func (r rowsWrapper) Next() bool { _ = "STUB: not implemented"; return false }

func (r rowsWrapper) Scan(dest ...any) error { _ = "STUB: not implemented"; return nil }

type colWrapper struct {
	ct *sql.ColumnType
}

func (c colWrapper) Name() string { _ = "STUB: not implemented"; return "" }
