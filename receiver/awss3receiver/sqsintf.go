// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awss3receiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awss3receiver"

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// sqsClient defines the SQS operations used by s3SQSNotificationReader
type sqsClient interface {
	ReceiveMessage(ctx context.Context, params *sqs.ReceiveMessageInput, optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error)
	DeleteMessage(ctx context.Context, params *sqs.DeleteMessageInput, optFns ...func(*sqs.Options)) (*sqs.DeleteMessageOutput, error)
}

// newSQSClient creates a new SQS client with the provided configuration
func newSQSClient(ctx context.Context, region, endpoint string) (sqsClient, error) {
	_ = "STUB: not implemented"
	return *new(sqsClient), nil
}
