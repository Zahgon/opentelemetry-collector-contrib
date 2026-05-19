// Copyright The OpenTelemetry Authors
// Portions of this file Copyright 2018-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package awsxray // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/smithy-go/middleware"
	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"
)

// Constant prefixes used to identify information in user-agent
const (
	agentPrefix   = "xray-otel-exporter/"
	execEnvPrefix = " exec-env/"
	osPrefix      = " OS/"
)

// XRayClient contains the AWS XRay API calls that exporter/awsxrayexporter uses
type XRayClient interface {
	PutTraceSegments(ctx context.Context, params *xray.PutTraceSegmentsInput, optFns ...func(*xray.Options)) (*xray.PutTraceSegmentsOutput, error)
	PutTelemetryRecords(ctx context.Context, params *xray.PutTelemetryRecordsInput, optFns ...func(*xray.Options)) (*xray.PutTelemetryRecordsOutput, error)
}

func getModVersion() string { _ = "STUB: not implemented"; return "" }

// NewXRayClient creates a new instance of the XRay client with an AWS configuration and session.
func NewXRayClient(_ *zap.Logger, cfg aws.Config, buildInfo component.BuildInfo) XRayClient {
	_ = "STUB: not implemented"
	return *new(XRayClient)
}

type withTimestampRequestHeader struct{}

var _ middleware.FinalizeMiddleware = (*withTimestampRequestHeader)(nil)

func (*withTimestampRequestHeader) ID() string { _ = "STUB: not implemented"; return "" }

func (*withTimestampRequestHeader) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	middleware.FinalizeOutput, middleware.Metadata, error,
) {
	_ = "STUB: not implemented"
	return *new(middleware.FinalizeOutput), *new(middleware.Metadata), nil
}

func WithTimestampRequestHeader(pos middleware.RelativePosition) func(options *xray.Options) {
	_ = "STUB: not implemented"
	return nil
}

type addToUserAgentHeader struct {
	id, val string
}

var _ middleware.SerializeMiddleware = (*addToUserAgentHeader)(nil)

func (a *addToUserAgentHeader) ID() string { _ = "STUB: not implemented"; return "" }

func (a *addToUserAgentHeader) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (out middleware.SerializeOutput, metadata middleware.Metadata, err error) {
	_ = "STUB: not implemented"
	return *new(middleware.SerializeOutput), *new(middleware.Metadata), nil
}

func AddToUserAgentHeader(id, val string, pos middleware.RelativePosition) func(options *xray.Options) {
	_ = "STUB: not implemented"
	return nil
}
