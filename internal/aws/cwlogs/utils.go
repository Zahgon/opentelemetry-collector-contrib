// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cwlogs // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/cwlogs"

// Added function to check if value is an accepted number of log retention days
func ValidateRetentionValue(input int32) error { _ = "STUB: not implemented"; return nil }

// Check if the tags input is valid
func ValidateTagsInput(input map[string]string) error { _ = "STUB: not implemented"; return nil }

// The regex for the Key and Value requires "alphanumerics, whitespace, and _.:/=+-!" as noted here: https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateLogGroup.html#:~:text=%5E(%5B%5Cp%7BL%7D%5Cp%7BZ%7D%5Cp%7BN%7D_.%3A/%3D%2B%5C%2D%40%5D%2B)%24
