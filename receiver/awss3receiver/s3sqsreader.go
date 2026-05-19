// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awss3receiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awss3receiver"

import (
	"context"

	"go.uber.org/zap"
)

// S3 event notification structure from AWS
// See: https://docs.aws.amazon.com/AmazonS3/latest/userguide/notification-content-structure.html

// s3ObjectData represents an S3 object in the notification
type s3ObjectData struct {
	Key string `json:"key"`
}

// s3BucketData represents an S3 bucket in the notification
type s3BucketData struct {
	Name string `json:"name"`
}

// s3Data represents the S3 specific data in the notification
type s3Data struct {
	Bucket s3BucketData `json:"bucket"`
	Object s3ObjectData `json:"object"`
}

// s3EventRecord represents a single record in an S3 event notification
type s3EventRecord struct {
	EventSource string `json:"eventSource"`
	EventName   string `json:"eventName"`
	S3          s3Data `json:"s3"`
}

// s3EventNotification is the top-level structure for S3 event notifications
type s3EventNotification struct {
	Event   string          `json:"Event"`
	Records []s3EventRecord `json:"Records"`
}

// snsMessage represents the structure of an SNS notification message
type snsMessage struct {
	Type    string `json:"Type"`
	Message string `json:"Message"`
}

// s3SQSNotificationReader listens for SNS notifications about new S3 objects
type s3SQSNotificationReader struct {
	logger                     *zap.Logger
	s3Client                   SingleObjectAPI
	sqsClient                  sqsClient
	queueURL                   string
	s3Bucket                   string
	s3Prefix                   string
	maxNumberOfMessages        int32
	waitTimeSeconds            int32
	tagObjectAfterIngestion    bool
	skipIngestingTaggedObjects bool
}

func newS3SQSReader(ctx context.Context, logger *zap.Logger, cfg *Config) (*s3SQSNotificationReader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use configured values or defaults for SQS polling parameters
// Default to 10 messages

// Default to 20 seconds

func (r *s3SQSNotificationReader) readAll(ctx context.Context, _ string, callback s3ObjectCallback) error {
	_ = "STUB: not implemented"
	return nil
}

// Add a small sleep to avoid tight loops on persistent errors

// First try to parse as direct S3 event notification

// If direct parsing failed, try to extract from SNS notification format

// Track whether all records were successfully processed.
// Only delete the message if all records succeed to prevent data loss.

// Process each S3 object notification

// Decode the URL-encoded S3 key

// Swallow no such key errors as nothing more can be done

// Swallow no such key errors as nothing more can be done

// Don't mark as failed as the object was processed successfully

// Only delete the message if all records were successfully processed.
// If any record failed, leave the message in the queue for retry.
