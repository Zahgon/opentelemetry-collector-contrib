// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awss3receiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awss3receiver"

import (
	"time"

	"go.opentelemetry.io/collector/component"
)

// S3DownloaderConfig contains aws s3 downloader related config to controls things
// like bucket, prefix, batching, connections, retries, etc.
type S3DownloaderConfig struct {
	Region                         string `mapstructure:"region"`
	S3Bucket                       string `mapstructure:"s3_bucket"`
	S3Prefix                       string `mapstructure:"s3_prefix"`
	S3PartitionFormat              string `mapstructure:"s3_partition_format"`
	S3PartitionTimezone            string `mapstructure:"s3_partition_timezone"`
	FilePrefix                     string `mapstructure:"file_prefix"`
	FilePrefixIncludeTelemetryType bool   `mapstructure:"file_prefix_include_telemetry_type"`
	Endpoint                       string `mapstructure:"endpoint"`
	EndpointPartitionID            string `mapstructure:"endpoint_partition_id"`
	S3ForcePathStyle               bool   `mapstructure:"s3_force_path_style"`
	TagObjectAfterIngestion        bool   `mapstructure:"tag_object_after_ingestion"`
	SkipIngestingTaggedObjects     bool   `mapstructure:"skip_ingesting_tagged_objects"`
}

// SQSConfig holds SQS queue configuration for receiving object change notifications.
type SQSConfig struct {
	// QueueURL is the URL of the SQS queue to receive S3 notifications.
	QueueURL string `mapstructure:"queue_url"`
	// Region specifies the AWS region of the SQS queue.
	Region string `mapstructure:"region"`
	// Endpoint is the optional custom endpoint for SQS (useful for testing).
	Endpoint string `mapstructure:"endpoint"`
	// WaitTimeSeconds specifies the duration (in seconds) for long polling SQS messages.
	// Maximum is 20 seconds. Default is 20 seconds.
	WaitTimeSeconds *int64 `mapstructure:"wait_time_seconds"`
	// MaxNumberOfMessages specifies the maximum number of messages to receive in a single poll.
	// Valid values: 1-10. Default is 10.
	MaxNumberOfMessages *int64 `mapstructure:"max_number_of_messages"`
}

// Notifications groups optional notification sources.
type Notifications struct {
	OpAMP *component.ID `mapstructure:"opampextension"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// Encoding defines the encoding configuration for the file receiver.
type Encoding struct {
	Extension component.ID `mapstructure:"extension"`
	Suffix    string       `mapstructure:"suffix"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// Config defines the configuration for the file receiver.
type Config struct {
	S3Downloader  S3DownloaderConfig `mapstructure:"s3downloader"`
	StartTime     string             `mapstructure:"starttime"`
	EndTime       string             `mapstructure:"endtime"`
	Encodings     []Encoding         `mapstructure:"encodings"`
	Notifications Notifications      `mapstructure:"notifications"`
	// SQS configures receiving S3 object change notifications via an SQS queue.
	SQS *SQSConfig `mapstructure:"sqs"`
}

const (
	s3PartitionFormatDefault = "year=%Y/month=%m/day=%d/hour=%H/minute=%M"
)

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func (c Config) Validate() error { _ = "STUB: not implemented"; return nil }

// Check for valid time-based configuration

// If one of StartTime/EndTime is specified, the other must also be specified

// StartTime and SQS cannot be specified together

// Validate StartTime format if specified

// Validate EndTime format if specified

// Validate SQS notifications if configured

// Validate wait time seconds

// Validate max number of messages

func parseTime(timeStr, configName string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}
