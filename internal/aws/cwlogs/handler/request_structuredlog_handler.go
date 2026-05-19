// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package handler // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/cwlogs/handler"

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/smithy-go/middleware"
)

type withStructuredLogHeader struct{}

var _ middleware.SerializeMiddleware = (*withStructuredLogHeader)(nil)

func (*withStructuredLogHeader) ID() string { _ = "STUB: not implemented"; return "" }

func (*withStructuredLogHeader) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (middleware.SerializeOutput, middleware.Metadata, error) {
	_ = "STUB: not implemented"
	return *new(middleware.SerializeOutput), *new(middleware.Metadata), nil
}

// WithStructuredLogHeader sets the `x-amzn-logs-format` header to `json/emf`
func WithStructuredLogHeader(pos middleware.RelativePosition) func(options *cloudwatchlogs.Options) {
	_ = "STUB: not implemented"
	return nil
}
