// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package storagetest // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/storagetest"

import (
	"context"
	"errors"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension/xextension/storage"
)

var errClientClosed = errors.New("client closed")

type TestClient struct {
	cache    map[string][]byte
	cacheMux sync.Mutex

	kind component.Kind
	id   component.ID
	name string

	storageFile string

	closed bool
}

// NewInMemoryClient creates a storage.Client that functions as a map[string][]byte
// This is useful for tests that do not involve collector restart behavior.
func NewInMemoryClient(kind component.Kind, id component.ID, name string) *TestClient {
	_ = "STUB: not implemented"
	return nil
}

// NewFileBackedClient creates a storage.Client that will load previous
// storage contents upon creation and save storage contents when closed.
// It also has metadata which may be used to validate test expectations.
func NewFileBackedClient(kind component.Kind, id component.ID, name, storageDir string) *TestClient {
	_ = "STUB: not implemented"
	return nil
}

// Attempt to load previous storage content

// Assume no previous storage content

// Assume no previous storage content

func (p *TestClient) Get(_ context.Context, key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *TestClient) Set(_ context.Context, key string, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *TestClient) Delete(_ context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *TestClient) Batch(_ context.Context, ops ...*storage.Operation) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *TestClient) Close(_ context.Context) error { _ = "STUB: not implemented"; return nil }

const clientCreatorID = "client_creator_id"

func setCreatorID(ctx context.Context, client storage.Client, creatorID component.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// CreatorID is the component.ID of the extension that created the component
func CreatorID(ctx context.Context, client storage.Client) (component.ID, error) {
	_ = "STUB: not implemented"
	return *new(component.ID), nil
}
