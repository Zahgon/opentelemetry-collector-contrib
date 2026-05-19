// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awslambdareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awslambdareceiver"

import (
	"go.opentelemetry.io/collector/component"
)

const s3ARNPrefix = "arn:aws:s3:::"

// defaultS3PathPatterns maps known encoding names to their default S3 path patterns.
// "*" matches exactly one path segment (the AWS account ID in standard AWS log paths).
var defaultS3PathPatterns = map[string]string{
	"vpcflow":         "AWSLogs/*/vpcflowlogs",
	"cloudtrail":      "AWSLogs/*/CloudTrail",
	"elbaccess":       "AWSLogs/*/elasticloadbalancing",
	"waf":             "AWSLogs/*/WAFLogs",
	"networkfirewall": "AWSLogs/*/network-firewall",
}

// S3Encoding defines one entry in the S3 multi-encoding routing table.
type S3Encoding struct {
	// Name identifies the encoding. For known names (vpcflow, cloudtrail, elbaccess, waf,
	// networkfirewall) the default path_pattern is applied automatically.
	Name string `mapstructure:"name"`

	// Encoding is the extension ID for decoding (e.g. "awslogs_encoding/vpcflow").
	// If empty, content is passed through as-is using the built-in raw decoder.
	Encoding string `mapstructure:"encoding"`

	// PathPattern is matched as a prefix against the S3 object key.
	// "*" matches exactly one path segment. Example: "AWSLogs/*/vpcflowlogs"
	// If empty, the default pattern for a known Name is used.
	PathPattern string `mapstructure:"path_pattern"`
}

// resolvePathPattern returns the effective path pattern for this encoding entry.
// Returns the configured PathPattern if set, else the default for known names.
func (e *S3Encoding) resolvePathPattern() string { _ = "STUB: not implemented"; return "" }

// Validate validates an S3Encoding entry.
func (e *S3Encoding) Validate() error { _ = "STUB: not implemented"; return nil }

// Unknown name without an explicit path_pattern is not routable.

// sharedConfig defines configuration options shared between Lambda trigger types.
type sharedConfig struct {
	// Encoding defines the encoding to decode incoming Lambda invocation data.
	Encoding string `mapstructure:"encoding"`
}

// S3Config defines configuration options for the S3 Lambda trigger.
// It supersedes the sharedConfig for the s3 key.
// sharedConfig is embedded with mapstructure:",squash" so the "encoding" key
// remains at the top level of the s3 block in YAML — fully backwards-compatible.
type S3Config struct {
	sharedConfig `mapstructure:",squash"`

	// Encodings defines multiple encoding entries for S3 path-based routing (multi-format mode).
	// Each entry maps a path_pattern prefix to an encoding extension.
	// Mutually exclusive with sharedConfig.Encoding.
	//
	// Only supported for logs signal type. Metrics receivers reject configs that set this field.
	Encodings []S3Encoding `mapstructure:"encodings"`
}

// Validate validates the S3Config.
func (c *S3Config) Validate() error { _ = "STUB: not implemented"; return nil }

// sortedEncodings returns a copy of Encodings sorted by path pattern specificity:
// more-specific patterns first, catch-all "*" last.
// This makes matching order-independent — users can list encodings in any order.
func (c *S3Config) sortedEncodings() []S3Encoding { _ = "STUB: not implemented"; return nil }

// Pre-split patterns once; keyed by pattern string so the map stays correct
// as the sort swaps elements.

// Config is the top-level configuration for the awslambda receiver.
type Config struct {
	// S3 defines configuration for the S3 Lambda trigger.
	S3 S3Config `mapstructure:"s3"`

	// CloudWatch defines configuration for the CloudWatch Logs Lambda trigger.
	CloudWatch sharedConfig `mapstructure:"cloudwatch"`

	// FailureBucketARN is the ARN of the S3 bucket used to store failed Lambda event records.
	FailureBucketARN string `mapstructure:"failure_bucket_arn"`

	_ struct{} // Prevent unkeyed literal initialization.
}

var _ component.Config = (*Config)(nil)

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// getBucketNameFromARN extracts the S3 bucket name from an ARN.
// Example: "arn:aws:s3:::myBucket/folderA" => "myBucket"
func getBucketNameFromARN(arn string) (string, error) { _ = "STUB: not implemented"; return "", nil }
