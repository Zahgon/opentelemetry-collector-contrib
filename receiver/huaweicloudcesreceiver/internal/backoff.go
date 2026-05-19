// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/huaweicloudcesreceiver/internal"

import (
	"context"

	"github.com/cenkalti/backoff/v4"
	"go.opentelemetry.io/collector/config/configretry"
	"go.uber.org/zap"
)

func NewExponentialBackOff(backOffConfig *configretry.BackOffConfig) *backoff.ExponentialBackOff {
	_ = "STUB: not implemented"
	return nil
}

// Generic function to make an API call with exponential backoff and context cancellation handling.
func MakeAPICallWithRetry[T any](
	ctx context.Context,
	shutdownChan chan struct{},
	logger *zap.Logger,
	apiCall func() (*T, error),
	isThrottlingError func(error) bool,
	backOffConfig *backoff.ExponentialBackOff,
) (*T, error) {
	_ = "STUB: not implemented"
	// Immediately check for context cancellation or server shutdown.
	return nil, nil
}

// Make the initial API call.

// If the error is not due to request throttling, return the error.

// Initialize the backoff mechanism for retrying the API call.

// Retry loop for handling throttling errors.

// Handle context cancellation or shutdown before retrying.

// Retry the API call.
