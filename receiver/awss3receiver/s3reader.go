// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awss3receiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awss3receiver"

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type s3TimeBasedReader struct {
	logger *zap.Logger

	listObjectsClient              ListObjectsAPI
	singleObjectClient             SingleObjectAPI
	s3Bucket                       string
	s3Prefix                       string
	s3PartitionFormat              string
	S3PartitionTimeLocation        *time.Location
	filePrefix                     string
	filePrefixIncludeTelemetryType bool
	startTime                      time.Time
	endTime                        time.Time
	notifier                       statusNotifier
	tagObjectAfterIngestion        bool
	skipIngestingTaggedObjects     bool
}

func newS3TimeBasedReader(ctx context.Context, notifier statusNotifier, logger *zap.Logger, cfg *Config) (*s3TimeBasedReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// readAll implements the s3Reader interface
func (s3Reader *s3TimeBasedReader) readAll(ctx context.Context, telemetryType string, dataCallback s3ObjectCallback) error {
	_ = "STUB: not implemented"
	return nil
}

func (s3Reader *s3TimeBasedReader) readTelemetryForTime(ctx context.Context, t time.Time, telemetryType string, dataCallback s3ObjectCallback) error {
	_ = "STUB: not implemented"
	return nil
}

// Don't return error as the object was processed successfully

func (s3Reader *s3TimeBasedReader) getObjectPrefixForTime(t time.Time, telemetryType string) string {
	_ = "STUB: not implemented"
	return ""
}

// Retrieve the configured S3 prefix (may be empty, "/", "//", "logs/", etc.)

// Case 1: No prefix provided → use only timeKey + filePrefix

// Case 2: Prefix contains only slashes (e.g., "/", "//", "///")
// Keep the exact number of slashes and directly append timeKey without adding an extra "/"

// Case 3: Normal prefix (e.g., "logs", "logs/", "/logs/", "//raw//")
// Always add a "/" between prefix and timeKey to build a valid S3 path

func (s3Reader *s3TimeBasedReader) sendStatus(ctx context.Context, status statusNotification) {
	_ = "STUB: not implemented"
	return
}

func getTimeKey(partitionFormat string, t time.Time, location *time.Location) string {
	_ = "STUB: not implemented"
	return ""
}

func determineTimestep(partitionFormat string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}
