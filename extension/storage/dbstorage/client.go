// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dbstorage // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/dbstorage"

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib" // Postgres driver
	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.uber.org/zap"
)

type dbStorageClient struct {
	logger  *zap.Logger
	db      *sql.DB
	dialect *dbDialect
}

func newClient(ctx context.Context, logger *zap.Logger, db *sql.DB, driverName, tableName string) (*dbStorageClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create storage table if not exists

// Set Prepared Statements for some regularly used queries for performance optimization

// Get will retrieve data from storage that corresponds to the specified key
func (c *dbStorageClient) Get(ctx context.Context, key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Set will store data. The data can be retrieved using the same key
		nil
}

func (c *dbStorageClient) Set(ctx context.Context, key string, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete will delete data associated with the specified key
func (c *dbStorageClient) Delete(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// Batch executes the specified operations in order. Get operation results are updated in place
func (c *dbStorageClient) Batch(ctx context.Context, ops ...*storage.Operation) error {
	_ = "STUB: not implemented"
	// Start a new transaction
	return nil
}

// In case of any error we should roll back whole transaction to keep DB in consistent state
// In case of successful commit - tx.Rollback() will be a no-op here as tx is already closed

// We should ignore error related already finished transaction here
// It might happened, for example, if Context was canceled outside of Batch() function
// in this case whole transaction will be rolled back by sql package and we'll receive ErrTxDone here,
// which is actually not an issue because transaction was correctly closed with rollback

// Batch optimization when we have set of operations with the same OpType
// It might give us big performance improvement with lower resource utilization on big batches

// Close will close all prepared statements
// Caller is responsible of closing used DB connection after this
func (c *dbStorageClient) Close(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *dbStorageClient) get(ctx context.Context, key string, tx *sql.Tx) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *dbStorageClient) set(ctx context.Context, key string, value []byte, tx *sql.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dbStorageClient) delete(ctx context.Context, key string, tx *sql.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *dbStorageClient) aggregatedBatch(ctx context.Context, tx *sql.Tx, squashOp storage.OpType, ops ...*storage.Operation) error {
	_ = "STUB: not implemented"
	return nil
}

// Iterate over operations in chunks with length = c.dialect.MaxAggregationSize

func (c *dbStorageClient) batchGet(ctx context.Context, tx *sql.Tx, ops ...*storage.Operation) error {
	_ = "STUB: not implemented"
	return nil

	// Form a multi-row SELECT Query
}

// Create helper structs for passing data to query and getting result back

// Make sure that all associated resource are closed on func exit

// Iterate over returned rows

// If we haven't received row for some key - just skip it and leave op.Value as is

// Check for any errors happened during rows scan

func (c *dbStorageClient) batchSet(ctx context.Context, tx *sql.Tx, ops ...*storage.Operation) error {
	_ = "STUB: not implemented"
	return nil

	// Form a multi-row INSERT Query
}

func (c *dbStorageClient) batchDelete(ctx context.Context, tx *sql.Tx, ops ...*storage.Operation) error {
	_ = "STUB: not implemented"
	return nil

	// Form a multi-row DELETE Query
}

// wrapTx will wrap provided Statement in Transaction if it is provided
func wrapTx(stmt *sql.Stmt, tx *sql.Tx) *sql.Stmt { _ = "STUB: not implemented"; return nil }

// canAggregate checks is possible to aggregate batch of Operations into single query
// As for now it's possible to aggregate only whole batch with a single Operation Type
// Returns aggregated Operation Type and aggregation flag
func canAggregate(ops ...*storage.Operation) (storage.OpType, bool) {
	_ = "STUB: not implemented"
	// Empty case
	return *new(storage.OpType), false
}

// Nothing to squash

// Initial value

// If next Operation is different - exit loop

// Continue...

// generatePlaceholders creates SQL placeholder for parametrized queries
// Positional placeholders, like "$N" is used as they are supported by all used SQL drivers
// By default, when `groupSize = 0` will generate N monotonic placeholders, i.e. "$1, $2, ... $n"
// If `groupSize > 0` - will generate placeholders groups with size = `groupSize`,
// i.e for `groupSize = 2` - "($1, $2), ($3, $4), ... ($n*groupSize-1, $n*groupSize)"
func generatePlaceholders(n, groupSize int) string { _ = "STUB: not implemented"; return "" }

// Simple case when we don'e need to group placeholders

// Complex case when we need to group placeholders with defined group size
