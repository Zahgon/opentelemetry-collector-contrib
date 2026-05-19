// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package redisstorageextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/redisstorageextension"

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.uber.org/zap"
)

type redisStorage struct {
	cfg    *Config
	logger *zap.Logger
	client *redis.Client
}

// Ensure this storage extension implements the appropriate interface
var _ storage.Extension = (*redisStorage)(nil)

func newRedisStorage(logger *zap.Logger, config *Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}

// Start runs cleanup if configured
func (rs *redisStorage) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown will close any open databases
func (rs *redisStorage) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

type redisClient struct {
	client     *redis.Client
	prefix     string
	expiration time.Duration
}

var _ storage.Client = redisClient{}

func (rc redisClient) Get(ctx context.Context, key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rc redisClient) Set(ctx context.Context, key string, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (rc redisClient) Delete(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (rc redisClient) Batch(ctx context.Context, ops ...*storage.Operation) error {
	_ = "STUB: not implemented"
	return nil
}

// once the pipeline has been executed, we need to fetch all the values
// and set them on the op

// the output of Bucket.Get is only valid within a transaction, so we need to make a copy
// to be able to return the value

func (redisClient) Close(context.Context) error {
	_ = "STUB: not implemented"

	// GetClient returns a storage client for an individual component
	return nil
}

func (rs *redisStorage) GetClient(_ context.Context, kind component.Kind, ent component.ID, name string) (storage.Client, error) {
	_ = "STUB: not implemented"
	return *new(storage.Client), nil
}

func (rs *redisStorage) getPrefix(ent component.ID, kind, name string) string {
	_ = "STUB: not implemented"
	return ""
}

func kindString(k component.Kind) string { _ = "STUB: not implemented"; return "" }

// not expected
