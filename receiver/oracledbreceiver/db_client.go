// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oracledbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver"

import (
	"context"
	"database/sql"

	// register db driver
	_ "github.com/sijms/go-ora/v2"
	"go.uber.org/zap"
)

type dbClient interface {
	metricRows(ctx context.Context, args ...any) ([]metricRow, error)
}

type metricRow map[string]string

type dbSQLClient struct {
	db     *sql.DB
	logger *zap.Logger
	sql    string
}

func newDbClient(db *sql.DB, sql string, logger *zap.Logger) dbClient {
	_ = "STUB: not implemented"
	return *new(dbClient)
}

func (cl dbSQLClient) metricRows(ctx context.Context, args ...any) ([]metricRow, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The Postgres driver returns a []uint8 (a string) for decimal and numeric types,
// which we want to render as strings. e.g. "4.1" instead of "[52, 46, 49]".
// Other slice types get the same treatment.

type reusableRow struct {
	attrs    map[string]func() string
	scanDest []any
}

func (row reusableRow) toMetricRow() metricRow { _ = "STUB: not implemented"; return *new(metricRow) }
