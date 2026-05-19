// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsinfo // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/ecsInfo"

import (
	"context"
	"net/http"
)

const (
	ecsAgentEndpoint         = "http://%s:51678/v1/metadata"
	ecsAgentTaskInfoEndpoint = "http://%s:51678/v1/tasks"
	taskStatusRunning        = "RUNNING"
	maxHTTPResponseLength    = 5 * 1024 * 1024 // 5MB
)

// There are two formats of ContainerInstance ARN (https://docs.aws.amazon.com/AmazonECS/latest/userguide/ecs-account-settings.html#ecs-resource-ids)
// arn:aws:ecs:region:aws_account_id:container-instance/container-instance-id
// arn:aws:ecs:region:aws_account_id:container-instance/cluster-name/container-instance-id
// This function will return "container-instance-id" for both ARN format

func GetContainerInstanceIDFromArn(arn string) (containerInstanceID string, err error) {
	_ = "STUB: not implemented"
	// When splitting the ARN with ":", the 6th segments could be either:
	// container-instance/47c0ab6e-2c2c-475e-9c30-b878fa7a8c3d or
	// container-instance/cluster-name/47c0ab6e-2c2c-475e-9c30-b878fa7a8c3d
	return "", nil
}

// Further splitting tmpResult with "/", it could be splitted into either 2 or 3
// Characters of "cluster-name" is only allowed to be letters, numbers and hyphens

// Check the channel is closed or not.
func isClosed(ch <-chan bool) bool { _ = "STUB: not implemented"; return false }

type doer interface {
	Do(request *http.Request) (*http.Response, error)
}

func request(ctx context.Context, endpoint string, client doer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// value -1 indicates that the length is unknown, see https://golang.org/src/net/http/response.go
// In this case, we read until the limit is reached
// This might happen with chunked responses from ECS Introspection API

func clientGet(ctx context.Context, url string, client doer) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
