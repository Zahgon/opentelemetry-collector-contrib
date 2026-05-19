// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package credentials // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/sumologicextension/internal/credentials"

import (
	"go.uber.org/zap"
)

const (
	DefaultCollectorDataDirectory = ".sumologic-otel-collector/"
)

func GetDefaultCollectorCredentialsDirectory() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// LocalFsStore implements Store interface and can be used to store and retrieve
// collector credentials from local file system.
//
// Files are stored locally in collectorCredentialsDirectory.
type LocalFsStore struct {
	collectorCredentialsDirectory string
	logger                        *zap.Logger
}

type LocalFsStoreOpt func(*LocalFsStore)

func WithLogger(l *zap.Logger) LocalFsStoreOpt {
	_ = "STUB: not implemented"
	return *new(LocalFsStoreOpt)
}

func WithCredentialsDirectory(dir string) LocalFsStoreOpt {
	_ = "STUB: not implemented"
	return *new(LocalFsStoreOpt)
}

func NewLocalFsStore(opts ...LocalFsStoreOpt) (Store, error) {
	_ = "STUB: not implemented"
	return *new(Store), nil
}

// Check checks if collector credentials can be found under a name being a hash
// of provided key inside collectorCredentialsDirectory.
func (cr LocalFsStore) Check(key string) bool { _ = "STUB: not implemented"; return false }

// Get retrieves collector credentials stored in local file system and then
// decrypts it using a hash of provided key.
func (cr LocalFsStore) Get(key string) (CollectorCredentials, error) {
	_ = "STUB: not implemented"
	return *new(CollectorCredentials), nil
}

// Store stores collector credentials in a file in directory as specified
// in CollectorCredentialsDirectory.
// The credentials are encrypted using the provided key.
func (cr LocalFsStore) Store(key string, creds CollectorCredentials) error {
	_ = "STUB: not implemented"
	return nil
}

func (cr LocalFsStore) Delete(key string) error { _ = "STUB: not implemented"; return nil }

// Validate checks if the store is operating correctly
// This mostly means file permissions and the like
func (cr LocalFsStore) Validate() error { _ = "STUB: not implemented"; return nil }

// ensureDir checks if the specified directory exists and has the right permissions
// if it doesn't then it tries to create it.
func ensureDir(path string) error { _ = "STUB: not implemented"; return nil }

// If the directory doesn't have the execution bit then
// set it so that we can 'exec' into it.
