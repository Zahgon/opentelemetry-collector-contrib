// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mongodbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbreceiver"

import (
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/scraper/scrapererror"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbreceiver/internal/metadata"
)

var errKeyNotFound = errors.New("could not find key for metric")

var operationsMap = map[string]metadata.AttributeOperation{
	"insert":   metadata.AttributeOperationInsert,
	"queries":  metadata.AttributeOperationQuery,
	"update":   metadata.AttributeOperationUpdate,
	"remove":   metadata.AttributeOperationDelete,
	"getmore":  metadata.AttributeOperationGetmore,
	"commands": metadata.AttributeOperationCommand,
}

var documentMap = map[string]metadata.AttributeOperation{
	"inserted": metadata.AttributeOperationInsert,
	"updated":  metadata.AttributeOperationUpdate,
	"deleted":  metadata.AttributeOperationDelete,
}

var lockTypeMap = map[string]metadata.AttributeLockType{
	"ParallelBatchWriterMode":    metadata.AttributeLockTypeParallelBatchWriteMode,
	"ReplicationStateTransition": metadata.AttributeLockTypeReplicationStateTransition,
	"Global":                     metadata.AttributeLockTypeGlobal,
	"Database":                   metadata.AttributeLockTypeDatabase,
	"Collection":                 metadata.AttributeLockTypeCollection,
	"Mutex":                      metadata.AttributeLockTypeMutex,
	"Metadata":                   metadata.AttributeLockTypeMetadata,
	"oplog":                      metadata.AttributeLockTypeOplog,
}

var lockModeMap = map[string]metadata.AttributeLockMode{
	"R": metadata.AttributeLockModeShared,
	"W": metadata.AttributeLockModeExclusive,
	"r": metadata.AttributeLockModeIntentShared,
	"w": metadata.AttributeLockModeIntentExclusive,
}

const (
	collectMetricError          = "failed to collect metric %s: %w"
	collectMetricWithAttributes = "failed to collect metric %s with attribute(s) %s: %w"
)

// DBStats
func (s *mongodbScraper) recordCollections(now pcommon.Timestamp, doc bson.M, dbName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordDataSize(now pcommon.Timestamp, doc bson.M, dbName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordStorageSize(now pcommon.Timestamp, doc bson.M, dbName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordObjectCount(now pcommon.Timestamp, doc bson.M, dbName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordIndexCount(now pcommon.Timestamp, doc bson.M, dbName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordIndexSize(now pcommon.Timestamp, doc bson.M, dbName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordExtentCount(now pcommon.Timestamp, doc bson.M, dbName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	// Mongo version 4.4+ no longer returns numExtents since it is part of the obsolete MMAPv1
	// https://www.mongodb.com/docs/manual/release-notes/4.4-compatibility/#mmapv1-cleanup
	return
}

// ServerStatus
func (s *mongodbScraper) recordConnections(now pcommon.Timestamp, doc bson.M, dbName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordMemoryUsage(now pcommon.Timestamp, doc bson.M, dbName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// convert from mebibytes to bytes

func (s *mongodbScraper) recordDocumentOperations(now pcommon.Timestamp, doc bson.M, dbName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordSessionCount(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// mongodb is using a different storage engine and this metric cannot be collected

func (s *mongodbScraper) recordLatencyTime(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// Admin Stats
func (s *mongodbScraper) recordOperations(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// For telegraf metrics to get QPS for opcounters
// Store current counts for next iteration

func (s *mongodbScraper) recordOperationsRepl(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordReplOperationPerSecond(now pcommon.Timestamp, operationVal string, currentCount int64) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordFlushesPerSecond(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordOperationPerSecond(now pcommon.Timestamp, operationVal string, currentCount int64) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordActiveWrites(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordActiveReads(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordWTCacheBytes(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordPageFaults(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordCacheOperations(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// mongodb is using a different storage engine and this metric cannot be collected

func (s *mongodbScraper) recordGlobalLockTime(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordCursorCount(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordCursorTimeoutCount(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordNetworkCount(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordUptime(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordHealth(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// Lock Metrics are only supported by MongoDB v3.2+
func (s *mongodbScraper) recordLockAcquireCounts(now pcommon.Timestamp, doc bson.M, dBName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// Continue if the lock type is not supported by current server's MongoDB version

// MongoDB only publishes this lock metric is it is available.
// Do not raise error when key is not found

func (s *mongodbScraper) recordLockAcquireWaitCounts(now pcommon.Timestamp, doc bson.M, dBName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// Continue if the lock type is not supported by current server's MongoDB version

// MongoDB only publishes this lock metric is it is available.
// Do not raise error when key is not found

func (s *mongodbScraper) recordLockTimeAcquiringMicros(now pcommon.Timestamp, doc bson.M, dBName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// Continue if the lock type is not supported by current server's MongoDB version

// MongoDB only publishes this lock metric is it is available.
// Do not raise error when key is not found

func (s *mongodbScraper) recordLockDeadlockCount(now pcommon.Timestamp, doc bson.M, dBName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// Continue if the lock type is not supported by current server's MongoDB version

// MongoDB only publishes this lock metric is it is available.
// Do not raise error when key is not found

// Index Stats
func (s *mongodbScraper) recordIndexAccess(now pcommon.Timestamp, documents []bson.M, dbName, collectionName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// Top Stats
func (s *mongodbScraper) recordOperationTime(now pcommon.Timestamp, doc bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func aggregateOperationTimeValues(document bson.M, collectionPathNames []string, operationMap map[string]metadata.AttributeOperation) (map[string]int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getOperationTimeValues(document bson.M, collectionPathName, operation string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func digForCollectionPathNames(document bson.M) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func collectMetric(document bson.M, path []string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func dig(document bson.M, path []string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func digBsonD(value bson.D, remainingPath []string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func parseInt(val any) (int64, error) { _ = "STUB: not implemented"; return 0, nil }
