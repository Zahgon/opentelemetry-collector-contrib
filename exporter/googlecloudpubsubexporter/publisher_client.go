// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudpubsubexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudpubsubexporter"

import (
	"context"

	"cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"
	"github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
)

// publisherClient subset of `pubsub.PublisherClient`
type publisherClient interface {
	Close() error
	Publish(ctx context.Context, in *pubsubpb.PublishRequest, opts ...gax.CallOption) (*pubsubpb.PublishResponse, error)
}

// wrappedPublisherClient allows to override the close function
type wrappedPublisherClient struct {
	publisherClient
	closeFn func() error
}

func (c *wrappedPublisherClient) Close() error { _ = "STUB: not implemented"; return nil }

func newPublisherClient(ctx context.Context, config *Config, userAgent string) (publisherClient, error) {
	_ = "STUB: not implemented"
	return *new(publisherClient), nil
}

// In new v2 client the Publish is moved to the TopicAdmin client

func generateClientOptions(config *Config, userAgent string) ([]option.ClientOption, func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// we need to be able to properly close the grpc client otherwise it'll leak goroutines
