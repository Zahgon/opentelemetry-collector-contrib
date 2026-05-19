// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal/metadata"

import (
	"go.mongodb.org/atlas/mongodbatlas"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// metricRecordFunc records the data point to the metric builder at the supplied timestamp
type metricRecordFunc func(*MetricsBuilder, *mongodbatlas.DataPoints, pcommon.Timestamp)

// getRecordFunc returns the metricRecordFunc that matches the metric name. Nil if none is found.
func getRecordFunc(metricName string) metricRecordFunc {
	_ = "STUB: not implemented"

	// MongoDB CPU usage. For hosts with more than one CPU core, these values can exceed 100%.
	return *new(metricRecordFunc)
}

// MongoDB CPU usage scaled to a range of 0% to 100%. Atlas computes this value by dividing by the number of CPU cores.

// Context: Process

// Rate of asserts for a MongoDB process found in the asserts document that the serverStatus command generates.

// Amount of data flushed in the background.

// Amount of bytes in the WiredTiger storage engine cache and tickets found in the wiredTiger.cache and wiredTiger.concurrentTransactions documents that the serverStatus command generates.

// Number of connections to a MongoDB process found in the connections document that the serverStatus command generates.

// Number of cursors for a MongoDB process found in the metrics.cursor document that the serverStatus command generates.

// Numbers of Memory Issues and Page Faults for a MongoDB process.

// Number of operations waiting on locks for the MongoDB process that the serverStatus command generates. Cloud Manager computes these values based on the type of storage engine.

// Number of index btree operations.

// Number of journaling operations.

// Amount of memory for a MongoDB process found in the mem document that the serverStatus command collects.

// Amount of throughput for MongoDB process found in the network document that the serverStatus command collects.

// Durations and throughput of the MongoDB process' oplog.

// Number of database operations on a MongoDB process since the process last started.

// Rate of database operations on a MongoDB process since the process last started found in the opcounters document that the serverStatus command collects.

// Rate of database operations on MongoDB secondaries found in the opcountersRepl document that the serverStatus command collects.

// Average rate of documents returned, inserted, updated, or deleted per second during a selected time period.

// Average rate for operations per second during a selected time period that perform a sort but cannot perform the sort using an index.

// Average execution time in milliseconds per read, write, or command operation during a selected time period.

// Number of times the host restarted within the previous hour.

// Average rate per second to scan index items during queries and query-plan evaluations found in the value of totalKeysExamined from the explain command.

// Average rate of documents scanned per second during queries and query-plan evaluations found in the value of totalDocsExamined from the explain command.

// Ratio of the number of index items scanned to the number of documents returned.

// Ratio of the number of documents scanned to the number of documents returned.

// CPU usage of processes on the host. For hosts with more than one CPU core, this value can exceed 100%.

// CPU usage of processes on the host scaled to a range of 0 to 100% by dividing by the number of CPU cores.

// Physical memory usage, in bytes, that the host uses.

// Average rate of physical bytes per second that the eth0 network interface received and transmitted.

// Total amount of memory that swap uses.

// Total amount of memory written and read from swap.

// Memory usage, in bytes, that Atlas Search processes use.

// Disk space, in bytes, that Atlas Search indexes use.
// FTS_DISK_UTILIZATION is the documented field name, but FTS_DISK_USAGE is what is returned from the API.
// Including both so if the API changes to match the documentation this metric is still collected.

// Percentage of CPU that Atlas Search processes use.

// Process Disk Measurements (https://docs.atlas.mongodb.com/reference/api/process-disks-measurements/)

// Measures throughput of I/O operations for the disk partition used for MongoDB.

// Measures throughput of data read and written to the disk partition (not cache) used by MongoDB.

// This is a calculated metric that is the sum of the read and write throughput.

// Measures the queue depth of the disk partition used by MongoDB.

// Measures latency per operation type of the disk partition used by MongoDB.

// The percentage of time during which requests are being issued to and serviced by the partition.
// This includes requests from any process, not just MongoDB processes.

// Measures the free disk space and used disk space on the disk partition used by MongoDB.

// Process Database Measurements (https://docs.atlas.mongodb.com/reference/api/process-disks-measurements/)

func MeasurementsToMetric(mb *MetricsBuilder, meas *mongodbatlas.Measurements) error {
	_ = "STUB: not implemented"
	return nil
}

func addDataPoint(mb *MetricsBuilder, meas *mongodbatlas.Measurements, recordFunc metricRecordFunc) error {
	_ = "STUB: not implemented"
	return nil
}
