// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudpubsubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver"

import (
	"context"

	"google.golang.org/api/option"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver/internal"
)

// wrappedSubscriberClient allows to override the close function
type wrappedSubscriberClient struct {
	internal.SubscriberClient
	closeFn func() error
}

func (c *wrappedSubscriberClient) Close() error { _ = "STUB: not implemented"; return nil }

func newSubscriberClient(ctx context.Context, config *Config, userAgent string) (internal.SubscriberClient, error) {
	_ = "STUB: not implemented"
	return *new(internal.SubscriberClient), nil
}

func generateClientOptions(config *Config, userAgent string) ([]option.ClientOption, func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// we need to be able to properly close the grpc client otherwise it'll leak goroutines
