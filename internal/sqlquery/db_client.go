// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlquery // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/sqlquery"

import (
	"context"

	// Do not register any Db drivers here: users should register the ones that are applicable to them.
	"go.uber.org/zap"
)

type StringMap map[string]string

type DbClient interface {
	QueryRows(ctx context.Context, args ...any) ([]StringMap, error)
}

type DbSQLClient struct {
	Db        Db
	Logger    *zap.Logger
	Telemetry TelemetryConfig
	SQL       string
}

func NewDbClient(db Db, sql string, logger *zap.Logger, telemetry TelemetryConfig) DbClient {
	_ = "STUB: not implemented"
	return *new(DbClient)
}

func (cl DbSQLClient) QueryRows(ctx context.Context, args ...any) ([]StringMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cl DbSQLClient) prepareQueryFields(sql string, args []any) []zap.Field {
	_ = "STUB: not implemented"
	return nil
}

// This is only used for testing, but need to be exposed to other packages.
type FakeDBClient struct {
	RequestCounter      int
	StringMaps          [][]StringMap
	Err                 error
	ErrNullValueWarning bool
}

func (c *FakeDBClient) QueryRows(context.Context, ...any) ([]StringMap, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
