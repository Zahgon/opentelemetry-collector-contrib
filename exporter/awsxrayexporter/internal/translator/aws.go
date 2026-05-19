// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter/internal/translator"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
)

func makeAws(attributes map[string]pcommon.Value, resource pcommon.Resource, logGroupNames []string) (map[string]pcommon.Value, *awsxray.AWSData) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deterministically handled with if else above

// Deterministically handled with if else above

// not AWS so return nil

// Favor Semantic Conventions for specific SQS and DynamoDB attributes.

// EC2 - add ec2 metadata to xray request if
//       1. cloud.platfrom is set to "aws_ec2" or
//       2. there is an non-blank host/instance id found

// ECS

// Beanstalk

// EKS or native Kubernetes

// Since we must couple log group ARNs and Log Group Names in the same CWLogs object, we first try to derive the
// names from the ARN, then fall back to recording the names, if they do not exist in the resource
// then pull from them from config.

// Convention for SDK name for xray SDK information is e.g., `X-Ray SDK for Java`, `X-Ray for Go`.
// We fill in with e.g, `opentelemetry for java` by using the conventionsv112

func getLogGroupNamesOrArns(logGroupNamesOrArns string) []string {
	_ = "STUB: not implemented"
	// Split the input string by '&'
	return nil
}

// Filter out empty strings

// Normalize value to slice.
// 1. String values are converted to a slice so that we can also handle resource
// attributes that are set using the OTEL_RESOURCE_ATTRIBUTES
// (multiple log group names or arns are separate by & like this "log-group1&log-group2&log-group3")
// 2. Slices are kept as they are
// 3. Other types will result in a empty slice so that we avoid panic.
func normalizeToSlice(v pcommon.Value) pcommon.Slice {
	_ = "STUB: not implemented"
	return *new(pcommon.Slice)
}

// Given an array of log group ARNs, create a corresponding amount of LogGroupMetadata objects with log_group and arn
// populated, or given an array of just log group names, create the LogGroupMetadata objects with arn omitted
func getLogGroupMetadata(logGroups pcommon.Slice, isArn bool) []awsxray.LogGroupMetadata {
	_ = "STUB: not implemented"
	return nil
}

// Log group name will always be in the 7th position of the ARN
// https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/iam-access-control-overview-cwl.html#CWL_ARN_Format
func parseLogGroup(arn string) string { _ = "STUB: not implemented"; return "" }
