// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filestorage // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage"

import (
	"context"
	"sync"
	"time"

	"go.etcd.io/bbolt"
	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.uber.org/zap"
)

var defaultBucket = []byte(`default`)

const (
	TempDbPrefix = "tempdb"

	elapsedKey       = "elapsed"
	directoryKey     = "directory"
	tempDirectoryKey = "tempDirectory"

	oneMiB = 1048576
)

type fileStorageClient struct {
	logger          *zap.Logger
	compactionMutex sync.RWMutex
	db              *bbolt.DB
	compactionCfg   *CompactionConfig
	openTimeout     time.Duration
	stopCh          chan struct{}
	wg              sync.WaitGroup
	closed          bool
}

func bboltOptions(timeout time.Duration, noSync bool) *bbolt.Options {
	_ = "STUB: not implemented"
	return nil
}

func newClient(logger *zap.Logger, filePath string, timeout time.Duration, compactionCfg *CompactionConfig, noSync bool) (*fileStorageClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get will retrieve data from storage that corresponds to the specified key
func (c *fileStorageClient) Get(ctx context.Context, key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set will store data. The data can be retrieved using the same key
func (c *fileStorageClient) Set(ctx context.Context, key string, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete will delete data associated with the specified key
func (c *fileStorageClient) Delete(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// Batch executes the specified operations in order. Get operation results are updated in place
func (c *fileStorageClient) Batch(_ context.Context, ops ...*storage.Operation) error {
	_ = "STUB: not implemented"
	return nil
}

// the output of Bucket.Get is only valid within a transaction, so we need to make a copy
// to be able to return the value

// Close will close the database
func (c *fileStorageClient) Close(_ context.Context) error { _ = "STUB: not implemented"; return nil }

// Compact database. Use temporary file as helper as we cannot replace database in-place
func (c *fileStorageClient) Compact(compactionDirectory string, timeout time.Duration, maxTransactionSize int64) error {
	_ = "STUB: not implemented"
	return nil
}

// create temporary file in compactionDirectory

// File still exists and needs to be removed

// use temporary file as compaction target

// cannot reuse newClient as db shouldn't contain any bucket

// replace current db file with compacted db file
// we reopen the DB file irrespective of the success of the replace, as we can't leave it closed

// Leave c.db pointing at the old (closed) DB so that callers get
// errors from the closed DB instead of panicking on a nil pointer.

// if we only failed the remove, we're mostly ok and should just log a warning

// startCompactionLoop provides asynchronous compaction function
func (c *fileStorageClient) startCompactionLoop() { _ = "STUB: not implemented"; return }

// shouldCompact checks whether the conditions for online compaction are met
func (c *fileStorageClient) shouldCompact() bool { _ = "STUB: not implemented"; return false }

func (c *fileStorageClient) getDbSize() (totalSizeResult, dataSizeResult int64, errResult error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// moveFileWithFallback is the equivalent of os.Rename, except it falls back to
// a non-atomic Truncate and Copy if the arguments are on different filesystems
func moveFileWithFallback(src, dest string) error { _ = "STUB: not implemented"; return nil }

// EXDEV is the error code for linking cross-device, we want to continue if we encounter it
// other errors, we simply return as-is

// if we're trying to rename across devices, try truncate and copy instead
// assuming the file isn't too big
