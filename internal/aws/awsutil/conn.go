// Copyright The OpenTelemetry Authors
// Portions of this file Copyright 2018-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package awsutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil"

import (
	"context"
	"net/http"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"go.uber.org/zap"
)

// newHTTPClient returns new HTTP client instance with provided configuration.
func newHTTPClient(logger *zap.Logger, maxIdle, requestTimeout int, noVerify bool,
	proxyAddress string,
) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// is not enabled by default as we configure TLSClientConfig for supporting SSL to data plane.
// http2.ConfigureTransport will setup transport layer to use HTTP2

// GetProxyFunc returns a proxy function for use in http.Transport.
// When an explicit proxy address is configured, it returns http.ProxyURL.
// Otherwise, it returns http.ProxyFromEnvironment which respects NO_PROXY.
func GetProxyFunc(proxyAddress string) (func(*http.Request) (*url.URL, error), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetAWSConfig(ctx context.Context, logger *zap.Logger, settings *AWSSessionSettings) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}

// getAWSConfig returns AWS config instance. This is separated from GetAWSConfig to allow
// easier testing with mocked AssumeRoleAPIClient.
func getAWSConfig(ctx context.Context, logger *zap.Logger, settings *AWSSessionSettings, getAssumeRoleAPIClient func(getAssumeRoleAPIClient aws.Config) stscreds.AssumeRoleAPIClient) (aws.Config, error) {
	_ = "STUB: not implemented"
	return *new(aws.Config), nil
}
